# Logstash exporter   [![Build Status](https://travis-ci.org/sagap/logstash_exporter.svg?branch=master)](https://travis-ci.org/sagap/logstash_exporter)
Prometheus exporter for the metrics available in Logstash version 6.4.  
  
## Usage  
  
```bash  
go get -u github.com/sagap/logstash_exporter  
cd $GOPATH/src/github.com/sagap/logstash_exporter  
make  
./logstash_exporter -listen_port=":1234"
```  
  
In order to take advantage of its functionality you need to invoke a curl command:  
```
curl http://localhost:1234/monitoring\?monitoringPort\=123  
```
or configure the prometheus.yml as below:    
```
  - job_name: 'logstash-testflow'
    metrics_path: /monitoring
    static_configs:
      - targets:
        - localhost:1234
        labels:
          instace: 'logstash-testflow'
    relabel_configs:
      - source_labels: [__param_]
        target_label: __param_monitoringPort
        replacement: http://localhost:9600
```
  
### Flags  
Flag | Description | Default  
-----|-------------|---------  
-listen_port | Exporter bind address | ":1234"
  
## Implemented metrics  
* Events in
* Events out
* Events filtered
* JVM mem heap used percent
* Process CPU percent
