// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mocksgen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/tenderly/stylus/go-ethereum"
	"github.com/tenderly/stylus/go-ethereum/accounts/abi"
	"github.com/tenderly/stylus/go-ethereum/accounts/abi/bind"
	"github.com/tenderly/stylus/go-ethereum/common"
	"github.com/tenderly/stylus/go-ethereum/core/types"
	"github.com/tenderly/stylus/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ChallengeLibChallenge is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibChallenge struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}

// ChallengeLibParticipant is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibParticipant struct {
	Addr     common.Address
	TimeLeft *big.Int
}

// ChallengeLibSegmentSelection is an auto generated low-level Go binding around an user-defined struct.
type ChallengeLibSegmentSelection struct {
	OldSegmentsStart  *big.Int
	OldSegmentsLength *big.Int
	OldSegments       [][32]byte
	ChallengePosition *big.Int
}

// GlobalState is an auto generated low-level Go binding around an user-defined struct.
type GlobalState struct {
	Bytes32Vals [2][32]byte
	U64Vals     [2]uint64
}

// ISequencerInboxMaxTimeVariation is an auto generated low-level Go binding around an user-defined struct.
type ISequencerInboxMaxTimeVariation struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}

// ISequencerInboxTimeBounds is an auto generated low-level Go binding around an user-defined struct.
type ISequencerInboxTimeBounds struct {
	MinTimestamp   uint64
	MaxTimestamp   uint64
	MinBlockNumber uint64
	MaxBlockNumber uint64
}

// BenchmarksMetaData contains all meta data concerning the Benchmarks contract.
var BenchmarksMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"fillBlockAdd\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillBlockHash\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillBlockMulMod\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillBlockQuickStep\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fillBlockRecover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061028c806100206000396000f3fe60806040526004361061004a5760003560e01c80630b39d8e51461004f578063142ddcf9146100595780636488e93014610061578063ea01a66f14610069578063fb721c1814610071575b600080fd5b610057610079565b005b61005761019a565b6100576101b2565b6100576101f3565b6100576101fd565b60008051602061023783398151915273361594f5429d23ece0a88e4fbe529e1c49d524d8601b7fc6178c2de1078cd36c3bd302cde755340d7f17fcb3fcc0b9c333ba03b217029f7f5fdbcefe2675e96219cdae57a7894280bf80fd40d44ce146a35e169ea6a78fd35b60408051600081526020810180835287905260ff85169181019190915260608101839052608081018290526001600160a01b0385169060019060a0016020604051602081039080840390855afa158015610140573d6000803e3d6000fd5b505050602060405103516001600160a01b0316146101955760405162461bcd60e51b815260206004820152600e60248201526d15d493d391d7d05490925390555560921b604482015260640160405180910390fd5b6100e2565b60005b6000805160206102378339815191520161019d565b6000805160206102378339815191525b6401000003d0197fc6178c2de1078cd36c3bd302cde755340d7f17fcb3fcc0b9c333ba03b217029f820990506101c2565b60005b50346101f6565b6000805160206102378339815191525b60408051602081018390520160405160208183030381529060405280519060200120905061020d56feeddecf107b5740cef7f5a01e3ea7e287665c4e75a8eb6afae2fda2e3d4367786a2646970667358221220c81f48c14bed617c8b19fa929d3e00b9ad66e1923b1eb2c6d1e7ad5ea93d94f264736f6c63430008090033",
}

// BenchmarksABI is the input ABI used to generate the binding from.
// Deprecated: Use BenchmarksMetaData.ABI instead.
var BenchmarksABI = BenchmarksMetaData.ABI

// BenchmarksBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BenchmarksMetaData.Bin instead.
var BenchmarksBin = BenchmarksMetaData.Bin

// DeployBenchmarks deploys a new Ethereum contract, binding an instance of Benchmarks to it.
func DeployBenchmarks(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Benchmarks, error) {
	parsed, err := BenchmarksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BenchmarksBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Benchmarks{BenchmarksCaller: BenchmarksCaller{contract: contract}, BenchmarksTransactor: BenchmarksTransactor{contract: contract}, BenchmarksFilterer: BenchmarksFilterer{contract: contract}}, nil
}

// Benchmarks is an auto generated Go binding around an Ethereum contract.
type Benchmarks struct {
	BenchmarksCaller     // Read-only binding to the contract
	BenchmarksTransactor // Write-only binding to the contract
	BenchmarksFilterer   // Log filterer for contract events
}

// BenchmarksCaller is an auto generated read-only Go binding around an Ethereum contract.
type BenchmarksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchmarksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BenchmarksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchmarksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BenchmarksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BenchmarksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BenchmarksSession struct {
	Contract     *Benchmarks       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BenchmarksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BenchmarksCallerSession struct {
	Contract *BenchmarksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BenchmarksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BenchmarksTransactorSession struct {
	Contract     *BenchmarksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BenchmarksRaw is an auto generated low-level Go binding around an Ethereum contract.
type BenchmarksRaw struct {
	Contract *Benchmarks // Generic contract binding to access the raw methods on
}

// BenchmarksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BenchmarksCallerRaw struct {
	Contract *BenchmarksCaller // Generic read-only contract binding to access the raw methods on
}

// BenchmarksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BenchmarksTransactorRaw struct {
	Contract *BenchmarksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBenchmarks creates a new instance of Benchmarks, bound to a specific deployed contract.
func NewBenchmarks(address common.Address, backend bind.ContractBackend) (*Benchmarks, error) {
	contract, err := bindBenchmarks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Benchmarks{BenchmarksCaller: BenchmarksCaller{contract: contract}, BenchmarksTransactor: BenchmarksTransactor{contract: contract}, BenchmarksFilterer: BenchmarksFilterer{contract: contract}}, nil
}

// NewBenchmarksCaller creates a new read-only instance of Benchmarks, bound to a specific deployed contract.
func NewBenchmarksCaller(address common.Address, caller bind.ContractCaller) (*BenchmarksCaller, error) {
	contract, err := bindBenchmarks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BenchmarksCaller{contract: contract}, nil
}

// NewBenchmarksTransactor creates a new write-only instance of Benchmarks, bound to a specific deployed contract.
func NewBenchmarksTransactor(address common.Address, transactor bind.ContractTransactor) (*BenchmarksTransactor, error) {
	contract, err := bindBenchmarks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BenchmarksTransactor{contract: contract}, nil
}

// NewBenchmarksFilterer creates a new log filterer instance of Benchmarks, bound to a specific deployed contract.
func NewBenchmarksFilterer(address common.Address, filterer bind.ContractFilterer) (*BenchmarksFilterer, error) {
	contract, err := bindBenchmarks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BenchmarksFilterer{contract: contract}, nil
}

// bindBenchmarks binds a generic wrapper to an already deployed contract.
func bindBenchmarks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BenchmarksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Benchmarks *BenchmarksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Benchmarks.Contract.BenchmarksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Benchmarks *BenchmarksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benchmarks.Contract.BenchmarksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Benchmarks *BenchmarksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Benchmarks.Contract.BenchmarksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Benchmarks *BenchmarksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Benchmarks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Benchmarks *BenchmarksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benchmarks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Benchmarks *BenchmarksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Benchmarks.Contract.contract.Transact(opts, method, params...)
}

// FillBlockAdd is a paid mutator transaction binding the contract method 0x142ddcf9.
//
// Solidity: function fillBlockAdd() payable returns()
func (_Benchmarks *BenchmarksTransactor) FillBlockAdd(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benchmarks.contract.Transact(opts, "fillBlockAdd")
}

// FillBlockAdd is a paid mutator transaction binding the contract method 0x142ddcf9.
//
// Solidity: function fillBlockAdd() payable returns()
func (_Benchmarks *BenchmarksSession) FillBlockAdd() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockAdd(&_Benchmarks.TransactOpts)
}

// FillBlockAdd is a paid mutator transaction binding the contract method 0x142ddcf9.
//
// Solidity: function fillBlockAdd() payable returns()
func (_Benchmarks *BenchmarksTransactorSession) FillBlockAdd() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockAdd(&_Benchmarks.TransactOpts)
}

// FillBlockHash is a paid mutator transaction binding the contract method 0xfb721c18.
//
// Solidity: function fillBlockHash() payable returns()
func (_Benchmarks *BenchmarksTransactor) FillBlockHash(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benchmarks.contract.Transact(opts, "fillBlockHash")
}

// FillBlockHash is a paid mutator transaction binding the contract method 0xfb721c18.
//
// Solidity: function fillBlockHash() payable returns()
func (_Benchmarks *BenchmarksSession) FillBlockHash() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockHash(&_Benchmarks.TransactOpts)
}

// FillBlockHash is a paid mutator transaction binding the contract method 0xfb721c18.
//
// Solidity: function fillBlockHash() payable returns()
func (_Benchmarks *BenchmarksTransactorSession) FillBlockHash() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockHash(&_Benchmarks.TransactOpts)
}

// FillBlockMulMod is a paid mutator transaction binding the contract method 0x6488e930.
//
// Solidity: function fillBlockMulMod() payable returns()
func (_Benchmarks *BenchmarksTransactor) FillBlockMulMod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benchmarks.contract.Transact(opts, "fillBlockMulMod")
}

// FillBlockMulMod is a paid mutator transaction binding the contract method 0x6488e930.
//
// Solidity: function fillBlockMulMod() payable returns()
func (_Benchmarks *BenchmarksSession) FillBlockMulMod() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockMulMod(&_Benchmarks.TransactOpts)
}

// FillBlockMulMod is a paid mutator transaction binding the contract method 0x6488e930.
//
// Solidity: function fillBlockMulMod() payable returns()
func (_Benchmarks *BenchmarksTransactorSession) FillBlockMulMod() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockMulMod(&_Benchmarks.TransactOpts)
}

// FillBlockQuickStep is a paid mutator transaction binding the contract method 0xea01a66f.
//
// Solidity: function fillBlockQuickStep() payable returns()
func (_Benchmarks *BenchmarksTransactor) FillBlockQuickStep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benchmarks.contract.Transact(opts, "fillBlockQuickStep")
}

// FillBlockQuickStep is a paid mutator transaction binding the contract method 0xea01a66f.
//
// Solidity: function fillBlockQuickStep() payable returns()
func (_Benchmarks *BenchmarksSession) FillBlockQuickStep() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockQuickStep(&_Benchmarks.TransactOpts)
}

// FillBlockQuickStep is a paid mutator transaction binding the contract method 0xea01a66f.
//
// Solidity: function fillBlockQuickStep() payable returns()
func (_Benchmarks *BenchmarksTransactorSession) FillBlockQuickStep() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockQuickStep(&_Benchmarks.TransactOpts)
}

// FillBlockRecover is a paid mutator transaction binding the contract method 0x0b39d8e5.
//
// Solidity: function fillBlockRecover() payable returns()
func (_Benchmarks *BenchmarksTransactor) FillBlockRecover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Benchmarks.contract.Transact(opts, "fillBlockRecover")
}

// FillBlockRecover is a paid mutator transaction binding the contract method 0x0b39d8e5.
//
// Solidity: function fillBlockRecover() payable returns()
func (_Benchmarks *BenchmarksSession) FillBlockRecover() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockRecover(&_Benchmarks.TransactOpts)
}

// FillBlockRecover is a paid mutator transaction binding the contract method 0x0b39d8e5.
//
// Solidity: function fillBlockRecover() payable returns()
func (_Benchmarks *BenchmarksTransactorSession) FillBlockRecover() (*types.Transaction, error) {
	return _Benchmarks.Contract.FillBlockRecover(&_Benchmarks.TransactOpts)
}

// BridgeStubMetaData contains all meta data concerning the BridgeStub contract.
var BridgeStubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerMessageNumber\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerReportedSubMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"batchPoster\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c86806100206000396000f3fe6080604052600436106101295760003560e01c80639e5d4c49116100ab578063cee3d7281161006f578063cee3d72814610367578063d5719dc214610382578063e76f5c8d146103a2578063e77145f4146101c2578063eca067ad146103c2578063ee35f327146103d757600080fd5b80639e5d4c49146102a8578063ab5d8943146102d6578063ae60bd13146102f6578063c4d66de814610332578063cb23bcb51461035257600080fd5b80635fca4a16116100f25780635fca4a16146101e45780637a88b107146101fa57806386598a561461021d5780638db5993b1461025d578063945e11471461027057600080fd5b806284120c1461012e57806316bf557914610152578063413b35bd1461017257806347fb24c5146101a25780634f61f850146101c4575b600080fd5b34801561013a57600080fd5b506005545b6040519081526020015b60405180910390f35b34801561015e57600080fd5b5061013f61016d3660046109b6565b6103f7565b34801561017e57600080fd5b5061019261018d3660046109e7565b610418565b6040519015158152602001610149565b3480156101ae57600080fd5b506101c26101bd366004610a0b565b61043b565b005b3480156101d057600080fd5b506101c26101df3660046109e7565b610652565b3480156101f057600080fd5b5061013f60075481565b34801561020657600080fd5b5061013f610215366004610a49565b600092915050565b34801561022957600080fd5b5061023d610238366004610a75565b6106a6565b604080519485526020850193909352918301526060820152608001610149565b61013f61026b366004610aa7565b6107dd565b34801561027c57600080fd5b5061029061028b3660046109b6565b610846565b6040516001600160a01b039091168152602001610149565b3480156102b457600080fd5b506102c86102c3366004610aee565b610870565b604051610149929190610b77565b3480156102e257600080fd5b50600354610290906001600160a01b031681565b34801561030257600080fd5b506101926103113660046109e7565b6001600160a01b031660009081526020819052604090206001015460ff1690565b34801561033e57600080fd5b506101c261034d3660046109e7565b61088c565b34801561035e57600080fd5b50610290610418565b34801561037357600080fd5b506101c261034d366004610a0b565b34801561038e57600080fd5b5061013f61039d3660046109b6565b6108a4565b3480156103ae57600080fd5b506102906103bd3660046109b6565b6108b4565b3480156103ce57600080fd5b5060045461013f565b3480156103e357600080fd5b50600654610290906001600160a01b031681565b6005818154811061040757600080fd5b600091825260209091200154905081565b600060405162461bcd60e51b815260040161043290610bd6565b60405180910390fd5b6001600160a01b03821660008181526020818152604091829020600181015492518515158152909360ff90931692917f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521910160405180910390a282151581151514156104a75750505050565b82156105335760408051808201825260018054825260208083018281526001600160a01b0389166000818152928390529482209351845551928201805460ff1916931515939093179092558054808201825591527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b031916909117905561064c565b60018054610542908290610bff565b8154811061055257610552610c24565b6000918252602090912001548254600180546001600160a01b0390931692909190811061058157610581610c24565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000015460008060018560000154815481106105ce576105ce610c24565b60009182526020808320909101546001600160a01b03168352820192909252604001902055600180548061060457610604610c3a565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038616825281905260408120908155600101805460ff191690555b50505050565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a9060200160405180910390a150565b60008060008085600754141580156106bd57508515155b80156106ca575060075415155b156106f65760075460405163e2051feb60e01b8152600481019190915260248101879052604401610432565b600785905560055493508315610734576005805461071690600190610bff565b8154811061072657610726610c24565b906000526020600020015492505b8615610765576004610747600189610bff565b8154811061075757610757610c24565b906000526020600020015491505b60408051602081018590529081018990526060810183905260800160408051601f198184030181529190528051602090910120600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0018190559398929750909550919350915050565b3360009081526020819052604081206001015460ff166108305760405162461bcd60e51b815260206004820152600e60248201526d09c9ea8be8ca49e9abe929c849eb60931b6044820152606401610432565b61083e8484434248876108c4565b949350505050565b6002818154811061085657600080fd5b6000918252602090912001546001600160a01b0316905081565b6000606060405162461bcd60e51b815260040161043290610bd6565b60405162461bcd60e51b815260040161043290610bd6565b6004818154811061040757600080fd5b6001818154811061085657600080fd5b60045460408051600060208083018290526021830182905260358301829052603d8301829052604583018290526065830182905260858084018790528451808503909101815260a5909301909352815191909201209091906000821561094f576004610931600185610bff565b8154811061094157610941610c24565b906000526020600020015490505b6040805160208082019390935280820193909352805180840382018152606090930190528151910120600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0155979650505050505050565b6000602082840312156109c857600080fd5b5035919050565b6001600160a01b03811681146109e457600080fd5b50565b6000602082840312156109f957600080fd5b8135610a04816109cf565b9392505050565b60008060408385031215610a1e57600080fd5b8235610a29816109cf565b915060208301358015158114610a3e57600080fd5b809150509250929050565b60008060408385031215610a5c57600080fd5b8235610a67816109cf565b946020939093013593505050565b60008060008060808587031215610a8b57600080fd5b5050823594602084013594506040840135936060013592509050565b600080600060608486031215610abc57600080fd5b833560ff81168114610acd57600080fd5b92506020840135610add816109cf565b929592945050506040919091013590565b60008060008060608587031215610b0457600080fd5b8435610b0f816109cf565b935060208501359250604085013567ffffffffffffffff80821115610b3357600080fd5b818701915087601f830112610b4757600080fd5b813581811115610b5657600080fd5b886020828501011115610b6857600080fd5b95989497505060200194505050565b821515815260006020604081840152835180604085015260005b81811015610bad57858101830151858201606001528201610b91565b81811115610bbf576000606083870101525b50601f01601f191692909201606001949350505050565b6020808252600f908201526e1393d517d253541311535153951151608a1b604082015260600190565b600082821015610c1f57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fdfea2646970667358221220acb36ff1631207b3749629a85cf946ae3805ba0ada7fdc02c7b686bc400f48ca64736f6c63430008090033",
}

// BridgeStubABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeStubMetaData.ABI instead.
var BridgeStubABI = BridgeStubMetaData.ABI

// BridgeStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeStubMetaData.Bin instead.
var BridgeStubBin = BridgeStubMetaData.Bin

// DeployBridgeStub deploys a new Ethereum contract, binding an instance of BridgeStub to it.
func DeployBridgeStub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeStub, error) {
	parsed, err := BridgeStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeStubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeStub{BridgeStubCaller: BridgeStubCaller{contract: contract}, BridgeStubTransactor: BridgeStubTransactor{contract: contract}, BridgeStubFilterer: BridgeStubFilterer{contract: contract}}, nil
}

// BridgeStub is an auto generated Go binding around an Ethereum contract.
type BridgeStub struct {
	BridgeStubCaller     // Read-only binding to the contract
	BridgeStubTransactor // Write-only binding to the contract
	BridgeStubFilterer   // Log filterer for contract events
}

// BridgeStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeStubSession struct {
	Contract     *BridgeStub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeStubCallerSession struct {
	Contract *BridgeStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BridgeStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeStubTransactorSession struct {
	Contract     *BridgeStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BridgeStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeStubRaw struct {
	Contract *BridgeStub // Generic contract binding to access the raw methods on
}

// BridgeStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeStubCallerRaw struct {
	Contract *BridgeStubCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeStubTransactorRaw struct {
	Contract *BridgeStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeStub creates a new instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStub(address common.Address, backend bind.ContractBackend) (*BridgeStub, error) {
	contract, err := bindBridgeStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeStub{BridgeStubCaller: BridgeStubCaller{contract: contract}, BridgeStubTransactor: BridgeStubTransactor{contract: contract}, BridgeStubFilterer: BridgeStubFilterer{contract: contract}}, nil
}

// NewBridgeStubCaller creates a new read-only instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubCaller(address common.Address, caller bind.ContractCaller) (*BridgeStubCaller, error) {
	contract, err := bindBridgeStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeStubCaller{contract: contract}, nil
}

// NewBridgeStubTransactor creates a new write-only instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeStubTransactor, error) {
	contract, err := bindBridgeStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeStubTransactor{contract: contract}, nil
}

// NewBridgeStubFilterer creates a new log filterer instance of BridgeStub, bound to a specific deployed contract.
func NewBridgeStubFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeStubFilterer, error) {
	contract, err := bindBridgeStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeStubFilterer{contract: contract}, nil
}

