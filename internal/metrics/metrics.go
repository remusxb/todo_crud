package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

// DefaultCollectors will collect just about everything.
func DefaultCollectors() []prometheus.Collector {
	return []prometheus.Collector{
		collectors.NewBuildInfoCollector(),
		collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics()),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	}
}
