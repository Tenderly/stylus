package state

import "github.com/tenderly/stylus/go-ethereum/common"

type wasmActivation struct {
	moduleHash common.Hash
}

func (ch wasmActivation) revert(s *StateDB) {
	delete(s.activatedWasms, ch.moduleHash)
}

func (ch wasmActivation) dirtied() *common.Address {
	return nil
}
