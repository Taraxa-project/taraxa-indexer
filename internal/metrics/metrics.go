package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	StorageAddCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "taraxa_indexer_storage_add_counter",
		Help: "The total number of Add operations to storage",
	})
	StorageGetCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "taraxa_indexer_storage_get_counter",
		Help: "The total number of Get operations to storage",
	})
	RpcCallsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "taraxa_indexer_rpc_calls_counter",
		Help: "The total number of RPC calls to taraxa-node",
	})
	IndexedBlocksCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "taraxa_indexer_blocks_indexed_counter",
		Help: "Number of indexed blocks by indexer since restart",
	})
	IndexedBlocksTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "taraxa_indexer_blocks_indexed_total",
		Help: "Total number of indexed blocks",
	})
)

func RunPrometheusServer() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
