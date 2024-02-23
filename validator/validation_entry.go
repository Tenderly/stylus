package validator

import (
	"github.com/tenderly/stylus/go-ethereum/common"
	"github.com/tenderly/stylus/go-ethereum/core/state"
)

type BatchInfo struct {
	Number uint64
	Data   []byte
}

type ValidationInput struct {
	Id            uint64
	HasDelayedMsg bool
	DelayedMsgNr  uint64
	Preimages     map[common.Hash][]byte
	UserWasms     state.UserWasms
	BatchInfo     []BatchInfo
	DelayedMsg    []byte
	StartState    GoGlobalState
	DebugChain    bool
}
