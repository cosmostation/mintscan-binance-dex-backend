package prometheus

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// ExporterMetrics wraps metrics for exporter
type ExporterMetrics struct{
	BlockNumber   *prometheus.GaugeVec
}

// NewMetricsForExporter creates new metrics for exporter
func NewMetricsForExporter(cfg config.PrometheusConfig) ExporterMetrics {
	blockNumber := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Name: "block_height",
			Help: "Height of block",
		}, []string{})
	// set initial value
	blockNumber.WithLabelValues().Set(float64(0))

	return ExporterMetrics{blockNumber}
}

// RegisterMetricsForExporter registers metrics for exporter
func RegisterMetricsForExporter(metrics ExporterMetrics){
	prometheus.Register(metrics.BlockNumber)
}

// StartMetricScraping starts metrics scraping in metrics path
func StartMetricScraping(cfg config.PrometheusConfig) {
	http.Handle(cfg.Path, promhttp.Handler())
	http.ListenAndServe(":" + cfg.Port, nil)
}