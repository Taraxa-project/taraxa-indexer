package metrics

import (
	"net/http"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

const metricPrefix = "taraxa_indexer"

var (
	StorageCommitCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: metricPrefix + "_storage_commit_counter",
		Help: "The total number of commit operations to storage",
	})
	RpcCallsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: metricPrefix + "_rpc_calls_counter",
		Help: "The total number of RPC calls to taraxa-node",
	})
	IndexedBlocksCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: metricPrefix + "_blocks_indexed_counter",
		Help: "Number of indexed blocks by indexer since start",
	})
	IndexedTransactionsLastProcessedBlock = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_indexed_transactions_last_block",
		Help: "Number of indexed transactions in last processed block",
	})
	IndexedDagBlocksLastProcessedBlock = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_indexed_dags_last_block",
		Help: "Number of indexed DAG blocks in last processed block",
	})
	IndexedTransactions = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_indexed_transactions",
		Help: "Number of indexed transactions since start",
	})
	IndexedDagBlocks = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_indexed_dags",
		Help: "Number of indexed DAG blocks since start",
	})
	BlockProcessingTimeMilisec = promauto.NewSummary(prometheus.SummaryOpts{
		Name: metricPrefix + "_block_processing_time",
		Help: "Time of processing block in milisec",
		// Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})
	LastProcessedBlockTimestamp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_time_last_block",
		Help: "Timestamp of last processed block",
	})
	IndexedBlocksTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_blocks_indexed_total",
		Help: "Total number of indexed blocks",
	})
	IndexedTransactionsTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_transactions_indexed_total",
		Help: "Total number of indexed transactions",
	})
	IndexedDagsTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: metricPrefix + "_dags_indexed_total",
		Help: "Total number of indexed DAG blocks",
	})
)

func RunPrometheusServer(listenAddr string) {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.WithError(err).Fatal("Can't start http server")
	}
}

func Save(start_processing time.Time, dags_count, trx_count uint64, finalized *common.FinalizationData) {
	BlockProcessingTimeMilisec.Observe(float64(time.Since(start_processing).Milliseconds()))
	IndexedBlocksCounter.Inc()

	IndexedDagBlocksLastProcessedBlock.Set(float64(dags_count))
	IndexedTransactionsLastProcessedBlock.Set(float64(trx_count))

	IndexedDagBlocks.Add(float64(dags_count))
	IndexedTransactions.Add(float64(trx_count))

	IndexedBlocksTotal.Set(float64(finalized.PbftCount))
	IndexedDagsTotal.Set(float64(finalized.DagCount))
	IndexedTransactionsTotal.Set(float64(finalized.TrxCount))

	LastProcessedBlockTimestamp.SetToCurrentTime()
}
