package prometheus

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// ExporterMetrics wraps metrics for chain-exporter
type ExporterMetrics struct{
	BlockNumber   *prometheus.GaugeVec
}

// NewMetricsForExporter creates mertircs for chain-exporter
func NewMetricsForExporter(cfg config.PrometheusConfig) ExporterMetrics {

	blockNumber := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: cfg.Namespace,
			Name: "block_height",
			Help: "Height of block",
		},
		[]string{})

	blockNumber.WithLabelValues().Set(float64(0))

	return ExporterMetrics{blockNumber}
}

// RegisterMetricForExporter registers metrics for chain-exporter
func RegisterMetricForExporter(metrics ExporterMetrics){
	prometheus.Register(metrics.BlockNumber)
}

// StartMetricsScraping starts scraping metrics for chain-exporter
func StartMetricsScraping(cfg config.PrometheusConfig){
	http.Handle(cfg.Path, promhttp.Handler())
	http.ListenAndServe(":" + cfg.Port, nil)
}