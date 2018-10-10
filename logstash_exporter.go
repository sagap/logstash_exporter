package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"github.com/sagap/logstash_exporter/collector"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	scrapeDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: collector.Namespace,
			Subsystem: "exporter",
			Name:      "scrape_duration_seconds",
			Help:      "logstash_exporter: Duration of a scrape job.",
		},
		[]string{"collector", "result"},
	)
)

// LogstashCollector collector type
type LogstashCollector struct {
	collectors map[string]collector.Collector
}

// NewLogstashCollector register a logstash collector
func NewLogstashCollector(logstashEndpoint string) *LogstashCollector {
	nodeMonitoringCollector, err := collector.NewNodeMonitoringCollector(logstashEndpoint)
	if err != nil {
		log.Fatalf("Cannot register a new collector: %v", err)
	}
	return &LogstashCollector{
		collectors: map[string]collector.Collector{
			"monitoring": nodeMonitoringCollector,
		},
	}
}

func listen(exporterBindAddress string) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/monitoring",
		func(writer http.ResponseWriter, request *http.Request) {
			logstashEndpoint := request.URL.Query()["monitoringPort"]
			logstashCollector := NewLogstashCollector(logstashEndpoint[0])
			prometheus.MustRegister(logstashCollector)
			promhttp.Handler().ServeHTTP(writer, request)
			prometheus.Unregister(logstashCollector)
		})
	if err := http.ListenAndServe(exporterBindAddress, nil); err != nil {
		log.Fatalf("Cannot start Logstash exporter: %s", err)
	}
}

// Describe logstash metrics
func (coll LogstashCollector) Describe(ch chan<- *prometheus.Desc) {
	scrapeDurations.Describe(ch)
}

// Collect logstash metrics
func (coll LogstashCollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(coll.collectors))
	for name, c := range coll.collectors {
		go func(name string, c collector.Collector) {
			execute(name, c, ch)
			wg.Done()
		}(name, c)
	}
	wg.Wait()
	scrapeDurations.Collect(ch)
}

func execute(name string, c collector.Collector, ch chan<- prometheus.Metric) {
	begin := time.Now()
	err := c.Collect(ch)
	duration := time.Since(begin)
	var result string
	if err != nil {
		log.Errorf("ERROR: %s collector failed after %fs: %s", name, duration.Seconds(), err)
		result = "error"
	} else {
		log.Debugf("OK: %s collector succeeded after %fs.", name, duration.Seconds())
		result = "success"
	}
	scrapeDurations.WithLabelValues(name, result).Observe(duration.Seconds())
}

func init() {
	prometheus.MustRegister(version.NewCollector("logstash_exporter"))
}

func main() {
	exporterBindAddress := flag.String("listen_port", ":1234", "Address on which to expose metrics and web interface.")
	flag.Parse()
	log.Infoln("Starting Logstash exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())
	listen(*exporterBindAddress)
}
