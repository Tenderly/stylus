// Copyright 2023-2024, Offchain Labs, Inc.
// For license information, see https://github.com/tenderly/stylus/blob/master/LICENSE

package programs

import (
	"errors"
	"math/big"

	"github.com/holiman/uint256"
	"github.com/tenderly/stylus/arbos/util"
	"github.com/tenderly/stylus/go-ethereum/common"
	"github.com/tenderly/stylus/go-ethereum/core/types"
	"github.com/tenderly/stylus/go-ethereum/core/vm"
	"github.com/tenderly/stylus/go-ethereum/log"
	"github.com/tenderly/stylus/go-ethereum/params"
	am "github.com/tenderly/stylus/util/arbmath"
)

type getBytes32Type func(key common.Hash) (value common.Hash, cost uint64)
type setBytes32Type func(key, value common.Hash) (cost uint64, err error)
type contractCallType func(
	contract common.Address, calldata []byte, gas uint64, value *big.Int) (
	retdata_len uint32, cost uint64, err error,
)
type delegateCallType func(
	contract common.Address, calldata []byte, gas uint64) (
	retdata_len uint32, cost uint64, err error,
)
type staticCallType func(
	contract common.Address, calldata []byte, gas uint64) (
	retdata_len uint32, cost uint64, err error,
)
type create1Type func(
	code []byte, endowment *big.Int, gas uint64) (
	addr common.Address, retdata_len uint32, cost uint64, err error,
)
type create2Type func(
	code []byte, salt, endowment *big.Int, gas uint64) (
	addr common.Address, retdata_len uint32, cost uint64, err error,
)
type getReturnDataType func(offset uint32, size uint32) []byte
type emitLogType func(data []byte, topics uint32) error
type accountBalanceType func(address common.Address) (value common.Hash, cost uint64)
type accountCodeType func(address common.Address, offset, size uint32, gas uint64) ([]byte, uint64)
type accountCodeSizeType func(address common.Address, gas uint64) (uint32, uint64)
type accountCodehashType func(address common.Address) (value common.Hash, cost uint64)
type addPagesType func(pages uint16) (cost uint64)
type captureHostioType func(name string, args, outs []byte, startInk, endInk uint64)

type goClosures struct {
	getBytes32      getBytes32Type
	setBytes32      setBytes32Type
	contractCall    contractCallType
	delegateCall    delegateCallType
	staticCall      staticCallType
	create1         create1Type
	create2         create2Type
	getReturnData   getReturnDataType
	emitLog         emitLogType
	accountBalance  accountBalanceType
	accountCode     accountCodeType
	accountCodeSize accountCodeSizeType
	accountCodeHash accountCodehashType
	addPages        addPagesType
	captureHostio   captureHostioType
}

