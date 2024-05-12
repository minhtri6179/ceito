package metric

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	m    prometheus.Gauge
	info *prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		m: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "test",
			Help: "test",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "info",
			Help: "info",
		}, []string{"version"}),
	}
	reg.MustRegister(m.m, m.info)
	return m
}
