package prometheus

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type ExporterMetrics struct{
	BlockNumber   *prometheus.GaugeVec
}

func NewMetricsForExporter() ExporterMetrics {

	blockNumber := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "block_number",
			Help: "Number of block",
		},
		[]string{})

	blockNumber.WithLabelValues().Set(float64(0))

	return ExporterMetrics{blockNumber}
}

func RegisterMetricsForExporter(metrics ExporterMetrics){
	prometheus.Register(metrics.BlockNumber)
}

func StartMetricScraping(cfg config.PrometheusConfig) {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":" + cfg.Port, nil)
}