func newApiClosures(
	interpreter *vm.EVMInterpreter,
	tracingInfo *util.TracingInfo,
	scope *vm.ScopeContext,
	memoryModel *MemoryModel,
) *goClosures {
	contract := scope.Contract
	actingAddress := contract.Address() // not necessarily WASM
	readOnly := interpreter.ReadOnly()
	evm := interpreter.Evm()
	depth := evm.Depth()
	db := evm.StateDB

	getBytes32 := func(key common.Hash) (common.Hash, uint64) {
		if tracingInfo != nil {
			tracingInfo.RecordStorageGet(key)
		}
		cost := vm.WasmStateLoadCost(db, actingAddress, key)
		return db.GetState(actingAddress, key), cost
	}
	setBytes32 := func(key, value common.Hash) (uint64, error) {
		if tracingInfo != nil {
			tracingInfo.RecordStorageSet(key, value)
		}
		if readOnly {
			return 0, vm.ErrWriteProtection
		}
		cost := vm.WasmStateStoreCost(db, actingAddress, key, value)
		db.SetState(actingAddress, key, value)
		return cost, nil
	}
	doCall := func(
		contract common.Address, opcode vm.OpCode, input []byte, gas uint64, value *big.Int,
	) (uint32, uint64, error) {
		// This closure can perform each kind of contract call based on the opcode passed in.
		// The implementation for each should match that of the EVM.
		//
		// Note that while the Yellow Paper is authoritative, the following go-ethereum
		// functions provide corresponding implementations in the vm package.
		//     - operations_acl.go makeCallVariantGasCallEIP2929()
		//     - gas_table.go      gasCall() gasDelegateCall() gasStaticCall()
		//     - instructions.go   opCall()  opDelegateCall()  opStaticCall()
		//

		// read-only calls are not payable (opCall)
		if readOnly && value.Sign() != 0 {
			return 0, 0, vm.ErrWriteProtection
		}

		startGas := gas

		// computes makeCallVariantGasCallEIP2929 and gasCall/gasDelegateCall/gasStaticCall
		baseCost, err := vm.WasmCallCost(db, contract, value, startGas)
		if err != nil {
			return 0, gas, err
		}
		gas -= baseCost

		// apply the 63/64ths rule
		one64th := gas / 64
		gas -= one64th

		// Tracing: emit the call (value transfer is done later in evm.Call)
		if tracingInfo != nil {
			tracingInfo.Tracer.CaptureState(0, opcode, startGas, baseCost+gas, scope, []byte{}, depth, nil)
		}

		// EVM rule: calls that pay get a stipend (opCall)
		if value.Sign() != 0 {
			gas = am.SaturatingUAdd(gas, params.CallStipend)
		}

		var ret []byte
		var returnGas uint64

		switch opcode {
		case vm.CALL:
			ret, returnGas, err = evm.Call(scope.Contract, contract, input, gas, value)
		case vm.DELEGATECALL:
			ret, returnGas, err = evm.DelegateCall(scope.Contract, contract, input, gas)
		case vm.STATICCALL:
			ret, returnGas, err = evm.StaticCall(scope.Contract, contract, input, gas)
		default:
			log.Crit("unsupported call type", "opcode", opcode)
		}

		interpreter.SetReturnData(ret)
		cost := am.SaturatingUSub(startGas, returnGas+one64th) // user gets 1/64th back
		return uint32(len(ret)), cost, err
	}
	contractCall := func(contract common.Address, input []byte, gas uint64, value *big.Int) (uint32, uint64, error) {
		return doCall(contract, vm.CALL, input, gas, value)
	}
	delegateCall := func(contract common.Address, input []byte, gas uint64) (uint32, uint64, error) {
		return doCall(contract, vm.DELEGATECALL, input, gas, common.Big0)
	}
	staticCall := func(contract common.Address, input []byte, gas uint64) (uint32, uint64, error) {
		return doCall(contract, vm.STATICCALL, input, gas, common.Big0)
	}
	create := func(code []byte, endowment, salt *big.Int, gas uint64) (common.Address, uint32, uint64, error) {
		// This closure can perform both kinds of contract creation based on the salt passed in.
		// The implementation for each should match that of the EVM.
		//
		// Note that while the Yellow Paper is authoritative, the following go-ethereum
		// functions provide corresponding implementations in the vm package.
		//     - instructions.go opCreate() opCreate2()
		//     - gas_table.go    gasCreate() gasCreate2()
		//

		opcode := vm.CREATE
		if salt != nil {
			opcode = vm.CREATE2
		}
		zeroAddr := common.Address{}
		startGas := gas

		if readOnly {
			return zeroAddr, 0, 0, vm.ErrWriteProtection
		}

		// pay for static and dynamic costs (gasCreate and gasCreate2)
		baseCost := params.CreateGas
		if opcode == vm.CREATE2 {
			keccakWords := am.WordsForBytes(uint64(len(code)))
			keccakCost := am.SaturatingUMul(params.Keccak256WordGas, keccakWords)
			baseCost = am.SaturatingUAdd(baseCost, keccakCost)
		}
		if gas < baseCost {
			return zeroAddr, 0, gas, vm.ErrOutOfGas
		}
		gas -= baseCost

		// apply the 63/64ths rule
		one64th := gas / 64
		gas -= one64th

		// Tracing: emit the create
		if tracingInfo != nil {
			tracingInfo.Tracer.CaptureState(0, opcode, startGas, baseCost+gas, scope, []byte{}, depth, nil)
		}

		var res []byte
		var addr common.Address // zero on failure
		var returnGas uint64
		var suberr error

		if opcode == vm.CREATE {
			res, addr, returnGas, suberr = evm.Create(contract, code, gas, endowment)
		} else {
			salt256, _ := uint256.FromBig(salt)
			res, addr, returnGas, suberr = evm.Create2(contract, code, gas, endowment, salt256)
		}
		if suberr != nil {
			addr = zeroAddr
		}
		if !errors.Is(vm.ErrExecutionReverted, suberr) {
			res = nil // returnData is only provided in the revert case (opCreate)
		}
		interpreter.SetReturnData(res)
		cost := am.SaturatingUSub(startGas, returnGas+one64th) // user gets 1/64th back
		return addr, uint32(len(res)), cost, nil
	}
	create1 := func(code []byte, endowment *big.Int, gas uint64) (common.Address, uint32, uint64, error) {
		return create(code, endowment, nil, gas)
	}
	create2 := func(code []byte, endowment, salt *big.Int, gas uint64) (common.Address, uint32, uint64, error) {
		return create(code, endowment, salt, gas)
	}
	getReturnData := func(start uint32, size uint32) []byte {
		end := am.SaturatingUAdd(start, size)
		return am.SliceWithRunoff(interpreter.GetReturnData(), start, end)
	}
	emitLog := func(data []byte, topics uint32) error {
		if readOnly {
			return vm.ErrWriteProtection
		}
		hashes := make([]common.Hash, topics)
		for i := uint32(0); i < topics; i++ {
			hashes[i] = common.BytesToHash(data[:(i+1)*32])
		}
		event := &types.Log{
			Address:     actingAddress,
			Topics:      hashes,
			Data:        data[32*topics:],
			BlockNumber: evm.Context.BlockNumber.Uint64(),
			// Geth will set other fields
		}
		db.AddLog(event)
		return nil
	}
	accountBalance := func(address common.Address) (common.Hash, uint64) {
		cost := vm.WasmAccountTouchCost(evm.StateDB, address, false)
		balance := evm.StateDB.GetBalance(address)
		return common.BigToHash(balance), cost
	}
	accountCode := func(address common.Address, offset, size uint32, gas uint64) ([]byte, uint64) {
		// In the future it'll be possible to know the size of a contract before loading it.
		// For now, require the worst case before doing the load.

		cost := vm.WasmAccountTouchCost(evm.StateDB, address, true)
		if gas < cost {
			return []byte{}, cost
		}
		end := am.SaturatingUAdd(offset, size)
		code := am.SliceWithRunoff(evm.StateDB.GetCode(address), offset, end)
		return code, cost
	}
	accountCodeSize := func(address common.Address, gas uint64) (uint32, uint64) {
		cost := vm.WasmAccountTouchCost(evm.StateDB, address, true)
		if gas < cost {
			return 0, cost
		}
		size := evm.StateDB.GetCodeSize(address)
		return uint32(size), cost
	}
	accountCodehash := func(address common.Address) (common.Hash, uint64) {
		cost := vm.WasmAccountTouchCost(evm.StateDB, address, false)
		return evm.StateDB.GetCodeHash(address), cost
	}
	addPages := func(pages uint16) uint64 {
		open, ever := db.AddStylusPages(pages)
		return memoryModel.GasCost(pages, open, ever)
	}
	captureHostio := func(name string, args, outs []byte, startInk, endInk uint64) {
		tracingInfo.Tracer.CaptureStylusHostio(name, args, outs, startInk, endInk)
	}

	return &goClosures{
		getBytes32:      getBytes32,
		setBytes32:      setBytes32,
		contractCall:    contractCall,
		delegateCall:    delegateCall,
		staticCall:      staticCall,
		create1:         create1,
		create2:         create2,
		getReturnData:   getReturnData,
		emitLog:         emitLog,
		accountBalance:  accountBalance,
		accountCode:     accountCode,
		accountCodeSize: accountCodeSize,
		accountCodeHash: accountCodehash,
		addPages:        addPages,
		captureHostio:   captureHostio,
	}
}
