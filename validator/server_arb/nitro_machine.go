// Copyright 2021-2023, Offchain Labs, Inc.
// For license information, see https://github.com/tenderly/stylus/blob/master/LICENSE

package server_arb

/*
#cgo CFLAGS: -g -Wall -I../../target/include/
#include "arbitrator.h"
#include <stdlib.h>
*/
import "C"
import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/tenderly/stylus/go-ethereum/common"
	"github.com/tenderly/stylus/go-ethereum/log"
	"github.com/tenderly/stylus/validator/server_common"
)

func createArbMachine(ctx context.Context, locator *server_common.MachineLocator, config *ArbitratorMachineConfig, moduleRoot common.Hash) (*arbMachines, error) {
	binPath := filepath.Join(locator.GetMachinePath(moduleRoot), config.WavmBinaryPath)
	cBinPath := C.CString(binPath)
	defer C.free(unsafe.Pointer(cBinPath))
	log.Info("creating nitro machine", "binpath", binPath)

	result := &arbMachines{}
	result.zeroStep.Freeze()

	// We try to store/load state before first host_io to a file.
	// We will chicken out of that if something fails, but still try to calculate the machine
	statePath := filepath.Join(locator.GetMachinePath(moduleRoot), config.UntilHostIoStatePath)
	_, err := os.Stat(statePath)
	if err == nil {
		log.Info("found cached machine until host io state", "moduleRoot", moduleRoot)

		result.hostIo.Freeze()
		return result, nil

	} else if errors.Is(err, os.ErrNotExist) {
		log.Info("didn't find cached machine until host io state", "path", statePath)
	} else {
		log.Warn("error checking if machine until host io state is cached", "path", statePath, "err", err)
	}

	result.hostIo.Freeze()
	return result, nil
}
