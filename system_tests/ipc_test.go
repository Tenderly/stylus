// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package arbtest

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/tenderly/stylus/cmd/genericconf"
	"github.com/tenderly/stylus/go-ethereum/ethclient"
)

func TestIpcRpc(t *testing.T) {
	ipcPath := filepath.Join(t.TempDir(), "test.ipc")

	ipcConfig := genericconf.IPCConfigDefault
	ipcConfig.Path = ipcPath

	stackConf := stackConfigForTest(t)
	ipcConfig.Apply(stackConf)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, l2node, _, _, _, _, l1stack := createTestNodeOnL1WithConfig(t, ctx, true, nil, nil, stackConf)
	defer requireClose(t, l1stack)
	defer l2node.StopAndWait()

	_, err := ethclient.Dial(ipcPath)
	Require(t, err)
}
