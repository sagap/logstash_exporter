package collector

import "fmt"

// monitoring response
type NodeMonitoringResponse struct {
	Jvm struct {
		Mem struct {
			HeapUsedPercent float64 `json:"heap_used_percent"`
		} `json:"mem"`
	} `json:"jvm"`
	Process struct {
		CPU struct {
			Percent float64 `json:"percent"`
		} `json:"cpu"`
	} `json:"process"`
	Events struct {
		In       int `json:"in"`
		Filtered int `json:"filtered"`
		Out      int `json:"out"`
	} `json:"events"`
}

// invoking Logstash instance to get metrics
func NodeMonitoring(endpoint string) (NodeMonitoringResponse, error) {
	var response NodeMonitoringResponse
	fmt.Println("Endpoint: ", endpoint)
	handler := &HTTPHandler{
		Endpoint: endpoint + "/_node/stats",
	}
	err := getMetrics(handler, &response)
	return response, err
}

//TODO function
func SetEndpointsToInvoke(endpoint <-chan string) {

}
