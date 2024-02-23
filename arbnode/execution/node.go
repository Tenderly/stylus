package execution

import (
	"errors"

	"github.com/tenderly/stylus/go-ethereum/arbitrum"
	"github.com/tenderly/stylus/go-ethereum/core"
	"github.com/tenderly/stylus/go-ethereum/eth/filters"
	"github.com/tenderly/stylus/go-ethereum/ethdb"
	"github.com/tenderly/stylus/go-ethereum/node"
	"github.com/tenderly/stylus/util/headerreader"
)

type ExecutionNode struct {
	ChainDB      ethdb.Database
	Backend      *arbitrum.Backend
	FilterSystem *filters.FilterSystem
	ArbInterface *ArbInterface
	ExecEngine   *ExecutionEngine
	Recorder     *BlockRecorder
	Sequencer    *Sequencer // either nil or same as TxPublisher
	TxPublisher  TransactionPublisher
}

func CreateExecutionNode(
	stack *node.Node,
	chainDB ethdb.Database,
	l2BlockChain *core.BlockChain,
	l1Reader *headerreader.HeaderReader,
	syncMonitor arbitrum.SyncProgressBackend,
	fwTarget string,
	fwConfig *ForwarderConfig,
	rpcConfig arbitrum.Config,
	recordingDbConfig *arbitrum.RecordingDatabaseConfig,
	seqConfigFetcher SequencerConfigFetcher,
	precheckConfigFetcher TxPreCheckerConfigFetcher,
) (*ExecutionNode, error) {
	execEngine, err := NewExecutionEngine(l2BlockChain)
	if err != nil {
		return nil, err
	}
	recorder := NewBlockRecorder(recordingDbConfig, execEngine, chainDB)
	var txPublisher TransactionPublisher
	var sequencer *Sequencer
	seqConfig := seqConfigFetcher()
	if seqConfig.Enable {
		if fwTarget != "" {
			return nil, errors.New("sequencer and forwarding target both set")
		}
		sequencer, err = NewSequencer(execEngine, l1Reader, seqConfigFetcher)
		if err != nil {
			return nil, err
		}
		txPublisher = sequencer
	} else {
		if fwConfig.RedisUrl != "" {
			txPublisher = NewRedisTxForwarder(fwTarget, fwConfig)
		} else if fwTarget == "" {
			txPublisher = NewTxDropper()
		} else {
			txPublisher = NewForwarder(fwTarget, fwConfig)
		}
	}

	txPublisher = NewTxPreChecker(txPublisher, l2BlockChain, precheckConfigFetcher)
	arbInterface, err := NewArbInterface(execEngine, txPublisher)
	if err != nil {
		return nil, err
	}
	filterConfig := filters.Config{
		LogCacheSize: rpcConfig.FilterLogCacheSize,
		Timeout:      rpcConfig.FilterTimeout,
	}
	backend, filterSystem, err := arbitrum.NewBackend(stack, &rpcConfig, chainDB, arbInterface, syncMonitor, filterConfig)
	if err != nil {
		return nil, err
	}

	return &ExecutionNode{
		ChainDB:      chainDB,
		Backend:      backend,
		FilterSystem: filterSystem,
		ArbInterface: arbInterface,
		ExecEngine:   execEngine,
		Recorder:     recorder,
		Sequencer:    sequencer,
		TxPublisher:  txPublisher,
	}, nil

}
