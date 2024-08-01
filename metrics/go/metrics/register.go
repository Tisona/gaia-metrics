package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

func RegisterMetrics(latest_block_height float64, desync float64, n_peers float64) *prometheus.Registry {
	r := prometheus.NewRegistry()

	g_latest_block_height := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gaia_latest_block_height",
		Help: "gaia_latest_block_height",
	})
	g_desync := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gaia_desync",
		Help: "gaia_desync",
	})
	g_n_peers := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gaia_n_peers",
		Help: "gaia_n_peers",
	})

	g_latest_block_height.Set(latest_block_height)
	g_desync.Set(desync)
	g_n_peers.Set(n_peers)

	r.MustRegister(g_latest_block_height)
	r.MustRegister(g_desync)
	r.MustRegister(g_n_peers)

	return r
}
