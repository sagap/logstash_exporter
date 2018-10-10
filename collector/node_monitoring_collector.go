package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type NodeMonitoringCollector struct {
	endpoint              string
	JVMMemHeapUsedPercent *prometheus.Desc
	ProcessCPUPercent     *prometheus.Desc
	EventsIn              *prometheus.Desc
	EventsFiltered        *prometheus.Desc
	EventsOut             *prometheus.Desc
}

func NewNodeMonitoringCollector(logstashEndpoint string) (Collector, error) {
	const subsystem = "monitoring"

	return &NodeMonitoringCollector{
		endpoint: logstashEndpoint,

		JVMMemHeapUsedPercent: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "jvm_mem_heap_used_percent"),
			"jvm_mem_heap_used_percent",
			nil,
			nil,
		),
		ProcessCPUPercent: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "process_cpu_percent"),
			"process_cpu_percent",
			nil,
			nil,
		),
		EventsIn: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "events_in"),
			"events_in",
			nil,
			nil,
		),
		EventsFiltered: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "events_filtered"),
			"events_filtered",
			nil,
			nil,
		),
		EventsOut: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, subsystem, "events_out"),
			"events_out",
			nil,
			nil,
		),
	}, nil
}

// implements Collect of interface Collector
func (nodeCollector *NodeMonitoringCollector) Collect(ch chan<- prometheus.Metric) error {
	if desc, err := nodeCollector.collect(ch); err != nil {
		log.Error("Failed collecting node metrics", desc, err)
		return err
	}
	return nil
}

func (nodeCollector *NodeMonitoringCollector) collect(ch chan<- prometheus.Metric) (*prometheus.Desc, error) {
	statsResponse, err := NodeMonitoring(nodeCollector.endpoint)
	if err != nil {
		return nil, err
	}
	ch <- prometheus.MustNewConstMetric(
		nodeCollector.JVMMemHeapUsedPercent,
		prometheus.GaugeValue,
		statsResponse.Jvm.Mem.HeapUsedPercent,
	)
	ch <- prometheus.MustNewConstMetric(
		nodeCollector.ProcessCPUPercent,
		prometheus.GaugeValue,
		statsResponse.Process.CPU.Percent,
	)
	ch <- prometheus.MustNewConstMetric(
		nodeCollector.EventsIn,
		prometheus.GaugeValue,
		float64(statsResponse.Events.In),
	)
	ch <- prometheus.MustNewConstMetric(
		nodeCollector.EventsFiltered,
		prometheus.GaugeValue,
		float64(statsResponse.Events.Filtered),
	)
	ch <- prometheus.MustNewConstMetric(
		nodeCollector.EventsOut,
		prometheus.GaugeValue,
		float64(statsResponse.Events.Out),
	)
	return nil, nil
}
