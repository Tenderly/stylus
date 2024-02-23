package server_jit

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync/atomic"

	flag "github.com/spf13/pflag"

	"github.com/tenderly/stylus/go-ethereum/common"

	"github.com/tenderly/stylus/util/stopwaiter"
	"github.com/tenderly/stylus/validator"
	"github.com/tenderly/stylus/validator/server_common"
)

type JitSpawnerConfig struct {
	Workers   int  `koanf:"workers" reload:"hot"`
	Cranelift bool `koanf:"cranelift"`
}

type JitSpawnerConfigFecher func() *JitSpawnerConfig

var DefaultJitSpawnerConfig = JitSpawnerConfig{
	Workers:   0,
	Cranelift: true,
}

func JitSpawnerConfigAddOptions(prefix string, f *flag.FlagSet) {
	f.Int(prefix+".workers", DefaultJitSpawnerConfig.Workers, "number of concurrent validation threads")
	f.Bool(prefix+".cranelift", DefaultJitSpawnerConfig.Cranelift, "use Cranelift instead of LLVM when validating blocks using the jit-accelerated block validator")
}

type JitSpawner struct {
	stopwaiter.StopWaiter
	count         int32
	locator       *server_common.MachineLocator
	machineLoader *JitMachineLoader
	config        JitSpawnerConfigFecher
}

func NewJitSpawner(locator *server_common.MachineLocator, config JitSpawnerConfigFecher, fatalErrChan chan error) (*JitSpawner, error) {
	// TODO - preload machines
	machineConfig := DefaultJitMachineConfig
	machineConfig.JitCranelift = config().Cranelift
	loader, err := NewJitMachineLoader(&machineConfig, locator, fatalErrChan)
	if err != nil {
		return nil, err
	}
	spawner := &JitSpawner{
		locator:       locator,
		machineLoader: loader,
		config:        config,
	}
	return spawner, nil
}

func (v *JitSpawner) Start(ctx_in context.Context) error {
	v.StopWaiter.Start(ctx_in, v)
	return nil
}

func (v *JitSpawner) execute(
	ctx context.Context, entry *validator.ValidationInput, moduleRoot common.Hash,
) (validator.GoGlobalState, error) {
	machine, err := v.machineLoader.GetMachine(ctx, moduleRoot)
	if err != nil {
		return validator.GoGlobalState{}, fmt.Errorf("unabled to get WASM machine: %w", err)
	}

	resolver := func(hash common.Hash) ([]byte, error) {
		// Check if it's a known preimage
		if preimage, ok := entry.Preimages[hash]; ok {
			return preimage, nil
		}
		return nil, errors.New("preimage not found")
	}
	state, err := machine.prove(ctx, entry, resolver)
	return state, err
}

func (s *JitSpawner) Name() string {
	if s.config().Cranelift {
		return "jit-cranelift"
	}
	return "jit"
}

func (v *JitSpawner) Launch(entry *validator.ValidationInput, moduleRoot common.Hash) validator.ValidationRun {
	atomic.AddInt32(&v.count, 1)
	promise := stopwaiter.LaunchPromiseThread[validator.GoGlobalState](v, func(ctx context.Context) (validator.GoGlobalState, error) {
		defer atomic.AddInt32(&v.count, -1)
		return v.execute(ctx, entry, moduleRoot)
	})
	return server_common.NewValRun(promise, moduleRoot)
}

func (v *JitSpawner) Room() int {
	avail := v.config().Workers
	if avail == 0 {
		avail = runtime.NumCPU()
	}
	return avail
}

func (v *JitSpawner) Stop() {
	v.StopOnly()
	v.machineLoader.Stop()
}