// bindBridgeStub binds a generic wrapper to an already deployed contract.
func bindBridgeStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeStubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeStub *BridgeStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeStub.Contract.BridgeStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeStub *BridgeStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.Contract.BridgeStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeStub *BridgeStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeStub.Contract.BridgeStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeStub *BridgeStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeStub *BridgeStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeStub *BridgeStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeStub.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubSession) ActiveOutbox() (common.Address, error) {
	return _BridgeStub.Contract.ActiveOutbox(&_BridgeStub.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeStub *BridgeStubCallerSession) ActiveOutbox() (common.Address, error) {
	return _BridgeStub.Contract.ActiveOutbox(&_BridgeStub.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxes(&_BridgeStub.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeStub *BridgeStubCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedDelayedInboxes(&_BridgeStub.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedOutboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeStub *BridgeStubCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeStub.Contract.AllowedOutboxList(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubCaller) AllowedOutboxes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "allowedOutboxes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubSession) AllowedOutboxes(arg0 common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedOutboxes(&_BridgeStub.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address ) pure returns(bool)
func (_BridgeStub *BridgeStubCallerSession) AllowedOutboxes(arg0 common.Address) (bool, error) {
	return _BridgeStub.Contract.AllowedOutboxes(&_BridgeStub.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.DelayedInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.DelayedInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.DelayedMessageCount(&_BridgeStub.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.DelayedMessageCount(&_BridgeStub.CallOpts)
}

// ExecuteCall is a free data retrieval call binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address , uint256 , bytes ) pure returns(bool, bytes)
func (_BridgeStub *BridgeStubCaller) ExecuteCall(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (bool, []byte, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "executeCall", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// ExecuteCall is a free data retrieval call binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address , uint256 , bytes ) pure returns(bool, bytes)
func (_BridgeStub *BridgeStubSession) ExecuteCall(arg0 common.Address, arg1 *big.Int, arg2 []byte) (bool, []byte, error) {
	return _BridgeStub.Contract.ExecuteCall(&_BridgeStub.CallOpts, arg0, arg1, arg2)
}

// ExecuteCall is a free data retrieval call binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address , uint256 , bytes ) pure returns(bool, bytes)
func (_BridgeStub *BridgeStubCallerSession) ExecuteCall(arg0 common.Address, arg1 *big.Int, arg2 []byte) (bool, []byte, error) {
	return _BridgeStub.Contract.ExecuteCall(&_BridgeStub.CallOpts, arg0, arg1, arg2)
}

// Initialize is a free data retrieval call binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) pure returns()
func (_BridgeStub *BridgeStubCaller) Initialize(opts *bind.CallOpts, arg0 common.Address) error {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "initialize", arg0)

	if err != nil {
		return err
	}

	return err

}

// Initialize is a free data retrieval call binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) pure returns()
func (_BridgeStub *BridgeStubSession) Initialize(arg0 common.Address) error {
	return _BridgeStub.Contract.Initialize(&_BridgeStub.CallOpts, arg0)
}

// Initialize is a free data retrieval call binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address ) pure returns()
func (_BridgeStub *BridgeStubCallerSession) Initialize(arg0 common.Address) error {
	return _BridgeStub.Contract.Initialize(&_BridgeStub.CallOpts, arg0)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubSession) Rollup() (common.Address, error) {
	return _BridgeStub.Contract.Rollup(&_BridgeStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() pure returns(address)
func (_BridgeStub *BridgeStubCallerSession) Rollup() (common.Address, error) {
	return _BridgeStub.Contract.Rollup(&_BridgeStub.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubSession) SequencerInbox() (common.Address, error) {
	return _BridgeStub.Contract.SequencerInbox(&_BridgeStub.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeStub *BridgeStubCallerSession) SequencerInbox() (common.Address, error) {
	return _BridgeStub.Contract.SequencerInbox(&_BridgeStub.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.SequencerInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeStub *BridgeStubCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeStub.Contract.SequencerInboxAccs(&_BridgeStub.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerMessageCount(&_BridgeStub.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerMessageCount(&_BridgeStub.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCaller) SequencerReportedSubMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "sequencerReportedSubMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerReportedSubMessageCount(&_BridgeStub.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeStub *BridgeStubCallerSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeStub.Contract.SequencerReportedSubMessageCount(&_BridgeStub.CallOpts)
}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubCaller) SetOutbox(opts *bind.CallOpts, arg0 common.Address, arg1 bool) error {
	var out []interface{}
	err := _BridgeStub.contract.Call(opts, &out, "setOutbox", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubSession) SetOutbox(arg0 common.Address, arg1 bool) error {
	return _BridgeStub.Contract.SetOutbox(&_BridgeStub.CallOpts, arg0, arg1)
}

// SetOutbox is a free data retrieval call binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address , bool ) pure returns()
func (_BridgeStub *BridgeStubCallerSession) SetOutbox(arg0 common.Address, arg1 bool) error {
	return _BridgeStub.Contract.SetOutbox(&_BridgeStub.CallOpts, arg0, arg1)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeStub.Contract.AcceptFundsFromOldBridge(&_BridgeStub.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeStub *BridgeStubTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeStub.Contract.AcceptFundsFromOldBridge(&_BridgeStub.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueDelayedMessage(&_BridgeStub.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeStub *BridgeStubTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueDelayedMessage(&_BridgeStub.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueSequencerMessage(&_BridgeStub.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeStub *BridgeStubTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeStub.Contract.EnqueueSequencerMessage(&_BridgeStub.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetDelayedInbox(&_BridgeStub.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeStub *BridgeStubTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetDelayedInbox(&_BridgeStub.TransactOpts, inbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetSequencerInbox(&_BridgeStub.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeStub *BridgeStubTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeStub.Contract.SetSequencerInbox(&_BridgeStub.TransactOpts, _sequencerInbox)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.contract.Transact(opts, "submitBatchSpendingReport", batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.SubmitBatchSpendingReport(&_BridgeStub.TransactOpts, batchPoster, dataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address batchPoster, bytes32 dataHash) returns(uint256)
func (_BridgeStub *BridgeStubTransactorSession) SubmitBatchSpendingReport(batchPoster common.Address, dataHash [32]byte) (*types.Transaction, error) {
	return _BridgeStub.Contract.SubmitBatchSpendingReport(&_BridgeStub.TransactOpts, batchPoster, dataHash)
}

// BridgeStubBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the BridgeStub contract.
type BridgeStubBridgeCallTriggeredIterator struct {
	Event *BridgeStubBridgeCallTriggered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeStubBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubBridgeCallTriggered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeStubBridgeCallTriggered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeStubBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubBridgeCallTriggered represents a BridgeCallTriggered event raised by the BridgeStub contract.
type BridgeStubBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeStubBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubBridgeCallTriggeredIterator{contract: _BridgeStub.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeStubBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubBridgeCallTriggered)
				if err := _BridgeStub.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeCallTriggered is a log parse operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeStub *BridgeStubFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeStubBridgeCallTriggered, error) {
	event := new(BridgeStubBridgeCallTriggered)
	if err := _BridgeStub.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the BridgeStub contract.
type BridgeStubInboxToggleIterator struct {
	Event *BridgeStubInboxToggle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeStubInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubInboxToggle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeStubInboxToggle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeStubInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubInboxToggle represents a InboxToggle event raised by the BridgeStub contract.
type BridgeStubInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeStubInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubInboxToggleIterator{contract: _BridgeStub.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeStubInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubInboxToggle)
				if err := _BridgeStub.contract.UnpackLog(event, "InboxToggle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboxToggle is a log parse operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) ParseInboxToggle(log types.Log) (*BridgeStubInboxToggle, error) {
	event := new(BridgeStubInboxToggle)
	if err := _BridgeStub.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the BridgeStub contract.
type BridgeStubMessageDeliveredIterator struct {
	Event *BridgeStubMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeStubMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeStubMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeStubMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubMessageDelivered represents a MessageDelivered event raised by the BridgeStub contract.
type BridgeStubMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	BaseFeeL1       *big.Int
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeStubMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubMessageDeliveredIterator{contract: _BridgeStub.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeStubMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubMessageDelivered)
				if err := _BridgeStub.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeStub *BridgeStubFilterer) ParseMessageDelivered(log types.Log) (*BridgeStubMessageDelivered, error) {
	event := new(BridgeStubMessageDelivered)
	if err := _BridgeStub.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the BridgeStub contract.
type BridgeStubOutboxToggleIterator struct {
	Event *BridgeStubOutboxToggle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeStubOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubOutboxToggle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeStubOutboxToggle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeStubOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubOutboxToggle represents a OutboxToggle event raised by the BridgeStub contract.
type BridgeStubOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeStubOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeStubOutboxToggleIterator{contract: _BridgeStub.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeStubOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubOutboxToggle)
				if err := _BridgeStub.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutboxToggle is a log parse operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeStub *BridgeStubFilterer) ParseOutboxToggle(log types.Log) (*BridgeStubOutboxToggle, error) {
	event := new(BridgeStubOutboxToggle)
	if err := _BridgeStub.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeStubSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the BridgeStub contract.
type BridgeStubSequencerInboxUpdatedIterator struct {
	Event *BridgeStubSequencerInboxUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeStubSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeStubSequencerInboxUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeStubSequencerInboxUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeStubSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeStubSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeStubSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the BridgeStub contract.
type BridgeStubSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeStubSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _BridgeStub.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeStubSequencerInboxUpdatedIterator{contract: _BridgeStub.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeStubSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeStub.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeStubSequencerInboxUpdated)
				if err := _BridgeStub.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerInboxUpdated is a log parse operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeStub *BridgeStubFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeStubSequencerInboxUpdated, error) {
	event := new(BridgeStubSequencerInboxUpdated)
	if err := _BridgeStub.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedMetaData contains all meta data concerning the BridgeUnproxied contract.
var BridgeUnproxiedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerMessageNumber\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"InvalidOutboxSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotDelayedInbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotOutbox\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotRollupOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotSequencerInbox\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"BridgeCallTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"InboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeInboxAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"MessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"OutboxToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSequencerInbox\",\"type\":\"address\"}],\"name\":\"SequencerInboxUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptFundsFromOldBridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeOutbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedDelayedInboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"}],\"name\":\"allowedDelayedInboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allowedOutboxList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"}],\"name\":\"allowedOutboxes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delayedInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"enqueueDelayedMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"enqueueSequencerMessage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"seqMessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"acc\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"rollup_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sequencerInboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerReportedSubMessageCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setDelayedInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"outbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"setSequencerInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMsgCount\",\"type\":\"uint256\"}],\"name\":\"setSequencerReportedSubMessageCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchSpendingReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50600580546001600160a01b03199081166001600160a01b0317909155600880549091163317905560805161165c6100576000396000610cc1015261165c6000f3fe6080604052600436106101345760003560e01c8063ab5d8943116100ab578063d5719dc21161006f578063d5719dc2146103a1578063e76f5c8d146103c1578063e77145f4146101e9578063eca067ad146103e1578063ee35f327146103f6578063f81ff3b31461041657600080fd5b8063ab5d8943146102ef578063ae60bd1314610304578063c4d66de814610341578063cb23bcb514610361578063cee3d7281461038157600080fd5b80635fca4a16116100fd5780635fca4a161461020b5780637a88b1071461022157806386598a56146102415780638db5993b14610281578063945e1147146102945780639e5d4c49146102c157600080fd5b806284120c1461013957806316bf55791461015d578063413b35bd1461017d57806347fb24c5146101c95780634f61f850146101eb575b600080fd5b34801561014557600080fd5b506007545b6040519081526020015b60405180910390f35b34801561016957600080fd5b5061014a610178366004611322565b610436565b34801561018957600080fd5b506101b9610198366004611353565b6001600160a01b031660009081526002602052604090206001015460ff1690565b6040519015158152602001610154565b3480156101d557600080fd5b506101e96101e4366004611370565b610457565b005b3480156101f757600080fd5b506101e9610206366004611353565b61074d565b34801561021757600080fd5b5061014a600a5481565b34801561022d57600080fd5b5061014a61023c3660046113ae565b610872565b34801561024d57600080fd5b5061026161025c3660046113da565b6108b8565b604080519485526020850193909352918301526060820152608001610154565b61014a61028f36600461140c565b610a1f565b3480156102a057600080fd5b506102b46102af366004611322565b610a6a565b6040516101549190611453565b3480156102cd57600080fd5b506102e16102dc366004611467565b610a94565b6040516101549291906114f0565b3480156102fb57600080fd5b506102b4610bec565b34801561031057600080fd5b506101b961031f366004611353565b6001600160a01b03166000908152600160208190526040909120015460ff1690565b34801561034d57600080fd5b506101e961035c366004611353565b610c12565b34801561036d57600080fd5b506008546102b4906001600160a01b031681565b34801561038d57600080fd5b506101e961039c366004611370565b610d86565b3480156103ad57600080fd5b5061014a6103bc366004611322565b61109f565b3480156103cd57600080fd5b506102b46103dc366004611322565b6110af565b3480156103ed57600080fd5b5060065461014a565b34801561040257600080fd5b506009546102b4906001600160a01b031681565b34801561042257600080fd5b506101e9610431366004611322565b6110bf565b6007818154811061044657600080fd5b600091825260209091200154905081565b6008546001600160a01b0316331461052f5760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b1580156104ae57600080fd5b505afa1580156104c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e6919061154f565b9050336001600160a01b0382161461052d57600854604051630739600760e01b81526105249133916001600160a01b0390911690849060040161156c565b60405180910390fd5b505b6001600160a01b0382166000818152600160208181526040928390209182015492518515158152919360ff90931692917f6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521910160405180910390a2821515811515141561059c5750505050565b821561062a57604080518082018252600380548252600160208084018281526001600160a01b038a166000818152928490529582209451855551938201805460ff1916941515949094179093558154908101825591527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319169091179055610746565b6003805461063a9060019061158f565b8154811061064a5761064a6115b4565b6000918252602090912001548254600380546001600160a01b03909316929091908110610679576106796115b4565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081600001546001600060038560000154815481106106c7576106c76115b4565b60009182526020808320909101546001600160a01b0316835282019290925260400190205560038054806106fd576106fd6115ca565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526001908190526040822091825501805460ff191690555b50505b5050565b6008546001600160a01b0316331461081c5760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b1580156107a457600080fd5b505afa1580156107b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107dc919061154f565b9050336001600160a01b0382161461081a57600854604051630739600760e01b81526105249133916001600160a01b0390911690849060040161156c565b505b600980546001600160a01b0319166001600160a01b0383161790556040517f8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a90610867908390611453565b60405180910390a150565b6009546000906001600160a01b031633146108a2573360405163223e13c160e21b81526004016105249190611453565b6108b1600d8443424887611193565b9392505050565b6009546000908190819081906001600160a01b031633146108ee573360405163223e13c160e21b81526004016105249190611453565b85600a54141580156108ff57508515155b801561090c5750600a5415155b1561093857600a5460405163e2051feb60e01b8152600481019190915260248101879052604401610524565b600a8590556007549350831561097657600780546109589060019061158f565b81548110610968576109686115b4565b906000526020600020015492505b86156109a757600661098960018961158f565b81548110610999576109996115b4565b906000526020600020015491505b60408051602081018590529081018990526060810183905260800160408051601f198184030181529190528051602090910120600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c688018190559398929750909550919350915050565b3360009081526001602081905260408220015460ff16610a54573360405163b6c60ea360e01b81526004016105249190611453565b610a62848443424887611193565b949350505050565b60048181548110610a7a57600080fd5b6000918252602090912001546001600160a01b0316905081565b3360009081526002602052604081206001015460609060ff16610acc57336040516332ea82ab60e01b81526004016105249190611453565b8215801590610ae357506001600160a01b0386163b155b15610b03578560405163b5cf5b8f60e01b81526004016105249190611453565b600580546001600160a01b0319811633179091556040516001600160a01b03918216918816908790610b3890889088906115e0565b60006040518083038185875af1925050503d8060008114610b75576040519150601f19603f3d011682016040523d82523d6000602084013e610b7a565b606091505b50600580546001600160a01b0319166001600160a01b038581169190911790915560405192955090935088169033907f2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d46690610bda908a908a908a906115f0565b60405180910390a35094509492505050565b6005546000906001600160a01b0390811690811415610c0d57600091505090565b919050565b600054610100900460ff16610c2d5760005460ff1615610c31565b303b155b610c945760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610524565b600054610100900460ff16158015610cb6576000805461ffff19166101011790555b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161415610d445760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608401610524565b600580546001600160a01b036001600160a01b03199182168117909255600880549091169184169190911790558015610749576000805461ff00191690555050565b6008546001600160a01b03163314610e555760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b158015610ddd57600080fd5b505afa158015610df1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e15919061154f565b9050336001600160a01b03821614610e5357600854604051630739600760e01b81526105249133916001600160a01b0390911690849060040161156c565b505b6001600160a01b038281161415610e81578160405163077abed160e41b81526004016105249190611453565b6001600160a01b038216600081815260026020908152604091829020600181015492518515158152909360ff90931692917f49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa910160405180910390a28215158115151415610eef5750505050565b8215610f7e57604080518082018252600480548252600160208084018281526001600160a01b038a16600081815260029093529582209451855551938201805460ff1916941515949094179093558154908101825591527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b0180546001600160a01b0319169091179055610746565b60048054610f8e9060019061158f565b81548110610f9e57610f9e6115b4565b6000918252602090912001548254600480546001600160a01b03909316929091908110610fcd57610fcd6115b4565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550816000015460026000600485600001548154811061101b5761101b6115b4565b60009182526020808320909101546001600160a01b031683528201929092526040019020556004805480611051576110516115ca565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03861682526002905260408120908155600101805460ff1916905550505050565b6006818154811061044657600080fd5b60038181548110610a7a57600080fd5b6008546001600160a01b0316331461118e5760085460408051638da5cb5b60e01b815290516000926001600160a01b031691638da5cb5b916004808301926020929190829003018186803b15801561111657600080fd5b505afa15801561112a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061114e919061154f565b9050336001600160a01b0382161461118c57600854604051630739600760e01b81526105249133916001600160a01b0390911690849060040161156c565b505b600a55565b600654604080516001600160f81b031960f88a901b166020808301919091526bffffffffffffffffffffffff1960608a901b1660218301526001600160c01b031960c089811b8216603585015288901b16603d830152604582018490526065820186905260858083018690528351808403909101815260a59092019092528051910120600091906000821561124d57600661122f60018561158f565b8154811061123f5761123f6115b4565b906000526020600020015490505b6040805160208082018490528183018590528251808303840181526060830180855281519190920120600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f015533905260ff8c1660808201526001600160a01b038b1660a082015260c0810187905260e0810188905267ffffffffffffffff89166101008201529051829185917f5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1918190036101200190a3509098975050505050505050565b60006020828403121561133457600080fd5b5035919050565b6001600160a01b038116811461135057600080fd5b50565b60006020828403121561136557600080fd5b81356108b18161133b565b6000806040838503121561138357600080fd5b823561138e8161133b565b9150602083013580151581146113a357600080fd5b809150509250929050565b600080604083850312156113c157600080fd5b82356113cc8161133b565b946020939093013593505050565b600080600080608085870312156113f057600080fd5b5050823594602084013594506040840135936060013592509050565b60008060006060848603121561142157600080fd5b833560ff8116811461143257600080fd5b925060208401356114428161133b565b929592945050506040919091013590565b6001600160a01b0391909116815260200190565b6000806000806060858703121561147d57600080fd5b84356114888161133b565b935060208501359250604085013567ffffffffffffffff808211156114ac57600080fd5b818701915087601f8301126114c057600080fd5b8135818111156114cf57600080fd5b8860208285010111156114e157600080fd5b95989497505060200194505050565b821515815260006020604081840152835180604085015260005b818110156115265785810183015185820160600152820161150a565b81811115611538576000606083870101525b50601f01601f191692909201606001949350505050565b60006020828403121561156157600080fd5b81516108b18161133b565b6001600160a01b0393841681529183166020830152909116604082015260600190565b6000828210156115af57634e487b7160e01b600052601160045260246000fd5b500390565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fd5b8183823760009101908152919050565b83815260406020820152816040820152818360608301376000818301606090810191909152601f909201601f191601019291505056fea2646970667358221220d495ef370ff856296bcae98df4579bcda278197ceb7f398b03f50f0ab205e45564736f6c63430008090033",
}

// BridgeUnproxiedABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeUnproxiedMetaData.ABI instead.
var BridgeUnproxiedABI = BridgeUnproxiedMetaData.ABI

// BridgeUnproxiedBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeUnproxiedMetaData.Bin instead.
var BridgeUnproxiedBin = BridgeUnproxiedMetaData.Bin

// DeployBridgeUnproxied deploys a new Ethereum contract, binding an instance of BridgeUnproxied to it.
func DeployBridgeUnproxied(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeUnproxied, error) {
	parsed, err := BridgeUnproxiedMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeUnproxiedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeUnproxied{BridgeUnproxiedCaller: BridgeUnproxiedCaller{contract: contract}, BridgeUnproxiedTransactor: BridgeUnproxiedTransactor{contract: contract}, BridgeUnproxiedFilterer: BridgeUnproxiedFilterer{contract: contract}}, nil
}

// BridgeUnproxied is an auto generated Go binding around an Ethereum contract.
type BridgeUnproxied struct {
	BridgeUnproxiedCaller     // Read-only binding to the contract
	BridgeUnproxiedTransactor // Write-only binding to the contract
	BridgeUnproxiedFilterer   // Log filterer for contract events
}

// BridgeUnproxiedCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeUnproxiedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUnproxiedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeUnproxiedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUnproxiedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeUnproxiedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeUnproxiedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeUnproxiedSession struct {
	Contract     *BridgeUnproxied  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeUnproxiedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeUnproxiedCallerSession struct {
	Contract *BridgeUnproxiedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BridgeUnproxiedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeUnproxiedTransactorSession struct {
	Contract     *BridgeUnproxiedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BridgeUnproxiedRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeUnproxiedRaw struct {
	Contract *BridgeUnproxied // Generic contract binding to access the raw methods on
}

// BridgeUnproxiedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeUnproxiedCallerRaw struct {
	Contract *BridgeUnproxiedCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeUnproxiedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeUnproxiedTransactorRaw struct {
	Contract *BridgeUnproxiedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeUnproxied creates a new instance of BridgeUnproxied, bound to a specific deployed contract.
func NewBridgeUnproxied(address common.Address, backend bind.ContractBackend) (*BridgeUnproxied, error) {
	contract, err := bindBridgeUnproxied(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxied{BridgeUnproxiedCaller: BridgeUnproxiedCaller{contract: contract}, BridgeUnproxiedTransactor: BridgeUnproxiedTransactor{contract: contract}, BridgeUnproxiedFilterer: BridgeUnproxiedFilterer{contract: contract}}, nil
}

// NewBridgeUnproxiedCaller creates a new read-only instance of BridgeUnproxied, bound to a specific deployed contract.
func NewBridgeUnproxiedCaller(address common.Address, caller bind.ContractCaller) (*BridgeUnproxiedCaller, error) {
	contract, err := bindBridgeUnproxied(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedCaller{contract: contract}, nil
}

// NewBridgeUnproxiedTransactor creates a new write-only instance of BridgeUnproxied, bound to a specific deployed contract.
func NewBridgeUnproxiedTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeUnproxiedTransactor, error) {
	contract, err := bindBridgeUnproxied(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedTransactor{contract: contract}, nil
}

// NewBridgeUnproxiedFilterer creates a new log filterer instance of BridgeUnproxied, bound to a specific deployed contract.
func NewBridgeUnproxiedFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeUnproxiedFilterer, error) {
	contract, err := bindBridgeUnproxied(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedFilterer{contract: contract}, nil
}

// bindBridgeUnproxied binds a generic wrapper to an already deployed contract.
func bindBridgeUnproxied(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeUnproxiedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeUnproxied *BridgeUnproxiedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeUnproxied.Contract.BridgeUnproxiedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeUnproxied *BridgeUnproxiedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.BridgeUnproxiedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeUnproxied *BridgeUnproxiedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.BridgeUnproxiedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeUnproxied *BridgeUnproxiedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeUnproxied.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeUnproxied *BridgeUnproxiedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeUnproxied *BridgeUnproxiedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.contract.Transact(opts, method, params...)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) ActiveOutbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "activeOutbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) ActiveOutbox() (common.Address, error) {
	return _BridgeUnproxied.Contract.ActiveOutbox(&_BridgeUnproxied.CallOpts)
}

// ActiveOutbox is a free data retrieval call binding the contract method 0xab5d8943.
//
// Solidity: function activeOutbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) ActiveOutbox() (common.Address, error) {
	return _BridgeUnproxied.Contract.ActiveOutbox(&_BridgeUnproxied.CallOpts)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) AllowedDelayedInboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "allowedDelayedInboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeUnproxied.Contract.AllowedDelayedInboxList(&_BridgeUnproxied.CallOpts, arg0)
}

// AllowedDelayedInboxList is a free data retrieval call binding the contract method 0xe76f5c8d.
//
// Solidity: function allowedDelayedInboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) AllowedDelayedInboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeUnproxied.Contract.AllowedDelayedInboxList(&_BridgeUnproxied.CallOpts, arg0)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedCaller) AllowedDelayedInboxes(opts *bind.CallOpts, inbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "allowedDelayedInboxes", inbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeUnproxied.Contract.AllowedDelayedInboxes(&_BridgeUnproxied.CallOpts, inbox)
}

// AllowedDelayedInboxes is a free data retrieval call binding the contract method 0xae60bd13.
//
// Solidity: function allowedDelayedInboxes(address inbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) AllowedDelayedInboxes(inbox common.Address) (bool, error) {
	return _BridgeUnproxied.Contract.AllowedDelayedInboxes(&_BridgeUnproxied.CallOpts, inbox)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) AllowedOutboxList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "allowedOutboxList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeUnproxied.Contract.AllowedOutboxList(&_BridgeUnproxied.CallOpts, arg0)
}

// AllowedOutboxList is a free data retrieval call binding the contract method 0x945e1147.
//
// Solidity: function allowedOutboxList(uint256 ) view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) AllowedOutboxList(arg0 *big.Int) (common.Address, error) {
	return _BridgeUnproxied.Contract.AllowedOutboxList(&_BridgeUnproxied.CallOpts, arg0)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedCaller) AllowedOutboxes(opts *bind.CallOpts, outbox common.Address) (bool, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "allowedOutboxes", outbox)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeUnproxied.Contract.AllowedOutboxes(&_BridgeUnproxied.CallOpts, outbox)
}

// AllowedOutboxes is a free data retrieval call binding the contract method 0x413b35bd.
//
// Solidity: function allowedOutboxes(address outbox) view returns(bool)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) AllowedOutboxes(outbox common.Address) (bool, error) {
	return _BridgeUnproxied.Contract.AllowedOutboxes(&_BridgeUnproxied.CallOpts, outbox)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedCaller) DelayedInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "delayedInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeUnproxied.Contract.DelayedInboxAccs(&_BridgeUnproxied.CallOpts, arg0)
}

// DelayedInboxAccs is a free data retrieval call binding the contract method 0xd5719dc2.
//
// Solidity: function delayedInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) DelayedInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeUnproxied.Contract.DelayedInboxAccs(&_BridgeUnproxied.CallOpts, arg0)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCaller) DelayedMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "delayedMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.DelayedMessageCount(&_BridgeUnproxied.CallOpts)
}

// DelayedMessageCount is a free data retrieval call binding the contract method 0xeca067ad.
//
// Solidity: function delayedMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) DelayedMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.DelayedMessageCount(&_BridgeUnproxied.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) Rollup() (common.Address, error) {
	return _BridgeUnproxied.Contract.Rollup(&_BridgeUnproxied.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) Rollup() (common.Address, error) {
	return _BridgeUnproxied.Contract.Rollup(&_BridgeUnproxied.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedSession) SequencerInbox() (common.Address, error) {
	return _BridgeUnproxied.Contract.SequencerInbox(&_BridgeUnproxied.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) SequencerInbox() (common.Address, error) {
	return _BridgeUnproxied.Contract.SequencerInbox(&_BridgeUnproxied.CallOpts)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedCaller) SequencerInboxAccs(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "sequencerInboxAccs", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeUnproxied.Contract.SequencerInboxAccs(&_BridgeUnproxied.CallOpts, arg0)
}

// SequencerInboxAccs is a free data retrieval call binding the contract method 0x16bf5579.
//
// Solidity: function sequencerInboxAccs(uint256 ) view returns(bytes32)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) SequencerInboxAccs(arg0 *big.Int) ([32]byte, error) {
	return _BridgeUnproxied.Contract.SequencerInboxAccs(&_BridgeUnproxied.CallOpts, arg0)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCaller) SequencerMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "sequencerMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.SequencerMessageCount(&_BridgeUnproxied.CallOpts)
}

// SequencerMessageCount is a free data retrieval call binding the contract method 0x0084120c.
//
// Solidity: function sequencerMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) SequencerMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.SequencerMessageCount(&_BridgeUnproxied.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCaller) SequencerReportedSubMessageCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeUnproxied.contract.Call(opts, &out, "sequencerReportedSubMessageCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.SequencerReportedSubMessageCount(&_BridgeUnproxied.CallOpts)
}

// SequencerReportedSubMessageCount is a free data retrieval call binding the contract method 0x5fca4a16.
//
// Solidity: function sequencerReportedSubMessageCount() view returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedCallerSession) SequencerReportedSubMessageCount() (*big.Int, error) {
	return _BridgeUnproxied.Contract.SequencerReportedSubMessageCount(&_BridgeUnproxied.CallOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) AcceptFundsFromOldBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "acceptFundsFromOldBridge")
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.AcceptFundsFromOldBridge(&_BridgeUnproxied.TransactOpts)
}

// AcceptFundsFromOldBridge is a paid mutator transaction binding the contract method 0xe77145f4.
//
// Solidity: function acceptFundsFromOldBridge() payable returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) AcceptFundsFromOldBridge() (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.AcceptFundsFromOldBridge(&_BridgeUnproxied.TransactOpts)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedTransactor) EnqueueDelayedMessage(opts *bind.TransactOpts, kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "enqueueDelayedMessage", kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.EnqueueDelayedMessage(&_BridgeUnproxied.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueDelayedMessage is a paid mutator transaction binding the contract method 0x8db5993b.
//
// Solidity: function enqueueDelayedMessage(uint8 kind, address sender, bytes32 messageDataHash) payable returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) EnqueueDelayedMessage(kind uint8, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.EnqueueDelayedMessage(&_BridgeUnproxied.TransactOpts, kind, sender, messageDataHash)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeUnproxied *BridgeUnproxiedTransactor) EnqueueSequencerMessage(opts *bind.TransactOpts, dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "enqueueSequencerMessage", dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeUnproxied *BridgeUnproxiedSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.EnqueueSequencerMessage(&_BridgeUnproxied.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// EnqueueSequencerMessage is a paid mutator transaction binding the contract method 0x86598a56.
//
// Solidity: function enqueueSequencerMessage(bytes32 dataHash, uint256 afterDelayedMessagesRead, uint256 prevMessageCount, uint256 newMessageCount) returns(uint256 seqMessageIndex, bytes32 beforeAcc, bytes32 delayedAcc, bytes32 acc)
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) EnqueueSequencerMessage(dataHash [32]byte, afterDelayedMessagesRead *big.Int, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.EnqueueSequencerMessage(&_BridgeUnproxied.TransactOpts, dataHash, afterDelayedMessagesRead, prevMessageCount, newMessageCount)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeUnproxied *BridgeUnproxiedTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "executeCall", to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeUnproxied *BridgeUnproxiedSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.ExecuteCall(&_BridgeUnproxied.TransactOpts, to, value, data)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes data) returns(bool success, bytes returnData)
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) ExecuteCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.ExecuteCall(&_BridgeUnproxied.TransactOpts, to, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) Initialize(opts *bind.TransactOpts, rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "initialize", rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.Initialize(&_BridgeUnproxied.TransactOpts, rollup_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address rollup_) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) Initialize(rollup_ common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.Initialize(&_BridgeUnproxied.TransactOpts, rollup_)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SetDelayedInbox(opts *bind.TransactOpts, inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "setDelayedInbox", inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetDelayedInbox(&_BridgeUnproxied.TransactOpts, inbox, enabled)
}

// SetDelayedInbox is a paid mutator transaction binding the contract method 0x47fb24c5.
//
// Solidity: function setDelayedInbox(address inbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SetDelayedInbox(inbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetDelayedInbox(&_BridgeUnproxied.TransactOpts, inbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SetOutbox(opts *bind.TransactOpts, outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "setOutbox", outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetOutbox(&_BridgeUnproxied.TransactOpts, outbox, enabled)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xcee3d728.
//
// Solidity: function setOutbox(address outbox, bool enabled) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SetOutbox(outbox common.Address, enabled bool) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetOutbox(&_BridgeUnproxied.TransactOpts, outbox, enabled)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SetSequencerInbox(opts *bind.TransactOpts, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "setSequencerInbox", _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetSequencerInbox(&_BridgeUnproxied.TransactOpts, _sequencerInbox)
}

// SetSequencerInbox is a paid mutator transaction binding the contract method 0x4f61f850.
//
// Solidity: function setSequencerInbox(address _sequencerInbox) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SetSequencerInbox(_sequencerInbox common.Address) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetSequencerInbox(&_BridgeUnproxied.TransactOpts, _sequencerInbox)
}

// SetSequencerReportedSubMessageCount is a paid mutator transaction binding the contract method 0xf81ff3b3.
//
// Solidity: function setSequencerReportedSubMessageCount(uint256 newMsgCount) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SetSequencerReportedSubMessageCount(opts *bind.TransactOpts, newMsgCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "setSequencerReportedSubMessageCount", newMsgCount)
}

// SetSequencerReportedSubMessageCount is a paid mutator transaction binding the contract method 0xf81ff3b3.
//
// Solidity: function setSequencerReportedSubMessageCount(uint256 newMsgCount) returns()
func (_BridgeUnproxied *BridgeUnproxiedSession) SetSequencerReportedSubMessageCount(newMsgCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetSequencerReportedSubMessageCount(&_BridgeUnproxied.TransactOpts, newMsgCount)
}

// SetSequencerReportedSubMessageCount is a paid mutator transaction binding the contract method 0xf81ff3b3.
//
// Solidity: function setSequencerReportedSubMessageCount(uint256 newMsgCount) returns()
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SetSequencerReportedSubMessageCount(newMsgCount *big.Int) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SetSequencerReportedSubMessageCount(&_BridgeUnproxied.TransactOpts, newMsgCount)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedTransactor) SubmitBatchSpendingReport(opts *bind.TransactOpts, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.contract.Transact(opts, "submitBatchSpendingReport", sender, messageDataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedSession) SubmitBatchSpendingReport(sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SubmitBatchSpendingReport(&_BridgeUnproxied.TransactOpts, sender, messageDataHash)
}

// SubmitBatchSpendingReport is a paid mutator transaction binding the contract method 0x7a88b107.
//
// Solidity: function submitBatchSpendingReport(address sender, bytes32 messageDataHash) returns(uint256)
func (_BridgeUnproxied *BridgeUnproxiedTransactorSession) SubmitBatchSpendingReport(sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _BridgeUnproxied.Contract.SubmitBatchSpendingReport(&_BridgeUnproxied.TransactOpts, sender, messageDataHash)
}

// BridgeUnproxiedBridgeCallTriggeredIterator is returned from FilterBridgeCallTriggered and is used to iterate over the raw logs and unpacked data for BridgeCallTriggered events raised by the BridgeUnproxied contract.
type BridgeUnproxiedBridgeCallTriggeredIterator struct {
	Event *BridgeUnproxiedBridgeCallTriggered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeUnproxiedBridgeCallTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedBridgeCallTriggered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeUnproxiedBridgeCallTriggered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeUnproxiedBridgeCallTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedBridgeCallTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedBridgeCallTriggered represents a BridgeCallTriggered event raised by the BridgeUnproxied contract.
type BridgeUnproxiedBridgeCallTriggered struct {
	Outbox common.Address
	To     common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridgeCallTriggered is a free log retrieval operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterBridgeCallTriggered(opts *bind.FilterOpts, outbox []common.Address, to []common.Address) (*BridgeUnproxiedBridgeCallTriggeredIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedBridgeCallTriggeredIterator{contract: _BridgeUnproxied.contract, event: "BridgeCallTriggered", logs: logs, sub: sub}, nil
}

// WatchBridgeCallTriggered is a free log subscription operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchBridgeCallTriggered(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedBridgeCallTriggered, outbox []common.Address, to []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "BridgeCallTriggered", outboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedBridgeCallTriggered)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBridgeCallTriggered is a log parse operation binding the contract event 0x2d9d115ef3e4a606d698913b1eae831a3cdfe20d9a83d48007b0526749c3d466.
//
// Solidity: event BridgeCallTriggered(address indexed outbox, address indexed to, uint256 value, bytes data)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseBridgeCallTriggered(log types.Log) (*BridgeUnproxiedBridgeCallTriggered, error) {
	event := new(BridgeUnproxiedBridgeCallTriggered)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "BridgeCallTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedInboxToggleIterator is returned from FilterInboxToggle and is used to iterate over the raw logs and unpacked data for InboxToggle events raised by the BridgeUnproxied contract.
type BridgeUnproxiedInboxToggleIterator struct {
	Event *BridgeUnproxiedInboxToggle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeUnproxiedInboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedInboxToggle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeUnproxiedInboxToggle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeUnproxiedInboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedInboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedInboxToggle represents a InboxToggle event raised by the BridgeUnproxied contract.
type BridgeUnproxiedInboxToggle struct {
	Inbox   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInboxToggle is a free log retrieval operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterInboxToggle(opts *bind.FilterOpts, inbox []common.Address) (*BridgeUnproxiedInboxToggleIterator, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedInboxToggleIterator{contract: _BridgeUnproxied.contract, event: "InboxToggle", logs: logs, sub: sub}, nil
}

// WatchInboxToggle is a free log subscription operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchInboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedInboxToggle, inbox []common.Address) (event.Subscription, error) {

	var inboxRule []interface{}
	for _, inboxItem := range inbox {
		inboxRule = append(inboxRule, inboxItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "InboxToggle", inboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedInboxToggle)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "InboxToggle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboxToggle is a log parse operation binding the contract event 0x6675ce8882cb71637de5903a193d218cc0544be9c0650cb83e0955f6aa2bf521.
//
// Solidity: event InboxToggle(address indexed inbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseInboxToggle(log types.Log) (*BridgeUnproxiedInboxToggle, error) {
	event := new(BridgeUnproxiedInboxToggle)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "InboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedMessageDeliveredIterator is returned from FilterMessageDelivered and is used to iterate over the raw logs and unpacked data for MessageDelivered events raised by the BridgeUnproxied contract.
type BridgeUnproxiedMessageDeliveredIterator struct {
	Event *BridgeUnproxiedMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeUnproxiedMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeUnproxiedMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeUnproxiedMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedMessageDelivered represents a MessageDelivered event raised by the BridgeUnproxied contract.
type BridgeUnproxiedMessageDelivered struct {
	MessageIndex    *big.Int
	BeforeInboxAcc  [32]byte
	Inbox           common.Address
	Kind            uint8
	Sender          common.Address
	MessageDataHash [32]byte
	BaseFeeL1       *big.Int
	Timestamp       uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMessageDelivered is a free log retrieval operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterMessageDelivered(opts *bind.FilterOpts, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (*BridgeUnproxiedMessageDeliveredIterator, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedMessageDeliveredIterator{contract: _BridgeUnproxied.contract, event: "MessageDelivered", logs: logs, sub: sub}, nil
}

// WatchMessageDelivered is a free log subscription operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchMessageDelivered(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedMessageDelivered, messageIndex []*big.Int, beforeInboxAcc [][32]byte) (event.Subscription, error) {

	var messageIndexRule []interface{}
	for _, messageIndexItem := range messageIndex {
		messageIndexRule = append(messageIndexRule, messageIndexItem)
	}
	var beforeInboxAccRule []interface{}
	for _, beforeInboxAccItem := range beforeInboxAcc {
		beforeInboxAccRule = append(beforeInboxAccRule, beforeInboxAccItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "MessageDelivered", messageIndexRule, beforeInboxAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedMessageDelivered)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageDelivered is a log parse operation binding the contract event 0x5e3c1311ea442664e8b1611bfabef659120ea7a0a2cfc0667700bebc69cbffe1.
//
// Solidity: event MessageDelivered(uint256 indexed messageIndex, bytes32 indexed beforeInboxAcc, address inbox, uint8 kind, address sender, bytes32 messageDataHash, uint256 baseFeeL1, uint64 timestamp)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseMessageDelivered(log types.Log) (*BridgeUnproxiedMessageDelivered, error) {
	event := new(BridgeUnproxiedMessageDelivered)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "MessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedOutboxToggleIterator is returned from FilterOutboxToggle and is used to iterate over the raw logs and unpacked data for OutboxToggle events raised by the BridgeUnproxied contract.
type BridgeUnproxiedOutboxToggleIterator struct {
	Event *BridgeUnproxiedOutboxToggle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeUnproxiedOutboxToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedOutboxToggle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeUnproxiedOutboxToggle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeUnproxiedOutboxToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedOutboxToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedOutboxToggle represents a OutboxToggle event raised by the BridgeUnproxied contract.
type BridgeUnproxiedOutboxToggle struct {
	Outbox  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOutboxToggle is a free log retrieval operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterOutboxToggle(opts *bind.FilterOpts, outbox []common.Address) (*BridgeUnproxiedOutboxToggleIterator, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedOutboxToggleIterator{contract: _BridgeUnproxied.contract, event: "OutboxToggle", logs: logs, sub: sub}, nil
}

// WatchOutboxToggle is a free log subscription operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchOutboxToggle(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedOutboxToggle, outbox []common.Address) (event.Subscription, error) {

	var outboxRule []interface{}
	for _, outboxItem := range outbox {
		outboxRule = append(outboxRule, outboxItem)
	}

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "OutboxToggle", outboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedOutboxToggle)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutboxToggle is a log parse operation binding the contract event 0x49477e7356dbcb654ab85d7534b50126772d938130d1350e23e2540370c8dffa.
//
// Solidity: event OutboxToggle(address indexed outbox, bool enabled)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseOutboxToggle(log types.Log) (*BridgeUnproxiedOutboxToggle, error) {
	event := new(BridgeUnproxiedOutboxToggle)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "OutboxToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnproxiedSequencerInboxUpdatedIterator is returned from FilterSequencerInboxUpdated and is used to iterate over the raw logs and unpacked data for SequencerInboxUpdated events raised by the BridgeUnproxied contract.
type BridgeUnproxiedSequencerInboxUpdatedIterator struct {
	Event *BridgeUnproxiedSequencerInboxUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeUnproxiedSequencerInboxUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnproxiedSequencerInboxUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeUnproxiedSequencerInboxUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeUnproxiedSequencerInboxUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnproxiedSequencerInboxUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnproxiedSequencerInboxUpdated represents a SequencerInboxUpdated event raised by the BridgeUnproxied contract.
type BridgeUnproxiedSequencerInboxUpdated struct {
	NewSequencerInbox common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencerInboxUpdated is a free log retrieval operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) FilterSequencerInboxUpdated(opts *bind.FilterOpts) (*BridgeUnproxiedSequencerInboxUpdatedIterator, error) {

	logs, sub, err := _BridgeUnproxied.contract.FilterLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeUnproxiedSequencerInboxUpdatedIterator{contract: _BridgeUnproxied.contract, event: "SequencerInboxUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerInboxUpdated is a free log subscription operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) WatchSequencerInboxUpdated(opts *bind.WatchOpts, sink chan<- *BridgeUnproxiedSequencerInboxUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeUnproxied.contract.WatchLogs(opts, "SequencerInboxUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnproxiedSequencerInboxUpdated)
				if err := _BridgeUnproxied.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerInboxUpdated is a log parse operation binding the contract event 0x8c1e6003ed33ca6748d4ad3dd4ecc949065c89dceb31fdf546a5289202763c6a.
//
// Solidity: event SequencerInboxUpdated(address newSequencerInbox)
func (_BridgeUnproxied *BridgeUnproxiedFilterer) ParseSequencerInboxUpdated(log types.Log) (*BridgeUnproxiedSequencerInboxUpdated, error) {
	event := new(BridgeUnproxiedSequencerInboxUpdated)
	if err := _BridgeUnproxied.contract.UnpackLog(event, "SequencerInboxUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxStubMetaData contains all meta data concerning the InboxStub contract.
var InboxStubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"calculateRetryableSubmissionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransactionToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransactionToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sendWithdrawEthToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"unsafeCreateRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506109b4806100206000396000f3fe60806040526004361061011e5760003560e01c80636e6e8a6a116100a0578063b75436bb11610064578063b75436bb14610261578063c474d2c514610281578063e6bd12cf14610202578063e78cea921461029f578063ee35f327146102d757600080fd5b80636e6e8a6a146101f457806370665f14146102105780638456cb59146101765780638a631aa61461022b578063a66b327d1461024657600080fd5b80635075788b116100e75780635075788b146101235780635c975abb146101b55780635e916758146101e6578063679b6ded146101f457806367ef3ab81461020257600080fd5b8062f72382146101235780631fe927cf146101565780633f4ba83a14610176578063439370b11461018d578063485cc95514610195575b600080fd5b34801561012f57600080fd5b5061014361013e3660046105ca565b6102f7565b6040519081526020015b60405180910390f35b34801561016257600080fd5b50610143610171366004610646565b610339565b34801561018257600080fd5b5061018b6103d2565b005b6101436102f7565b3480156101a157600080fd5b5061018b6101b0366004610687565b61040c565b3480156101c157600080fd5b506001546101d690600160a01b900460ff1681565b604051901515815260200161014d565b61014361013e3660046106c0565b61014361013e366004610729565b61014361013e3660046107cd565b34801561021c57600080fd5b5061014361013e36600461083f565b34801561023757600080fd5b5061014361013e36600461088c565b34801561025257600080fd5b5061014361013e3660046108e0565b34801561026d57600080fd5b5061014361027c366004610646565b610477565b34801561028d57600080fd5b5061018b61029c366004610902565b50565b3480156102ab57600080fd5b506000546102bf906001600160a01b031681565b6040516001600160a01b03909116815260200161014d565b3480156102e357600080fd5b506001546102bf906001600160a01b031681565b60405162461bcd60e51b815260206004820152600f60248201526e1393d517d253541311535153951151608a1b60448201526000906064015b60405180910390fd5b60003332146103785760405162461bcd60e51b815260206004820152600b60248201526a6f726967696e206f6e6c7960a81b6044820152606401610330565b600061039d6003338686604051610390929190610926565b60405180910390206104d3565b60405190915081907fab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c90600090a29392505050565b60405162461bcd60e51b815260206004820152600f60248201526e1393d5081253541311535153951151608a1b6044820152606401610330565b6000546001600160a01b0316156104545760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b6044820152606401610330565b50600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806104906003338686604051610390929190610926565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b85856040516104c4929190610936565b60405180910390a29392505050565b60008054604051638db5993b60e01b815260ff861660048201526001600160a01b0385811660248301526044820185905290911690638db5993b9034906064016020604051808303818588803b15801561052c57600080fd5b505af1158015610540573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906105659190610965565b949350505050565b6001600160a01b038116811461029c57600080fd5b60008083601f84011261059457600080fd5b5081356001600160401b038111156105ab57600080fd5b6020830191508360208285010111156105c357600080fd5b9250929050565b600080600080600080600060c0888a0312156105e557600080fd5b87359650602088013595506040880135945060608801356106058161056d565b93506080880135925060a08801356001600160401b0381111561062757600080fd5b6106338a828b01610582565b989b979a50959850939692959293505050565b6000806020838503121561065957600080fd5b82356001600160401b0381111561066f57600080fd5b61067b85828601610582565b90969095509350505050565b6000806040838503121561069a57600080fd5b82356106a58161056d565b915060208301356106b58161056d565b809150509250929050565b6000806000806000608086880312156106d857600080fd5b853594506020860135935060408601356106f18161056d565b925060608601356001600160401b0381111561070c57600080fd5b61071888828901610582565b969995985093965092949392505050565b60008060008060008060008060006101008a8c03121561074857600080fd5b89356107538161056d565b985060208a0135975060408a0135965060608a01356107718161056d565b955060808a01356107818161056d565b945060a08a0135935060c08a0135925060e08a01356001600160401b038111156107aa57600080fd5b6107b68c828d01610582565b915080935050809150509295985092959850929598565b60008060008060008060a087890312156107e657600080fd5b86359550602087013594506040870135935060608701356108068161056d565b925060808701356001600160401b0381111561082157600080fd5b61082d89828a01610582565b979a9699509497509295939492505050565b600080600080600060a0868803121561085757600080fd5b85359450602086013593506040860135925060608601359150608086013561087e8161056d565b809150509295509295909350565b60008060008060008060a087890312156108a557600080fd5b863595506020870135945060408701356108be8161056d565b93506060870135925060808701356001600160401b0381111561082157600080fd5b600080604083850312156108f357600080fd5b50508035926020909101359150565b60006020828403121561091457600080fd5b813561091f8161056d565b9392505050565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60006020828403121561097757600080fd5b505191905056fea26469706673582212204325f46f621b22c1c1fca290f32aa4aa3e79602f4adb8b6d24907cebd3e3dd6264736f6c63430008090033",
}

// InboxStubABI is the input ABI used to generate the binding from.
// Deprecated: Use InboxStubMetaData.ABI instead.
var InboxStubABI = InboxStubMetaData.ABI

// InboxStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InboxStubMetaData.Bin instead.
var InboxStubBin = InboxStubMetaData.Bin

// DeployInboxStub deploys a new Ethereum contract, binding an instance of InboxStub to it.
func DeployInboxStub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InboxStub, error) {
	parsed, err := InboxStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InboxStubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InboxStub{InboxStubCaller: InboxStubCaller{contract: contract}, InboxStubTransactor: InboxStubTransactor{contract: contract}, InboxStubFilterer: InboxStubFilterer{contract: contract}}, nil
}

// InboxStub is an auto generated Go binding around an Ethereum contract.
type InboxStub struct {
	InboxStubCaller     // Read-only binding to the contract
	InboxStubTransactor // Write-only binding to the contract
	InboxStubFilterer   // Log filterer for contract events
}

// InboxStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxStubSession struct {
	Contract     *InboxStub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxStubCallerSession struct {
	Contract *InboxStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// InboxStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxStubTransactorSession struct {
	Contract     *InboxStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InboxStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxStubRaw struct {
	Contract *InboxStub // Generic contract binding to access the raw methods on
}

// InboxStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxStubCallerRaw struct {
	Contract *InboxStubCaller // Generic read-only contract binding to access the raw methods on
}

// InboxStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxStubTransactorRaw struct {
	Contract *InboxStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInboxStub creates a new instance of InboxStub, bound to a specific deployed contract.
func NewInboxStub(address common.Address, backend bind.ContractBackend) (*InboxStub, error) {
	contract, err := bindInboxStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InboxStub{InboxStubCaller: InboxStubCaller{contract: contract}, InboxStubTransactor: InboxStubTransactor{contract: contract}, InboxStubFilterer: InboxStubFilterer{contract: contract}}, nil
}

// NewInboxStubCaller creates a new read-only instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubCaller(address common.Address, caller bind.ContractCaller) (*InboxStubCaller, error) {
	contract, err := bindInboxStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxStubCaller{contract: contract}, nil
}

// NewInboxStubTransactor creates a new write-only instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxStubTransactor, error) {
	contract, err := bindInboxStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxStubTransactor{contract: contract}, nil
}

// NewInboxStubFilterer creates a new log filterer instance of InboxStub, bound to a specific deployed contract.
func NewInboxStubFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxStubFilterer, error) {
	contract, err := bindInboxStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxStubFilterer{contract: contract}, nil
}

// bindInboxStub binds a generic wrapper to an already deployed contract.
func bindInboxStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InboxStubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxStub *InboxStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxStub.Contract.InboxStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxStub *InboxStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.Contract.InboxStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxStub *InboxStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxStub.Contract.InboxStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InboxStub *InboxStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InboxStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InboxStub *InboxStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InboxStub *InboxStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InboxStub.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubSession) Bridge() (common.Address, error) {
	return _InboxStub.Contract.Bridge(&_InboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_InboxStub *InboxStubCallerSession) Bridge() (common.Address, error) {
	return _InboxStub.Contract.Bridge(&_InboxStub.CallOpts)
}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 , uint256 ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) CalculateRetryableSubmissionFee(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "calculateRetryableSubmissionFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 , uint256 ) pure returns(uint256)
func (_InboxStub *InboxStubSession) CalculateRetryableSubmissionFee(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _InboxStub.Contract.CalculateRetryableSubmissionFee(&_InboxStub.CallOpts, arg0, arg1)
}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 , uint256 ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) CalculateRetryableSubmissionFee(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _InboxStub.Contract.CalculateRetryableSubmissionFee(&_InboxStub.CallOpts, arg0, arg1)
}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() pure returns()
func (_InboxStub *InboxStubCaller) Pause(opts *bind.CallOpts) error {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "pause")

	if err != nil {
		return err
	}

	return err

}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() pure returns()
func (_InboxStub *InboxStubSession) Pause() error {
	return _InboxStub.Contract.Pause(&_InboxStub.CallOpts)
}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() pure returns()
func (_InboxStub *InboxStubCallerSession) Pause() error {
	return _InboxStub.Contract.Pause(&_InboxStub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubSession) Paused() (bool, error) {
	return _InboxStub.Contract.Paused(&_InboxStub.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_InboxStub *InboxStubCallerSession) Paused() (bool, error) {
	return _InboxStub.Contract.Paused(&_InboxStub.CallOpts)
}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendContractTransaction(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendContractTransaction", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendContractTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendContractTransaction is a free data retrieval call binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendContractTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendUnsignedTransaction(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendUnsignedTransaction", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendUnsignedTransaction is a free data retrieval call binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransaction(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendUnsignedTransactionToFork is a free data retrieval call binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendUnsignedTransactionToFork(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendUnsignedTransactionToFork", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendUnsignedTransactionToFork is a free data retrieval call binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendUnsignedTransactionToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransactionToFork(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendUnsignedTransactionToFork is a free data retrieval call binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , uint256 , bytes ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendUnsignedTransactionToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*big.Int, error) {
	return _InboxStub.Contract.SendUnsignedTransactionToFork(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// SendWithdrawEthToFork is a free data retrieval call binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 , uint256 , uint256 , uint256 , address ) pure returns(uint256)
func (_InboxStub *InboxStubCaller) SendWithdrawEthToFork(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sendWithdrawEthToFork", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SendWithdrawEthToFork is a free data retrieval call binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 , uint256 , uint256 , uint256 , address ) pure returns(uint256)
func (_InboxStub *InboxStubSession) SendWithdrawEthToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address) (*big.Int, error) {
	return _InboxStub.Contract.SendWithdrawEthToFork(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendWithdrawEthToFork is a free data retrieval call binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 , uint256 , uint256 , uint256 , address ) pure returns(uint256)
func (_InboxStub *InboxStubCallerSession) SendWithdrawEthToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 common.Address) (*big.Int, error) {
	return _InboxStub.Contract.SendWithdrawEthToFork(&_InboxStub.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_InboxStub *InboxStubCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_InboxStub *InboxStubSession) SequencerInbox() (common.Address, error) {
	return _InboxStub.Contract.SequencerInbox(&_InboxStub.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_InboxStub *InboxStubCallerSession) SequencerInbox() (common.Address, error) {
	return _InboxStub.Contract.SequencerInbox(&_InboxStub.CallOpts)
}

// Unpause is a free data retrieval call binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() pure returns()
func (_InboxStub *InboxStubCaller) Unpause(opts *bind.CallOpts) error {
	var out []interface{}
	err := _InboxStub.contract.Call(opts, &out, "unpause")

	if err != nil {
		return err
	}

	return err

}

// Unpause is a free data retrieval call binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() pure returns()
func (_InboxStub *InboxStubSession) Unpause() error {
	return _InboxStub.Contract.Unpause(&_InboxStub.CallOpts)
}

// Unpause is a free data retrieval call binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() pure returns()
func (_InboxStub *InboxStubCallerSession) Unpause() error {
	return _InboxStub.Contract.Unpause(&_InboxStub.CallOpts)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) CreateRetryableTicket(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "createRetryableTicket", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) CreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.CreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) CreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.CreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubTransactor) DepositEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "depositEth")
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubSession) DepositEth() (*types.Transaction, error) {
	return _InboxStub.Contract.DepositEth(&_InboxStub.TransactOpts)
}

// DepositEth is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) DepositEth() (*types.Transaction, error) {
	return _InboxStub.Contract.DepositEth(&_InboxStub.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address ) returns()
func (_InboxStub *InboxStubTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "initialize", _bridge, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address ) returns()
func (_InboxStub *InboxStubSession) Initialize(_bridge common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.Initialize(&_InboxStub.TransactOpts, _bridge, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address ) returns()
func (_InboxStub *InboxStubTransactorSession) Initialize(_bridge common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.Initialize(&_InboxStub.TransactOpts, _bridge, arg1)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubTransactor) PostUpgradeInit(opts *bind.TransactOpts, _bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "postUpgradeInit", _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.PostUpgradeInit(&_InboxStub.TransactOpts, _bridge)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address _bridge) returns()
func (_InboxStub *InboxStubTransactorSession) PostUpgradeInit(_bridge common.Address) (*types.Transaction, error) {
	return _InboxStub.Contract.PostUpgradeInit(&_InboxStub.TransactOpts, _bridge)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL1FundedContractTransaction", arg0, arg1, arg2, arg3)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) SendL1FundedContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedContractTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL1FundedContractTransaction(arg0 *big.Int, arg1 *big.Int, arg2 common.Address, arg3 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedContractTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL1FundedUnsignedTransaction", arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) SendL1FundedUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL1FundedUnsignedTransaction(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransaction(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL1FundedUnsignedTransactionToFork(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL1FundedUnsignedTransactionToFork", arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) SendL1FundedUnsignedTransactionToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransactionToFork(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 , uint256 , uint256 , address , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL1FundedUnsignedTransactionToFork(arg0 *big.Int, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL1FundedUnsignedTransactionToFork(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2Message(&_InboxStub.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2Message(&_InboxStub.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2MessageFromOrigin(&_InboxStub.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_InboxStub *InboxStubTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.SendL2MessageFromOrigin(&_InboxStub.TransactOpts, messageData)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactor) UnsafeCreateRetryableTicket(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.contract.Transact(opts, "unsafeCreateRetryableTicket", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubSession) UnsafeCreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.UnsafeCreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address , uint256 , uint256 , address , address , uint256 , uint256 , bytes ) payable returns(uint256)
func (_InboxStub *InboxStubTransactorSession) UnsafeCreateRetryableTicket(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 common.Address, arg4 common.Address, arg5 *big.Int, arg6 *big.Int, arg7 []byte) (*types.Transaction, error) {
	return _InboxStub.Contract.UnsafeCreateRetryableTicket(&_InboxStub.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// InboxStubInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredIterator struct {
	Event *InboxStubInboxMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InboxStubInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxStubInboxMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InboxStubInboxMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InboxStubInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxStubInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxStubInboxMessageDelivered represents a InboxMessageDelivered event raised by the InboxStub contract.
type InboxStubInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxStubInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxStubInboxMessageDeliveredIterator{contract: _InboxStub.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxStubInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxStubInboxMessageDelivered)
				if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_InboxStub *InboxStubFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxStubInboxMessageDelivered, error) {
	event := new(InboxStubInboxMessageDelivered)
	if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxStubInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxStubInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxStubInboxMessageDeliveredFromOrigin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InboxStubInboxMessageDeliveredFromOrigin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxStubInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxStubInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the InboxStub contract.
type InboxStubInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxStubInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxStubInboxMessageDeliveredFromOriginIterator{contract: _InboxStub.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxStubInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _InboxStub.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxStubInboxMessageDeliveredFromOrigin)
				if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_InboxStub *InboxStubFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxStubInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxStubInboxMessageDeliveredFromOrigin)
	if err := _InboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockResultReceiverMetaData contains all meta data concerning the MockResultReceiver contract.
var MockResultReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"manager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"challengeIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"ChallengeCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"challengeIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeIndex_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"winner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser_\",\"type\":\"address\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"contractIChallengeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161050f38038061050f83398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b61047c806100936000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630357aa491461006757806314eab5e71461007c578063481c6a75146100ad578063d6853748146100d8578063dfbf53ae146100ef578063e82898b314610102575b600080fd5b61007a610075366004610235565b610115565b005b61008f61008a36600461028a565b61017b565b60405167ffffffffffffffff90911681526020015b60405180910390f35b6000546100c0906001600160a01b031681565b6040516001600160a01b0390911681526020016100a4565b6100e160035481565b6040519081526020016100a4565b6001546100c0906001600160a01b031681565b6002546100c0906001600160a01b031681565b600180546001600160a01b03199081166001600160a01b0385811691821790935560028054909216928416928317909155600385905560405185907f88cb1f3fe351f3ac338db9c36bff1ece1750423c7ae6dfc427cd194b1c69b12790600090a4505050565b600080546040516314eab5e760e01b81526001600160a01b03909116906314eab5e7906101ba908c908c908c908c908c908c908c908c90600401610395565b602060405180830381600087803b1580156101d457600080fd5b505af11580156101e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020c9190610422565b9998505050505050505050565b80356001600160a01b038116811461023057600080fd5b919050565b60008060006060848603121561024a57600080fd5b8335925061025a60208501610219565b915061026860408501610219565b90509250925092565b67ffffffffffffffff8116811461028757600080fd5b50565b600080600080600080600080610200898b0312156102a757600080fd5b88359750606089018a8111156102bc57600080fd5b60208a0197506101608a018b8111156102d457600080fd5b909650356102e181610271565b94506102f06101808a01610219565b93506102ff6101a08a01610219565b92506101c089013591506101e089013590509295985092959890939650565b806000805b6002808210610332575061038e565b6040808588378681018481529085019084905b8382101561037757823561035881610271565b67ffffffffffffffff1681526020928301926001929092019101610345565b505050608095860195939093019250600101610323565b5050505050565b888152610200810160208083018a6000805b60028110156103d0578235600481106103be578283fd5b845292840192918401916001016103a7565b50505050506103e2606083018961031e565b67ffffffffffffffff969096166101608201526001600160a01b03948516610180820152929093166101a08301526101c08201526101e001529392505050565b60006020828403121561043457600080fd5b815161043f81610271565b939250505056fea264697066735822122095e79d59e3d07c05fbcb2702e29a3c39566a759880063e3fe82f4333efbf25a564736f6c63430008090033",
}

// MockResultReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MockResultReceiverMetaData.ABI instead.
var MockResultReceiverABI = MockResultReceiverMetaData.ABI

// MockResultReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockResultReceiverMetaData.Bin instead.
var MockResultReceiverBin = MockResultReceiverMetaData.Bin

// DeployMockResultReceiver deploys a new Ethereum contract, binding an instance of MockResultReceiver to it.
func DeployMockResultReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, manager_ common.Address) (common.Address, *types.Transaction, *MockResultReceiver, error) {
	parsed, err := MockResultReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockResultReceiverBin), backend, manager_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockResultReceiver{MockResultReceiverCaller: MockResultReceiverCaller{contract: contract}, MockResultReceiverTransactor: MockResultReceiverTransactor{contract: contract}, MockResultReceiverFilterer: MockResultReceiverFilterer{contract: contract}}, nil
}

// MockResultReceiver is an auto generated Go binding around an Ethereum contract.
type MockResultReceiver struct {
	MockResultReceiverCaller     // Read-only binding to the contract
	MockResultReceiverTransactor // Write-only binding to the contract
	MockResultReceiverFilterer   // Log filterer for contract events
}

// MockResultReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockResultReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockResultReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockResultReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockResultReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockResultReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockResultReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockResultReceiverSession struct {
	Contract     *MockResultReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MockResultReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockResultReceiverCallerSession struct {
	Contract *MockResultReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MockResultReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockResultReceiverTransactorSession struct {
	Contract     *MockResultReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MockResultReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockResultReceiverRaw struct {
	Contract *MockResultReceiver // Generic contract binding to access the raw methods on
}

// MockResultReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockResultReceiverCallerRaw struct {
	Contract *MockResultReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MockResultReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockResultReceiverTransactorRaw struct {
	Contract *MockResultReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockResultReceiver creates a new instance of MockResultReceiver, bound to a specific deployed contract.
func NewMockResultReceiver(address common.Address, backend bind.ContractBackend) (*MockResultReceiver, error) {
	contract, err := bindMockResultReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiver{MockResultReceiverCaller: MockResultReceiverCaller{contract: contract}, MockResultReceiverTransactor: MockResultReceiverTransactor{contract: contract}, MockResultReceiverFilterer: MockResultReceiverFilterer{contract: contract}}, nil
}

// NewMockResultReceiverCaller creates a new read-only instance of MockResultReceiver, bound to a specific deployed contract.
func NewMockResultReceiverCaller(address common.Address, caller bind.ContractCaller) (*MockResultReceiverCaller, error) {
	contract, err := bindMockResultReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiverCaller{contract: contract}, nil
}

// NewMockResultReceiverTransactor creates a new write-only instance of MockResultReceiver, bound to a specific deployed contract.
func NewMockResultReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MockResultReceiverTransactor, error) {
	contract, err := bindMockResultReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiverTransactor{contract: contract}, nil
}

// NewMockResultReceiverFilterer creates a new log filterer instance of MockResultReceiver, bound to a specific deployed contract.
func NewMockResultReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MockResultReceiverFilterer, error) {
	contract, err := bindMockResultReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiverFilterer{contract: contract}, nil
}

// bindMockResultReceiver binds a generic wrapper to an already deployed contract.
func bindMockResultReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockResultReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockResultReceiver *MockResultReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockResultReceiver.Contract.MockResultReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockResultReceiver *MockResultReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.MockResultReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockResultReceiver *MockResultReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.MockResultReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockResultReceiver *MockResultReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockResultReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockResultReceiver *MockResultReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockResultReceiver *MockResultReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.contract.Transact(opts, method, params...)
}

// ChallengeIndex is a free data retrieval call binding the contract method 0xd6853748.
//
// Solidity: function challengeIndex() view returns(uint256)
func (_MockResultReceiver *MockResultReceiverCaller) ChallengeIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MockResultReceiver.contract.Call(opts, &out, "challengeIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeIndex is a free data retrieval call binding the contract method 0xd6853748.
//
// Solidity: function challengeIndex() view returns(uint256)
func (_MockResultReceiver *MockResultReceiverSession) ChallengeIndex() (*big.Int, error) {
	return _MockResultReceiver.Contract.ChallengeIndex(&_MockResultReceiver.CallOpts)
}

// ChallengeIndex is a free data retrieval call binding the contract method 0xd6853748.
//
// Solidity: function challengeIndex() view returns(uint256)
func (_MockResultReceiver *MockResultReceiverCallerSession) ChallengeIndex() (*big.Int, error) {
	return _MockResultReceiver.Contract.ChallengeIndex(&_MockResultReceiver.CallOpts)
}

// Loser is a free data retrieval call binding the contract method 0xe82898b3.
//
// Solidity: function loser() view returns(address)
func (_MockResultReceiver *MockResultReceiverCaller) Loser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockResultReceiver.contract.Call(opts, &out, "loser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Loser is a free data retrieval call binding the contract method 0xe82898b3.
//
// Solidity: function loser() view returns(address)
func (_MockResultReceiver *MockResultReceiverSession) Loser() (common.Address, error) {
	return _MockResultReceiver.Contract.Loser(&_MockResultReceiver.CallOpts)
}

// Loser is a free data retrieval call binding the contract method 0xe82898b3.
//
// Solidity: function loser() view returns(address)
func (_MockResultReceiver *MockResultReceiverCallerSession) Loser() (common.Address, error) {
	return _MockResultReceiver.Contract.Loser(&_MockResultReceiver.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_MockResultReceiver *MockResultReceiverCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockResultReceiver.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_MockResultReceiver *MockResultReceiverSession) Manager() (common.Address, error) {
	return _MockResultReceiver.Contract.Manager(&_MockResultReceiver.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_MockResultReceiver *MockResultReceiverCallerSession) Manager() (common.Address, error) {
	return _MockResultReceiver.Contract.Manager(&_MockResultReceiver.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_MockResultReceiver *MockResultReceiverCaller) Winner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockResultReceiver.contract.Call(opts, &out, "winner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_MockResultReceiver *MockResultReceiverSession) Winner() (common.Address, error) {
	return _MockResultReceiver.Contract.Winner(&_MockResultReceiver.CallOpts)
}

// Winner is a free data retrieval call binding the contract method 0xdfbf53ae.
//
// Solidity: function winner() view returns(address)
func (_MockResultReceiver *MockResultReceiverCallerSession) Winner() (common.Address, error) {
	return _MockResultReceiver.Contract.Winner(&_MockResultReceiver.CallOpts)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex_, address winner_, address loser_) returns()
func (_MockResultReceiver *MockResultReceiverTransactor) CompleteChallenge(opts *bind.TransactOpts, challengeIndex_ *big.Int, winner_ common.Address, loser_ common.Address) (*types.Transaction, error) {
	return _MockResultReceiver.contract.Transact(opts, "completeChallenge", challengeIndex_, winner_, loser_)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex_, address winner_, address loser_) returns()
func (_MockResultReceiver *MockResultReceiverSession) CompleteChallenge(challengeIndex_ *big.Int, winner_ common.Address, loser_ common.Address) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.CompleteChallenge(&_MockResultReceiver.TransactOpts, challengeIndex_, winner_, loser_)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x0357aa49.
//
// Solidity: function completeChallenge(uint256 challengeIndex_, address winner_, address loser_) returns()
func (_MockResultReceiver *MockResultReceiverTransactorSession) CompleteChallenge(challengeIndex_ *big.Int, winner_ common.Address, loser_ common.Address) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.CompleteChallenge(&_MockResultReceiver.TransactOpts, challengeIndex_, winner_, loser_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_MockResultReceiver *MockResultReceiverTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _MockResultReceiver.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_MockResultReceiver *MockResultReceiverSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.CreateChallenge(&_MockResultReceiver.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_MockResultReceiver *MockResultReceiverTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _MockResultReceiver.Contract.CreateChallenge(&_MockResultReceiver.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// MockResultReceiverChallengeCompletedIterator is returned from FilterChallengeCompleted and is used to iterate over the raw logs and unpacked data for ChallengeCompleted events raised by the MockResultReceiver contract.
type MockResultReceiverChallengeCompletedIterator struct {
	Event *MockResultReceiverChallengeCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MockResultReceiverChallengeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockResultReceiverChallengeCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MockResultReceiverChallengeCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MockResultReceiverChallengeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockResultReceiverChallengeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockResultReceiverChallengeCompleted represents a ChallengeCompleted event raised by the MockResultReceiver contract.
type MockResultReceiverChallengeCompleted struct {
	ChallengeIndex *big.Int
	Winner         common.Address
	Loser          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeCompleted is a free log retrieval operation binding the contract event 0x88cb1f3fe351f3ac338db9c36bff1ece1750423c7ae6dfc427cd194b1c69b127.
//
// Solidity: event ChallengeCompleted(uint256 indexed challengeIndex, address indexed winner, address indexed loser)
func (_MockResultReceiver *MockResultReceiverFilterer) FilterChallengeCompleted(opts *bind.FilterOpts, challengeIndex []*big.Int, winner []common.Address, loser []common.Address) (*MockResultReceiverChallengeCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var loserRule []interface{}
	for _, loserItem := range loser {
		loserRule = append(loserRule, loserItem)
	}

	logs, sub, err := _MockResultReceiver.contract.FilterLogs(opts, "ChallengeCompleted", challengeIndexRule, winnerRule, loserRule)
	if err != nil {
		return nil, err
	}
	return &MockResultReceiverChallengeCompletedIterator{contract: _MockResultReceiver.contract, event: "ChallengeCompleted", logs: logs, sub: sub}, nil
}

// WatchChallengeCompleted is a free log subscription operation binding the contract event 0x88cb1f3fe351f3ac338db9c36bff1ece1750423c7ae6dfc427cd194b1c69b127.
//
// Solidity: event ChallengeCompleted(uint256 indexed challengeIndex, address indexed winner, address indexed loser)
func (_MockResultReceiver *MockResultReceiverFilterer) WatchChallengeCompleted(opts *bind.WatchOpts, sink chan<- *MockResultReceiverChallengeCompleted, challengeIndex []*big.Int, winner []common.Address, loser []common.Address) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var loserRule []interface{}
	for _, loserItem := range loser {
		loserRule = append(loserRule, loserItem)
	}

	logs, sub, err := _MockResultReceiver.contract.WatchLogs(opts, "ChallengeCompleted", challengeIndexRule, winnerRule, loserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockResultReceiverChallengeCompleted)
				if err := _MockResultReceiver.contract.UnpackLog(event, "ChallengeCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeCompleted is a log parse operation binding the contract event 0x88cb1f3fe351f3ac338db9c36bff1ece1750423c7ae6dfc427cd194b1c69b127.
//
// Solidity: event ChallengeCompleted(uint256 indexed challengeIndex, address indexed winner, address indexed loser)
func (_MockResultReceiver *MockResultReceiverFilterer) ParseChallengeCompleted(log types.Log) (*MockResultReceiverChallengeCompleted, error) {
	event := new(MockResultReceiverChallengeCompleted)
	if err := _MockResultReceiver.contract.UnpackLog(event, "ChallengeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProgramTestMetaData contains all meta data concerning the ProgramTest contract.
var ProgramTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"Hash\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callKeccak\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"expected\",\"type\":\"bytes\"}],\"name\":\"checkRevertData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fundedAccount\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"staticcallEvmData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"program\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"staticcallProgram\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e37806100206000396000f3fe60806040526004361061003f5760003560e01c80631d00bae4146100445780633fdd58e21461006657806396ec12e51461009c578063aba8c4ba146100af575b600080fd5b34801561005057600080fd5b5061006461005f366004610b53565b6100cf565b005b34801561007257600080fd5b50610086610081366004610b53565b6101d5565b6040516100939190610bf2565b60405180910390f35b6100866100aa366004610c0c565b61025f565b3480156100bb57600080fd5b506100866100ca366004610c8c565b610409565b600080846001600160a01b031684846040516100ec929190610cee565b6000604051808303816000865af19150503d8060008114610129576040519150601f19603f3d011682016040523d82523d6000602084013e61012e565b606091505b5091509150816101595760405162461bcd60e51b815260040161015090610cfe565b60405180910390fd5b600061016482610d23565b90507f224c8d9ad1bbf0f44a61d7bd8e7e9049b1a320e04b047da9910945675c31ba438160405161019791815260200190565b60405180910390a16101ac8460018188610d4a565b6040516101ba929190610cee565b604051809103902081146101cd57600080fd5b505050505050565b6060600080856001600160a01b031685856040516101f4929190610cee565b600060405180830381855afa9150503d806000811461022f576040519150601f19603f3d011682016040523d82523d6000602084013e610234565b606091505b5091509150816102565760405162461bcd60e51b815260040161015090610cfe565b95945050505050565b6060600080876001600160a01b031634888860405161027f929190610cee565b60006040518083038185875af1925050503d80600081146102bc576040519150601f19603f3d011682016040523d82523d6000602084013e6102c1565b606091505b509150915081156103095760405162461bcd60e51b8152602060048201526012602482015271756e6578706563746564207375636365737360701b6044820152606401610150565b805184146103545760405162461bcd60e51b81526020600482015260186024820152770eee4dedcce40e4caeccae4e840c8c2e8c240d8cadccee8d60431b6044820152606401610150565b60005b81518110156103fd5785858281811061037257610372610d74565b9050013560f81c60f81b6001600160f81b03191682828151811061039857610398610d74565b01602001516001600160f81b031916146103eb5760405162461bcd60e51b81526020600482015260146024820152730e4caeccae4e840c8c2e8c240dad2e6dac2e8c6d60631b6044820152606401610150565b806103f581610da0565b915050610357565b50979650505050505050565b6060600080876001600160a01b0316866001600160401b03168686604051610432929190610cee565b6000604051808303818686fa925050503d806000811461046e576040519150601f19603f3d011682016040523d82523d6000602084013e610473565b606091505b5060408051808201909152600d81526c0313637b1b590373ab6b132b91609d1b602082015291935091506069906001906104b89084906104b38443610dbb565b6109e5565b92506104ea836040518060400160405280600d81526020016c031b430b4b71034b2101010101609d1b815250466109e5565b925061051c836040518060400160405280600d81526020016c03130b9b2903332b2901010101609d1b815250486109e5565b925061054e836040518060400160405280600d81526020016c033b0b990383934b1b29010101609d1b8152503a6109e5565b9250610580836040518060400160405280600d81526020016c033b0b9903634b6b4ba1010101609d1b815250456109e5565b92506105b3836040518060400160405280600d81526020016c03b30b63ab2901010101010101609d1b81525060006109e5565b92506105e5836040518060400160405280600d81526020016c03a34b6b2b9ba30b6b81010101609d1b815250426109e5565b9250610621836040518060400160405280600d81526020016c03130b630b731b290101010101609d1b8152508b6001600160a01b0316316109e5565b925061065c836040518060400160405280600d81526020016c0393ab9ba1030b2323932b9b99609d1b8152508c6001600160a01b03166109e5565b9250610697836040518060400160405280600d81526020016c039b2b73232b91010101010101609d1b815250306001600160a01b03166109e5565b92506106d2836040518060400160405280600d81526020016c037b934b3b4b71010101010101609d1b815250326001600160a01b03166109e5565b925061070d836040518060400160405280600d81526020016c031b7b4b73130b9b2901010101609d1b815250416001600160a01b03166109e5565b925061074c836040518060400160405280600d81526020016c0e4eae6e840c6dec8cad0c2e6d609b1b8152508c6001600160a01b03163f60001c6109e5565b925061078b836040518060400160405280600d81526020016c030b9311031b7b232b430b9b41609d1b815250846001600160a01b03163f60001c6109e5565b92506107ca836040518060400160405280600d81526020016c032ba341031b7b232b430b9b41609d1b815250836001600160a01b03163f60001c6109e5565b925060008a6001600160a01b03163b6001600160401b038111156107f0576107f0610dd2565b6040519080825280601f01601f19166020018201604052801561081a576020820181803683370190505b50905060005b8b6001600160a01b03163b81101561088a5784818151811061084457610844610d74565b602001015160f81c60f81b82828151811061086157610861610d74565b60200101906001600160f81b031916908160001a9053508061088281610da0565b915050610820565b508a6001600160a01b0316803b806020016040519081016040528181526000908060200190933c805190602001208180519060200120146108f65760405162461bcd60e51b815260040161015090602080825260049082015263636f646560e01b604082015260600190565b60008b6001600160a01b03163b855161090f9190610dbb565b6001600160401b0381111561092657610926610dd2565b6040519080825280601f01601f191660200182016040528015610950576020820181803683370190505b5090506001600160a01b038c163b5b85518110156109d55785818151811061097a5761097a610d74565b602001015160f81c60f81b828e6001600160a01b03163b8361099c9190610dbb565b815181106109ac576109ac610d74565b60200101906001600160f81b031916908160001a905350806109cd81610da0565b91505061095f565b509b9a5050505050505050505050565b60606000848060200190518101906109fd9190610de8565b905083838214610a205760405162461bcd60e51b81526004016101509190610bf2565b50600060208651610a319190610dbb565b6001600160401b03811115610a4857610a48610dd2565b6040519080825280601f01601f191660200182016040528015610a72576020820181803683370190505b50905060205b8651811015610ae557868181518110610a9357610a93610d74565b602001015160f81c60f81b82602083610aac9190610dbb565b81518110610abc57610abc610d74565b60200101906001600160f81b031916908160001a90535080610add81610da0565b915050610a78565b5095945050505050565b80356001600160a01b0381168114610b0657600080fd5b919050565b60008083601f840112610b1d57600080fd5b5081356001600160401b03811115610b3457600080fd5b602083019150836020828501011115610b4c57600080fd5b9250929050565b600080600060408486031215610b6857600080fd5b610b7184610aef565b925060208401356001600160401b03811115610b8c57600080fd5b610b9886828701610b0b565b9497909650939450505050565b6000815180845260005b81811015610bcb57602081850181015186830182015201610baf565b81811115610bdd576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610c056020830184610ba5565b9392505050565b600080600080600060608688031215610c2457600080fd5b610c2d86610aef565b945060208601356001600160401b0380821115610c4957600080fd5b610c5589838a01610b0b565b90965094506040880135915080821115610c6e57600080fd5b50610c7b88828901610b0b565b969995985093965092949392505050565b600080600080600060808688031215610ca457600080fd5b610cad86610aef565b9450610cbb60208701610aef565b935060408601356001600160401b038082168214610cd857600080fd5b90935060608701359080821115610c6e57600080fd5b8183823760009101908152919050565b6020808252600b908201526a18d85b1b0819985a5b195960aa1b604082015260600190565b80516020808301519190811015610d44576000198160200360031b1b821691505b50919050565b60008085851115610d5a57600080fd5b83861115610d6757600080fd5b5050820193919092039150565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415610db457610db4610d8a565b5060010190565b600082821015610dcd57610dcd610d8a565b500390565b634e487b7160e01b600052604160045260246000fd5b600060208284031215610dfa57600080fd5b505191905056fea26469706673582212204e4d172112162ceeb68339d89e4407f76059a3ca36743a8aa92d78b00f4f9ad464736f6c63430008090033",
}

// ProgramTestABI is the input ABI used to generate the binding from.
// Deprecated: Use ProgramTestMetaData.ABI instead.
var ProgramTestABI = ProgramTestMetaData.ABI

// ProgramTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProgramTestMetaData.Bin instead.
var ProgramTestBin = ProgramTestMetaData.Bin

// DeployProgramTest deploys a new Ethereum contract, binding an instance of ProgramTest to it.
func DeployProgramTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProgramTest, error) {
	parsed, err := ProgramTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProgramTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProgramTest{ProgramTestCaller: ProgramTestCaller{contract: contract}, ProgramTestTransactor: ProgramTestTransactor{contract: contract}, ProgramTestFilterer: ProgramTestFilterer{contract: contract}}, nil
}

// ProgramTest is an auto generated Go binding around an Ethereum contract.
type ProgramTest struct {
	ProgramTestCaller     // Read-only binding to the contract
	ProgramTestTransactor // Write-only binding to the contract
	ProgramTestFilterer   // Log filterer for contract events
}

// ProgramTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProgramTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProgramTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProgramTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProgramTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProgramTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProgramTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProgramTestSession struct {
	Contract     *ProgramTest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProgramTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProgramTestCallerSession struct {
	Contract *ProgramTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ProgramTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProgramTestTransactorSession struct {
	Contract     *ProgramTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ProgramTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProgramTestRaw struct {
	Contract *ProgramTest // Generic contract binding to access the raw methods on
}

// ProgramTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProgramTestCallerRaw struct {
	Contract *ProgramTestCaller // Generic read-only contract binding to access the raw methods on
}

// ProgramTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProgramTestTransactorRaw struct {
	Contract *ProgramTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProgramTest creates a new instance of ProgramTest, bound to a specific deployed contract.
func NewProgramTest(address common.Address, backend bind.ContractBackend) (*ProgramTest, error) {
	contract, err := bindProgramTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProgramTest{ProgramTestCaller: ProgramTestCaller{contract: contract}, ProgramTestTransactor: ProgramTestTransactor{contract: contract}, ProgramTestFilterer: ProgramTestFilterer{contract: contract}}, nil
}

// NewProgramTestCaller creates a new read-only instance of ProgramTest, bound to a specific deployed contract.
func NewProgramTestCaller(address common.Address, caller bind.ContractCaller) (*ProgramTestCaller, error) {
	contract, err := bindProgramTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProgramTestCaller{contract: contract}, nil
}

// NewProgramTestTransactor creates a new write-only instance of ProgramTest, bound to a specific deployed contract.
func NewProgramTestTransactor(address common.Address, transactor bind.ContractTransactor) (*ProgramTestTransactor, error) {
	contract, err := bindProgramTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProgramTestTransactor{contract: contract}, nil
}

// NewProgramTestFilterer creates a new log filterer instance of ProgramTest, bound to a specific deployed contract.
func NewProgramTestFilterer(address common.Address, filterer bind.ContractFilterer) (*ProgramTestFilterer, error) {
	contract, err := bindProgramTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProgramTestFilterer{contract: contract}, nil
}

// bindProgramTest binds a generic wrapper to an already deployed contract.
func bindProgramTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProgramTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProgramTest *ProgramTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProgramTest.Contract.ProgramTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProgramTest *ProgramTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProgramTest.Contract.ProgramTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProgramTest *ProgramTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProgramTest.Contract.ProgramTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProgramTest *ProgramTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProgramTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProgramTest *ProgramTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProgramTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProgramTest *ProgramTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProgramTest.Contract.contract.Transact(opts, method, params...)
}

// StaticcallEvmData is a free data retrieval call binding the contract method 0xaba8c4ba.
//
// Solidity: function staticcallEvmData(address program, address fundedAccount, uint64 gas, bytes data) view returns(bytes)
func (_ProgramTest *ProgramTestCaller) StaticcallEvmData(opts *bind.CallOpts, program common.Address, fundedAccount common.Address, gas uint64, data []byte) ([]byte, error) {
	var out []interface{}
	err := _ProgramTest.contract.Call(opts, &out, "staticcallEvmData", program, fundedAccount, gas, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// StaticcallEvmData is a free data retrieval call binding the contract method 0xaba8c4ba.
//
// Solidity: function staticcallEvmData(address program, address fundedAccount, uint64 gas, bytes data) view returns(bytes)
func (_ProgramTest *ProgramTestSession) StaticcallEvmData(program common.Address, fundedAccount common.Address, gas uint64, data []byte) ([]byte, error) {
	return _ProgramTest.Contract.StaticcallEvmData(&_ProgramTest.CallOpts, program, fundedAccount, gas, data)
}

// StaticcallEvmData is a free data retrieval call binding the contract method 0xaba8c4ba.
//
// Solidity: function staticcallEvmData(address program, address fundedAccount, uint64 gas, bytes data) view returns(bytes)
func (_ProgramTest *ProgramTestCallerSession) StaticcallEvmData(program common.Address, fundedAccount common.Address, gas uint64, data []byte) ([]byte, error) {
	return _ProgramTest.Contract.StaticcallEvmData(&_ProgramTest.CallOpts, program, fundedAccount, gas, data)
}

// StaticcallProgram is a free data retrieval call binding the contract method 0x3fdd58e2.
//
// Solidity: function staticcallProgram(address program, bytes data) view returns(bytes)
func (_ProgramTest *ProgramTestCaller) StaticcallProgram(opts *bind.CallOpts, program common.Address, data []byte) ([]byte, error) {
	var out []interface{}
	err := _ProgramTest.contract.Call(opts, &out, "staticcallProgram", program, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// StaticcallProgram is a free data retrieval call binding the contract method 0x3fdd58e2.
//
// Solidity: function staticcallProgram(address program, bytes data) view returns(bytes)
func (_ProgramTest *ProgramTestSession) StaticcallProgram(program common.Address, data []byte) ([]byte, error) {
	return _ProgramTest.Contract.StaticcallProgram(&_ProgramTest.CallOpts, program, data)
}

// StaticcallProgram is a free data retrieval call binding the contract method 0x3fdd58e2.
//
// Solidity: function staticcallProgram(address program, bytes data) view returns(bytes)
func (_ProgramTest *ProgramTestCallerSession) StaticcallProgram(program common.Address, data []byte) ([]byte, error) {
	return _ProgramTest.Contract.StaticcallProgram(&_ProgramTest.CallOpts, program, data)
}

// CallKeccak is a paid mutator transaction binding the contract method 0x1d00bae4.
//
// Solidity: function callKeccak(address program, bytes data) returns()
func (_ProgramTest *ProgramTestTransactor) CallKeccak(opts *bind.TransactOpts, program common.Address, data []byte) (*types.Transaction, error) {
	return _ProgramTest.contract.Transact(opts, "callKeccak", program, data)
}

// CallKeccak is a paid mutator transaction binding the contract method 0x1d00bae4.
//
// Solidity: function callKeccak(address program, bytes data) returns()
func (_ProgramTest *ProgramTestSession) CallKeccak(program common.Address, data []byte) (*types.Transaction, error) {
	return _ProgramTest.Contract.CallKeccak(&_ProgramTest.TransactOpts, program, data)
}

// CallKeccak is a paid mutator transaction binding the contract method 0x1d00bae4.
//
// Solidity: function callKeccak(address program, bytes data) returns()
func (_ProgramTest *ProgramTestTransactorSession) CallKeccak(program common.Address, data []byte) (*types.Transaction, error) {
	return _ProgramTest.Contract.CallKeccak(&_ProgramTest.TransactOpts, program, data)
}

// CheckRevertData is a paid mutator transaction binding the contract method 0x96ec12e5.
//
// Solidity: function checkRevertData(address program, bytes data, bytes expected) payable returns(bytes)
func (_ProgramTest *ProgramTestTransactor) CheckRevertData(opts *bind.TransactOpts, program common.Address, data []byte, expected []byte) (*types.Transaction, error) {
	return _ProgramTest.contract.Transact(opts, "checkRevertData", program, data, expected)
}

// CheckRevertData is a paid mutator transaction binding the contract method 0x96ec12e5.
//
// Solidity: function checkRevertData(address program, bytes data, bytes expected) payable returns(bytes)
func (_ProgramTest *ProgramTestSession) CheckRevertData(program common.Address, data []byte, expected []byte) (*types.Transaction, error) {
	return _ProgramTest.Contract.CheckRevertData(&_ProgramTest.TransactOpts, program, data, expected)
}

// CheckRevertData is a paid mutator transaction binding the contract method 0x96ec12e5.
//
// Solidity: function checkRevertData(address program, bytes data, bytes expected) payable returns(bytes)
func (_ProgramTest *ProgramTestTransactorSession) CheckRevertData(program common.Address, data []byte, expected []byte) (*types.Transaction, error) {
	return _ProgramTest.Contract.CheckRevertData(&_ProgramTest.TransactOpts, program, data, expected)
}

// ProgramTestHashIterator is returned from FilterHash and is used to iterate over the raw logs and unpacked data for Hash events raised by the ProgramTest contract.
type ProgramTestHashIterator struct {
	Event *ProgramTestHash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProgramTestHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProgramTestHash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProgramTestHash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProgramTestHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProgramTestHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProgramTestHash represents a Hash event raised by the ProgramTest contract.
type ProgramTestHash struct {
	Result [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHash is a free log retrieval operation binding the contract event 0x224c8d9ad1bbf0f44a61d7bd8e7e9049b1a320e04b047da9910945675c31ba43.
//
// Solidity: event Hash(bytes32 result)
func (_ProgramTest *ProgramTestFilterer) FilterHash(opts *bind.FilterOpts) (*ProgramTestHashIterator, error) {

	logs, sub, err := _ProgramTest.contract.FilterLogs(opts, "Hash")
	if err != nil {
		return nil, err
	}
	return &ProgramTestHashIterator{contract: _ProgramTest.contract, event: "Hash", logs: logs, sub: sub}, nil
}

// WatchHash is a free log subscription operation binding the contract event 0x224c8d9ad1bbf0f44a61d7bd8e7e9049b1a320e04b047da9910945675c31ba43.
//
// Solidity: event Hash(bytes32 result)
func (_ProgramTest *ProgramTestFilterer) WatchHash(opts *bind.WatchOpts, sink chan<- *ProgramTestHash) (event.Subscription, error) {

	logs, sub, err := _ProgramTest.contract.WatchLogs(opts, "Hash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProgramTestHash)
				if err := _ProgramTest.contract.UnpackLog(event, "Hash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHash is a log parse operation binding the contract event 0x224c8d9ad1bbf0f44a61d7bd8e7e9049b1a320e04b047da9910945675c31ba43.
//
// Solidity: event Hash(bytes32 result)
func (_ProgramTest *ProgramTestFilterer) ParseHash(log types.Log) (*ProgramTestHash, error) {
	event := new(ProgramTestHash)
	if err := _ProgramTest.contract.UnpackLog(event, "Hash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyAdminForBindingMetaData contains all meta data concerning the ProxyAdminForBinding contract.
var ProxyAdminForBindingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"getProxyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTransparentUpgradeableProxy\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6107408061007e6000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100a65780637eff275e146100bd5780638da5cb5b146100dd5780639623609d146100f257806399a88ec414610105578063f2fde38b14610125578063f3b7dead14610145575b600080fd5b34801561007c57600080fd5b5061009061008b3660046104f6565b610165565b60405161009d919061051a565b60405180910390f35b3480156100b257600080fd5b506100bb6101f6565b005b3480156100c957600080fd5b506100bb6100d836600461052e565b61023a565b3480156100e957600080fd5b506100906102cb565b6100bb61010036600461057d565b6102da565b34801561011157600080fd5b506100bb61012036600461052e565b610370565b34801561013157600080fd5b506100bb6101403660046104f6565b6103cb565b34801561015157600080fd5b506100906101603660046104f6565b61046b565b6000806000836001600160a01b031660405161018b90635c60da1b60e01b815260040190565b600060405180830381855afa9150503d80600081146101c6576040519150601f19603f3d011682016040523d82523d6000602084013e6101cb565b606091505b5091509150816101da57600080fd5b808060200190518101906101ee9190610653565b949350505050565b336101ff6102cb565b6001600160a01b03161461022e5760405162461bcd60e51b815260040161022590610670565b60405180910390fd5b6102386000610491565b565b336102436102cb565b6001600160a01b0316146102695760405162461bcd60e51b815260040161022590610670565b6040516308f2839760e41b81526001600160a01b03831690638f2839709061029590849060040161051a565b600060405180830381600087803b1580156102af57600080fd5b505af11580156102c3573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b336102e36102cb565b6001600160a01b0316146103095760405162461bcd60e51b815260040161022590610670565b60405163278f794360e11b81526001600160a01b03841690634f1ef28690349061033990869086906004016106a5565b6000604051808303818588803b15801561035257600080fd5b505af1158015610366573d6000803e3d6000fd5b5050505050505050565b336103796102cb565b6001600160a01b03161461039f5760405162461bcd60e51b815260040161022590610670565b604051631b2ce7f360e11b81526001600160a01b03831690633659cfe69061029590849060040161051a565b336103d46102cb565b6001600160a01b0316146103fa5760405162461bcd60e51b815260040161022590610670565b6001600160a01b03811661045f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610225565b61046881610491565b50565b6000806000836001600160a01b031660405161018b906303e1469160e61b815260040190565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116811461046857600080fd5b60006020828403121561050857600080fd5b8135610513816104e1565b9392505050565b6001600160a01b0391909116815260200190565b6000806040838503121561054157600080fd5b823561054c816104e1565b9150602083013561055c816104e1565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60008060006060848603121561059257600080fd5b833561059d816104e1565b925060208401356105ad816104e1565b9150604084013567ffffffffffffffff808211156105ca57600080fd5b818601915086601f8301126105de57600080fd5b8135818111156105f0576105f0610567565b604051601f8201601f19908116603f0116810190838211818310171561061857610618610567565b8160405282815289602084870101111561063157600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60006020828403121561066557600080fd5b8151610513816104e1565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60018060a01b038316815260006020604081840152835180604085015260005b818110156106e1578581018301518582016060015282016106c5565b818111156106f3576000606083870101525b50601f01601f19169290920160600194935050505056fea264697066735822122030876067dd5282d3eb5b6a87529c77a035be5809eec6b650765b6feed77af1b464736f6c63430008090033",
}

// ProxyAdminForBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyAdminForBindingMetaData.ABI instead.
var ProxyAdminForBindingABI = ProxyAdminForBindingMetaData.ABI

// ProxyAdminForBindingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProxyAdminForBindingMetaData.Bin instead.
var ProxyAdminForBindingBin = ProxyAdminForBindingMetaData.Bin

// DeployProxyAdminForBinding deploys a new Ethereum contract, binding an instance of ProxyAdminForBinding to it.
func DeployProxyAdminForBinding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProxyAdminForBinding, error) {
	parsed, err := ProxyAdminForBindingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyAdminForBindingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProxyAdminForBinding{ProxyAdminForBindingCaller: ProxyAdminForBindingCaller{contract: contract}, ProxyAdminForBindingTransactor: ProxyAdminForBindingTransactor{contract: contract}, ProxyAdminForBindingFilterer: ProxyAdminForBindingFilterer{contract: contract}}, nil
}

// ProxyAdminForBinding is an auto generated Go binding around an Ethereum contract.
type ProxyAdminForBinding struct {
	ProxyAdminForBindingCaller     // Read-only binding to the contract
	ProxyAdminForBindingTransactor // Write-only binding to the contract
	ProxyAdminForBindingFilterer   // Log filterer for contract events
}

// ProxyAdminForBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyAdminForBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminForBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyAdminForBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminForBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyAdminForBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyAdminForBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxyAdminForBindingSession struct {
	Contract     *ProxyAdminForBinding // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ProxyAdminForBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyAdminForBindingCallerSession struct {
	Contract *ProxyAdminForBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ProxyAdminForBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyAdminForBindingTransactorSession struct {
	Contract     *ProxyAdminForBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ProxyAdminForBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyAdminForBindingRaw struct {
	Contract *ProxyAdminForBinding // Generic contract binding to access the raw methods on
}

// ProxyAdminForBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyAdminForBindingCallerRaw struct {
	Contract *ProxyAdminForBindingCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyAdminForBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyAdminForBindingTransactorRaw struct {
	Contract *ProxyAdminForBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxyAdminForBinding creates a new instance of ProxyAdminForBinding, bound to a specific deployed contract.
func NewProxyAdminForBinding(address common.Address, backend bind.ContractBackend) (*ProxyAdminForBinding, error) {
	contract, err := bindProxyAdminForBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminForBinding{ProxyAdminForBindingCaller: ProxyAdminForBindingCaller{contract: contract}, ProxyAdminForBindingTransactor: ProxyAdminForBindingTransactor{contract: contract}, ProxyAdminForBindingFilterer: ProxyAdminForBindingFilterer{contract: contract}}, nil
}

// NewProxyAdminForBindingCaller creates a new read-only instance of ProxyAdminForBinding, bound to a specific deployed contract.
func NewProxyAdminForBindingCaller(address common.Address, caller bind.ContractCaller) (*ProxyAdminForBindingCaller, error) {
	contract, err := bindProxyAdminForBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminForBindingCaller{contract: contract}, nil
}

// NewProxyAdminForBindingTransactor creates a new write-only instance of ProxyAdminForBinding, bound to a specific deployed contract.
func NewProxyAdminForBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyAdminForBindingTransactor, error) {
	contract, err := bindProxyAdminForBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminForBindingTransactor{contract: contract}, nil
}

// NewProxyAdminForBindingFilterer creates a new log filterer instance of ProxyAdminForBinding, bound to a specific deployed contract.
func NewProxyAdminForBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyAdminForBindingFilterer, error) {
	contract, err := bindProxyAdminForBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminForBindingFilterer{contract: contract}, nil
}

// bindProxyAdminForBinding binds a generic wrapper to an already deployed contract.
func bindProxyAdminForBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProxyAdminForBindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdminForBinding *ProxyAdminForBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdminForBinding.Contract.ProxyAdminForBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdminForBinding *ProxyAdminForBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.ProxyAdminForBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdminForBinding *ProxyAdminForBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.ProxyAdminForBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProxyAdminForBinding *ProxyAdminForBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProxyAdminForBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.contract.Transact(opts, method, params...)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingCaller) GetProxyAdmin(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdminForBinding.contract.Call(opts, &out, "getProxyAdmin", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _ProxyAdminForBinding.Contract.GetProxyAdmin(&_ProxyAdminForBinding.CallOpts, proxy)
}

// GetProxyAdmin is a free data retrieval call binding the contract method 0xf3b7dead.
//
// Solidity: function getProxyAdmin(address proxy) view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingCallerSession) GetProxyAdmin(proxy common.Address) (common.Address, error) {
	return _ProxyAdminForBinding.Contract.GetProxyAdmin(&_ProxyAdminForBinding.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingCaller) GetProxyImplementation(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdminForBinding.contract.Call(opts, &out, "getProxyImplementation", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _ProxyAdminForBinding.Contract.GetProxyImplementation(&_ProxyAdminForBinding.CallOpts, proxy)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x204e1c7a.
//
// Solidity: function getProxyImplementation(address proxy) view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingCallerSession) GetProxyImplementation(proxy common.Address) (common.Address, error) {
	return _ProxyAdminForBinding.Contract.GetProxyImplementation(&_ProxyAdminForBinding.CallOpts, proxy)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProxyAdminForBinding.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingSession) Owner() (common.Address, error) {
	return _ProxyAdminForBinding.Contract.Owner(&_ProxyAdminForBinding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProxyAdminForBinding *ProxyAdminForBindingCallerSession) Owner() (common.Address, error) {
	return _ProxyAdminForBinding.Contract.Owner(&_ProxyAdminForBinding.CallOpts)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactor) ChangeProxyAdmin(opts *bind.TransactOpts, proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.contract.Transact(opts, "changeProxyAdmin", proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.ChangeProxyAdmin(&_ProxyAdminForBinding.TransactOpts, proxy, newAdmin)
}

// ChangeProxyAdmin is a paid mutator transaction binding the contract method 0x7eff275e.
//
// Solidity: function changeProxyAdmin(address proxy, address newAdmin) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactorSession) ChangeProxyAdmin(proxy common.Address, newAdmin common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.ChangeProxyAdmin(&_ProxyAdminForBinding.TransactOpts, proxy, newAdmin)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProxyAdminForBinding.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.RenounceOwnership(&_ProxyAdminForBinding.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.RenounceOwnership(&_ProxyAdminForBinding.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.TransferOwnership(&_ProxyAdminForBinding.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.TransferOwnership(&_ProxyAdminForBinding.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactor) Upgrade(opts *bind.TransactOpts, proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.contract.Transact(opts, "upgrade", proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.Upgrade(&_ProxyAdminForBinding.TransactOpts, proxy, implementation)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address proxy, address implementation) returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactorSession) Upgrade(proxy common.Address, implementation common.Address) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.Upgrade(&_ProxyAdminForBinding.TransactOpts, proxy, implementation)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactor) UpgradeAndCall(opts *bind.TransactOpts, proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdminForBinding.contract.Transact(opts, "upgradeAndCall", proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.UpgradeAndCall(&_ProxyAdminForBinding.TransactOpts, proxy, implementation, data)
}

// UpgradeAndCall is a paid mutator transaction binding the contract method 0x9623609d.
//
// Solidity: function upgradeAndCall(address proxy, address implementation, bytes data) payable returns()
func (_ProxyAdminForBinding *ProxyAdminForBindingTransactorSession) UpgradeAndCall(proxy common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProxyAdminForBinding.Contract.UpgradeAndCall(&_ProxyAdminForBinding.TransactOpts, proxy, implementation, data)
}

// ProxyAdminForBindingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProxyAdminForBinding contract.
type ProxyAdminForBindingOwnershipTransferredIterator struct {
	Event *ProxyAdminForBindingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ProxyAdminForBindingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyAdminForBindingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ProxyAdminForBindingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ProxyAdminForBindingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyAdminForBindingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyAdminForBindingOwnershipTransferred represents a OwnershipTransferred event raised by the ProxyAdminForBinding contract.
type ProxyAdminForBindingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdminForBinding *ProxyAdminForBindingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProxyAdminForBindingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdminForBinding.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyAdminForBindingOwnershipTransferredIterator{contract: _ProxyAdminForBinding.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdminForBinding *ProxyAdminForBindingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProxyAdminForBindingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProxyAdminForBinding.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyAdminForBindingOwnershipTransferred)
				if err := _ProxyAdminForBinding.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProxyAdminForBinding *ProxyAdminForBindingFilterer) ParseOwnershipTransferred(log types.Log) (*ProxyAdminForBindingOwnershipTransferred, error) {
	event := new(ProxyAdminForBindingOwnershipTransferred)
	if err := _ProxyAdminForBinding.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SdkStorageMetaData contains all meta data concerning the SdkStorage contract.
var SdkStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"populate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061148c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063a7f437791461003b578063f809f20514610045575b600080fd5b61004361004d565b005b61004361041d565b6007805461005a906111b8565b159050610121576007805461006e816111b8565b8061008957634e487b7160e01b600052603160045260246000fd5b601f811180156100a057600181146100c257610118565b6001826021036101000a03600183039250600283028482191617935050610118565b8360005260206000208260208114610105576020600019808601828104949094018054601f959095169092036101000a0119909216825560011990940193610115565b81546000835560ff1916603e1794505b50505b5050905561004d565b601060088054610130906111b8565b905011156101f85760088054610145816111b8565b8061016057634e487b7160e01b600052603160045260246000fd5b601f811180156101775760018114610199576101ef565b6001826021036101000a036001830392506002830284821916179350506101ef565b83600052602060002082602081146101dc576020600019808601828104949094018054601f959095169092036101000a01199092168255600119909401936101ec565b81546000835560ff1916603e1794505b50505b50509055610121565b60408051808201909152600f8082526e7761736d2069732063757465203c3360881b602090920191825261022e91600991610e45565b505b6005541561027f576005805480610249576102496111f3565b60008281526020902060046000199092019182040180546001600160401b03600860038516026101000a02191690559055610230565b600654600110156102be57600680548061029b5761029b6111f3565b6001900381819060005260206000200160006102b79190610ec9565b905561027f565b60005b60088110156102f8576000818152600a6020526040902080546001600160a01b0319169055806102f08161121f565b9150506102c1565b5060086000908152600a60209081527f2c1fd36ba11b13b555f58753742999069764391f450ca8727fe8a3eeffe6777580546001600160a01b03191690911790555b6004816001600160a01b0316101561037f576001600160a01b0381166000908152600b6020526040812061036d91610eee565b806103778161123a565b91505061033a565b506004805480610391576103916111f3565b60008281526020812060026000199390930192830201805465ffffffffffff19168155600101819055915560188190556019819055601a819055601b556103da601c6000610f13565b6103e6601d6000610f34565b6021805465ffffffffffff19908116909155600060228190556023805483169055602481905560258054909216909155602655565b565b600080546001600160a81b0319166170011781556001805460306001600160a01b03199091161790556002805465ffffffffffff1916657fffffff002017905560406003555b6020816001600160401b031610156104e457600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db06004820401805460039092166008026101000a6001600160401b038181021990931692841602919091179055806104dc81611261565b915050610463565b50604d60056007815481106104fb576104fb61127e565b90600052602060002090600491828204019190066008026101000a8154816001600160401b0302191690836001600160401b0316021790555060005b600a81101561063f576006816001600160401b0381111561055a5761055a611294565b604051908082528060200260200182016040528015610583578160200160208202803683370190505b508154600181018355600092835260209283902082516105a99491909201920190610f70565b5060005b8181101561062c5780600683815481106105c9576105c961127e565b9060005260206000200182815481106105e4576105e461127e565b90600052602060002090600691828204019190066005026101000a81548164ffffffffff021916908364ffffffffff16021790555080806106249061121f565b9150506105ad565b50806106378161121f565b915050610537565b5060005b600a8110156107015760005b818110156106ee5760026006838154811061066c5761066c61127e565b9060005260206000200182815481106106875761068761127e565b90600052602060002090600691828204019190066005028282829054906101000a900464ffffffffff166106bb91906112aa565b92506101000a81548164ffffffffff021916908364ffffffffff16021790555080806106e69061121f565b91505061064f565b50806106f98161121f565b915050610643565b5060005b601f8160ff16101561075c576007816040516020016107259291906112d7565b60405160208183030381529060405260079080519060200190610749929190610e45565b50806107548161138c565b915050610705565b5060005b60508160ff1610156107b7576008816040516020016107809291906112d7565b604051602081830303815290604052600890805190602001906107a4929190610e45565b50806107af8161138c565b915050610760565b5060408051808201909152600f8082526e617262697472756d207374796c757360881b60209092019182526107ee91600991610e45565b5060005b6010811015610834576000818152600a6020526040902080546001600160a01b0319166001600160a01b0383161790558061082c8161121f565b9150506107f2565b5060005b6004816001600160a01b031610156108f85760408051600080825260208083018085526001600160a01b0386168352600b909152929020905161087b929061101d565b5060005b816001600160a01b031681116108e5576001600160a01b0382166000908152600b60209081526040822080546001810182559083529181902090820401805460ff601f9093166101000a92830219169091179055806108dd8161121f565b91505061087f565b50806108f08161123a565b915050610838565b5060005b60048160030b121561097b57600c8054600101808255600082905263ffffffff83169190829081106109305761093061127e565b60009182526020808320600386900b84529190910190526040902080546001600160a01b0319166001600160a01b039290921691909117905580610973816113ac565b9150506108fc565b5060005b60048160ff1610156109ed576109968160016113c6565b6001600160f81b031960f883901b166000908152600d6020526040812060ff9290921691906109c66002856113eb565b60ff16158152602081019190915260400160002055806109e58161138c565b91505061097f565b50604051657374796c757360d01b8152600290600e906006019081526040519081900360200190208154815461ffff19811661ffff9092169182178355835463ffffffff6201000091829004160265ffffffffffff1990911690911717815560019182015491015560005b6004811015610b0c5760048054600181018255600091909152600280547f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b928202928301805461ffff90921661ffff19831681178255925465ffffffffffff19909216909217620100009182900463ffffffff169091021790556003547f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c9091015580610b048161121f565b915050610a58565b5060408051808201909152600e8082526d4c3220697320666f7220796f752160901b6020909201918252610b4291601191610e45565b5060005b6005811015610b9d578060148160058110610b6357610b6361127e565b600a91828204019190066003026101000a81548162ffffff021916908362ffffff1602179055508080610b959061121f565b915050610b46565b5060005b6002811015610bef578060168160028110610bbe57610bbe61127e565b0180546001600160a01b0319166001600160a01b039290921691909117905580610be78161121f565b915050610ba1565b5060005b6004811015610c85578060188160048110610c1057610c1061127e565b01805460ff191660ff9290921691909117905560188160048110610c3657610c3661127e565b0154610c469060ff1660016113c6565b60188260048110610c5957610c5961127e565b01805460ff929092166101000261ff001990921691909117905580610c7d8161121f565b915050610bf3565b5060005b6003811015610d2f57610c9a6110b0565b60005b6004811215610cd85780828260048110610cb957610cb961127e565b600b9290920b6020909202015280610cd08161141b565b915050610c9d565b50601c8054600181018255600091909152610d1a906002027f0e4562a10381dec21b205ed72637e6b1b523bdd0e4d4d50af5cd23dd4500a211018260046110ce565b50508080610d279061121f565b915050610c89565b5060005b6004811015610dc25760005b600481600b0b1215610daf57601d8260048110610d5e57610d5e61127e565b01805460018082018355600092835260209092206002820401805492909116600c026101000a6001600160601b03818102199093169284160291909117905580610da781611434565b915050610d3f565b5080610dba8161121f565b915050610d33565b5060005b6003811015610e4257600260218260038110610de457610de461127e565b825460029190910291909101805461ffff90921661ffff19831681178255835465ffffffffffff1990931617620100009283900463ffffffff1690920291909117815560019182015491015580610e3a8161121f565b915050610dc6565b50565b828054610e51906111b8565b90600052602060002090601f016020900481019282610e735760008555610eb9565b82601f10610e8c57805160ff1916838001178555610eb9565b82800160010185558215610eb9579182015b82811115610eb9578251825591602001919060010190610e9e565b50610ec5929150611164565b5090565b508054600082556005016006900490600052602060002090810190610e429190611164565b50805460008255601f016020900490600052602060002090810190610e429190611164565b5080546000825560020290600052602060002090810190610e429190611179565b506000610f418282611193565b506001016000610f518282611193565b506001016000610f618282611193565b5061041b906001016000611193565b82805482825590600052602060002090600501600690048101928215610eb95791602002820160005b83821115610fdf57835183826101000a81548164ffffffffff021916908364ffffffffff1602179055509260200192600501602081600401049283019260010302610f99565b80156110105782816101000a81549064ffffffffff0219169055600501602081600401049283019260010302610fdf565b5050610ec5929150611164565b82805482825590600052602060002090601f01602090048101928215610eb95791602002820160005b8382111561108357835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302611046565b80156110105782816101000a81549060ff0219169055600101602081600001049283019260010302611083565b60405180608001604052806004906020820280368337509192915050565b600283019183908215610eb95791602002820160005b8382111561113157835183826101000a8154816001600160601b030219169083600b0b6001600160601b031602179055509260200192600c01602081600b010492830192600103026110e4565b80156110105782816101000a8154906001600160601b030219169055600c01602081600b01049283019260010302611131565b5b80821115610ec55760008155600101611165565b80821115610ec55760008082556001820155600201611179565b508054600082556001016002900490600052602060002090810190610e429190611164565b600181811c908216806111cc57607f821691505b602082108114156111ed57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561123357611233611209565b5060010190565b60006001600160a01b038281168082141561125757611257611209565b6001019392505050565b60006001600160401b038083168181141561125757611257611209565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b600064ffffffffff808316818516818304811182151516156112ce576112ce611209565b02949350505050565b600080845481600182811c9150808316806112f357607f831692505b602080841082141561131357634e487b7160e01b86526022600452602486fd5b818015611327576001811461133857611365565b60ff19861689528489019650611365565b60008b81526020902060005b8681101561135d5781548b820152908501908301611344565b505084890196505b50505061137e848860f81b6001600160f81b0319169052565b929092019695505050505050565b600060ff821660ff8114156113a3576113a3611209565b60010192915050565b60008160030b637fffffff8114156113a3576113a3611209565b600060ff821660ff84168060ff038211156113e3576113e3611209565b019392505050565b600060ff83168061140c57634e487b7160e01b600052601260045260246000fd5b8060ff84160691505092915050565b60006001600160ff1b0382141561123357611233611209565b600081600b0b6b7fffffffffffffffffffffff8114156113a3576113a361120956fea26469706673582212207fff0c6884ce2f8f45ede07cfe79337eaccb3448f321aae188fca64654463c7b64736f6c63430008090033",
}

// SdkStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SdkStorageMetaData.ABI instead.
var SdkStorageABI = SdkStorageMetaData.ABI

// SdkStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SdkStorageMetaData.Bin instead.
var SdkStorageBin = SdkStorageMetaData.Bin

// DeploySdkStorage deploys a new Ethereum contract, binding an instance of SdkStorage to it.
func DeploySdkStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SdkStorage, error) {
	parsed, err := SdkStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SdkStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SdkStorage{SdkStorageCaller: SdkStorageCaller{contract: contract}, SdkStorageTransactor: SdkStorageTransactor{contract: contract}, SdkStorageFilterer: SdkStorageFilterer{contract: contract}}, nil
}

// SdkStorage is an auto generated Go binding around an Ethereum contract.
type SdkStorage struct {
	SdkStorageCaller     // Read-only binding to the contract
	SdkStorageTransactor // Write-only binding to the contract
	SdkStorageFilterer   // Log filterer for contract events
}

// SdkStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SdkStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdkStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SdkStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdkStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SdkStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SdkStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SdkStorageSession struct {
	Contract     *SdkStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SdkStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SdkStorageCallerSession struct {
	Contract *SdkStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SdkStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SdkStorageTransactorSession struct {
	Contract     *SdkStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SdkStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SdkStorageRaw struct {
	Contract *SdkStorage // Generic contract binding to access the raw methods on
}

// SdkStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SdkStorageCallerRaw struct {
	Contract *SdkStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SdkStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SdkStorageTransactorRaw struct {
	Contract *SdkStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSdkStorage creates a new instance of SdkStorage, bound to a specific deployed contract.
func NewSdkStorage(address common.Address, backend bind.ContractBackend) (*SdkStorage, error) {
	contract, err := bindSdkStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SdkStorage{SdkStorageCaller: SdkStorageCaller{contract: contract}, SdkStorageTransactor: SdkStorageTransactor{contract: contract}, SdkStorageFilterer: SdkStorageFilterer{contract: contract}}, nil
}

// NewSdkStorageCaller creates a new read-only instance of SdkStorage, bound to a specific deployed contract.
func NewSdkStorageCaller(address common.Address, caller bind.ContractCaller) (*SdkStorageCaller, error) {
	contract, err := bindSdkStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SdkStorageCaller{contract: contract}, nil
}

// NewSdkStorageTransactor creates a new write-only instance of SdkStorage, bound to a specific deployed contract.
func NewSdkStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SdkStorageTransactor, error) {
	contract, err := bindSdkStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SdkStorageTransactor{contract: contract}, nil
}

// NewSdkStorageFilterer creates a new log filterer instance of SdkStorage, bound to a specific deployed contract.
func NewSdkStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SdkStorageFilterer, error) {
	contract, err := bindSdkStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SdkStorageFilterer{contract: contract}, nil
}

// bindSdkStorage binds a generic wrapper to an already deployed contract.
func bindSdkStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SdkStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SdkStorage *SdkStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SdkStorage.Contract.SdkStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SdkStorage *SdkStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdkStorage.Contract.SdkStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SdkStorage *SdkStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SdkStorage.Contract.SdkStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SdkStorage *SdkStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SdkStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SdkStorage *SdkStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdkStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SdkStorage *SdkStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SdkStorage.Contract.contract.Transact(opts, method, params...)
}

// Populate is a paid mutator transaction binding the contract method 0xf809f205.
//
// Solidity: function populate() returns()
func (_SdkStorage *SdkStorageTransactor) Populate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdkStorage.contract.Transact(opts, "populate")
}

// Populate is a paid mutator transaction binding the contract method 0xf809f205.
//
// Solidity: function populate() returns()
func (_SdkStorage *SdkStorageSession) Populate() (*types.Transaction, error) {
	return _SdkStorage.Contract.Populate(&_SdkStorage.TransactOpts)
}

// Populate is a paid mutator transaction binding the contract method 0xf809f205.
//
// Solidity: function populate() returns()
func (_SdkStorage *SdkStorageTransactorSession) Populate() (*types.Transaction, error) {
	return _SdkStorage.Contract.Populate(&_SdkStorage.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_SdkStorage *SdkStorageTransactor) Remove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SdkStorage.contract.Transact(opts, "remove")
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_SdkStorage *SdkStorageSession) Remove() (*types.Transaction, error) {
	return _SdkStorage.Contract.Remove(&_SdkStorage.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_SdkStorage *SdkStorageTransactorSession) Remove() (*types.Transaction, error) {
	return _SdkStorage.Contract.Remove(&_SdkStorage.TransactOpts)
}

// SequencerInboxStubMetaData contains all meta data concerning the SequencerInboxStub contract.
var SequencerInboxStubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"AlreadyValidDASKeyset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"BadSequencerNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DataNotAuthenticated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataLength\",\"type\":\"uint256\"}],\"name\":\"DataTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedBackwards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayedTooFar\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceIncludeBlockTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceIncludeTimeTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HadZeroInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectMessagePreimage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"NoSuchKeyset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBatchPoster\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotForked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrigin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidateKeyset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SequencerBatchData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchSequenceNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"beforeAcc\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"afterAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"delayedAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"minTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxBlockNumber\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structISequencerInbox.TimeBounds\",\"name\":\"timeBounds\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"enumISequencerInbox.BatchDataLocation\",\"name\":\"dataLocation\",\"type\":\"uint8\"}],\"name\":\"SequencerBatchDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keysetHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"SetValidKeyset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DATA_AUTHENTICATED_FLAG\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"addInitMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"addSequencerL2Batch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sequenceNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"afterDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"contractIGasRefunder\",\"name\":\"gasRefunder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"prevMessageCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMessageCount\",\"type\":\"uint256\"}],\"name\":\"addSequencerL2BatchFromOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"dasKeySetInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidKeyset\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"creationBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalDelayedMessagesRead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint64[2]\",\"name\":\"l1BlockAndTime\",\"type\":\"uint64[2]\"},{\"internalType\":\"uint256\",\"name\":\"baseFeeL1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageDataHash\",\"type\":\"bytes32\"}],\"name\":\"forceInclusion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"getKeysetCreationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"inboxAccs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"invalidateKeysetHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBatchPoster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ksHash\",\"type\":\"bytes32\"}],\"name\":\"isValidKeysetHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTimeVariation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeDelayAfterFork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"contractIOwnable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isBatchPoster_\",\"type\":\"bool\"}],\"name\":\"setIsBatchPoster\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSequencer_\",\"type\":\"bool\"}],\"name\":\"setIsSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"delayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureSeconds\",\"type\":\"uint256\"}],\"internalType\":\"structISequencerInbox.MaxTimeVariation\",\"name\":\"maxTimeVariation_\",\"type\":\"tuple\"}],\"name\":\"setMaxTimeVariation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"keysetBytes\",\"type\":\"bytes\"}],\"name\":\"setValidKeyset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelayedMessagesRead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e0604052306080524660a05262000022620000c1602090811b620019f717901c565b151560c0523480156200003457600080fd5b5060405162002b7b38038062002b7b833981016040819052620000579162000177565b600180546001600160a01b039485166001600160a01b03199182161782556002805490911633179055815160045560208083015160055560408084015160065560609093015160075592909316600090815260039092529020805460ff1916909117905562000268565b60408051600481526024810182526020810180516001600160e01b03166302881c7960e11b1790529051600091829182916064916200010191906200022a565b600060405180830381855afa9150503d80600081146200013e576040519150601f19603f3d011682016040523d82523d6000602084013e62000143565b606091505b509150915081801562000157575080516020145b9250505090565b6001600160a01b03811681146200017457600080fd5b50565b600080600083850360c08112156200018e57600080fd5b84516200019b816200015e565b6020860151909450620001ae816200015e565b92506080603f1982011215620001c357600080fd5b50604051608081016001600160401b0381118282101715620001f557634e487b7160e01b600052604160045260246000fd5b806040525060408501518152606085015160208201526080850151604082015260a08501516060820152809150509250925092565b6000825160005b818110156200024d576020818601810151858301520162000231565b818111156200025d576000828501525b509190910192915050565b60805160a05160c0516128dc6200029f600039600081816112ec0152611c3701526000610fe40152600061048301526128dc6000f3fe608060405234801561001057600080fd5b50600436106101635760003560e01c80637fa3a40e116100ce578063d1ce8da811610087578063d1ce8da81461033f578063d9dd67ab14610352578063e0bc972914610365578063e5a358c814610378578063e78cea921461039c578063ebea461d146103af578063f1981578146103e557600080fd5b80637fa3a40e146102ca57806384420860146102d35780638f111f3c146102e657806396cc5c78146102f9578063b31761f814610301578063cb23bcb51461031457600080fd5b80636633ae85116101205780636633ae85146101f95780636d46e9871461020c5780636e7df3e71461022f5780636f12b0c914610242578063715ea34b1461025557806371c3e6fe146102a757600080fd5b806306f13056146101685780631637be48146101835780631f7a92b2146101b65780631f956632146101cb578063258f0495146101de57806327957a49146101f1575b600080fd5b6101706103f8565b6040519081526020015b60405180910390f35b6101a6610191366004612107565b60009081526008602052604090205460ff1690565b604051901515815260200161017a565b6101c96101c4366004612138565b610478565b005b6101c96101d9366004612187565b610630565b6101706101ec366004612107565b610736565b610170602881565b6101c9610207366004612107565b61079f565b6101a661021a3660046121c0565b60096020526000908152604090205460ff1681565b6101c961023d366004612187565b61098a565b6101c961025036600461222c565b610a90565b610288610263366004612107565b60086020526000908152604090205460ff81169061010090046001600160401b031682565b6040805192151583526001600160401b0390911660208301520161017a565b6101a66102b53660046121c0565b60036020526000908152604090205460ff1681565b61017060005481565b6101c96102e1366004612107565b610c7f565b6101c96102f4366004612296565b610dd2565b6101c9610fe1565b6101c961030f366004612312565b611058565b600254610327906001600160a01b031681565b6040516001600160a01b03909116815260200161017a565b6101c961034d366004612385565b611158565b610170610360366004612107565b61143c565b6101c9610373366004612296565b6114bf565b610383600160fe1b81565b6040516001600160f81b0319909116815260200161017a565b600154610327906001600160a01b031681565b6004546005546006546007546103c59392919084565b60408051948552602085019390935291830152606082015260800161017a565b6101c96103f33660046123c6565b611614565b600154604080516221048360e21b815290516000926001600160a01b0316916284120c916004808301926020929190829003018186803b15801561043b57600080fd5b505afa15801561044f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104739190612436565b905090565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016141561050b5760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b60648201526084015b60405180910390fd5b6001546001600160a01b03161561053557604051633bcd329760e21b815260040160405180910390fd5b6001600160a01b03821661055c57604051631ad0f74360e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0384169081179091556040805163cb23bcb560e01b8152905163cb23bcb591600480820192602092909190829003018186803b1580156105b057600080fd5b505afa1580156105c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e8919061244f565b600280546001600160a01b0319166001600160a01b03929092169190911790558035600490815560208201356005556040820135600655606082013560075581905b50505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561067e57600080fd5b505afa158015610692573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b6919061244f565b6001600160a01b0316336001600160a01b0316146106f857600254604051631194af8760e11b81526105029133916001600160a01b039091169060040161246c565b6001600160a01b038216600090815260096020526040808220805460ff19168415151790555160049160008051602061288783398151915291a25050565b600081815260086020908152604080832081518083019092525460ff81161515825261010090046001600160401b03169181018290529061078c5760405162f20c5d60e01b815260048101849052602401610502565b602001516001600160401b031692915050565b6000816040516020016107b491815260200190565b60408051808303601f190181529082905260015481516020830120638db5993b60e01b8452600b6004850152600060248501819052604485019190915291935090916001600160a01b0390911690638db5993b90606401602060405180830381600087803b15801561082557600080fd5b505af1158015610839573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085d9190612436565b905080156108a45760405162461bcd60e51b81526020600482015260146024820152731053149150511657d111531056515117d253925560621b6044820152606401610502565b807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b836040516108d491906124b2565b60405180910390a26000806108e96001611a8f565b915091506000806000806109038660016000806001611abb565b93509350935093508360001461094e5760405162461bcd60e51b815260206004820152601060248201526f1053149150511657d4d15457d253925560821b6044820152606401610502565b808385600080516020612867833981519152856000548a600260405161097794939291906124e5565b60405180910390a4505050505050505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156109d857600080fd5b505afa1580156109ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a10919061244f565b6001600160a01b0316336001600160a01b031614610a5257600254604051631194af8760e11b81526105029133916001600160a01b039091169060040161246c565b6001600160a01b038216600090815260036020526040808220805460ff19168415151790555160019160008051602061288783398151915291a25050565b8060005a9050333214610ab65760405163feb3d07160e01b815260040160405180910390fd5b3360009081526003602052604090205460ff16610ae657604051632dd9fc9760e01b815260040160405180910390fd5b600080610af4888888611ebb565b90925090506000808080610b0b868b8d8480611abb565b93509350935093508c8414610b3d5760405163ac7411c960e01b815260048101859052602481018e9052604401610502565b80838e600080516020612867833981519152856000548a6000604051610b6694939291906124e5565b60405180910390a4505050506001600160a01b038416159150610c769050573660006020610b9583601f61256f565b610b9f9190612587565b9050610200610baf60028361268d565b610bb99190612587565b610bc482600661269c565b610bce919061256f565b610bd8908461256f565b9250333214610be657600091505b836001600160a01b031663e3db8a49335a610c0190876126bb565b856040518463ffffffff1660e01b8152600401610c20939291906126d2565b602060405180830381600087803b158015610c3a57600080fd5b505af1158015610c4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7291906126f3565b5050505b50505050505050565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610ccd57600080fd5b505afa158015610ce1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d05919061244f565b6001600160a01b0316336001600160a01b031614610d4757600254604051631194af8760e11b81526105029133916001600160a01b039091169060040161246c565b60008181526008602052604090205460ff16610d785760405162f20c5d60e01b815260048101829052602401610502565b600081815260086020526040808220805460ff191690555182917f5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a91a260405160039060008051602061288783398151915290600090a250565b8260005a9050333214610df85760405163feb3d07160e01b815260040160405180910390fd5b3360009081526003602052604090205460ff16610e2857604051632dd9fc9760e01b815260040160405180910390fd5b600080610e368a8a8a611ebb565b90925090508a81838b8b8a8a6000808080610e5489888a8989611abb565b93509350935093508a8414158015610e6e57506000198b14155b15610e965760405163ac7411c960e01b815260048101859052602481018c9052604401610502565b808385600080516020612867833981519152856000548f6000604051610ebf94939291906124e5565b60405180910390a4505050506001600160a01b038b16159850610fd6975050505050505050573660006020610ef583601f61256f565b610eff9190612587565b9050610200610f0f60028361268d565b610f199190612587565b610f2482600661269c565b610f2e919061256f565b610f38908461256f565b9250333214610f4657600091505b836001600160a01b031663e3db8a49335a610f6190876126bb565b856040518463ffffffff1660e01b8152600401610f80939291906126d2565b602060405180830381600087803b158015610f9a57600080fd5b505af1158015610fae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fd291906126f3565b5050505b505050505050505050565b467f0000000000000000000000000000000000000000000000000000000000000000141561102257604051635180dd8360e11b815260040160405180910390fd5b60408051608081018252600180825260208201819052918101829052606001819052600481905560058190556006819055600755565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156110a657600080fd5b505afa1580156110ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110de919061244f565b6001600160a01b0316336001600160a01b03161461112057600254604051631194af8760e11b81526105029133916001600160a01b039091169060040161246c565b80516004556020810151600555604080820151600655606082015160075551600090600080516020612887833981519152908290a250565b600260009054906101000a90046001600160a01b03166001600160a01b0316638da5cb5b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156111a657600080fd5b505afa1580156111ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111de919061244f565b6001600160a01b0316336001600160a01b03161461122057600254604051631194af8760e11b81526105029133916001600160a01b039091169060040161246c565b60008282604051611232929190612710565b604051908190038120607f60f91b6020830152602182015260410160408051601f1981840301815291905280516020909101209050600160ff1b81186201000083106112b65760405162461bcd60e51b81526020600482015260136024820152726b657973657420697320746f6f206c6172676560681b6044820152606401610502565b60008181526008602052604090205460ff16156112e957604051637d17eeed60e11b815260048101829052602401610502565b437f0000000000000000000000000000000000000000000000000000000000000000156113855760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561134a57600080fd5b505afa15801561135e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113829190612436565b90505b604080518082018252600181526001600160401b0383811660208084019182526000878152600890915284902092518354915168ffffffffffffffffff1990921690151568ffffffffffffffff0019161761010091909216021790555182907fabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722906114139088908890612720565b60405180910390a260405160029060008051602061288783398151915290600090a25050505050565b6001546040516316bf557960e01b8152600481018390526000916001600160a01b0316906316bf55799060240160206040518083038186803b15801561148157600080fd5b505afa158015611495573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b99190612436565b92915050565b8260005a3360009081526003602052604090205490915060ff161580156114f157506002546001600160a01b03163314155b1561150f57604051632dd9fc9760e01b815260040160405180910390fd5b60008061151d8a8a8a611ebb565b909250905060008b82848b8a8a86808061153a8787838888611abb565b929c5090945092509050888a1480159061155657506000198914155b1561157e5760405163ac7411c960e01b8152600481018b9052602481018a9052604401610502565b80838b600080516020612867833981519152856000548d60016040516115a794939291906124e5565b60405180910390a4505050505050505050807ffe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c208c8c6040516115ea929190612720565b60405180910390a25050506001600160a01b03821615610fd6573660006020610ef583601f61256f565b600054861161163657604051633eb9f37d60e11b815260040160405180910390fd5b60006116e6868461164a6020890189612765565b61165a60408a0160208b01612765565b61166560018d6126bb565b6040805160f89690961b6001600160f81b03191660208088019190915260609590951b6001600160601b031916602187015260c093841b6001600160c01b031990811660358801529290931b909116603d85015260458401526065830188905260858084018790528151808503909101815260a59093019052815191012090565b60045490915043906116fb6020880188612765565b6001600160401b031661170e919061256f565b1061172c5760405163ad3515d960e01b815260040160405180910390fd5b60065442906117416040880160208901612765565b6001600160401b0316611754919061256f565b106117725760405163c76d17e560e01b815260040160405180910390fd5b6000600188111561180a576001546001600160a01b031663d5719dc261179960028b6126bb565b6040518263ffffffff1660e01b81526004016117b791815260200190565b60206040518083038186803b1580156117cf57600080fd5b505afa1580156117e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118079190612436565b90505b60408051602080820184905281830185905282518083038401815260609092019092528051910120600180546001600160a01b03169063d5719dc290611850908c6126bb565b6040518263ffffffff1660e01b815260040161186e91815260200190565b60206040518083038186803b15801561188657600080fd5b505afa15801561189a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118be9190612436565b146118dc576040516313947fd760e01b815260040160405180910390fd5b6000806118e88a611a8f565b9150915060008a90506000600160009054906101000a90046001600160a01b03166001600160a01b0316635fca4a166040518163ffffffff1660e01b815260040160206040518083038186803b15801561194157600080fd5b505afa158015611955573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119799190612436565b9050600080548d8361198b919061256f565b61199591906126bb565b90506000806000806119ab898860008989611abb565b9350935093509350808385600080516020612867833981519152856000548d60026040516119dc94939291906124e5565b60405180910390a45050505050505050505050505050505050565b60408051600481526024810182526020810180516001600160e01b03166302881c7960e11b179052905160009182918291606491611a35919061278e565b600060405180830381855afa9150503d8060008114611a70576040519150601f19603f3d011682016040523d82523d6000602084013e611a75565b606091505b5091509150818015611a88575080516020145b9250505090565b6000611a996120e0565b600080611aa58561202e565b8151602090920191909120969095509350505050565b600080600080600054881015611ae457604051633eb9f37d60e11b815260040160405180910390fd5b600160009054906101000a90046001600160a01b03166001600160a01b031663eca067ad6040518163ffffffff1660e01b815260040160206040518083038186803b158015611b3257600080fd5b505afa158015611b46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b6a9190612436565b881115611b8a5760405163925f8bd360e01b815260040160405180910390fd5b60015460405163432cc52b60e11b8152600481018b9052602481018a905260448101889052606481018790526001600160a01b03909116906386598a5690608401608060405180830381600087803b158015611be557600080fd5b505af1158015611bf9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c1d91906127aa565b60008c9055929650909450925090508615611eaf573360607f000000000000000000000000000000000000000000000000000000000000000015611d8a576000606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c9757600080fd5b505afa158015611cab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ccf9190612436565b90506000611cdd4883612587565b90506001600160401b03811115611d2a5760405162461bcd60e51b8152602060048201526011602482015270130c57d1d054d7d393d517d55253950d8d607a1b6044820152606401610502565b60408051426020820152606086901b6001600160601b03191681830152605481018f9052607481018a905248609482015260c09290921b6001600160c01b03191660b48301528051609c81840301815260bc90920190529150611dd39050565b604080514260208201526001600160601b0319606085901b1691810191909152605481018c90526074810187905248609482015260b40160405160208183030381529060405290505b60015481516020830120604051637a88b10760e01b81526000926001600160a01b031691637a88b10791611e1f9187916004016001600160a01b03929092168252602082015260400190565b602060405180830381600087803b158015611e3957600080fd5b505af1158015611e4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e719190612436565b9050807fff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b83604051611ea391906124b2565b60405180910390a25050505b95509550955095915050565b6000611ec56120e0565b84846000611ed482602861256f565b90506201cccc811115611f0657604051634634691b60e01b8152600481018290526201cccc6024820152604401610502565b8115801590611f3e5750600160fe1b808484600081611f2757611f2761274f565b9050013560f81c60f81b166001600160f81b031916145b15611f5c57604051631f97007f60e01b815260040160405180910390fd5b60218210801590611f8a575082826000818110611f7b57611f7b61274f565b90910135600160ff1b16151590505b15611fde576000611f9f6021600185876127e0565b611fa89161280a565b60008181526008602052604090205490915060ff16611fdc5760405162f20c5d60e01b815260048101829052602401610502565b505b600080611fea8861202e565b915091506000828b8b60405160200161200593929190612828565b60408051808303601f1901815291905280516020909101209b919a509098505050505050505050565b60606120386120e0565b60006120426120d5565b905060008160000151826020015183604001518460600151886040516020016120aa95949392919060c095861b6001600160c01b0319908116825294861b8516600882015292851b8416601084015290841b8316601883015290921b16602082015260280190565b604051602081830303815290604052905060288151146120cc576120cc612850565b94909350915050565b6120dd6120e0565b90565b60408051608081018252600080825260208201819052918101829052606081019190915290565b60006020828403121561211957600080fd5b5035919050565b6001600160a01b038116811461213557600080fd5b50565b60008082840360a081121561214c57600080fd5b833561215781612120565b92506080601f198201121561216b57600080fd5b506020830190509250929050565b801515811461213557600080fd5b6000806040838503121561219a57600080fd5b82356121a581612120565b915060208301356121b581612179565b809150509250929050565b6000602082840312156121d257600080fd5b81356121dd81612120565b9392505050565b60008083601f8401126121f657600080fd5b5081356001600160401b0381111561220d57600080fd5b60208301915083602082850101111561222557600080fd5b9250929050565b60008060008060006080868803121561224457600080fd5b8535945060208601356001600160401b0381111561226157600080fd5b61226d888289016121e4565b90955093505060408601359150606086013561228881612120565b809150509295509295909350565b600080600080600080600060c0888a0312156122b157600080fd5b8735965060208801356001600160401b038111156122ce57600080fd5b6122da8a828b016121e4565b9097509550506040880135935060608801356122f581612120565b969995985093969295946080840135945060a09093013592915050565b60006080828403121561232457600080fd5b604051608081018181106001600160401b038211171561235457634e487b7160e01b600052604160045260246000fd5b8060405250823581526020830135602082015260408301356040820152606083013560608201528091505092915050565b6000806020838503121561239857600080fd5b82356001600160401b038111156123ae57600080fd5b6123ba858286016121e4565b90969095509350505050565b60008060008060008060e087890312156123df57600080fd5b86359550602087013560ff811681146123f757600080fd5b9450608087018881111561240a57600080fd5b60408801945035925060a087013561242181612120565b8092505060c087013590509295509295509295565b60006020828403121561244857600080fd5b5051919050565b60006020828403121561246157600080fd5b81516121dd81612120565b6001600160a01b0392831681529116602082015260400190565b60005b838110156124a1578181015183820152602001612489565b8381111561062a5750506000910152565b60208152600082518060208401526124d1816040850160208701612486565b601f01601f19169190910160400192915050565b600060e0820190508582528460208301526001600160401b038085511660408401528060208601511660608401528060408601511660808401528060608601511660a0840152506003831061254a57634e487b7160e01b600052602160045260246000fd5b8260c083015295945050505050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561258257612582612559565b500190565b6000826125a457634e487b7160e01b600052601260045260246000fd5b500490565b600181815b808511156125e45781600019048211156125ca576125ca612559565b808516156125d757918102915b93841c93908002906125ae565b509250929050565b6000826125fb575060016114b9565b81612608575060006114b9565b816001811461261e576002811461262857612644565b60019150506114b9565b60ff84111561263957612639612559565b50506001821b6114b9565b5060208310610133831016604e8410600b8410161715612667575081810a6114b9565b61267183836125a9565b806000190482111561268557612685612559565b029392505050565b60006121dd60ff8416836125ec565b60008160001904831182151516156126b6576126b6612559565b500290565b6000828210156126cd576126cd612559565b500390565b6001600160a01b039390931683526020830191909152604082015260600190565b60006020828403121561270557600080fd5b81516121dd81612179565b8183823760009101908152919050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561277757600080fd5b81356001600160401b03811681146121dd57600080fd5b600082516127a0818460208701612486565b9190910192915050565b600080600080608085870312156127c057600080fd5b505082516020840151604085015160609095015191969095509092509050565b600080858511156127f057600080fd5b838611156127fd57600080fd5b5050820193919092039150565b803560208310156114b957600019602084900360031b1b1692915050565b6000845161283a818460208901612486565b8201838582376000930192835250909392505050565b634e487b7160e01b600052600160045260246000fdfe7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7ea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456ea2646970667358221220e2c5185529a80b7b57964883bbd082991aaf5d5aa81aa0dd246a6ac286f8bd0264736f6c63430008090033",
}

// SequencerInboxStubABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerInboxStubMetaData.ABI instead.
var SequencerInboxStubABI = SequencerInboxStubMetaData.ABI

// SequencerInboxStubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SequencerInboxStubMetaData.Bin instead.
var SequencerInboxStubBin = SequencerInboxStubMetaData.Bin

// DeploySequencerInboxStub deploys a new Ethereum contract, binding an instance of SequencerInboxStub to it.
func DeploySequencerInboxStub(auth *bind.TransactOpts, backend bind.ContractBackend, bridge_ common.Address, sequencer_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (common.Address, *types.Transaction, *SequencerInboxStub, error) {
	parsed, err := SequencerInboxStubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SequencerInboxStubBin), backend, bridge_, sequencer_, maxTimeVariation_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SequencerInboxStub{SequencerInboxStubCaller: SequencerInboxStubCaller{contract: contract}, SequencerInboxStubTransactor: SequencerInboxStubTransactor{contract: contract}, SequencerInboxStubFilterer: SequencerInboxStubFilterer{contract: contract}}, nil
}

// SequencerInboxStub is an auto generated Go binding around an Ethereum contract.
type SequencerInboxStub struct {
	SequencerInboxStubCaller     // Read-only binding to the contract
	SequencerInboxStubTransactor // Write-only binding to the contract
	SequencerInboxStubFilterer   // Log filterer for contract events
}

// SequencerInboxStubCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerInboxStubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerInboxStubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerInboxStubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerInboxStubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerInboxStubSession struct {
	Contract     *SequencerInboxStub // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SequencerInboxStubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerInboxStubCallerSession struct {
	Contract *SequencerInboxStubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SequencerInboxStubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerInboxStubTransactorSession struct {
	Contract     *SequencerInboxStubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SequencerInboxStubRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerInboxStubRaw struct {
	Contract *SequencerInboxStub // Generic contract binding to access the raw methods on
}

// SequencerInboxStubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerInboxStubCallerRaw struct {
	Contract *SequencerInboxStubCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerInboxStubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerInboxStubTransactorRaw struct {
	Contract *SequencerInboxStubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerInboxStub creates a new instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStub(address common.Address, backend bind.ContractBackend) (*SequencerInboxStub, error) {
	contract, err := bindSequencerInboxStub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStub{SequencerInboxStubCaller: SequencerInboxStubCaller{contract: contract}, SequencerInboxStubTransactor: SequencerInboxStubTransactor{contract: contract}, SequencerInboxStubFilterer: SequencerInboxStubFilterer{contract: contract}}, nil
}

// NewSequencerInboxStubCaller creates a new read-only instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubCaller(address common.Address, caller bind.ContractCaller) (*SequencerInboxStubCaller, error) {
	contract, err := bindSequencerInboxStub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubCaller{contract: contract}, nil
}

// NewSequencerInboxStubTransactor creates a new write-only instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerInboxStubTransactor, error) {
	contract, err := bindSequencerInboxStub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubTransactor{contract: contract}, nil
}

// NewSequencerInboxStubFilterer creates a new log filterer instance of SequencerInboxStub, bound to a specific deployed contract.
func NewSequencerInboxStubFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerInboxStubFilterer, error) {
	contract, err := bindSequencerInboxStub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubFilterer{contract: contract}, nil
}

// bindSequencerInboxStub binds a generic wrapper to an already deployed contract.
func bindSequencerInboxStub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerInboxStubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInboxStub *SequencerInboxStubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInboxStub.Contract.SequencerInboxStubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInboxStub *SequencerInboxStubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SequencerInboxStubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInboxStub *SequencerInboxStubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SequencerInboxStubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerInboxStub *SequencerInboxStubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerInboxStub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerInboxStub *SequencerInboxStubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerInboxStub *SequencerInboxStubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.contract.Transact(opts, method, params...)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCaller) DATAAUTHENTICATEDFLAG(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "DATA_AUTHENTICATED_FLAG")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInboxStub.CallOpts)
}

// DATAAUTHENTICATEDFLAG is a free data retrieval call binding the contract method 0xe5a358c8.
//
// Solidity: function DATA_AUTHENTICATED_FLAG() view returns(bytes1)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) DATAAUTHENTICATEDFLAG() ([1]byte, error) {
	return _SequencerInboxStub.Contract.DATAAUTHENTICATEDFLAG(&_SequencerInboxStub.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) HEADERLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "HEADER_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInboxStub.Contract.HEADERLENGTH(&_SequencerInboxStub.CallOpts)
}

// HEADERLENGTH is a free data retrieval call binding the contract method 0x27957a49.
//
// Solidity: function HEADER_LENGTH() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) HEADERLENGTH() (*big.Int, error) {
	return _SequencerInboxStub.Contract.HEADERLENGTH(&_SequencerInboxStub.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) BatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "batchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) BatchCount() (*big.Int, error) {
	return _SequencerInboxStub.Contract.BatchCount(&_SequencerInboxStub.CallOpts)
}

// BatchCount is a free data retrieval call binding the contract method 0x06f13056.
//
// Solidity: function batchCount() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) BatchCount() (*big.Int, error) {
	return _SequencerInboxStub.Contract.BatchCount(&_SequencerInboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) Bridge() (common.Address, error) {
	return _SequencerInboxStub.Contract.Bridge(&_SequencerInboxStub.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) Bridge() (common.Address, error) {
	return _SequencerInboxStub.Contract.Bridge(&_SequencerInboxStub.CallOpts)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubCaller) DasKeySetInfo(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "dasKeySetInfo", arg0)

	outstruct := new(struct {
		IsValidKeyset bool
		CreationBlock uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidKeyset = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CreationBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInboxStub.Contract.DasKeySetInfo(&_SequencerInboxStub.CallOpts, arg0)
}

// DasKeySetInfo is a free data retrieval call binding the contract method 0x715ea34b.
//
// Solidity: function dasKeySetInfo(bytes32 ) view returns(bool isValidKeyset, uint64 creationBlock)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) DasKeySetInfo(arg0 [32]byte) (struct {
	IsValidKeyset bool
	CreationBlock uint64
}, error) {
	return _SequencerInboxStub.Contract.DasKeySetInfo(&_SequencerInboxStub.CallOpts, arg0)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) GetKeysetCreationBlock(opts *bind.CallOpts, ksHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "getKeysetCreationBlock", ksHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInboxStub.Contract.GetKeysetCreationBlock(&_SequencerInboxStub.CallOpts, ksHash)
}

// GetKeysetCreationBlock is a free data retrieval call binding the contract method 0x258f0495.
//
// Solidity: function getKeysetCreationBlock(bytes32 ksHash) view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) GetKeysetCreationBlock(ksHash [32]byte) (*big.Int, error) {
	return _SequencerInboxStub.Contract.GetKeysetCreationBlock(&_SequencerInboxStub.CallOpts, ksHash)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubCaller) InboxAccs(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "inboxAccs", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInboxStub.Contract.InboxAccs(&_SequencerInboxStub.CallOpts, index)
}

// InboxAccs is a free data retrieval call binding the contract method 0xd9dd67ab.
//
// Solidity: function inboxAccs(uint256 index) view returns(bytes32)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) InboxAccs(index *big.Int) ([32]byte, error) {
	return _SequencerInboxStub.Contract.InboxAccs(&_SequencerInboxStub.CallOpts, index)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsBatchPoster(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isBatchPoster", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsBatchPoster(&_SequencerInboxStub.CallOpts, arg0)
}

// IsBatchPoster is a free data retrieval call binding the contract method 0x71c3e6fe.
//
// Solidity: function isBatchPoster(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsBatchPoster(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsBatchPoster(&_SequencerInboxStub.CallOpts, arg0)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isSequencer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsSequencer(&_SequencerInboxStub.CallOpts, arg0)
}

// IsSequencer is a free data retrieval call binding the contract method 0x6d46e987.
//
// Solidity: function isSequencer(address ) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsSequencer(arg0 common.Address) (bool, error) {
	return _SequencerInboxStub.Contract.IsSequencer(&_SequencerInboxStub.CallOpts, arg0)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCaller) IsValidKeysetHash(opts *bind.CallOpts, ksHash [32]byte) (bool, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "isValidKeysetHash", ksHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInboxStub.Contract.IsValidKeysetHash(&_SequencerInboxStub.CallOpts, ksHash)
}

// IsValidKeysetHash is a free data retrieval call binding the contract method 0x1637be48.
//
// Solidity: function isValidKeysetHash(bytes32 ksHash) view returns(bool)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) IsValidKeysetHash(ksHash [32]byte) (bool, error) {
	return _SequencerInboxStub.Contract.IsValidKeysetHash(&_SequencerInboxStub.CallOpts, ksHash)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256 delayBlocks, uint256 futureBlocks, uint256 delaySeconds, uint256 futureSeconds)
func (_SequencerInboxStub *SequencerInboxStubCaller) MaxTimeVariation(opts *bind.CallOpts) (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "maxTimeVariation")

	outstruct := new(struct {
		DelayBlocks   *big.Int
		FutureBlocks  *big.Int
		DelaySeconds  *big.Int
		FutureSeconds *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelayBlocks = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FutureBlocks = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DelaySeconds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FutureSeconds = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256 delayBlocks, uint256 futureBlocks, uint256 delaySeconds, uint256 futureSeconds)
func (_SequencerInboxStub *SequencerInboxStubSession) MaxTimeVariation() (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	return _SequencerInboxStub.Contract.MaxTimeVariation(&_SequencerInboxStub.CallOpts)
}

// MaxTimeVariation is a free data retrieval call binding the contract method 0xebea461d.
//
// Solidity: function maxTimeVariation() view returns(uint256 delayBlocks, uint256 futureBlocks, uint256 delaySeconds, uint256 futureSeconds)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) MaxTimeVariation() (struct {
	DelayBlocks   *big.Int
	FutureBlocks  *big.Int
	DelaySeconds  *big.Int
	FutureSeconds *big.Int
}, error) {
	return _SequencerInboxStub.Contract.MaxTimeVariation(&_SequencerInboxStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubSession) Rollup() (common.Address, error) {
	return _SequencerInboxStub.Contract.Rollup(&_SequencerInboxStub.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) Rollup() (common.Address, error) {
	return _SequencerInboxStub.Contract.Rollup(&_SequencerInboxStub.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCaller) TotalDelayedMessagesRead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SequencerInboxStub.contract.Call(opts, &out, "totalDelayedMessagesRead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInboxStub.Contract.TotalDelayedMessagesRead(&_SequencerInboxStub.CallOpts)
}

// TotalDelayedMessagesRead is a free data retrieval call binding the contract method 0x7fa3a40e.
//
// Solidity: function totalDelayedMessagesRead() view returns(uint256)
func (_SequencerInboxStub *SequencerInboxStubCallerSession) TotalDelayedMessagesRead() (*big.Int, error) {
	return _SequencerInboxStub.Contract.TotalDelayedMessagesRead(&_SequencerInboxStub.CallOpts)
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x6633ae85.
//
// Solidity: function addInitMessage(uint256 chainId) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddInitMessage(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addInitMessage", chainId)
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x6633ae85.
//
// Solidity: function addInitMessage(uint256 chainId) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddInitMessage(chainId *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddInitMessage(&_SequencerInboxStub.TransactOpts, chainId)
}

// AddInitMessage is a paid mutator transaction binding the contract method 0x6633ae85.
//
// Solidity: function addInitMessage(uint256 chainId) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddInitMessage(chainId *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddInitMessage(&_SequencerInboxStub.TransactOpts, chainId)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2Batch(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2Batch", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2Batch(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2Batch is a paid mutator transaction binding the contract method 0xe0bc9729.
//
// Solidity: function addSequencerL2Batch(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2Batch(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2Batch(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2BatchFromOrigin(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2BatchFromOrigin", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchFromOrigin(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin is a paid mutator transaction binding the contract method 0x6f12b0c9.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2BatchFromOrigin(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder)
}

// AddSequencerL2BatchFromOrigin0 is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) AddSequencerL2BatchFromOrigin0(opts *bind.TransactOpts, sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "addSequencerL2BatchFromOrigin0", sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOrigin0 is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) AddSequencerL2BatchFromOrigin0(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin0(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// AddSequencerL2BatchFromOrigin0 is a paid mutator transaction binding the contract method 0x8f111f3c.
//
// Solidity: function addSequencerL2BatchFromOrigin(uint256 sequenceNumber, bytes data, uint256 afterDelayedMessagesRead, address gasRefunder, uint256 prevMessageCount, uint256 newMessageCount) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) AddSequencerL2BatchFromOrigin0(sequenceNumber *big.Int, data []byte, afterDelayedMessagesRead *big.Int, gasRefunder common.Address, prevMessageCount *big.Int, newMessageCount *big.Int) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.AddSequencerL2BatchFromOrigin0(&_SequencerInboxStub.TransactOpts, sequenceNumber, data, afterDelayedMessagesRead, gasRefunder, prevMessageCount, newMessageCount)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) ForceInclusion(opts *bind.TransactOpts, _totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "forceInclusion", _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.ForceInclusion(&_SequencerInboxStub.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// ForceInclusion is a paid mutator transaction binding the contract method 0xf1981578.
//
// Solidity: function forceInclusion(uint256 _totalDelayedMessagesRead, uint8 kind, uint64[2] l1BlockAndTime, uint256 baseFeeL1, address sender, bytes32 messageDataHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) ForceInclusion(_totalDelayedMessagesRead *big.Int, kind uint8, l1BlockAndTime [2]uint64, baseFeeL1 *big.Int, sender common.Address, messageDataHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.ForceInclusion(&_SequencerInboxStub.TransactOpts, _totalDelayedMessagesRead, kind, l1BlockAndTime, baseFeeL1, sender, messageDataHash)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) Initialize(opts *bind.TransactOpts, bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "initialize", bridge_, maxTimeVariation_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.Initialize(&_SequencerInboxStub.TransactOpts, bridge_, maxTimeVariation_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1f7a92b2.
//
// Solidity: function initialize(address bridge_, (uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) Initialize(bridge_ common.Address, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.Initialize(&_SequencerInboxStub.TransactOpts, bridge_, maxTimeVariation_)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) InvalidateKeysetHash(opts *bind.TransactOpts, ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "invalidateKeysetHash", ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.InvalidateKeysetHash(&_SequencerInboxStub.TransactOpts, ksHash)
}

// InvalidateKeysetHash is a paid mutator transaction binding the contract method 0x84420860.
//
// Solidity: function invalidateKeysetHash(bytes32 ksHash) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) InvalidateKeysetHash(ksHash [32]byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.InvalidateKeysetHash(&_SequencerInboxStub.TransactOpts, ksHash)
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) RemoveDelayAfterFork(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "removeDelayAfterFork")
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxStub *SequencerInboxStubSession) RemoveDelayAfterFork() (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.RemoveDelayAfterFork(&_SequencerInboxStub.TransactOpts)
}

// RemoveDelayAfterFork is a paid mutator transaction binding the contract method 0x96cc5c78.
//
// Solidity: function removeDelayAfterFork() returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) RemoveDelayAfterFork() (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.RemoveDelayAfterFork(&_SequencerInboxStub.TransactOpts)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetIsBatchPoster(opts *bind.TransactOpts, addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setIsBatchPoster", addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsBatchPoster(&_SequencerInboxStub.TransactOpts, addr, isBatchPoster_)
}

// SetIsBatchPoster is a paid mutator transaction binding the contract method 0x6e7df3e7.
//
// Solidity: function setIsBatchPoster(address addr, bool isBatchPoster_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetIsBatchPoster(addr common.Address, isBatchPoster_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsBatchPoster(&_SequencerInboxStub.TransactOpts, addr, isBatchPoster_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetIsSequencer(opts *bind.TransactOpts, addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setIsSequencer", addr, isSequencer_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetIsSequencer(addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsSequencer(&_SequencerInboxStub.TransactOpts, addr, isSequencer_)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address addr, bool isSequencer_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetIsSequencer(addr common.Address, isSequencer_ bool) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetIsSequencer(&_SequencerInboxStub.TransactOpts, addr, isSequencer_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetMaxTimeVariation(opts *bind.TransactOpts, maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setMaxTimeVariation", maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetMaxTimeVariation(&_SequencerInboxStub.TransactOpts, maxTimeVariation_)
}

// SetMaxTimeVariation is a paid mutator transaction binding the contract method 0xb31761f8.
//
// Solidity: function setMaxTimeVariation((uint256,uint256,uint256,uint256) maxTimeVariation_) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetMaxTimeVariation(maxTimeVariation_ ISequencerInboxMaxTimeVariation) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetMaxTimeVariation(&_SequencerInboxStub.TransactOpts, maxTimeVariation_)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactor) SetValidKeyset(opts *bind.TransactOpts, keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.contract.Transact(opts, "setValidKeyset", keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetValidKeyset(&_SequencerInboxStub.TransactOpts, keysetBytes)
}

// SetValidKeyset is a paid mutator transaction binding the contract method 0xd1ce8da8.
//
// Solidity: function setValidKeyset(bytes keysetBytes) returns()
func (_SequencerInboxStub *SequencerInboxStubTransactorSession) SetValidKeyset(keysetBytes []byte) (*types.Transaction, error) {
	return _SequencerInboxStub.Contract.SetValidKeyset(&_SequencerInboxStub.TransactOpts, keysetBytes)
}

// SequencerInboxStubInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredIterator struct {
	Event *SequencerInboxStubInboxMessageDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInboxMessageDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerInboxStubInboxMessageDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInboxMessageDelivered represents a InboxMessageDelivered event raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxStubInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInboxMessageDeliveredIterator{contract: _SequencerInboxStub.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInboxMessageDelivered)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInboxMessageDelivered(log types.Log) (*SequencerInboxStubInboxMessageDelivered, error) {
	event := new(SequencerInboxStubInboxMessageDelivered)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredFromOriginIterator struct {
	Event *SequencerInboxStubInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the SequencerInboxStub contract.
type SequencerInboxStubInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*SequencerInboxStubInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInboxMessageDeliveredFromOriginIterator{contract: _SequencerInboxStub.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*SequencerInboxStubInboxMessageDeliveredFromOrigin, error) {
	event := new(SequencerInboxStubInboxMessageDeliveredFromOrigin)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubInvalidateKeysetIterator is returned from FilterInvalidateKeyset and is used to iterate over the raw logs and unpacked data for InvalidateKeyset events raised by the SequencerInboxStub contract.
type SequencerInboxStubInvalidateKeysetIterator struct {
	Event *SequencerInboxStubInvalidateKeyset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerInboxStubInvalidateKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubInvalidateKeyset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerInboxStubInvalidateKeyset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerInboxStubInvalidateKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubInvalidateKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubInvalidateKeyset represents a InvalidateKeyset event raised by the SequencerInboxStub contract.
type SequencerInboxStubInvalidateKeyset struct {
	KeysetHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvalidateKeyset is a free log retrieval operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterInvalidateKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxStubInvalidateKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubInvalidateKeysetIterator{contract: _SequencerInboxStub.contract, event: "InvalidateKeyset", logs: logs, sub: sub}, nil
}

// WatchInvalidateKeyset is a free log subscription operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchInvalidateKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubInvalidateKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "InvalidateKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubInvalidateKeyset)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidateKeyset is a log parse operation binding the contract event 0x5cb4218b272fd214168ac43e90fb4d05d6c36f0b17ffb4c2dd07c234d744eb2a.
//
// Solidity: event InvalidateKeyset(bytes32 indexed keysetHash)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseInvalidateKeyset(log types.Log) (*SequencerInboxStubInvalidateKeyset, error) {
	event := new(SequencerInboxStubInvalidateKeyset)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "InvalidateKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the SequencerInboxStub contract.
type SequencerInboxStubOwnerFunctionCalledIterator struct {
	Event *SequencerInboxStubOwnerFunctionCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubOwnerFunctionCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerInboxStubOwnerFunctionCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the SequencerInboxStub contract.
type SequencerInboxStubOwnerFunctionCalled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts, id []*big.Int) (*SequencerInboxStubOwnerFunctionCalledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubOwnerFunctionCalledIterator{contract: _SequencerInboxStub.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubOwnerFunctionCalled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubOwnerFunctionCalled)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerFunctionCalled is a log parse operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseOwnerFunctionCalled(log types.Log) (*SequencerInboxStubOwnerFunctionCalled, error) {
	event := new(SequencerInboxStubOwnerFunctionCalled)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSequencerBatchDataIterator is returned from FilterSequencerBatchData and is used to iterate over the raw logs and unpacked data for SequencerBatchData events raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDataIterator struct {
	Event *SequencerInboxStubSequencerBatchData // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerInboxStubSequencerBatchDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSequencerBatchData)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerInboxStubSequencerBatchData)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerInboxStubSequencerBatchDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSequencerBatchDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSequencerBatchData represents a SequencerBatchData event raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchData struct {
	BatchSequenceNumber *big.Int
	Data                []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchData is a free log retrieval operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSequencerBatchData(opts *bind.FilterOpts, batchSequenceNumber []*big.Int) (*SequencerInboxStubSequencerBatchDataIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSequencerBatchDataIterator{contract: _SequencerInboxStub.contract, event: "SequencerBatchData", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchData is a free log subscription operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSequencerBatchData(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSequencerBatchData, batchSequenceNumber []*big.Int) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SequencerBatchData", batchSequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSequencerBatchData)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerBatchData is a log parse operation binding the contract event 0xfe325ca1efe4c5c1062c981c3ee74b781debe4ea9440306a96d2a55759c66c20.
//
// Solidity: event SequencerBatchData(uint256 indexed batchSequenceNumber, bytes data)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSequencerBatchData(log types.Log) (*SequencerInboxStubSequencerBatchData, error) {
	event := new(SequencerInboxStubSequencerBatchData)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSequencerBatchDeliveredIterator is returned from FilterSequencerBatchDelivered and is used to iterate over the raw logs and unpacked data for SequencerBatchDelivered events raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDeliveredIterator struct {
	Event *SequencerInboxStubSequencerBatchDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSequencerBatchDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerInboxStubSequencerBatchDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSequencerBatchDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSequencerBatchDelivered represents a SequencerBatchDelivered event raised by the SequencerInboxStub contract.
type SequencerInboxStubSequencerBatchDelivered struct {
	BatchSequenceNumber      *big.Int
	BeforeAcc                [32]byte
	AfterAcc                 [32]byte
	DelayedAcc               [32]byte
	AfterDelayedMessagesRead *big.Int
	TimeBounds               ISequencerInboxTimeBounds
	DataLocation             uint8
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterSequencerBatchDelivered is a free log retrieval operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSequencerBatchDelivered(opts *bind.FilterOpts, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (*SequencerInboxStubSequencerBatchDeliveredIterator, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}
	var afterAccRule []interface{}
	for _, afterAccItem := range afterAcc {
		afterAccRule = append(afterAccRule, afterAccItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSequencerBatchDeliveredIterator{contract: _SequencerInboxStub.contract, event: "SequencerBatchDelivered", logs: logs, sub: sub}, nil
}

// WatchSequencerBatchDelivered is a free log subscription operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSequencerBatchDelivered(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSequencerBatchDelivered, batchSequenceNumber []*big.Int, beforeAcc [][32]byte, afterAcc [][32]byte) (event.Subscription, error) {

	var batchSequenceNumberRule []interface{}
	for _, batchSequenceNumberItem := range batchSequenceNumber {
		batchSequenceNumberRule = append(batchSequenceNumberRule, batchSequenceNumberItem)
	}
	var beforeAccRule []interface{}
	for _, beforeAccItem := range beforeAcc {
		beforeAccRule = append(beforeAccRule, beforeAccItem)
	}
	var afterAccRule []interface{}
	for _, afterAccItem := range afterAcc {
		afterAccRule = append(afterAccRule, afterAccItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SequencerBatchDelivered", batchSequenceNumberRule, beforeAccRule, afterAccRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSequencerBatchDelivered)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencerBatchDelivered is a log parse operation binding the contract event 0x7394f4a19a13c7b92b5bb71033245305946ef78452f7b4986ac1390b5df4ebd7.
//
// Solidity: event SequencerBatchDelivered(uint256 indexed batchSequenceNumber, bytes32 indexed beforeAcc, bytes32 indexed afterAcc, bytes32 delayedAcc, uint256 afterDelayedMessagesRead, (uint64,uint64,uint64,uint64) timeBounds, uint8 dataLocation)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSequencerBatchDelivered(log types.Log) (*SequencerInboxStubSequencerBatchDelivered, error) {
	event := new(SequencerInboxStubSequencerBatchDelivered)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SequencerBatchDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SequencerInboxStubSetValidKeysetIterator is returned from FilterSetValidKeyset and is used to iterate over the raw logs and unpacked data for SetValidKeyset events raised by the SequencerInboxStub contract.
type SequencerInboxStubSetValidKeysetIterator struct {
	Event *SequencerInboxStubSetValidKeyset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SequencerInboxStubSetValidKeysetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SequencerInboxStubSetValidKeyset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SequencerInboxStubSetValidKeyset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SequencerInboxStubSetValidKeysetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SequencerInboxStubSetValidKeysetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SequencerInboxStubSetValidKeyset represents a SetValidKeyset event raised by the SequencerInboxStub contract.
type SequencerInboxStubSetValidKeyset struct {
	KeysetHash  [32]byte
	KeysetBytes []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetValidKeyset is a free log retrieval operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) FilterSetValidKeyset(opts *bind.FilterOpts, keysetHash [][32]byte) (*SequencerInboxStubSetValidKeysetIterator, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.FilterLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return &SequencerInboxStubSetValidKeysetIterator{contract: _SequencerInboxStub.contract, event: "SetValidKeyset", logs: logs, sub: sub}, nil
}

// WatchSetValidKeyset is a free log subscription operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) WatchSetValidKeyset(opts *bind.WatchOpts, sink chan<- *SequencerInboxStubSetValidKeyset, keysetHash [][32]byte) (event.Subscription, error) {

	var keysetHashRule []interface{}
	for _, keysetHashItem := range keysetHash {
		keysetHashRule = append(keysetHashRule, keysetHashItem)
	}

	logs, sub, err := _SequencerInboxStub.contract.WatchLogs(opts, "SetValidKeyset", keysetHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SequencerInboxStubSetValidKeyset)
				if err := _SequencerInboxStub.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetValidKeyset is a log parse operation binding the contract event 0xabca9b7986bc22ad0160eb0cb88ae75411eacfba4052af0b457a9335ef655722.
//
// Solidity: event SetValidKeyset(bytes32 indexed keysetHash, bytes keysetBytes)
func (_SequencerInboxStub *SequencerInboxStubFilterer) ParseSetValidKeyset(log types.Log) (*SequencerInboxStubSetValidKeyset, error) {
	event := new(SequencerInboxStubSetValidKeyset)
	if err := _SequencerInboxStub.contract.UnpackLog(event, "SetValidKeyset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleMetaData contains all meta data concerning the Simple contract.
var SimpleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"name\":\"CounterEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"}],\"name\":\"LogAndIncrementCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NullEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"}],\"name\":\"RedeemedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"checkBlockHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"useTopLevel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"directCase\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"staticCase\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"delegateCase\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"callcodeCase\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"callCase\",\"type\":\"bool\"}],\"name\":\"checkCalls\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"checkGasUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"useTopLevel\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"expected\",\"type\":\"bool\"}],\"name\":\"checkIsTopLevelOrWasAliased\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitNullEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementEmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementRedeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"logAndIncrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noop\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pleaseRevert\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d62806100206000396000f3fe608060405234801561001057600080fd5b50600436106100bf5760003560e01c80635dfc2e4a1161007c5780635dfc2e4a146100cc57806361bc221a146101195780638a390877146101445780639ff5ccac14610157578063b226a9641461015f578063d09de08a14610167578063ded5ecad1461016f57600080fd5b806305795f73146100c45780630e8c389f146100ce57806312e05dd1146100d65780631a2f8a92146100eb57806344c25fba146100fe5780635677c11e14610111575b600080fd5b6100cc610182565b005b6100cc6101c4565b445b6040519081526020015b60405180910390f35b6100d86100f9366004610abb565b6103aa565b6100cc61010c366004610b4d565b61042e565b6100d861080d565b60005461012c906001600160401b031681565b6040516001600160401b0390911681526020016100e2565b6100cc610152366004610bcf565b61086c565b6100cc6108f1565b6100cc61095e565b6100cc610989565b6100cc61017d366004610be8565b6109c8565b60405162461bcd60e51b8152602060048201526012602482015271534f4c49444954595f524556455254494e4760701b60448201526064015b60405180910390fd5b3332146102075760405162461bcd60e51b815260206004820152601160248201527029a2a72222a92fa727aa2fa7a924a3a4a760791b60448201526064016101bb565b60646001600160a01b031663175a260b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561024157600080fd5b505afa158015610255573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102799190610c21565b6102b35760405162461bcd60e51b815260206004820152600b60248201526a1393d517d053125054d15160aa1b60448201526064016101bb565b600080546001600160401b031690806102cb83610c5b565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550507f773c78bf96e65f61c1a2622b47d76e78bfe70dd59cf4f11470c4c121c315941333606e6001600160a01b031663de4ba2b36040518163ffffffff1660e01b815260040160206040518083038186803b15801561034c57600080fd5b505afa158015610360573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103849190610c82565b604080516001600160a01b039384168152929091166020830152015b60405180910390a1565b6000805a90506001600160a01b0385166103c661271083610c9f565b85856040516103d6929190610cb6565b6000604051808303818686fa925050503d8060008114610412576040519150601f19603f3d011682016040523d82523d6000602084013e610417565b606091505b5050505a6104259082610c9f565b95945050505050565b85156104cd5784151560646001600160a01b03166308bd624c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561047157600080fd5b505afa158015610485573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a99190610c21565b1515146104c85760405162461bcd60e51b81526004016101bb90610cc6565b610561565b84151560646001600160a01b031663175a260b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561050a57600080fd5b505afa15801561051e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105429190610c21565b1515146105615760405162461bcd60e51b81526004016101bb90610cc6565b60405163ded5ecad60e01b815286151560048201528415156024820152309063ded5ecad9060440160006040518083038186803b1580156105a157600080fd5b505afa1580156105b5573d6000803e3d6000fd5b505060408051891515602482015286151560448083019190915282518083039091018152606490910182526020810180516001600160e01b031663ded5ecad60e01b1790529051909250600091503090610610908490610cf1565b600060405180830381855af49150503d806000811461064b576040519150601f19603f3d011682016040523d82523d6000602084013e610650565b606091505b50509050806106985760405162461bcd60e51b81526020600482015260146024820152731111531151d0551157d0d0531317d1905253115160621b60448201526064016101bb565b6040805189151560248201528515156044808301919091528251808303909101815260649091019091526020810180516001600160e01b031663ded5ecad60e01b1781528151919350600091829182305af290508061072b5760405162461bcd60e51b815260206004820152600f60248201526e10d0531310d3d11157d19052531151608a1b60448201526064016101bb565b60408051891515602482015284151560448083019190915282518083039091018152606490910182526020810180516001600160e01b031663ded5ecad60e01b17905290519092503090610780908490610cf1565b6000604051808303816000865af19150503d80600081146107bd576040519150601f19603f3d011682016040523d82523d6000602084013e6107c2565b606091505b505080915050806108035760405162461bcd60e51b815260206004820152600b60248201526a10d0531317d1905253115160aa1b60448201526064016101bb565b5050505050505050565b600061081a600243610c9f565b40610826600143610c9f565b4014156108675760405162461bcd60e51b815260206004820152600f60248201526e0a6829a8abe84989e8696be9082a69608b1b60448201526064016101bb565b504390565b600054604080518381526001600160401b0390921660208301527f8df8e492f407b078593c5d8fd7e65ef68505999d911d5b99b017c0b7077398b9910160405180910390a1600080546001600160401b031690806108c983610c5b565b91906101000a8154816001600160401b0302191690836001600160401b031602179055505050565b600080546001600160401b0316908061090983610c5b565b82546101009290920a6001600160401b03818102199093169183160217909155600054604051911681527fa45d7e79cb3c6044f30c8dd891e6571301d6b8b6618df519c987905ec70742e791506020016103a0565b6040517f6f59c82101949290205a9ae9d0c657e6dae1a71c301ae76d385c2792294585fe90600090a1565b600080546001600160401b031690806109a183610c5b565b91906101000a8154816001600160401b0302191690836001600160401b0316021790555050565b8115610a665780151560646001600160a01b03166308bd624c6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a0b57600080fd5b505afa158015610a1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a439190610c21565b151514610a625760405162461bcd60e51b81526004016101bb90610cc6565b5050565b80151560646001600160a01b031663175a260b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a0b57600080fd5b6001600160a01b0381168114610ab857600080fd5b50565b600080600060408486031215610ad057600080fd5b8335610adb81610aa3565b925060208401356001600160401b0380821115610af757600080fd5b818601915086601f830112610b0b57600080fd5b813581811115610b1a57600080fd5b876020828501011115610b2c57600080fd5b6020830194508093505050509250925092565b8015158114610ab857600080fd5b60008060008060008060c08789031215610b6657600080fd5b8635610b7181610b3f565b95506020870135610b8181610b3f565b94506040870135610b9181610b3f565b93506060870135610ba181610b3f565b92506080870135610bb181610b3f565b915060a0870135610bc181610b3f565b809150509295509295509295565b600060208284031215610be157600080fd5b5035919050565b60008060408385031215610bfb57600080fd5b8235610c0681610b3f565b91506020830135610c1681610b3f565b809150509250929050565b600060208284031215610c3357600080fd5b8151610c3e81610b3f565b9392505050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380831681811415610c7857610c78610c45565b6001019392505050565b600060208284031215610c9457600080fd5b8151610c3e81610aa3565b600082821015610cb157610cb1610c45565b500390565b8183823760009101908152919050565b60208082526011908201527015539156141150d5115117d49154d55315607a1b604082015260600190565b6000825160005b81811015610d125760208186018101518583015201610cf8565b81811115610d21576000828501525b50919091019291505056fea26469706673582212208ffe181c1e22a5afaff2fe4fbc8dcdf836946229a9bb5206277ff8d4499fdce164736f6c63430008090033",
}

// SimpleABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleMetaData.ABI instead.
var SimpleABI = SimpleMetaData.ABI

// SimpleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleMetaData.Bin instead.
var SimpleBin = SimpleMetaData.Bin

// DeploySimple deploys a new Ethereum contract, binding an instance of Simple to it.
func DeploySimple(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simple, error) {
	parsed, err := SimpleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// Simple is an auto generated Go binding around an Ethereum contract.
type Simple struct {
	SimpleCaller     // Read-only binding to the contract
	SimpleTransactor // Write-only binding to the contract
	SimpleFilterer   // Log filterer for contract events
}

// SimpleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleSession struct {
	Contract     *Simple           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleCallerSession struct {
	Contract *SimpleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SimpleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleTransactorSession struct {
	Contract     *SimpleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRaw struct {
	Contract *Simple // Generic contract binding to access the raw methods on
}

// SimpleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleCallerRaw struct {
	Contract *SimpleCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleTransactorRaw struct {
	Contract *SimpleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimple creates a new instance of Simple, bound to a specific deployed contract.
func NewSimple(address common.Address, backend bind.ContractBackend) (*Simple, error) {
	contract, err := bindSimple(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simple{SimpleCaller: SimpleCaller{contract: contract}, SimpleTransactor: SimpleTransactor{contract: contract}, SimpleFilterer: SimpleFilterer{contract: contract}}, nil
}

// NewSimpleCaller creates a new read-only instance of Simple, bound to a specific deployed contract.
func NewSimpleCaller(address common.Address, caller bind.ContractCaller) (*SimpleCaller, error) {
	contract, err := bindSimple(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleCaller{contract: contract}, nil
}

// NewSimpleTransactor creates a new write-only instance of Simple, bound to a specific deployed contract.
func NewSimpleTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleTransactor, error) {
	contract, err := bindSimple(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleTransactor{contract: contract}, nil
}

// NewSimpleFilterer creates a new log filterer instance of Simple, bound to a specific deployed contract.
func NewSimpleFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleFilterer, error) {
	contract, err := bindSimple(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleFilterer{contract: contract}, nil
}

// bindSimple binds a generic wrapper to an already deployed contract.
func bindSimple(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.SimpleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.SimpleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simple *SimpleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simple.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simple *SimpleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simple *SimpleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simple.Contract.contract.Transact(opts, method, params...)
}

// CheckBlockHashes is a free data retrieval call binding the contract method 0x5677c11e.
//
// Solidity: function checkBlockHashes() view returns(uint256)
func (_Simple *SimpleCaller) CheckBlockHashes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "checkBlockHashes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckBlockHashes is a free data retrieval call binding the contract method 0x5677c11e.
//
// Solidity: function checkBlockHashes() view returns(uint256)
func (_Simple *SimpleSession) CheckBlockHashes() (*big.Int, error) {
	return _Simple.Contract.CheckBlockHashes(&_Simple.CallOpts)
}

// CheckBlockHashes is a free data retrieval call binding the contract method 0x5677c11e.
//
// Solidity: function checkBlockHashes() view returns(uint256)
func (_Simple *SimpleCallerSession) CheckBlockHashes() (*big.Int, error) {
	return _Simple.Contract.CheckBlockHashes(&_Simple.CallOpts)
}

// CheckGasUsed is a free data retrieval call binding the contract method 0x1a2f8a92.
//
// Solidity: function checkGasUsed(address to, bytes input) view returns(uint256)
func (_Simple *SimpleCaller) CheckGasUsed(opts *bind.CallOpts, to common.Address, input []byte) (*big.Int, error) {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "checkGasUsed", to, input)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckGasUsed is a free data retrieval call binding the contract method 0x1a2f8a92.
//
// Solidity: function checkGasUsed(address to, bytes input) view returns(uint256)
func (_Simple *SimpleSession) CheckGasUsed(to common.Address, input []byte) (*big.Int, error) {
	return _Simple.Contract.CheckGasUsed(&_Simple.CallOpts, to, input)
}

// CheckGasUsed is a free data retrieval call binding the contract method 0x1a2f8a92.
//
// Solidity: function checkGasUsed(address to, bytes input) view returns(uint256)
func (_Simple *SimpleCallerSession) CheckGasUsed(to common.Address, input []byte) (*big.Int, error) {
	return _Simple.Contract.CheckGasUsed(&_Simple.CallOpts, to, input)
}

// CheckIsTopLevelOrWasAliased is a free data retrieval call binding the contract method 0xded5ecad.
//
// Solidity: function checkIsTopLevelOrWasAliased(bool useTopLevel, bool expected) view returns()
func (_Simple *SimpleCaller) CheckIsTopLevelOrWasAliased(opts *bind.CallOpts, useTopLevel bool, expected bool) error {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "checkIsTopLevelOrWasAliased", useTopLevel, expected)

	if err != nil {
		return err
	}

	return err

}

// CheckIsTopLevelOrWasAliased is a free data retrieval call binding the contract method 0xded5ecad.
//
// Solidity: function checkIsTopLevelOrWasAliased(bool useTopLevel, bool expected) view returns()
func (_Simple *SimpleSession) CheckIsTopLevelOrWasAliased(useTopLevel bool, expected bool) error {
	return _Simple.Contract.CheckIsTopLevelOrWasAliased(&_Simple.CallOpts, useTopLevel, expected)
}

// CheckIsTopLevelOrWasAliased is a free data retrieval call binding the contract method 0xded5ecad.
//
// Solidity: function checkIsTopLevelOrWasAliased(bool useTopLevel, bool expected) view returns()
func (_Simple *SimpleCallerSession) CheckIsTopLevelOrWasAliased(useTopLevel bool, expected bool) error {
	return _Simple.Contract.CheckIsTopLevelOrWasAliased(&_Simple.CallOpts, useTopLevel, expected)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint64)
func (_Simple *SimpleCaller) Counter(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint64)
func (_Simple *SimpleSession) Counter() (uint64, error) {
	return _Simple.Contract.Counter(&_Simple.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint64)
func (_Simple *SimpleCallerSession) Counter() (uint64, error) {
	return _Simple.Contract.Counter(&_Simple.CallOpts)
}

// GetBlockDifficulty is a free data retrieval call binding the contract method 0x12e05dd1.
//
// Solidity: function getBlockDifficulty() view returns(uint256)
func (_Simple *SimpleCaller) GetBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "getBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockDifficulty is a free data retrieval call binding the contract method 0x12e05dd1.
//
// Solidity: function getBlockDifficulty() view returns(uint256)
func (_Simple *SimpleSession) GetBlockDifficulty() (*big.Int, error) {
	return _Simple.Contract.GetBlockDifficulty(&_Simple.CallOpts)
}

// GetBlockDifficulty is a free data retrieval call binding the contract method 0x12e05dd1.
//
// Solidity: function getBlockDifficulty() view returns(uint256)
func (_Simple *SimpleCallerSession) GetBlockDifficulty() (*big.Int, error) {
	return _Simple.Contract.GetBlockDifficulty(&_Simple.CallOpts)
}

// Noop is a free data retrieval call binding the contract method 0x5dfc2e4a.
//
// Solidity: function noop() pure returns()
func (_Simple *SimpleCaller) Noop(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "noop")

	if err != nil {
		return err
	}

	return err

}

// Noop is a free data retrieval call binding the contract method 0x5dfc2e4a.
//
// Solidity: function noop() pure returns()
func (_Simple *SimpleSession) Noop() error {
	return _Simple.Contract.Noop(&_Simple.CallOpts)
}

// Noop is a free data retrieval call binding the contract method 0x5dfc2e4a.
//
// Solidity: function noop() pure returns()
func (_Simple *SimpleCallerSession) Noop() error {
	return _Simple.Contract.Noop(&_Simple.CallOpts)
}

// PleaseRevert is a free data retrieval call binding the contract method 0x05795f73.
//
// Solidity: function pleaseRevert() pure returns()
func (_Simple *SimpleCaller) PleaseRevert(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Simple.contract.Call(opts, &out, "pleaseRevert")

	if err != nil {
		return err
	}

	return err

}

// PleaseRevert is a free data retrieval call binding the contract method 0x05795f73.
//
// Solidity: function pleaseRevert() pure returns()
func (_Simple *SimpleSession) PleaseRevert() error {
	return _Simple.Contract.PleaseRevert(&_Simple.CallOpts)
}

// PleaseRevert is a free data retrieval call binding the contract method 0x05795f73.
//
// Solidity: function pleaseRevert() pure returns()
func (_Simple *SimpleCallerSession) PleaseRevert() error {
	return _Simple.Contract.PleaseRevert(&_Simple.CallOpts)
}

// CheckCalls is a paid mutator transaction binding the contract method 0x44c25fba.
//
// Solidity: function checkCalls(bool useTopLevel, bool directCase, bool staticCase, bool delegateCase, bool callcodeCase, bool callCase) returns()
func (_Simple *SimpleTransactor) CheckCalls(opts *bind.TransactOpts, useTopLevel bool, directCase bool, staticCase bool, delegateCase bool, callcodeCase bool, callCase bool) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "checkCalls", useTopLevel, directCase, staticCase, delegateCase, callcodeCase, callCase)
}

// CheckCalls is a paid mutator transaction binding the contract method 0x44c25fba.
//
// Solidity: function checkCalls(bool useTopLevel, bool directCase, bool staticCase, bool delegateCase, bool callcodeCase, bool callCase) returns()
func (_Simple *SimpleSession) CheckCalls(useTopLevel bool, directCase bool, staticCase bool, delegateCase bool, callcodeCase bool, callCase bool) (*types.Transaction, error) {
	return _Simple.Contract.CheckCalls(&_Simple.TransactOpts, useTopLevel, directCase, staticCase, delegateCase, callcodeCase, callCase)
}

// CheckCalls is a paid mutator transaction binding the contract method 0x44c25fba.
//
// Solidity: function checkCalls(bool useTopLevel, bool directCase, bool staticCase, bool delegateCase, bool callcodeCase, bool callCase) returns()
func (_Simple *SimpleTransactorSession) CheckCalls(useTopLevel bool, directCase bool, staticCase bool, delegateCase bool, callcodeCase bool, callCase bool) (*types.Transaction, error) {
	return _Simple.Contract.CheckCalls(&_Simple.TransactOpts, useTopLevel, directCase, staticCase, delegateCase, callcodeCase, callCase)
}

// EmitNullEvent is a paid mutator transaction binding the contract method 0xb226a964.
//
// Solidity: function emitNullEvent() returns()
func (_Simple *SimpleTransactor) EmitNullEvent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "emitNullEvent")
}

// EmitNullEvent is a paid mutator transaction binding the contract method 0xb226a964.
//
// Solidity: function emitNullEvent() returns()
func (_Simple *SimpleSession) EmitNullEvent() (*types.Transaction, error) {
	return _Simple.Contract.EmitNullEvent(&_Simple.TransactOpts)
}

// EmitNullEvent is a paid mutator transaction binding the contract method 0xb226a964.
//
// Solidity: function emitNullEvent() returns()
func (_Simple *SimpleTransactorSession) EmitNullEvent() (*types.Transaction, error) {
	return _Simple.Contract.EmitNullEvent(&_Simple.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Simple *SimpleTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Simple *SimpleSession) Increment() (*types.Transaction, error) {
	return _Simple.Contract.Increment(&_Simple.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Simple *SimpleTransactorSession) Increment() (*types.Transaction, error) {
	return _Simple.Contract.Increment(&_Simple.TransactOpts)
}

// IncrementEmit is a paid mutator transaction binding the contract method 0x9ff5ccac.
//
// Solidity: function incrementEmit() returns()
func (_Simple *SimpleTransactor) IncrementEmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "incrementEmit")
}

// IncrementEmit is a paid mutator transaction binding the contract method 0x9ff5ccac.
//
// Solidity: function incrementEmit() returns()
func (_Simple *SimpleSession) IncrementEmit() (*types.Transaction, error) {
	return _Simple.Contract.IncrementEmit(&_Simple.TransactOpts)
}

// IncrementEmit is a paid mutator transaction binding the contract method 0x9ff5ccac.
//
// Solidity: function incrementEmit() returns()
func (_Simple *SimpleTransactorSession) IncrementEmit() (*types.Transaction, error) {
	return _Simple.Contract.IncrementEmit(&_Simple.TransactOpts)
}

// IncrementRedeem is a paid mutator transaction binding the contract method 0x0e8c389f.
//
// Solidity: function incrementRedeem() returns()
func (_Simple *SimpleTransactor) IncrementRedeem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "incrementRedeem")
}

// IncrementRedeem is a paid mutator transaction binding the contract method 0x0e8c389f.
//
// Solidity: function incrementRedeem() returns()
func (_Simple *SimpleSession) IncrementRedeem() (*types.Transaction, error) {
	return _Simple.Contract.IncrementRedeem(&_Simple.TransactOpts)
}

// IncrementRedeem is a paid mutator transaction binding the contract method 0x0e8c389f.
//
// Solidity: function incrementRedeem() returns()
func (_Simple *SimpleTransactorSession) IncrementRedeem() (*types.Transaction, error) {
	return _Simple.Contract.IncrementRedeem(&_Simple.TransactOpts)
}

// LogAndIncrement is a paid mutator transaction binding the contract method 0x8a390877.
//
// Solidity: function logAndIncrement(uint256 expected) returns()
func (_Simple *SimpleTransactor) LogAndIncrement(opts *bind.TransactOpts, expected *big.Int) (*types.Transaction, error) {
	return _Simple.contract.Transact(opts, "logAndIncrement", expected)
}

// LogAndIncrement is a paid mutator transaction binding the contract method 0x8a390877.
//
// Solidity: function logAndIncrement(uint256 expected) returns()
func (_Simple *SimpleSession) LogAndIncrement(expected *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.LogAndIncrement(&_Simple.TransactOpts, expected)
}

// LogAndIncrement is a paid mutator transaction binding the contract method 0x8a390877.
//
// Solidity: function logAndIncrement(uint256 expected) returns()
func (_Simple *SimpleTransactorSession) LogAndIncrement(expected *big.Int) (*types.Transaction, error) {
	return _Simple.Contract.LogAndIncrement(&_Simple.TransactOpts, expected)
}

// SimpleCounterEventIterator is returned from FilterCounterEvent and is used to iterate over the raw logs and unpacked data for CounterEvent events raised by the Simple contract.
type SimpleCounterEventIterator struct {
	Event *SimpleCounterEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleCounterEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleCounterEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SimpleCounterEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SimpleCounterEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleCounterEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleCounterEvent represents a CounterEvent event raised by the Simple contract.
type SimpleCounterEvent struct {
	Count uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCounterEvent is a free log retrieval operation binding the contract event 0xa45d7e79cb3c6044f30c8dd891e6571301d6b8b6618df519c987905ec70742e7.
//
// Solidity: event CounterEvent(uint64 count)
func (_Simple *SimpleFilterer) FilterCounterEvent(opts *bind.FilterOpts) (*SimpleCounterEventIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "CounterEvent")
	if err != nil {
		return nil, err
	}
	return &SimpleCounterEventIterator{contract: _Simple.contract, event: "CounterEvent", logs: logs, sub: sub}, nil
}

// WatchCounterEvent is a free log subscription operation binding the contract event 0xa45d7e79cb3c6044f30c8dd891e6571301d6b8b6618df519c987905ec70742e7.
//
// Solidity: event CounterEvent(uint64 count)
func (_Simple *SimpleFilterer) WatchCounterEvent(opts *bind.WatchOpts, sink chan<- *SimpleCounterEvent) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "CounterEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleCounterEvent)
				if err := _Simple.contract.UnpackLog(event, "CounterEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCounterEvent is a log parse operation binding the contract event 0xa45d7e79cb3c6044f30c8dd891e6571301d6b8b6618df519c987905ec70742e7.
//
// Solidity: event CounterEvent(uint64 count)
func (_Simple *SimpleFilterer) ParseCounterEvent(log types.Log) (*SimpleCounterEvent, error) {
	event := new(SimpleCounterEvent)
	if err := _Simple.contract.UnpackLog(event, "CounterEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleLogAndIncrementCalledIterator is returned from FilterLogAndIncrementCalled and is used to iterate over the raw logs and unpacked data for LogAndIncrementCalled events raised by the Simple contract.
type SimpleLogAndIncrementCalledIterator struct {
	Event *SimpleLogAndIncrementCalled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleLogAndIncrementCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleLogAndIncrementCalled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SimpleLogAndIncrementCalled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SimpleLogAndIncrementCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleLogAndIncrementCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleLogAndIncrementCalled represents a LogAndIncrementCalled event raised by the Simple contract.
type SimpleLogAndIncrementCalled struct {
	Expected *big.Int
	Have     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogAndIncrementCalled is a free log retrieval operation binding the contract event 0x8df8e492f407b078593c5d8fd7e65ef68505999d911d5b99b017c0b7077398b9.
//
// Solidity: event LogAndIncrementCalled(uint256 expected, uint256 have)
func (_Simple *SimpleFilterer) FilterLogAndIncrementCalled(opts *bind.FilterOpts) (*SimpleLogAndIncrementCalledIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "LogAndIncrementCalled")
	if err != nil {
		return nil, err
	}
	return &SimpleLogAndIncrementCalledIterator{contract: _Simple.contract, event: "LogAndIncrementCalled", logs: logs, sub: sub}, nil
}

// WatchLogAndIncrementCalled is a free log subscription operation binding the contract event 0x8df8e492f407b078593c5d8fd7e65ef68505999d911d5b99b017c0b7077398b9.
//
// Solidity: event LogAndIncrementCalled(uint256 expected, uint256 have)
func (_Simple *SimpleFilterer) WatchLogAndIncrementCalled(opts *bind.WatchOpts, sink chan<- *SimpleLogAndIncrementCalled) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "LogAndIncrementCalled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleLogAndIncrementCalled)
				if err := _Simple.contract.UnpackLog(event, "LogAndIncrementCalled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAndIncrementCalled is a log parse operation binding the contract event 0x8df8e492f407b078593c5d8fd7e65ef68505999d911d5b99b017c0b7077398b9.
//
// Solidity: event LogAndIncrementCalled(uint256 expected, uint256 have)
func (_Simple *SimpleFilterer) ParseLogAndIncrementCalled(log types.Log) (*SimpleLogAndIncrementCalled, error) {
	event := new(SimpleLogAndIncrementCalled)
	if err := _Simple.contract.UnpackLog(event, "LogAndIncrementCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleNullEventIterator is returned from FilterNullEvent and is used to iterate over the raw logs and unpacked data for NullEvent events raised by the Simple contract.
type SimpleNullEventIterator struct {
	Event *SimpleNullEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleNullEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleNullEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SimpleNullEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SimpleNullEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleNullEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleNullEvent represents a NullEvent event raised by the Simple contract.
type SimpleNullEvent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNullEvent is a free log retrieval operation binding the contract event 0x6f59c82101949290205a9ae9d0c657e6dae1a71c301ae76d385c2792294585fe.
//
// Solidity: event NullEvent()
func (_Simple *SimpleFilterer) FilterNullEvent(opts *bind.FilterOpts) (*SimpleNullEventIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "NullEvent")
	if err != nil {
		return nil, err
	}
	return &SimpleNullEventIterator{contract: _Simple.contract, event: "NullEvent", logs: logs, sub: sub}, nil
}

// WatchNullEvent is a free log subscription operation binding the contract event 0x6f59c82101949290205a9ae9d0c657e6dae1a71c301ae76d385c2792294585fe.
//
// Solidity: event NullEvent()
func (_Simple *SimpleFilterer) WatchNullEvent(opts *bind.WatchOpts, sink chan<- *SimpleNullEvent) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "NullEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleNullEvent)
				if err := _Simple.contract.UnpackLog(event, "NullEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNullEvent is a log parse operation binding the contract event 0x6f59c82101949290205a9ae9d0c657e6dae1a71c301ae76d385c2792294585fe.
//
// Solidity: event NullEvent()
func (_Simple *SimpleFilterer) ParseNullEvent(log types.Log) (*SimpleNullEvent, error) {
	event := new(SimpleNullEvent)
	if err := _Simple.contract.UnpackLog(event, "NullEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRedeemedEventIterator is returned from FilterRedeemedEvent and is used to iterate over the raw logs and unpacked data for RedeemedEvent events raised by the Simple contract.
type SimpleRedeemedEventIterator struct {
	Event *SimpleRedeemedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SimpleRedeemedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRedeemedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SimpleRedeemedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SimpleRedeemedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRedeemedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRedeemedEvent represents a RedeemedEvent event raised by the Simple contract.
type SimpleRedeemedEvent struct {
	Caller   common.Address
	Redeemer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeemedEvent is a free log retrieval operation binding the contract event 0x773c78bf96e65f61c1a2622b47d76e78bfe70dd59cf4f11470c4c121c3159413.
//
// Solidity: event RedeemedEvent(address caller, address redeemer)
func (_Simple *SimpleFilterer) FilterRedeemedEvent(opts *bind.FilterOpts) (*SimpleRedeemedEventIterator, error) {

	logs, sub, err := _Simple.contract.FilterLogs(opts, "RedeemedEvent")
	if err != nil {
		return nil, err
	}
	return &SimpleRedeemedEventIterator{contract: _Simple.contract, event: "RedeemedEvent", logs: logs, sub: sub}, nil
}

// WatchRedeemedEvent is a free log subscription operation binding the contract event 0x773c78bf96e65f61c1a2622b47d76e78bfe70dd59cf4f11470c4c121c3159413.
//
// Solidity: event RedeemedEvent(address caller, address redeemer)
func (_Simple *SimpleFilterer) WatchRedeemedEvent(opts *bind.WatchOpts, sink chan<- *SimpleRedeemedEvent) (event.Subscription, error) {

	logs, sub, err := _Simple.contract.WatchLogs(opts, "RedeemedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRedeemedEvent)
				if err := _Simple.contract.UnpackLog(event, "RedeemedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeemedEvent is a log parse operation binding the contract event 0x773c78bf96e65f61c1a2622b47d76e78bfe70dd59cf4f11470c4c121c3159413.
//
// Solidity: event RedeemedEvent(address caller, address redeemer)
func (_Simple *SimpleFilterer) ParseRedeemedEvent(log types.Log) (*SimpleRedeemedEvent, error) {
	event := new(SimpleRedeemedEvent)
	if err := _Simple.contract.UnpackLog(event, "RedeemedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleProxyMetaData contains all meta data concerning the SimpleProxy contract.
var SimpleProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"impl_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161011d38038061011d83398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b608051609561008860003960006017015260956000f3fe608060405236601057600e6013565b005b600e5b603a7f0000000000000000000000000000000000000000000000000000000000000000603c565b565b3660008037600080366000845af43d6000803e808015605a573d6000f35b3d6000fdfea26469706673582212207509cd70dccdfb4f725f4b8ea78e709940520b3ba4ea15a2b4b4870a7a3152ab64736f6c63430008090033",
}

// SimpleProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleProxyMetaData.ABI instead.
var SimpleProxyABI = SimpleProxyMetaData.ABI

// SimpleProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleProxyMetaData.Bin instead.
var SimpleProxyBin = SimpleProxyMetaData.Bin

// DeploySimpleProxy deploys a new Ethereum contract, binding an instance of SimpleProxy to it.
func DeploySimpleProxy(auth *bind.TransactOpts, backend bind.ContractBackend, impl_ common.Address) (common.Address, *types.Transaction, *SimpleProxy, error) {
	parsed, err := SimpleProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleProxyBin), backend, impl_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleProxy{SimpleProxyCaller: SimpleProxyCaller{contract: contract}, SimpleProxyTransactor: SimpleProxyTransactor{contract: contract}, SimpleProxyFilterer: SimpleProxyFilterer{contract: contract}}, nil
}

// SimpleProxy is an auto generated Go binding around an Ethereum contract.
type SimpleProxy struct {
	SimpleProxyCaller     // Read-only binding to the contract
	SimpleProxyTransactor // Write-only binding to the contract
	SimpleProxyFilterer   // Log filterer for contract events
}

// SimpleProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleProxySession struct {
	Contract     *SimpleProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleProxyCallerSession struct {
	Contract *SimpleProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SimpleProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleProxyTransactorSession struct {
	Contract     *SimpleProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SimpleProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleProxyRaw struct {
	Contract *SimpleProxy // Generic contract binding to access the raw methods on
}

// SimpleProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleProxyCallerRaw struct {
	Contract *SimpleProxyCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleProxyTransactorRaw struct {
	Contract *SimpleProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleProxy creates a new instance of SimpleProxy, bound to a specific deployed contract.
func NewSimpleProxy(address common.Address, backend bind.ContractBackend) (*SimpleProxy, error) {
	contract, err := bindSimpleProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleProxy{SimpleProxyCaller: SimpleProxyCaller{contract: contract}, SimpleProxyTransactor: SimpleProxyTransactor{contract: contract}, SimpleProxyFilterer: SimpleProxyFilterer{contract: contract}}, nil
}

// NewSimpleProxyCaller creates a new read-only instance of SimpleProxy, bound to a specific deployed contract.
func NewSimpleProxyCaller(address common.Address, caller bind.ContractCaller) (*SimpleProxyCaller, error) {
	contract, err := bindSimpleProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleProxyCaller{contract: contract}, nil
}

// NewSimpleProxyTransactor creates a new write-only instance of SimpleProxy, bound to a specific deployed contract.
func NewSimpleProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleProxyTransactor, error) {
	contract, err := bindSimpleProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleProxyTransactor{contract: contract}, nil
}

// NewSimpleProxyFilterer creates a new log filterer instance of SimpleProxy, bound to a specific deployed contract.
func NewSimpleProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleProxyFilterer, error) {
	contract, err := bindSimpleProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleProxyFilterer{contract: contract}, nil
}

// bindSimpleProxy binds a generic wrapper to an already deployed contract.
func bindSimpleProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleProxy *SimpleProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleProxy.Contract.SimpleProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleProxy *SimpleProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleProxy.Contract.SimpleProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleProxy *SimpleProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleProxy.Contract.SimpleProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleProxy *SimpleProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleProxy *SimpleProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleProxy *SimpleProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleProxy.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleProxy *SimpleProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SimpleProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleProxy *SimpleProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleProxy.Contract.Fallback(&_SimpleProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleProxy *SimpleProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleProxy.Contract.Fallback(&_SimpleProxy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleProxy *SimpleProxyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleProxy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleProxy *SimpleProxySession) Receive() (*types.Transaction, error) {
	return _SimpleProxy.Contract.Receive(&_SimpleProxy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleProxy *SimpleProxyTransactorSession) Receive() (*types.Transaction, error) {
	return _SimpleProxy.Contract.Receive(&_SimpleProxy.TransactOpts)
}

// SingleExecutionChallengeMetaData contains all meta data concerning the SingleExecutionChallenge contract.
var SingleExecutionChallengeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"},{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessagesRead_\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[2]\",\"name\":\"startAndEndHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"numSteps_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"newSegments\",\"type\":\"bytes32[]\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"machineStatuses\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"globalStateHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"numSteps\",\"type\":\"uint256\"}],\"name\":\"challengeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProveExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"osp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultReceiver\",\"outputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalChallengesCreated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a0604052306080523480156200001557600080fd5b50604051620036ff380380620036ff8339810160408190526200003891620002bc565b600580546001600160a01b03808c166001600160a01b03199283161790925560028054928b1692909116919091179055600080548190819062000084906001600160401b0316620003cd565b82546101009290920a6001600160401b03818102199093168284169182021790935560009283526001602090815260408085206007810180546001600160401b031916958f16959095179094558051600280825260608201835293965093949392918301908036833750508a518251929350918391506000906200010c576200010c62000403565b602090810291909101015288600160200201518160018151811062000135576200013562000403565b60200260200101818152505060006200015c60008a846200024360201b620019fb1760201c565b600684018190556040805180820182526001600160a01b038b811680835260209283018b90526002880180546001600160a01b03199081169092179055600388018b905583518085018552918c168083529190920189905286549091161785556001850187905542600486015560078501805460ff60401b1916680200000000000000001790555190915081906001600160401b038616907f86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d689062000228906000908e90889062000419565b60405180910390a350505050505050505050505050620004b4565b60008383836040516020016200025c9392919062000470565b6040516020818303038152906040528051906020012090509392505050565b6001600160a01b03811681146200029157600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b8051620002b7816200027b565b919050565b60008060008060008060008060006101408a8c031215620002dc57600080fd5b8951620002e9816200027b565b809950506020808b0151620002fe816200027b565b60408c01519099506001600160401b0380821682146200031d57600080fd5b8199508d607f8e01126200033057600080fd5b604051915060408201828110828211171562000350576200035062000294565b604052508060a08d018e8111156200036757600080fd5b60608e015b818110156200038557805183529184019184016200036c565b50519198509096506200039e91505060c08b01620002aa565b9350620003ae60e08b01620002aa565b92506101008a015191506101208a015190509295985092959850929598565b60006001600160401b0382811680821415620003f957634e487b7160e01b600052601160045260246000fd5b6001019392505050565b634e487b7160e01b600052603260045260246000fd5b6000606082018583526020858185015260606040850152818551808452608086019150828701935060005b81811015620004625784518352938301939183019160010162000444565b509098975050505050505050565b83815260006020848184015260408301845182860160005b82811015620004a65781518452928401929084019060010162000488565b509198975050505050505050565b60805161322f620004d06000396000611225015261322f6000f3fe608060405234801561001057600080fd5b50600436106100e05760003560e01c80639ede42b9116100875780639ede42b914610251578063a521b03214610274578063d248d12414610287578063e78cea921461029a578063ee35f327146102ad578063f26a62c6146102c0578063f8c8765e146102d3578063fb7be0a1146102e657600080fd5b806314eab5e7146100e55780631b45c86a1461011557806323a9ef231461012a5780633504f1d71461015557806356e9df97146101685780635ef489e61461017b5780637fd07a9c1461018e5780638f1d3776146101ae575b600080fd5b6100f86100f33660046127af565b6102f9565b6040516001600160401b0390911681526020015b60405180910390f35b610128610123366004612842565b610601565b005b61013d610138366004612842565b6106d1565b6040516001600160a01b03909116815260200161010c565b60025461013d906001600160a01b031681565b610128610176366004612842565b6106f5565b6000546100f8906001600160401b031681565b6101a161019c366004612842565b610863565b60405161010c919061289f565b61023e6101bc366004612911565b6001602081815260009283526040928390208351808501855281546001600160a01b0390811682529382015481840152845180860190955260028201549093168452600381015491840191909152600481015460058201546006830154600790930154939493919290916001600160401b03811690600160401b900460ff1687565b60405161010c979695949392919061292a565b61026461025f366004612842565b61093c565b604051901515815260200161010c565b610128610282366004612987565b610963565b610128610295366004612a2b565b610dd9565b60045461013d906001600160a01b031681565b60035461013d906001600160a01b031681565b60055461013d906001600160a01b031681565b6101286102e1366004612abd565b61121a565b6101286102f4366004612b19565b61138b565b6002546000906001600160a01b0316331461034e5760405162461bcd60e51b815260206004820152601060248201526f13d3931657d493d313155417d0d2105360821b60448201526064015b60405180910390fd5b6040805160028082526060820183526000926020830190803683370190505090506103a461037f60208b018b612bbd565b61039f8a60005b6080020180360381019061039a9190612c7c565b611a32565b611ab3565b816000815181106103b7576103b7612ba7565b60209081029190910101526103e68960016020020160208101906103db9190612bbd565b61039f8a6001610386565b816001815181106103f9576103f9612ba7565b6020908102919091010152600080548190819061041e906001600160401b0316612d2a565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060006001600160401b0316816001600160401b0316141561046657610466612d51565b6001600160401b0381166000908152600160205260408120600581018d9055906104a061049b368d90038d0160808e01612c7c565b611bd7565b905060026104b460408e0160208f01612bbd565b60038111156104c5576104c5612875565b14806104f3575060006104e86104e3368e90038e0160808f01612c7c565b611bec565b6001600160401b0316115b15610506578061050281612d2a565b9150505b6007820180546040805180820182526001600160a01b038d811680835260209283018d90526002880180546001600160a01b03199081169092179055600388018d905583518085018552918e16808352919092018b90528654909116178555600185018990554260048601556001600160401b0384811668ffffffffffffffffff1990931692909217600160401b179092559051908416907f76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a906105d0908e906080820190612db1565b60405180910390a26105ee8360008c6001600160401b031687611bfb565b5090925050505b98975050505050505050565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff16600281111561063a5761063a612875565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b8152509061067a5760405162461bcd60e51b81526004016103459190612dfd565b506106848161093c565b6106c35760405162461bcd60e51b815260206004820152601060248201526f54494d454f55545f444541444c494e4560801b6044820152606401610345565b6106ce816000611c91565b50565b6001600160401b03166000908152600160205260409020546001600160a01b031690565b6002546001600160a01b031633146107425760405162461bcd60e51b815260206004820152601060248201526f2727aa2fa922a9afa922a1a2a4ab22a960811b6044820152606401610345565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff16600281111561077b5761077b612875565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906107bb5760405162461bcd60e51b81526004016103459190612dfd565b506001600160401b038116600081815260016020819052604080832080546001600160a01b031990811682559281018490556002810180549093169092556003808301849055600483018490556005830184905560068301939093556007909101805468ffffffffffffffffff19169055517ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f409161085891612e30565b60405180910390a250565b61086b61270a565b6001600160401b0382811660009081526001602081815260409283902083516101208101855281546001600160a01b0390811660e0830190815294830154610100830152938152845180860186526002808401549095168152600383015481850152928101929092526004810154938201939093526005830154606082015260068301546080820152600783015493841660a08201529260c0840191600160401b90910460ff169081111561092257610922612875565b600281111561093357610933612875565b90525092915050565b6001600160401b038116600090815260016020526040812061095d90611dbf565b92915050565b6001600160401b038416600090815260016020526040812085918591610988846106d1565b6001600160a01b0316336001600160a01b0316146109b85760405162461bcd60e51b815260040161034590612e4a565b6109c18461093c565b156109de5760405162461bcd60e51b815260040161034590612e6f565b60008260028111156109f2576109f2612875565b1415610a605760006007820154600160401b900460ff166002811115610a1a57610a1a612875565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610a5a5760405162461bcd60e51b81526004016103459190612dfd565b50610b1f565b6001826002811115610a7457610a74612875565b1415610abe5760016007820154600160401b900460ff166002811115610a9c57610a9c612875565b14610ab95760405162461bcd60e51b815260040161034590612e96565b610b1f565b6002826002811115610ad257610ad2612875565b1415610b175760026007820154600160401b900460ff166002811115610afa57610afa612875565b14610ab95760405162461bcd60e51b815260040161034590612ebe565b610b1f612d51565b610b6d83356020850135610b366040870187612eea565b808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506119fb92505050565b816006015414610b8f5760405162461bcd60e51b815260040161034590612f3a565b6002610b9e6040850185612eea565b90501080610bc957506001610bb66040850185612eea565b610bc1929150612f5d565b836060013510155b15610be65760405162461bcd60e51b815260040161034590612f74565b600080610bf289611dd7565b9150915060018111610c325760405162461bcd60e51b81526020600482015260096024820152681513d3d7d4d213d49560ba1b6044820152606401610345565b806028811115610c40575060285b610c4b816001612f9f565b8814610c885760405162461bcd60e51b815260206004820152600c60248201526b57524f4e475f44454752454560a01b6044820152606401610345565b50610cd28989896000818110610ca057610ca0612ba7565b602002919091013590508a8a610cb7600182612f5d565b818110610cc657610cc6612ba7565b90506020020135611e68565b610d118a83838b8b80806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250611bfb92505050565b50600090505b6007820154600160401b900460ff166002811115610d3757610d37612875565b1415610d435750610dd0565b6040805180820190915281546001600160a01b03168152600182015460208201526004820154610d739042612f5d565b81602001818151610d849190612f5d565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b50505050505050565b6001600160401b038416600090815260016020526040902084908490600290610e01846106d1565b6001600160a01b0316336001600160a01b031614610e315760405162461bcd60e51b815260040161034590612e4a565b610e3a8461093c565b15610e575760405162461bcd60e51b815260040161034590612e6f565b6000826002811115610e6b57610e6b612875565b1415610ed95760006007820154600160401b900460ff166002811115610e9357610e93612875565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b81525090610ed35760405162461bcd60e51b81526004016103459190612dfd565b50610f98565b6001826002811115610eed57610eed612875565b1415610f375760016007820154600160401b900460ff166002811115610f1557610f15612875565b14610f325760405162461bcd60e51b815260040161034590612e96565b610f98565b6002826002811115610f4b57610f4b612875565b1415610f905760026007820154600160401b900460ff166002811115610f7357610f73612875565b14610f325760405162461bcd60e51b815260040161034590612ebe565b610f98612d51565b610faf83356020850135610b366040870187612eea565b816006015414610fd15760405162461bcd60e51b815260040161034590612f3a565b6002610fe06040850185612eea565b9050108061100b57506001610ff86040850185612eea565b611003929150612f5d565b836060013510155b156110285760405162461bcd60e51b815260040161034590612f74565b6001600160401b0388166000908152600160205260408120908061104b8a611dd7565b9092509050600181146110705760405162461bcd60e51b815260040161034590612fb7565b5060055460408051808201825260078501546001600160401b031681526004546001600160a01b0390811660208301526000931691635d3adcfb919085906110ba908f018f612eea565b8f606001358181106110ce576110ce612ba7565b905060200201358d8d6040518663ffffffff1660e01b81526004016110f7959493929190612fd9565b60206040518083038186803b15801561110f57600080fd5b505afa158015611123573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111479190613030565b905061115660408b018b612eea565b61116560608d01356001612f9f565b81811061117457611174612ba7565b905060200201358114156111b95760405162461bcd60e51b815260206004820152600c60248201526b14d0535157d3d4d417d1539160a21b6044820152606401610345565b6040516001600160401b038c16907fc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e90600090a261120e8b6001600160401b0316600090815260016020526040812060060155565b5060009150610d179050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614156112a85760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608401610345565b6002546001600160a01b0316156112f05760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b6044820152606401610345565b6001600160a01b03841661133b5760405162461bcd60e51b81526020600482015260126024820152712727afa922a9aaa62a2fa922a1a2a4ab22a960711b6044820152606401610345565b600280546001600160a01b039586166001600160a01b0319918216179091556003805494861694821694909417909355600480549285169284169290921790915560058054919093169116179055565b6001600160401b0385166000908152600160208190526040909120869186916113b3846106d1565b6001600160a01b0316336001600160a01b0316146113e35760405162461bcd60e51b815260040161034590612e4a565b6113ec8461093c565b156114095760405162461bcd60e51b815260040161034590612e6f565b600082600281111561141d5761141d612875565b141561148b5760006007820154600160401b900460ff16600281111561144557611445612875565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906114855760405162461bcd60e51b81526004016103459190612dfd565b5061154a565b600182600281111561149f5761149f612875565b14156114e95760016007820154600160401b900460ff1660028111156114c7576114c7612875565b146114e45760405162461bcd60e51b815260040161034590612e96565b61154a565b60028260028111156114fd576114fd612875565b14156115425760026007820154600160401b900460ff16600281111561152557611525612875565b146114e45760405162461bcd60e51b815260040161034590612ebe565b61154a612d51565b61156183356020850135610b366040870187612eea565b8160060154146115835760405162461bcd60e51b815260040161034590612f3a565b60026115926040850185612eea565b905010806115bd575060016115aa6040850185612eea565b6115b5929150612f5d565b836060013510155b156115da5760405162461bcd60e51b815260040161034590612f74565b60018510156116215760405162461bcd60e51b815260206004820152601360248201527210d2105313115391d157d513d3d7d4d213d495606a1b6044820152606401610345565b6508000000000085111561166c5760405162461bcd60e51b81526020600482015260126024820152714348414c4c454e47455f544f4f5f4c4f4e4760701b6044820152606401610345565b6116ae8861168e61168060208b018b612bbd565b8960005b6020020135611ab3565b6116a96116a160408c0160208d01612bbd565b8a6001611684565b611e68565b6001600160401b038916600090815260016020526040812090806116d18b611dd7565b91509150806001146116f55760405162461bcd60e51b815260040161034590612fb7565b600161170460208c018c612bbd565b600381111561171557611715612875565b146117cf5761172a60408b0160208c01612bbd565b600381111561173b5761173b612875565b61174860208c018c612bbd565b600381111561175957611759612875565b14801561176a5750883560208a0135145b6117a65760405162461bcd60e51b815260206004820152600d60248201526c48414c5445445f4348414e474560981b6044820152606401610345565b6117c78c6001600160401b0316600090815260016020526040812060060155565b505050611936565b60026117e160408c0160208d01612bbd565b60038111156117f2576117f2612875565b141561183b57883560208a01351461183b5760405162461bcd60e51b815260206004820152600c60248201526b4552524f525f4348414e474560a01b6044820152606401610345565b604080516002808252606082018352600092602083019080368337505050600585015490915061186d908b3590611f3d565b8160008151811061188057611880612ba7565b60209081029190910101526118ae8b60016020020160208101906118a49190612bbd565b60208c01356120ea565b816001815181106118c1576118c1612ba7565b602090810291909101015260078401805460ff60401b1916600160411b1790556118ee8d60008b84611bfb565b8c6001600160401b03167f24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db8460405161192991815260200190565b60405180910390a2505050505b60006007820154600160401b900460ff16600281111561195857611958612875565b141561196457506119f1565b6040805180820190915281546001600160a01b031681526001820154602082015260048201546119949042612f5d565b816020018181516119a59190612f5d565b90525060028201805483546001600160a01b038083166001600160a01b031992831617865560038601805460018801558551929093169116179091556020909101519055426004909101555b5050505050505050565b6000838383604051602001611a1293929190613049565b6040516020818303038152906040528051906020012090505b9392505050565b80518051602091820151828401518051908401516040516c23b637b130b61039ba30ba329d60991b95810195909552602d850193909352604d8401919091526001600160c01b031960c091821b8116606d85015291901b166075820152600090607d015b604051602081830303815290604052805190602001209050919050565b60006001836003811115611ac957611ac9612875565b1415611b0f576040516b213637b1b59039ba30ba329d60a11b6020820152602c8101839052604c015b60405160208183030381529060405280519060200120905061095d565b6002836003811115611b2357611b23612875565b1415611b595760405174213637b1b59039ba30ba32961032b93937b932b21d60591b602082015260358101839052605501611af2565b6003836003811115611b6d57611b6d612875565b1415611b9c5760405174213637b1b59039ba30ba3296103a37b7903330b91d60591b6020820152603501611af2565b60405162461bcd60e51b815260206004820152601060248201526f4241445f424c4f434b5f53544154555360801b6044820152606401610345565b6020810151600090815b602002015192915050565b60208101516000906001611be1565b6001821015611c0c57611c0c612d51565b600281511015611c1e57611c1e612d51565b6000611c2b8484846119fb565b6001600160401b038616600081815260016020526040908190206006018390555191925082917f86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d6890611c829088908890889061308b565b60405180910390a35050505050565b6001600160401b03821660008181526001602081905260408083206002808201805483546001600160a01b0319808216865596850188905595811690915560038301869055600480840187905560058401879055600684019690965560078301805468ffffffffffffffffff1916905590549251630357aa4960e01b8152948501959095526001600160a01b03948516602485018190529285166044850181905290949293909290911690630357aa4990606401600060405180830381600087803b158015611d5f57600080fd5b505af1158015611d73573d6000803e3d6000fd5b50505050846001600160401b03167ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f4085604051611db09190612e30565b60405180910390a25050505050565b6001810154600090611dd083612193565b1192915050565b600080806001611dea6040860186612eea565b611df5929150612f5d565b9050611e058160208601356130f6565b9150611e1560608501358361310a565b611e20908535612f9f565b92506002611e316040860186612eea565b611e3c929150612f5d565b84606001351415611e6257611e55816020860135613129565b611e5f9083612f9f565b91505b50915091565b81611e766040850185612eea565b8560600135818110611e8a57611e8a612ba7565b9050602002013514611ecc5760405162461bcd60e51b815260206004820152600b60248201526a15d493d391d7d4d510549560aa1b6044820152606401610345565b80611eda6040850185612eea565b611ee960608701356001612f9f565b818110611ef857611ef8612ba7565b905060200201351415611f385760405162461bcd60e51b815260206004820152600860248201526714d0535157d1539160c21b6044820152606401610345565b505050565b60408051600380825260808201909252600091829190816020015b6040805180820190915260008082526020820152815260200190600190039081611f58575050604080518082018252600080825260209182018190528251808401909352600483529082015290915081600081518110611fba57611fba612ba7565b6020026020010181905250611fcf60006121a5565b81600181518110611fe257611fe2612ba7565b6020026020010181905250611ff760006121a5565b8160028151811061200a5761200a612ba7565b602090810291909101810191909152604080518083018252838152815180830190925280825260009282019290925261205a60408051606080820183529181019182529081526000602082015290565b604080518082018252606080825260006020808401829052845180840186528381528082018390528086018390528551610140810187528381529182018890529481018690529182018390526080820184905260a082018b905260c0820181905260e0820181905261010082015261012081018990529091906120dc816121d8565b9a9950505050505050505050565b6000600183600381111561210057612100612875565b14156121175781604051602001611af2919061313d565b600283600381111561212b5761212b612875565b1415612155576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001611af2565b600383600381111561216957612169612875565b1415611b9c576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001611af2565b600081600401544261095d9190612f5d565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b600080825160038111156121ee576121ee612875565b141561233e576000612203836020015161242b565b612210846040015161242b565b61221d85606001516124b0565b60a086015160c087015160e0808901516101008a01516101208b01516040516f26b0b1b434b73290393ab73734b7339d60811b602082015260308101999099526050890197909752607088019590955260908701939093526001600160e01b031991811b821660b087015291821b811660b486015291901b1660b883015260bc82015260dc0160405160208183030381529060405290506122c18360800151612549565b80156122d35750826080015160400151155b156122e657805160209091012092915050565b608083015160400151600090156122fe5750600160f81b5b818161230d8660800151612560565b60405160200161231f93929190613162565b6040516020818303038152906040528051906020012092505050919050565b60018251600381111561235357612353612875565b141561236e578160a00151604051602001611a96919061313d565b60028251600381111561238357612383612875565b14156123ad576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001611a96565b6003825160038111156123c2576123c2612875565b14156123ec576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001611a96565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606401610345565b919050565b60208101518151515160005b818110156124a95783516124549061244f90836125ed565b612625565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c0160405160208183030381529060405280519060200120925080806124a1906131a9565b915050612437565b5050919050565b602081015160005b825151811015612543576124e8836000015182815181106124db576124db612ba7565b6020026020010151612642565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b6020820152603281019190915260528101839052607201604051602081830303815290604052805190602001209150808061253b906131a9565b9150506124b8565b50919050565b80515160009015801561095d575050602001511590565b602081015160005b825151811015612543576125988360000151828151811061258b5761258b612ba7565b60200260200101516126b2565b6040516b23bab0b9321039ba30b1b59d60a11b6020820152602c810191909152604c8101839052606c0160405160208183030381529060405280519060200120915080806125e5906131a9565b915050612568565b6040805180820190915260008082526020820152825180518390811061261557612615612ba7565b6020026020010151905092915050565b600081600001518260200151604051602001611a969291906131c4565b60006126518260000151612625565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401611a96565b60008160000151826020015183604001516126d08560600151612625565b6040516b22b93937b91033bab0b9321d60a11b6020820152602c810194909452604c840192909252606c830152608c82015260ac01611a96565b604080516101208101909152600060e0820181815261010083019190915281908152602001612749604080518082019091526000808252602082015290565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b806040810183101561095d57600080fd5b80356001600160401b038116811461242657600080fd5b6001600160a01b03811681146106ce57600080fd5b600080600080600080600080610200898b0312156127cc57600080fd5b883597506127dd8a60208b01612772565b965061016089018a8111156127f157600080fd5b60608a01965061280081612783565b9550506101808901356128128161279a565b93506101a08901356128238161279a565b979a96995094979396929592945050506101c0820135916101e0013590565b60006020828403121561285457600080fd5b611a2b82612783565b80516001600160a01b03168252602090810151910152565b634e487b7160e01b600052602160045260246000fd5b6003811061289b5761289b612875565b9052565b6000610120820190506128b382845161285d565b60208301516128c5604084018261285d565b5060408301516080830152606083015160a0830152608083015160c08301526001600160401b0360a08401511660e083015260c083015161290a61010084018261288b565b5092915050565b60006020828403121561292357600080fd5b5035919050565b6101208101612939828a61285d565b612946604083018961285d565b8660808301528560a08301528460c08301526001600160401b03841660e08301526105f561010083018461288b565b60006080828403121561254357600080fd5b6000806000806060858703121561299d57600080fd5b6129a685612783565b935060208501356001600160401b03808211156129c257600080fd5b6129ce88838901612975565b945060408701359150808211156129e457600080fd5b818701915087601f8301126129f857600080fd5b813581811115612a0757600080fd5b8860208260051b8501011115612a1c57600080fd5b95989497505060200194505050565b60008060008060608587031215612a4157600080fd5b612a4a85612783565b935060208501356001600160401b0380821115612a6657600080fd5b612a7288838901612975565b94506040870135915080821115612a8857600080fd5b818701915087601f830112612a9c57600080fd5b813581811115612aab57600080fd5b886020828501011115612a1c57600080fd5b60008060008060808587031215612ad357600080fd5b8435612ade8161279a565b93506020850135612aee8161279a565b92506040850135612afe8161279a565b91506060850135612b0e8161279a565b939692955090935050565b600080600080600060e08688031215612b3157600080fd5b612b3a86612783565b945060208601356001600160401b03811115612b5557600080fd5b612b6188828901612975565b945050612b718760408801612772565b9250612b808760808801612772565b9497939650919460c0013592915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060208284031215612bcf57600080fd5b813560048110611a2b57600080fd5b604080519081016001600160401b0381118282101715612c0057612c00612b91565b60405290565b600082601f830112612c1757600080fd5b604051604081018181106001600160401b0382111715612c3957612c39612b91565b8060405250806040840185811115612c5057600080fd5b845b81811015612c7157612c6381612783565b835260209283019201612c52565b509195945050505050565b600060808284031215612c8e57600080fd5b604051604081018181106001600160401b0382111715612cb057612cb0612b91565b604052601f83018413612cc257600080fd5b612cca612bde565b806040850186811115612cdc57600080fd5b855b81811015612cf6578035845260209384019301612cde565b50818452612d048782612c06565b6020850152509195945050505050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380831681811415612d4757612d47612d14565b6001019392505050565b634e487b7160e01b600052600160045260246000fd5b604081833760006040838101828152908301915b6002811015612daa576001600160401b03612d9584612783565b16825260209283019290910190600101612d7b565b5050505050565b6101008101612dc08285612d67565b611a2b6080830184612d67565b60005b83811015612de8578181015183820152602001612dd0565b83811115612df7576000848401525b50505050565b6020815260008251806020840152612e1c816040850160208701612dcd565b601f01601f19169190910160400192915050565b6020810160048310612e4457612e44612875565b91905290565b6020808252600b908201526a21a420a62fa9a2a72222a960a91b604082015260600190565b6020808252600d908201526c4348414c5f444541444c494e4560981b604082015260600190565b6020808252600e908201526d4348414c5f4e4f545f424c4f434b60901b604082015260600190565b60208082526012908201527121a420a62fa727aa2fa2ac22a1aaaa24a7a760711b604082015260600190565b6000808335601e19843603018112612f0157600080fd5b8301803591506001600160401b03821115612f1b57600080fd5b6020019150600581901b3603821315612f3357600080fd5b9250929050565b6020808252600990820152684249535f535441544560b81b604082015260600190565b600082821015612f6f57612f6f612d14565b500390565b6020808252601190820152704241445f4348414c4c454e47455f504f5360781b604082015260600190565b60008219821115612fb257612fb2612d14565b500190565b602080825260089082015267544f4f5f4c4f4e4760c01b604082015260600190565b8551815260018060a01b03602087015116602082015284604082015283606082015260a060808201528160a0820152818360c0830137600081830160c090810191909152601f909201601f19160101949350505050565b60006020828403121561304257600080fd5b5051919050565b83815260006020848184015260408301845182860160005b8281101561307d57815184529284019290840190600101613061565b509198975050505050505050565b6000606082018583526020858185015260606040850152818551808452608086019150828701935060005b818110156130d2578451835293830193918301916001016130b6565b509098975050505050505050565b634e487b7160e01b600052601260045260246000fd5b600082613105576131056130e0565b500490565b600081600019048311821515161561312457613124612d14565b500290565b600082613138576131386130e0565b500690565b7026b0b1b434b732903334b734b9b432b21d60791b8152601181019190915260310190565b60008451613174818460208901612dcd565b6b2bb4ba341033bab0b932399d60a11b9201918252506001600160f81b031992909216600c830152600d820152602d01919050565b60006000198214156131bd576131bd612d14565b5060010190565b652b30b63ab29d60d11b81526000600784106131e2576131e2612875565b5060f89290921b600683015260078201526027019056fea2646970667358221220567f8bb13dab5176b64d85e63ae5c98a933695340b8afaa42a44471ded21609e64736f6c63430008090033",
}

// SingleExecutionChallengeABI is the input ABI used to generate the binding from.
// Deprecated: Use SingleExecutionChallengeMetaData.ABI instead.
var SingleExecutionChallengeABI = SingleExecutionChallengeMetaData.ABI

// SingleExecutionChallengeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SingleExecutionChallengeMetaData.Bin instead.
var SingleExecutionChallengeBin = SingleExecutionChallengeMetaData.Bin

// DeploySingleExecutionChallenge deploys a new Ethereum contract, binding an instance of SingleExecutionChallenge to it.
func DeploySingleExecutionChallenge(auth *bind.TransactOpts, backend bind.ContractBackend, osp_ common.Address, resultReceiver_ common.Address, maxInboxMessagesRead_ uint64, startAndEndHashes [2][32]byte, numSteps_ *big.Int, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (common.Address, *types.Transaction, *SingleExecutionChallenge, error) {
	parsed, err := SingleExecutionChallengeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SingleExecutionChallengeBin), backend, osp_, resultReceiver_, maxInboxMessagesRead_, startAndEndHashes, numSteps_, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SingleExecutionChallenge{SingleExecutionChallengeCaller: SingleExecutionChallengeCaller{contract: contract}, SingleExecutionChallengeTransactor: SingleExecutionChallengeTransactor{contract: contract}, SingleExecutionChallengeFilterer: SingleExecutionChallengeFilterer{contract: contract}}, nil
}

// SingleExecutionChallenge is an auto generated Go binding around an Ethereum contract.
type SingleExecutionChallenge struct {
	SingleExecutionChallengeCaller     // Read-only binding to the contract
	SingleExecutionChallengeTransactor // Write-only binding to the contract
	SingleExecutionChallengeFilterer   // Log filterer for contract events
}

// SingleExecutionChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SingleExecutionChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleExecutionChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SingleExecutionChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleExecutionChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SingleExecutionChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SingleExecutionChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SingleExecutionChallengeSession struct {
	Contract     *SingleExecutionChallenge // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SingleExecutionChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SingleExecutionChallengeCallerSession struct {
	Contract *SingleExecutionChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// SingleExecutionChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SingleExecutionChallengeTransactorSession struct {
	Contract     *SingleExecutionChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// SingleExecutionChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SingleExecutionChallengeRaw struct {
	Contract *SingleExecutionChallenge // Generic contract binding to access the raw methods on
}

// SingleExecutionChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SingleExecutionChallengeCallerRaw struct {
	Contract *SingleExecutionChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// SingleExecutionChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SingleExecutionChallengeTransactorRaw struct {
	Contract *SingleExecutionChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSingleExecutionChallenge creates a new instance of SingleExecutionChallenge, bound to a specific deployed contract.
func NewSingleExecutionChallenge(address common.Address, backend bind.ContractBackend) (*SingleExecutionChallenge, error) {
	contract, err := bindSingleExecutionChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallenge{SingleExecutionChallengeCaller: SingleExecutionChallengeCaller{contract: contract}, SingleExecutionChallengeTransactor: SingleExecutionChallengeTransactor{contract: contract}, SingleExecutionChallengeFilterer: SingleExecutionChallengeFilterer{contract: contract}}, nil
}

// NewSingleExecutionChallengeCaller creates a new read-only instance of SingleExecutionChallenge, bound to a specific deployed contract.
func NewSingleExecutionChallengeCaller(address common.Address, caller bind.ContractCaller) (*SingleExecutionChallengeCaller, error) {
	contract, err := bindSingleExecutionChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeCaller{contract: contract}, nil
}

// NewSingleExecutionChallengeTransactor creates a new write-only instance of SingleExecutionChallenge, bound to a specific deployed contract.
func NewSingleExecutionChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*SingleExecutionChallengeTransactor, error) {
	contract, err := bindSingleExecutionChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeTransactor{contract: contract}, nil
}

// NewSingleExecutionChallengeFilterer creates a new log filterer instance of SingleExecutionChallenge, bound to a specific deployed contract.
func NewSingleExecutionChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*SingleExecutionChallengeFilterer, error) {
	contract, err := bindSingleExecutionChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeFilterer{contract: contract}, nil
}

// bindSingleExecutionChallenge binds a generic wrapper to an already deployed contract.
func bindSingleExecutionChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SingleExecutionChallengeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleExecutionChallenge *SingleExecutionChallengeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleExecutionChallenge.Contract.SingleExecutionChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleExecutionChallenge *SingleExecutionChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.SingleExecutionChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleExecutionChallenge *SingleExecutionChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.SingleExecutionChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SingleExecutionChallenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Bridge() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.Bridge(&_SingleExecutionChallenge.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) Bridge() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.Bridge(&_SingleExecutionChallenge.CallOpts)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex uint64) (ChallengeLibChallenge, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "challengeInfo", challengeIndex)

	if err != nil {
		return *new(ChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeLibChallenge)).(*ChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _SingleExecutionChallenge.Contract.ChallengeInfo(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _SingleExecutionChallenge.Contract.ChallengeInfo(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		Current            ChallengeLibParticipant
		Next               ChallengeLibParticipant
		LastMoveTimestamp  *big.Int
		WasmModuleRoot     [32]byte
		ChallengeStateHash [32]byte
		MaxInboxMessages   uint64
		Mode               uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(ChallengeLibParticipant)).(*ChallengeLibParticipant)
	outstruct.Next = *abi.ConvertType(out[1], new(ChallengeLibParticipant)).(*ChallengeLibParticipant)
	outstruct.LastMoveTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WasmModuleRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.ChallengeStateHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.MaxInboxMessages = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Mode = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _SingleExecutionChallenge.Contract.Challenges(&_SingleExecutionChallenge.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _SingleExecutionChallenge.Contract.Challenges(&_SingleExecutionChallenge.CallOpts, arg0)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _SingleExecutionChallenge.Contract.CurrentResponder(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _SingleExecutionChallenge.Contract.CurrentResponder(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) IsTimedOut(opts *bind.CallOpts, challengeIndex uint64) (bool, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "isTimedOut", challengeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _SingleExecutionChallenge.Contract.IsTimedOut(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 challengeIndex) view returns(bool)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) IsTimedOut(challengeIndex uint64) (bool, error) {
	return _SingleExecutionChallenge.Contract.IsTimedOut(&_SingleExecutionChallenge.CallOpts, challengeIndex)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) Osp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "osp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Osp() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.Osp(&_SingleExecutionChallenge.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) Osp() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.Osp(&_SingleExecutionChallenge.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) ResultReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "resultReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) ResultReceiver() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.ResultReceiver(&_SingleExecutionChallenge.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) ResultReceiver() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.ResultReceiver(&_SingleExecutionChallenge.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) SequencerInbox() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.SequencerInbox(&_SingleExecutionChallenge.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) SequencerInbox() (common.Address, error) {
	return _SingleExecutionChallenge.Contract.SequencerInbox(&_SingleExecutionChallenge.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeCaller) TotalChallengesCreated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SingleExecutionChallenge.contract.Call(opts, &out, "totalChallengesCreated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) TotalChallengesCreated() (uint64, error) {
	return _SingleExecutionChallenge.Contract.TotalChallengesCreated(&_SingleExecutionChallenge.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeCallerSession) TotalChallengesCreated() (uint64, error) {
	return _SingleExecutionChallenge.Contract.TotalChallengesCreated(&_SingleExecutionChallenge.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) BisectExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "bisectExecution", challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.BisectExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.BisectExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, newSegments)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) ChallengeExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "challengeExecution", challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.ChallengeExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.ChallengeExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "clearChallenge", challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.ClearChallenge(&_SingleExecutionChallenge.TransactOpts, challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.ClearChallenge(&_SingleExecutionChallenge.TransactOpts, challengeIndex)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.CreateChallenge(&_SingleExecutionChallenge.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.CreateChallenge(&_SingleExecutionChallenge.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.Initialize(&_SingleExecutionChallenge.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.Initialize(&_SingleExecutionChallenge.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) OneStepProveExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "oneStepProveExecution", challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.OneStepProveExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.OneStepProveExecution(&_SingleExecutionChallenge.TransactOpts, challengeIndex, selection, proof)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactor) Timeout(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.contract.Transact(opts, "timeout", challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.Timeout(&_SingleExecutionChallenge.TransactOpts, challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_SingleExecutionChallenge *SingleExecutionChallengeTransactorSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _SingleExecutionChallenge.Contract.Timeout(&_SingleExecutionChallenge.TransactOpts, challengeIndex)
}

// SingleExecutionChallengeBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeBisectedIterator struct {
	Event *SingleExecutionChallengeBisected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleExecutionChallengeBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeBisected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleExecutionChallengeBisected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleExecutionChallengeBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeBisected represents a Bisected event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeBisected struct {
	ChallengeIndex          uint64
	ChallengeRoot           [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	ChainHashes             [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*SingleExecutionChallengeBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeBisectedIterator{contract: _SingleExecutionChallenge.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeBisected)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBisected is a log parse operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseBisected(log types.Log) (*SingleExecutionChallengeBisected, error) {
	event := new(SingleExecutionChallengeBisected)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleExecutionChallengeChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeChallengeEndedIterator struct {
	Event *SingleExecutionChallengeChallengeEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleExecutionChallengeChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeChallengeEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleExecutionChallengeChallengeEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleExecutionChallengeChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeChallengeEnded represents a ChallengeEnded event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*SingleExecutionChallengeChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeChallengeEndedIterator{contract: _SingleExecutionChallenge.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeChallengeEnded)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeEnded is a log parse operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseChallengeEnded(log types.Log) (*SingleExecutionChallengeChallengeEnded, error) {
	event := new(SingleExecutionChallengeChallengeEnded)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleExecutionChallengeExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeExecutionChallengeBegunIterator struct {
	Event *SingleExecutionChallengeExecutionChallengeBegun // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleExecutionChallengeExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeExecutionChallengeBegun)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleExecutionChallengeExecutionChallengeBegun)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleExecutionChallengeExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*SingleExecutionChallengeExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeExecutionChallengeBegunIterator{contract: _SingleExecutionChallenge.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeExecutionChallengeBegun)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionChallengeBegun is a log parse operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseExecutionChallengeBegun(log types.Log) (*SingleExecutionChallengeExecutionChallengeBegun, error) {
	event := new(SingleExecutionChallengeExecutionChallengeBegun)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleExecutionChallengeInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeInitiatedChallengeIterator struct {
	Event *SingleExecutionChallengeInitiatedChallenge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleExecutionChallengeInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeInitiatedChallenge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleExecutionChallengeInitiatedChallenge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleExecutionChallengeInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeInitiatedChallenge represents a InitiatedChallenge event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*SingleExecutionChallengeInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeInitiatedChallengeIterator{contract: _SingleExecutionChallenge.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeInitiatedChallenge)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseInitiatedChallenge(log types.Log) (*SingleExecutionChallengeInitiatedChallenge, error) {
	event := new(SingleExecutionChallengeInitiatedChallenge)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SingleExecutionChallengeOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeOneStepProofCompletedIterator struct {
	Event *SingleExecutionChallengeOneStepProofCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SingleExecutionChallengeOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SingleExecutionChallengeOneStepProofCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SingleExecutionChallengeOneStepProofCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SingleExecutionChallengeOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SingleExecutionChallengeOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SingleExecutionChallengeOneStepProofCompleted represents a OneStepProofCompleted event raised by the SingleExecutionChallenge contract.
type SingleExecutionChallengeOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*SingleExecutionChallengeOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &SingleExecutionChallengeOneStepProofCompletedIterator{contract: _SingleExecutionChallenge.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *SingleExecutionChallengeOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _SingleExecutionChallenge.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SingleExecutionChallengeOneStepProofCompleted)
				if err := _SingleExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_SingleExecutionChallenge *SingleExecutionChallengeFilterer) ParseOneStepProofCompleted(log types.Log) (*SingleExecutionChallengeOneStepProofCompleted, error) {
	event := new(SingleExecutionChallengeOneStepProofCompleted)
	if err := _SingleExecutionChallenge.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedOutChallengeManagerMetaData contains all meta data concerning the TimedOutChallengeManager contract.
var TimedOutChallengeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"challengeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedSegmentLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"chainHashes\",\"type\":\"bytes32[]\"}],\"name\":\"Bisected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumIChallengeManager.ChallengeTerminationType\",\"name\":\"kind\",\"type\":\"uint8\"}],\"name\":\"ChallengeEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockSteps\",\"type\":\"uint256\"}],\"name\":\"ExecutionChallengeBegun\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"startState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"indexed\":false,\"internalType\":\"structGlobalState\",\"name\":\"endState\",\"type\":\"tuple\"}],\"name\":\"InitiatedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"newSegments\",\"type\":\"bytes32[]\"}],\"name\":\"bisectExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"machineStatuses\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"globalStateHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"numSteps\",\"type\":\"uint256\"}],\"name\":\"challengeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"challengeInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"internalType\":\"structChallengeLib.Challenge\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"current\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeLeft\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.Participant\",\"name\":\"next\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastMoveTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengeStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"maxInboxMessages\",\"type\":\"uint64\"},{\"internalType\":\"enumChallengeLib.ChallengeMode\",\"name\":\"mode\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"clearChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"wasmModuleRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"enumMachineStatus[2]\",\"name\":\"startAndEndMachineStatuses_\",\"type\":\"uint8[2]\"},{\"components\":[{\"internalType\":\"bytes32[2]\",\"name\":\"bytes32Vals\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint64[2]\",\"name\":\"u64Vals\",\"type\":\"uint64[2]\"}],\"internalType\":\"structGlobalState[2]\",\"name\":\"startAndEndGlobalStates_\",\"type\":\"tuple[2]\"},{\"internalType\":\"uint64\",\"name\":\"numBlocks\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"asserter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challenger_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"asserterTimeLeft_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"challengerTimeLeft_\",\"type\":\"uint256\"}],\"name\":\"createChallenge\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"currentResponder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"resultReceiver_\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"sequencerInbox_\",\"type\":\"address\"},{\"internalType\":\"contractIBridge\",\"name\":\"bridge_\",\"type\":\"address\"},{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"osp_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"isTimedOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"oldSegmentsStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldSegmentsLength\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"oldSegments\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"challengePosition\",\"type\":\"uint256\"}],\"internalType\":\"structChallengeLib.SegmentSelection\",\"name\":\"selection\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProveExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"osp\",\"outputs\":[{\"internalType\":\"contractIOneStepProofEntry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resultReceiver\",\"outputs\":[{\"internalType\":\"contractIChallengeResultReceiver\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"challengeIndex\",\"type\":\"uint64\"}],\"name\":\"timeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalChallengesCreated\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50608051611ee361003060003960006109d30152611ee36000f3fe608060405234801561001057600080fd5b50600436106100e05760003560e01c80639ede42b9116100875780639ede42b914610251578063a521b03214610275578063d248d12414610288578063e78cea921461029b578063ee35f327146102ae578063f26a62c6146102c1578063f8c8765e146102d4578063fb7be0a1146102e757600080fd5b806314eab5e7146100e55780631b45c86a1461011557806323a9ef231461012a5780633504f1d71461015557806356e9df97146101685780635ef489e61461017b5780637fd07a9c1461018e5780638f1d3776146101ae575b600080fd5b6100f86100f33660046116e4565b6102fa565b6040516001600160401b0390911681526020015b60405180910390f35b610128610123366004611777565b610602565b005b61013d610138366004611777565b61068a565b6040516001600160a01b03909116815260200161010c565b60025461013d906001600160a01b031681565b610128610176366004611777565b6106ae565b6000546100f8906001600160401b031681565b6101a161019c366004611777565b61081c565b60405161010c91906117d4565b61023e6101bc366004611846565b6001602081815260009283526040928390208351808501855281546001600160a01b0390811682529382015481840152845180860190955260028201549093168452600381015491840191909152600481015460058201546006830154600790930154939493919290916001600160401b03811690600160401b900460ff1687565b60405161010c979695949392919061185f565b61026561025f366004611777565b50600190565b604051901515815260200161010c565b6101286102833660046118bc565b6108f5565b610128610296366004611960565b6109a0565b60045461013d906001600160a01b031681565b60035461013d906001600160a01b031681565b60055461013d906001600160a01b031681565b6101286102e23660046119f2565b6109c8565b6101286102f5366004611a4e565b610b39565b6002546000906001600160a01b0316331461034f5760405162461bcd60e51b815260206004820152601060248201526f13d3931657d493d313155417d0d2105360821b60448201526064015b60405180910390fd5b6040805160028082526060820183526000926020830190803683370190505090506103a561038060208b018b611af2565b6103a08a60005b6080020180360381019061039b9190611bb1565b610b61565b610be2565b816000815181106103b8576103b8611adc565b60209081029190910101526103e78960016020020160208101906103dc9190611af2565b6103a08a6001610387565b816001815181106103fa576103fa611adc565b6020908102919091010152600080548190819061041f906001600160401b0316611c5f565b91906101000a8154816001600160401b0302191690836001600160401b031602179055905060006001600160401b0316816001600160401b0316141561046757610467611c86565b6001600160401b0381166000908152600160205260408120600581018d9055906104a161049c368d90038d0160808e01611bb1565b610d0c565b905060026104b560408e0160208f01611af2565b60038111156104c6576104c66117aa565b14806104f4575060006104e96104e4368e90038e0160808f01611bb1565b610d21565b6001600160401b0316115b15610507578061050381611c5f565b9150505b6007820180546040805180820182526001600160a01b038d811680835260209283018d90526002880180546001600160a01b03199081169092179055600388018d905583518085018552918e16808352919092018b90528654909116178555600185018990554260048601556001600160401b0384811668ffffffffffffffffff1990931692909217600160401b179092559051908416907f76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a906105d1908e906080820190611ce6565b60405180910390a26105ef8360008c6001600160401b031687610d30565b5090925050505b98975050505050505050565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff16600281111561063b5761063b6117aa565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b8152509061067b5760405162461bcd60e51b81526004016103469190611d32565b50610687816000610dc6565b50565b6001600160401b03166000908152600160205260409020546001600160a01b031690565b6002546001600160a01b031633146106fb5760405162461bcd60e51b815260206004820152601060248201526f2727aa2fa922a9afa922a1a2a4ab22a960811b6044820152606401610346565b60006001600160401b038216600090815260016020526040902060070154600160401b900460ff166002811115610734576107346117aa565b1415604051806040016040528060078152602001661393d7d0d2105360ca1b815250906107745760405162461bcd60e51b81526004016103469190611d32565b506001600160401b038116600081815260016020819052604080832080546001600160a01b031990811682559281018490556002810180549093169092556003808301849055600483018490556005830184905560068301939093556007909101805468ffffffffffffffffff19169055517ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f409161081191611d65565b60405180910390a250565b61082461163f565b6001600160401b0382811660009081526001602081815260409283902083516101208101855281546001600160a01b0390811660e0830190815294830154610100830152938152845180860186526002808401549095168152600383015481850152928101929092526004810154938201939093526005830154606082015260068301546080820152600783015493841660a08201529260c0840191600160401b90910460ff16908111156108db576108db6117aa565b60028111156108ec576108ec6117aa565b90525092915050565b6001600160401b03841660009081526001602052604081208591859161091a8461068a565b6001600160a01b0316336001600160a01b0316146109685760405162461bcd60e51b815260206004820152600b60248201526a21a420a62fa9a2a72222a960a91b6044820152606401610346565b60405162461bcd60e51b815260206004820152600d60248201526c4348414c5f444541444c494e4560981b6044820152606401610346565b6001600160401b03841660009081526001602052604090208490849060029061091a8461068a565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161415610a565760405162461bcd60e51b815260206004820152602c60248201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060448201526b19195b1959d85d1958d85b1b60a21b6064820152608401610346565b6002546001600160a01b031615610a9e5760405162461bcd60e51b815260206004820152600c60248201526b1053149150511657d253925560a21b6044820152606401610346565b6001600160a01b038416610ae95760405162461bcd60e51b81526020600482015260126024820152712727afa922a9aaa62a2fa922a1a2a4ab22a960711b6044820152606401610346565b600280546001600160a01b039586166001600160a01b0319918216179091556003805494861694821694909417909355600480549285169284169290921790915560058054919093169116179055565b6001600160401b03851660009081526001602081905260409091208691869161091a8461068a565b80518051602091820151828401518051908401516040516c23b637b130b61039ba30ba329d60991b95810195909552602d850193909352604d8401919091526001600160c01b031960c091821b8116606d85015291901b166075820152600090607d015b604051602081830303815290604052805190602001209050919050565b60006001836003811115610bf857610bf86117aa565b1415610c3e576040516b213637b1b59039ba30ba329d60a11b6020820152602c8101839052604c015b604051602081830303815290604052805190602001209050610d06565b6002836003811115610c5257610c526117aa565b1415610c885760405174213637b1b59039ba30ba32961032b93937b932b21d60591b602082015260358101839052605501610c21565b6003836003811115610c9c57610c9c6117aa565b1415610ccb5760405174213637b1b59039ba30ba3296103a37b7903330b91d60591b6020820152603501610c21565b60405162461bcd60e51b815260206004820152601060248201526f4241445f424c4f434b5f53544154555360801b6044820152606401610346565b92915050565b6020810151600090815b602002015192915050565b60208101516000906001610d16565b6001821015610d4157610d41611c86565b600281511015610d5357610d53611c86565b6000610d60848484610ef4565b6001600160401b038616600081815260016020526040908190206006018390555191925082917f86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d6890610db790889088908890611d7f565b60405180910390a35050505050565b6001600160401b03821660008181526001602081905260408083206002808201805483546001600160a01b0319808216865596850188905595811690915560038301869055600480840187905560058401879055600684019690965560078301805468ffffffffffffffffff1916905590549251630357aa4960e01b8152948501959095526001600160a01b03948516602485018190529285166044850181905290949293909290911690630357aa4990606401600060405180830381600087803b158015610e9457600080fd5b505af1158015610ea8573d6000803e3d6000fd5b50505050846001600160401b03167ffdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f4085604051610ee59190611d65565b60405180910390a25050505050565b6000838383604051602001610f0b93929190611dd4565b6040516020818303038152906040528051906020012090505b9392505050565b6040805180820190915260008082526020820152815260200190600190039081610f2b575050604080518082018252600080825260209182018190528251808401909352600483529082015290915081600081518110610f8d57610f8d611adc565b6020026020010181905250610fa260006110bd565b81600181518110610fb557610fb5611adc565b6020026020010181905250610fca60006110bd565b81600281518110610fdd57610fdd611adc565b602090810291909101810191909152604080518083018252838152815180830190925280825260009282019290925261102d60408051606080820183529181019182529081526000602082015290565b604080518082018252606080825260006020808401829052845180840186528381528082018390528086018390528551610140810187528381529182018890529481018690529182018390526080820184905260a082018b905260c0820181905260e0820181905261010082015261012081018990529091906110af816110f0565b9a9950505050505050505050565b604080518082019091526000808252602082015250604080518082019091526000815263ffffffff909116602082015290565b60008082516003811115611106576111066117aa565b141561125657600061111b8360200151611360565b6111288460400151611360565b61113585606001516113e5565b60a086015160c087015160e0808901516101008a01516101208b01516040516f26b0b1b434b73290393ab73734b7339d60811b602082015260308101999099526050890197909752607088019590955260908701939093526001600160e01b031991811b821660b087015291821b811660b486015291901b1660b883015260bc82015260dc0160405160208183030381529060405290506111d9836080015161147e565b80156111eb5750826080015160400151155b156111fe57805160209091012092915050565b608083015160400151600090156112165750600160f81b5b81816112258660800151611495565b60405160200161123793929190611e16565b6040516020818303038152906040528051906020012092505050919050565b60018251600381111561126b5761126b6117aa565b14156112a35760a08201516040517026b0b1b434b732903334b734b9b432b21d60791b60208201526031810191909152605101610bc5565b6002825160038111156112b8576112b86117aa565b14156112e2576040516f26b0b1b434b7329032b93937b932b21d60811b6020820152603001610bc5565b6003825160038111156112f7576112f76117aa565b1415611321576040516f26b0b1b434b732903a37b7903330b91d60811b6020820152603001610bc5565b60405162461bcd60e51b815260206004820152600f60248201526e4241445f4d4143485f53544154555360881b6044820152606401610346565b919050565b60208101518151515160005b818110156113de578351611389906113849083611522565b61155a565b6040516b2b30b63ab29039ba30b1b59d60a11b6020820152602c810191909152604c8101849052606c0160405160208183030381529060405280519060200120925080806113d690611e5d565b91505061136c565b5050919050565b602081015160005b8251518110156114785761141d8360000151828151811061141057611410611adc565b6020026020010151611577565b6040517129ba30b1b590333930b6b29039ba30b1b59d60711b6020820152603281019190915260528101839052607201604051602081830303815290604052805190602001209150808061147090611e5d565b9150506113ed565b50919050565b805151600090158015610d06575050602001511590565b602081015160005b825151811015611478576114cd836000015182815181106114c0576114c0611adc565b60200260200101516115e7565b6040516b23bab0b9321039ba30b1b59d60a11b6020820152602c810191909152604c8101839052606c01604051602081830303815290604052805190602001209150808061151a90611e5d565b91505061149d565b6040805180820190915260008082526020820152825180518390811061154a5761154a611adc565b6020026020010151905092915050565b600081600001518260200151604051602001610bc5929190611e78565b6000611586826000015161155a565b602080840151604080860151606087015191516b29ba30b1b590333930b6b29d60a11b94810194909452602c840194909452604c8301919091526001600160e01b031960e093841b8116606c840152921b9091166070820152607401610bc5565b6000816000015182602001518360400151611605856060015161155a565b6040516b22b93937b91033bab0b9321d60a11b6020820152602c810194909452604c840192909252606c830152608c82015260ac01610bc5565b604080516101208101909152600060e082018181526101008301919091528190815260200161167e604080518082019091526000808252602082015290565b815260006020820181905260408201819052606082018190526080820181905260a09091015290565b8060408101831015610d0657600080fd5b80356001600160401b038116811461135b57600080fd5b6001600160a01b038116811461068757600080fd5b600080600080600080600080610200898b03121561170157600080fd5b883597506117128a60208b016116a7565b965061016089018a81111561172657600080fd5b60608a019650611735816116b8565b955050610180890135611747816116cf565b93506101a0890135611758816116cf565b979a96995094979396929592945050506101c0820135916101e0013590565b60006020828403121561178957600080fd5b610f24826116b8565b80516001600160a01b03168252602090810151910152565b634e487b7160e01b600052602160045260246000fd5b600381106117d0576117d06117aa565b9052565b6000610120820190506117e8828451611792565b60208301516117fa6040840182611792565b5060408301516080830152606083015160a0830152608083015160c08301526001600160401b0360a08401511660e083015260c083015161183f6101008401826117c0565b5092915050565b60006020828403121561185857600080fd5b5035919050565b610120810161186e828a611792565b61187b6040830189611792565b8660808301528560a08301528460c08301526001600160401b03841660e08301526105f66101008301846117c0565b60006080828403121561147857600080fd5b600080600080606085870312156118d257600080fd5b6118db856116b8565b935060208501356001600160401b03808211156118f757600080fd5b611903888389016118aa565b9450604087013591508082111561191957600080fd5b818701915087601f83011261192d57600080fd5b81358181111561193c57600080fd5b8860208260051b850101111561195157600080fd5b95989497505060200194505050565b6000806000806060858703121561197657600080fd5b61197f856116b8565b935060208501356001600160401b038082111561199b57600080fd5b6119a7888389016118aa565b945060408701359150808211156119bd57600080fd5b818701915087601f8301126119d157600080fd5b8135818111156119e057600080fd5b88602082850101111561195157600080fd5b60008060008060808587031215611a0857600080fd5b8435611a13816116cf565b93506020850135611a23816116cf565b92506040850135611a33816116cf565b91506060850135611a43816116cf565b939692955090935050565b600080600080600060e08688031215611a6657600080fd5b611a6f866116b8565b945060208601356001600160401b03811115611a8a57600080fd5b611a96888289016118aa565b945050611aa687604088016116a7565b9250611ab587608088016116a7565b9497939650919460c0013592915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060208284031215611b0457600080fd5b813560048110610f2457600080fd5b604080519081016001600160401b0381118282101715611b3557611b35611ac6565b60405290565b600082601f830112611b4c57600080fd5b604051604081018181106001600160401b0382111715611b6e57611b6e611ac6565b8060405250806040840185811115611b8557600080fd5b845b81811015611ba657611b98816116b8565b835260209283019201611b87565b509195945050505050565b600060808284031215611bc357600080fd5b604051604081018181106001600160401b0382111715611be557611be5611ac6565b604052601f83018413611bf757600080fd5b611bff611b13565b806040850186811115611c1157600080fd5b855b81811015611c2b578035845260209384019301611c13565b50818452611c398782611b3b565b6020850152509195945050505050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380831681811415611c7c57611c7c611c49565b6001019392505050565b634e487b7160e01b600052600160045260246000fd5b604081833760006040838101828152908301915b6002811015611cdf576001600160401b03611cca846116b8565b16825260209283019290910190600101611cb0565b5050505050565b6101008101611cf58285611c9c565b610f246080830184611c9c565b60005b83811015611d1d578181015183820152602001611d05565b83811115611d2c576000848401525b50505050565b6020815260008251806020840152611d51816040850160208701611d02565b601f01601f19169190910160400192915050565b6020810160048310611d7957611d796117aa565b91905290565b6000606082018583526020858185015260606040850152818551808452608086019150828701935060005b81811015611dc657845183529383019391830191600101611daa565b509098975050505050505050565b83815260006020848184015260408301845182860160005b82811015611e0857815184529284019290840190600101611dec565b509198975050505050505050565b60008451611e28818460208901611d02565b6b2bb4ba341033bab0b932399d60a11b9201918252506001600160f81b031992909216600c830152600d820152602d01919050565b6000600019821415611e7157611e71611c49565b5060010190565b652b30b63ab29d60d11b8152600060078410611e9657611e966117aa565b5060f89290921b600683015260078201526027019056fea2646970667358221220af36d290cfdd8d3d3c36beaa5d4af8772e0299ec637be14a28a36cddc52d65f964736f6c63430008090033",
}

// TimedOutChallengeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TimedOutChallengeManagerMetaData.ABI instead.
var TimedOutChallengeManagerABI = TimedOutChallengeManagerMetaData.ABI

// TimedOutChallengeManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TimedOutChallengeManagerMetaData.Bin instead.
var TimedOutChallengeManagerBin = TimedOutChallengeManagerMetaData.Bin

// DeployTimedOutChallengeManager deploys a new Ethereum contract, binding an instance of TimedOutChallengeManager to it.
func DeployTimedOutChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TimedOutChallengeManager, error) {
	parsed, err := TimedOutChallengeManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TimedOutChallengeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TimedOutChallengeManager{TimedOutChallengeManagerCaller: TimedOutChallengeManagerCaller{contract: contract}, TimedOutChallengeManagerTransactor: TimedOutChallengeManagerTransactor{contract: contract}, TimedOutChallengeManagerFilterer: TimedOutChallengeManagerFilterer{contract: contract}}, nil
}

// TimedOutChallengeManager is an auto generated Go binding around an Ethereum contract.
type TimedOutChallengeManager struct {
	TimedOutChallengeManagerCaller     // Read-only binding to the contract
	TimedOutChallengeManagerTransactor // Write-only binding to the contract
	TimedOutChallengeManagerFilterer   // Log filterer for contract events
}

// TimedOutChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimedOutChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedOutChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimedOutChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedOutChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimedOutChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedOutChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimedOutChallengeManagerSession struct {
	Contract     *TimedOutChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TimedOutChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimedOutChallengeManagerCallerSession struct {
	Contract *TimedOutChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// TimedOutChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimedOutChallengeManagerTransactorSession struct {
	Contract     *TimedOutChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// TimedOutChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimedOutChallengeManagerRaw struct {
	Contract *TimedOutChallengeManager // Generic contract binding to access the raw methods on
}

// TimedOutChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimedOutChallengeManagerCallerRaw struct {
	Contract *TimedOutChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TimedOutChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimedOutChallengeManagerTransactorRaw struct {
	Contract *TimedOutChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimedOutChallengeManager creates a new instance of TimedOutChallengeManager, bound to a specific deployed contract.
func NewTimedOutChallengeManager(address common.Address, backend bind.ContractBackend) (*TimedOutChallengeManager, error) {
	contract, err := bindTimedOutChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManager{TimedOutChallengeManagerCaller: TimedOutChallengeManagerCaller{contract: contract}, TimedOutChallengeManagerTransactor: TimedOutChallengeManagerTransactor{contract: contract}, TimedOutChallengeManagerFilterer: TimedOutChallengeManagerFilterer{contract: contract}}, nil
}

// NewTimedOutChallengeManagerCaller creates a new read-only instance of TimedOutChallengeManager, bound to a specific deployed contract.
func NewTimedOutChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*TimedOutChallengeManagerCaller, error) {
	contract, err := bindTimedOutChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManagerCaller{contract: contract}, nil
}

// NewTimedOutChallengeManagerTransactor creates a new write-only instance of TimedOutChallengeManager, bound to a specific deployed contract.
func NewTimedOutChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TimedOutChallengeManagerTransactor, error) {
	contract, err := bindTimedOutChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManagerTransactor{contract: contract}, nil
}

// NewTimedOutChallengeManagerFilterer creates a new log filterer instance of TimedOutChallengeManager, bound to a specific deployed contract.
func NewTimedOutChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TimedOutChallengeManagerFilterer, error) {
	contract, err := bindTimedOutChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManagerFilterer{contract: contract}, nil
}

// bindTimedOutChallengeManager binds a generic wrapper to an already deployed contract.
func bindTimedOutChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimedOutChallengeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedOutChallengeManager *TimedOutChallengeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimedOutChallengeManager.Contract.TimedOutChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedOutChallengeManager *TimedOutChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.TimedOutChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedOutChallengeManager *TimedOutChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.TimedOutChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimedOutChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) Bridge() (common.Address, error) {
	return _TimedOutChallengeManager.Contract.Bridge(&_TimedOutChallengeManager.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) Bridge() (common.Address, error) {
	return _TimedOutChallengeManager.Contract.Bridge(&_TimedOutChallengeManager.CallOpts)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) ChallengeInfo(opts *bind.CallOpts, challengeIndex uint64) (ChallengeLibChallenge, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "challengeInfo", challengeIndex)

	if err != nil {
		return *new(ChallengeLibChallenge), err
	}

	out0 := *abi.ConvertType(out[0], new(ChallengeLibChallenge)).(*ChallengeLibChallenge)

	return out0, err

}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _TimedOutChallengeManager.Contract.ChallengeInfo(&_TimedOutChallengeManager.CallOpts, challengeIndex)
}

// ChallengeInfo is a free data retrieval call binding the contract method 0x7fd07a9c.
//
// Solidity: function challengeInfo(uint64 challengeIndex) view returns(((address,uint256),(address,uint256),uint256,bytes32,bytes32,uint64,uint8))
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) ChallengeInfo(challengeIndex uint64) (ChallengeLibChallenge, error) {
	return _TimedOutChallengeManager.Contract.ChallengeInfo(&_TimedOutChallengeManager.CallOpts, challengeIndex)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "challenges", arg0)

	outstruct := new(struct {
		Current            ChallengeLibParticipant
		Next               ChallengeLibParticipant
		LastMoveTimestamp  *big.Int
		WasmModuleRoot     [32]byte
		ChallengeStateHash [32]byte
		MaxInboxMessages   uint64
		Mode               uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(ChallengeLibParticipant)).(*ChallengeLibParticipant)
	outstruct.Next = *abi.ConvertType(out[1], new(ChallengeLibParticipant)).(*ChallengeLibParticipant)
	outstruct.LastMoveTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WasmModuleRoot = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.ChallengeStateHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.MaxInboxMessages = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.Mode = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _TimedOutChallengeManager.Contract.Challenges(&_TimedOutChallengeManager.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns((address,uint256) current, (address,uint256) next, uint256 lastMoveTimestamp, bytes32 wasmModuleRoot, bytes32 challengeStateHash, uint64 maxInboxMessages, uint8 mode)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) Challenges(arg0 *big.Int) (struct {
	Current            ChallengeLibParticipant
	Next               ChallengeLibParticipant
	LastMoveTimestamp  *big.Int
	WasmModuleRoot     [32]byte
	ChallengeStateHash [32]byte
	MaxInboxMessages   uint64
	Mode               uint8
}, error) {
	return _TimedOutChallengeManager.Contract.Challenges(&_TimedOutChallengeManager.CallOpts, arg0)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) CurrentResponder(opts *bind.CallOpts, challengeIndex uint64) (common.Address, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "currentResponder", challengeIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _TimedOutChallengeManager.Contract.CurrentResponder(&_TimedOutChallengeManager.CallOpts, challengeIndex)
}

// CurrentResponder is a free data retrieval call binding the contract method 0x23a9ef23.
//
// Solidity: function currentResponder(uint64 challengeIndex) view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) CurrentResponder(challengeIndex uint64) (common.Address, error) {
	return _TimedOutChallengeManager.Contract.CurrentResponder(&_TimedOutChallengeManager.CallOpts, challengeIndex)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 ) pure returns(bool)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) IsTimedOut(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "isTimedOut", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 ) pure returns(bool)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) IsTimedOut(arg0 uint64) (bool, error) {
	return _TimedOutChallengeManager.Contract.IsTimedOut(&_TimedOutChallengeManager.CallOpts, arg0)
}

// IsTimedOut is a free data retrieval call binding the contract method 0x9ede42b9.
//
// Solidity: function isTimedOut(uint64 ) pure returns(bool)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) IsTimedOut(arg0 uint64) (bool, error) {
	return _TimedOutChallengeManager.Contract.IsTimedOut(&_TimedOutChallengeManager.CallOpts, arg0)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) Osp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "osp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) Osp() (common.Address, error) {
	return _TimedOutChallengeManager.Contract.Osp(&_TimedOutChallengeManager.CallOpts)
}

// Osp is a free data retrieval call binding the contract method 0xf26a62c6.
//
// Solidity: function osp() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) Osp() (common.Address, error) {
	return _TimedOutChallengeManager.Contract.Osp(&_TimedOutChallengeManager.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) ResultReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "resultReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) ResultReceiver() (common.Address, error) {
	return _TimedOutChallengeManager.Contract.ResultReceiver(&_TimedOutChallengeManager.CallOpts)
}

// ResultReceiver is a free data retrieval call binding the contract method 0x3504f1d7.
//
// Solidity: function resultReceiver() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) ResultReceiver() (common.Address, error) {
	return _TimedOutChallengeManager.Contract.ResultReceiver(&_TimedOutChallengeManager.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) SequencerInbox() (common.Address, error) {
	return _TimedOutChallengeManager.Contract.SequencerInbox(&_TimedOutChallengeManager.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) SequencerInbox() (common.Address, error) {
	return _TimedOutChallengeManager.Contract.SequencerInbox(&_TimedOutChallengeManager.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCaller) TotalChallengesCreated(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TimedOutChallengeManager.contract.Call(opts, &out, "totalChallengesCreated")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) TotalChallengesCreated() (uint64, error) {
	return _TimedOutChallengeManager.Contract.TotalChallengesCreated(&_TimedOutChallengeManager.CallOpts)
}

// TotalChallengesCreated is a free data retrieval call binding the contract method 0x5ef489e6.
//
// Solidity: function totalChallengesCreated() view returns(uint64)
func (_TimedOutChallengeManager *TimedOutChallengeManagerCallerSession) TotalChallengesCreated() (uint64, error) {
	return _TimedOutChallengeManager.Contract.TotalChallengesCreated(&_TimedOutChallengeManager.CallOpts)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactor) BisectExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _TimedOutChallengeManager.contract.Transact(opts, "bisectExecution", challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.BisectExecution(&_TimedOutChallengeManager.TransactOpts, challengeIndex, selection, newSegments)
}

// BisectExecution is a paid mutator transaction binding the contract method 0xa521b032.
//
// Solidity: function bisectExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes32[] newSegments) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorSession) BisectExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, newSegments [][32]byte) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.BisectExecution(&_TimedOutChallengeManager.TransactOpts, challengeIndex, selection, newSegments)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactor) ChallengeExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _TimedOutChallengeManager.contract.Transact(opts, "challengeExecution", challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.ChallengeExecution(&_TimedOutChallengeManager.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ChallengeExecution is a paid mutator transaction binding the contract method 0xfb7be0a1.
//
// Solidity: function challengeExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, uint8[2] machineStatuses, bytes32[2] globalStateHashes, uint256 numSteps) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorSession) ChallengeExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, machineStatuses [2]uint8, globalStateHashes [2][32]byte, numSteps *big.Int) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.ChallengeExecution(&_TimedOutChallengeManager.TransactOpts, challengeIndex, selection, machineStatuses, globalStateHashes, numSteps)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactor) ClearChallenge(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _TimedOutChallengeManager.contract.Transact(opts, "clearChallenge", challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.ClearChallenge(&_TimedOutChallengeManager.TransactOpts, challengeIndex)
}

// ClearChallenge is a paid mutator transaction binding the contract method 0x56e9df97.
//
// Solidity: function clearChallenge(uint64 challengeIndex) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorSession) ClearChallenge(challengeIndex uint64) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.ClearChallenge(&_TimedOutChallengeManager.TransactOpts, challengeIndex)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactor) CreateChallenge(opts *bind.TransactOpts, wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _TimedOutChallengeManager.contract.Transact(opts, "createChallenge", wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.CreateChallenge(&_TimedOutChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0x14eab5e7.
//
// Solidity: function createChallenge(bytes32 wasmModuleRoot_, uint8[2] startAndEndMachineStatuses_, (bytes32[2],uint64[2])[2] startAndEndGlobalStates_, uint64 numBlocks, address asserter_, address challenger_, uint256 asserterTimeLeft_, uint256 challengerTimeLeft_) returns(uint64)
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorSession) CreateChallenge(wasmModuleRoot_ [32]byte, startAndEndMachineStatuses_ [2]uint8, startAndEndGlobalStates_ [2]GlobalState, numBlocks uint64, asserter_ common.Address, challenger_ common.Address, asserterTimeLeft_ *big.Int, challengerTimeLeft_ *big.Int) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.CreateChallenge(&_TimedOutChallengeManager.TransactOpts, wasmModuleRoot_, startAndEndMachineStatuses_, startAndEndGlobalStates_, numBlocks, asserter_, challenger_, asserterTimeLeft_, challengerTimeLeft_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactor) Initialize(opts *bind.TransactOpts, resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _TimedOutChallengeManager.contract.Transact(opts, "initialize", resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.Initialize(&_TimedOutChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address resultReceiver_, address sequencerInbox_, address bridge_, address osp_) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorSession) Initialize(resultReceiver_ common.Address, sequencerInbox_ common.Address, bridge_ common.Address, osp_ common.Address) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.Initialize(&_TimedOutChallengeManager.TransactOpts, resultReceiver_, sequencerInbox_, bridge_, osp_)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactor) OneStepProveExecution(opts *bind.TransactOpts, challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _TimedOutChallengeManager.contract.Transact(opts, "oneStepProveExecution", challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.OneStepProveExecution(&_TimedOutChallengeManager.TransactOpts, challengeIndex, selection, proof)
}

// OneStepProveExecution is a paid mutator transaction binding the contract method 0xd248d124.
//
// Solidity: function oneStepProveExecution(uint64 challengeIndex, (uint256,uint256,bytes32[],uint256) selection, bytes proof) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorSession) OneStepProveExecution(challengeIndex uint64, selection ChallengeLibSegmentSelection, proof []byte) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.OneStepProveExecution(&_TimedOutChallengeManager.TransactOpts, challengeIndex, selection, proof)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactor) Timeout(opts *bind.TransactOpts, challengeIndex uint64) (*types.Transaction, error) {
	return _TimedOutChallengeManager.contract.Transact(opts, "timeout", challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.Timeout(&_TimedOutChallengeManager.TransactOpts, challengeIndex)
}

// Timeout is a paid mutator transaction binding the contract method 0x1b45c86a.
//
// Solidity: function timeout(uint64 challengeIndex) returns()
func (_TimedOutChallengeManager *TimedOutChallengeManagerTransactorSession) Timeout(challengeIndex uint64) (*types.Transaction, error) {
	return _TimedOutChallengeManager.Contract.Timeout(&_TimedOutChallengeManager.TransactOpts, challengeIndex)
}

// TimedOutChallengeManagerBisectedIterator is returned from FilterBisected and is used to iterate over the raw logs and unpacked data for Bisected events raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerBisectedIterator struct {
	Event *TimedOutChallengeManagerBisected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TimedOutChallengeManagerBisectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedOutChallengeManagerBisected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TimedOutChallengeManagerBisected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TimedOutChallengeManagerBisectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedOutChallengeManagerBisectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedOutChallengeManagerBisected represents a Bisected event raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerBisected struct {
	ChallengeIndex          uint64
	ChallengeRoot           [32]byte
	ChallengedSegmentStart  *big.Int
	ChallengedSegmentLength *big.Int
	ChainHashes             [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBisected is a free log retrieval operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) FilterBisected(opts *bind.FilterOpts, challengeIndex []uint64, challengeRoot [][32]byte) (*TimedOutChallengeManagerBisectedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.FilterLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManagerBisectedIterator{contract: _TimedOutChallengeManager.contract, event: "Bisected", logs: logs, sub: sub}, nil
}

// WatchBisected is a free log subscription operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) WatchBisected(opts *bind.WatchOpts, sink chan<- *TimedOutChallengeManagerBisected, challengeIndex []uint64, challengeRoot [][32]byte) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}
	var challengeRootRule []interface{}
	for _, challengeRootItem := range challengeRoot {
		challengeRootRule = append(challengeRootRule, challengeRootItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.WatchLogs(opts, "Bisected", challengeIndexRule, challengeRootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedOutChallengeManagerBisected)
				if err := _TimedOutChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBisected is a log parse operation binding the contract event 0x86b34e9455464834eca718f62d4481437603bb929d8a78ccde5d1bc79fa06d68.
//
// Solidity: event Bisected(uint64 indexed challengeIndex, bytes32 indexed challengeRoot, uint256 challengedSegmentStart, uint256 challengedSegmentLength, bytes32[] chainHashes)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) ParseBisected(log types.Log) (*TimedOutChallengeManagerBisected, error) {
	event := new(TimedOutChallengeManagerBisected)
	if err := _TimedOutChallengeManager.contract.UnpackLog(event, "Bisected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedOutChallengeManagerChallengeEndedIterator is returned from FilterChallengeEnded and is used to iterate over the raw logs and unpacked data for ChallengeEnded events raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerChallengeEndedIterator struct {
	Event *TimedOutChallengeManagerChallengeEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TimedOutChallengeManagerChallengeEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedOutChallengeManagerChallengeEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TimedOutChallengeManagerChallengeEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TimedOutChallengeManagerChallengeEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedOutChallengeManagerChallengeEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedOutChallengeManagerChallengeEnded represents a ChallengeEnded event raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerChallengeEnded struct {
	ChallengeIndex uint64
	Kind           uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChallengeEnded is a free log retrieval operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) FilterChallengeEnded(opts *bind.FilterOpts, challengeIndex []uint64) (*TimedOutChallengeManagerChallengeEndedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.FilterLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManagerChallengeEndedIterator{contract: _TimedOutChallengeManager.contract, event: "ChallengeEnded", logs: logs, sub: sub}, nil
}

// WatchChallengeEnded is a free log subscription operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) WatchChallengeEnded(opts *bind.WatchOpts, sink chan<- *TimedOutChallengeManagerChallengeEnded, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.WatchLogs(opts, "ChallengeEnded", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedOutChallengeManagerChallengeEnded)
				if err := _TimedOutChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeEnded is a log parse operation binding the contract event 0xfdaece6c274a4b56af16761f83fd6b1062823192630ea08e019fdf9b2d747f40.
//
// Solidity: event ChallengeEnded(uint64 indexed challengeIndex, uint8 kind)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) ParseChallengeEnded(log types.Log) (*TimedOutChallengeManagerChallengeEnded, error) {
	event := new(TimedOutChallengeManagerChallengeEnded)
	if err := _TimedOutChallengeManager.contract.UnpackLog(event, "ChallengeEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedOutChallengeManagerExecutionChallengeBegunIterator is returned from FilterExecutionChallengeBegun and is used to iterate over the raw logs and unpacked data for ExecutionChallengeBegun events raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerExecutionChallengeBegunIterator struct {
	Event *TimedOutChallengeManagerExecutionChallengeBegun // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TimedOutChallengeManagerExecutionChallengeBegunIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedOutChallengeManagerExecutionChallengeBegun)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TimedOutChallengeManagerExecutionChallengeBegun)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TimedOutChallengeManagerExecutionChallengeBegunIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedOutChallengeManagerExecutionChallengeBegunIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedOutChallengeManagerExecutionChallengeBegun represents a ExecutionChallengeBegun event raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerExecutionChallengeBegun struct {
	ChallengeIndex uint64
	BlockSteps     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionChallengeBegun is a free log retrieval operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) FilterExecutionChallengeBegun(opts *bind.FilterOpts, challengeIndex []uint64) (*TimedOutChallengeManagerExecutionChallengeBegunIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.FilterLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManagerExecutionChallengeBegunIterator{contract: _TimedOutChallengeManager.contract, event: "ExecutionChallengeBegun", logs: logs, sub: sub}, nil
}

// WatchExecutionChallengeBegun is a free log subscription operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) WatchExecutionChallengeBegun(opts *bind.WatchOpts, sink chan<- *TimedOutChallengeManagerExecutionChallengeBegun, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.WatchLogs(opts, "ExecutionChallengeBegun", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedOutChallengeManagerExecutionChallengeBegun)
				if err := _TimedOutChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionChallengeBegun is a log parse operation binding the contract event 0x24e032e170243bbea97e140174b22dc7e54fb85925afbf52c70e001cd6af16db.
//
// Solidity: event ExecutionChallengeBegun(uint64 indexed challengeIndex, uint256 blockSteps)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) ParseExecutionChallengeBegun(log types.Log) (*TimedOutChallengeManagerExecutionChallengeBegun, error) {
	event := new(TimedOutChallengeManagerExecutionChallengeBegun)
	if err := _TimedOutChallengeManager.contract.UnpackLog(event, "ExecutionChallengeBegun", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedOutChallengeManagerInitiatedChallengeIterator is returned from FilterInitiatedChallenge and is used to iterate over the raw logs and unpacked data for InitiatedChallenge events raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerInitiatedChallengeIterator struct {
	Event *TimedOutChallengeManagerInitiatedChallenge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TimedOutChallengeManagerInitiatedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedOutChallengeManagerInitiatedChallenge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TimedOutChallengeManagerInitiatedChallenge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TimedOutChallengeManagerInitiatedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedOutChallengeManagerInitiatedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedOutChallengeManagerInitiatedChallenge represents a InitiatedChallenge event raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerInitiatedChallenge struct {
	ChallengeIndex uint64
	StartState     GlobalState
	EndState       GlobalState
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiatedChallenge is a free log retrieval operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) FilterInitiatedChallenge(opts *bind.FilterOpts, challengeIndex []uint64) (*TimedOutChallengeManagerInitiatedChallengeIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.FilterLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManagerInitiatedChallengeIterator{contract: _TimedOutChallengeManager.contract, event: "InitiatedChallenge", logs: logs, sub: sub}, nil
}

// WatchInitiatedChallenge is a free log subscription operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) WatchInitiatedChallenge(opts *bind.WatchOpts, sink chan<- *TimedOutChallengeManagerInitiatedChallenge, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.WatchLogs(opts, "InitiatedChallenge", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedOutChallengeManagerInitiatedChallenge)
				if err := _TimedOutChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitiatedChallenge is a log parse operation binding the contract event 0x76604fe17af46c9b5f53ffe99ff23e0f655dab91886b07ac1fc0254319f7145a.
//
// Solidity: event InitiatedChallenge(uint64 indexed challengeIndex, (bytes32[2],uint64[2]) startState, (bytes32[2],uint64[2]) endState)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) ParseInitiatedChallenge(log types.Log) (*TimedOutChallengeManagerInitiatedChallenge, error) {
	event := new(TimedOutChallengeManagerInitiatedChallenge)
	if err := _TimedOutChallengeManager.contract.UnpackLog(event, "InitiatedChallenge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedOutChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerOneStepProofCompletedIterator struct {
	Event *TimedOutChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TimedOutChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedOutChallengeManagerOneStepProofCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TimedOutChallengeManagerOneStepProofCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TimedOutChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedOutChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedOutChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the TimedOutChallengeManager contract.
type TimedOutChallengeManagerOneStepProofCompleted struct {
	ChallengeIndex uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, challengeIndex []uint64) (*TimedOutChallengeManagerOneStepProofCompletedIterator, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return &TimedOutChallengeManagerOneStepProofCompletedIterator{contract: _TimedOutChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *TimedOutChallengeManagerOneStepProofCompleted, challengeIndex []uint64) (event.Subscription, error) {

	var challengeIndexRule []interface{}
	for _, challengeIndexItem := range challengeIndex {
		challengeIndexRule = append(challengeIndexRule, challengeIndexItem)
	}

	logs, sub, err := _TimedOutChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", challengeIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedOutChallengeManagerOneStepProofCompleted)
				if err := _TimedOutChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xc2cc42e04ff8c36de71c6a2937ea9f161dd0dd9e175f00caa26e5200643c781e.
//
// Solidity: event OneStepProofCompleted(uint64 indexed challengeIndex)
func (_TimedOutChallengeManager *TimedOutChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*TimedOutChallengeManagerOneStepProofCompleted, error) {
	event := new(TimedOutChallengeManagerOneStepProofCompleted)
	if err := _TimedOutChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
