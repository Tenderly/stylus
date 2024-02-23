package genericconf

import (
	"fmt"
	"net/http"

	// Blank import pprof registers its HTTP handlers.
	_ "net/http/pprof" // #nosec G108

	"github.com/tenderly/stylus/go-ethereum/log"
	"github.com/tenderly/stylus/go-ethereum/metrics"
	"github.com/tenderly/stylus/go-ethereum/metrics/exp"
)

func StartPprof(address string) {
	exp.Exp(metrics.DefaultRegistry)
	log.Info("Starting metrics server with pprof", "addr", fmt.Sprintf("http://%s/debug/metrics", address))
	log.Info("Pprof endpoint", "addr", fmt.Sprintf("http://%s/debug/pprof", address))
	go func() {
		// #nosec G114
		if err := http.ListenAndServe(address, http.DefaultServeMux); err != nil {
			log.Error("Failure in running pprof server", "err", err)
		}
	}()
}
