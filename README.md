# Logstash exporter  
Prometheus exporter for the metrics available in Logstash version 6.4.  
  
## Usage  
  
```bash  
go get -u github.com/sagap/logstash_exporter  
cd $GOPATH/src/github.com/sagap/logstash_exporter  
make  
./logstash_exporter -exporter.bind_address :1234  
```  
  
In order to take advantage of its functionality you need to invoke a curl command:  
```
curl http://localhost:1234/monitoring\?monitoringPort\=123  
```
or configure the prometheus.yml as below:    
```
  - job_name: 'logstash-testflow'  
    metrics_path: /monitoring  
    params:  
            monitoringPort: [http://localhost:9600]  
    static_configs:  
    - targets:  
      - localhost:1234  
```

  
### Flags  
Flag | Description | Default  
-----|-------------|---------  
-exporter.bind_address | Exporter bind address | :1234
  
## Implemented metrics  
* Node metrics
