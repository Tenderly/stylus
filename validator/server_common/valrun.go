package server_common

import (
	"github.com/tenderly/stylus/go-ethereum/common"
	"github.com/tenderly/stylus/util/containers"
	"github.com/tenderly/stylus/validator"
)

type ValRun struct {
	containers.PromiseInterface[validator.GoGlobalState]
	root common.Hash
}

func (r *ValRun) WasmModuleRoot() common.Hash {
	return r.root
}

func NewValRun(promise containers.PromiseInterface[validator.GoGlobalState], root common.Hash) *ValRun {
	return &ValRun{
		PromiseInterface: promise,
		root:             root,
	}
}
