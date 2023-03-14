package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

const metricPrefix = "taraxa_indexer"

var (
	StorageAddCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: metricPrefix + "_storage_add_counter",
		Help: "The total number of Add operations to storage",
	})
	StorageGetCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: metricPrefix + "_storage_get_counter",
		Help: "The total number of Get operations to storage",
	})
	RpcCallsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: metricPrefix + "_rpc_calls_counter",
		Help: "The total number of RPC calls to taraxa-node",
	})
	IndexedBlocksCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: metricPrefix + "_blocks_indexed_counter",
		Help: "Number of indexed blocks by indexer since restart",
	})
	IndexedBlocksTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_blocks_indexed_total",
		Help: "Total number of indexed blocks",
	})
	IndexedDagBlocksTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_dag_blocks_indexed_total",
		Help: "Total number of indexed DAG blocks",
	})
	IndexedPbftBlocksTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_pbft_blocks_indexed_total",
		Help: "Total number of indexed PBFT blocks",
	})
	IndexedTransactionsTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_transactions_indexed_total",
		Help: "Total number of indexed transactions",
	})
	AverageBlockProcessingTimeMilisec = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_average_block_processing_time",
		Help: "Average time of processing block in milisec",
	})
)

func RunPrometheusServer() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		log.WithError(err).Fatal("Can't start http server")
	}
}
