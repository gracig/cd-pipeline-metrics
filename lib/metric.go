package lib

import (
	"fmt"
	"time"
)

//Adapter represents a generic metric adapter.
type Adapter struct {
	Verbose bool
}

//AddNewDataPoint would add a new data into a BatchManager object. The last one should be implemented for the desired backend
func (a *Adapter) AddNewDataPoint(ma MetricAdapter, name string, tags map[string]string, fields map[string]interface{}, t time.Time) error {
	if a.Verbose {
		dumpDataPoint(name, tags, fields, t)
	}
	return ma.AddDataPoint(name, tags, fields, t)
}

//WriteData would call a write method of a datasource using data from a Batch Manager. Both interfaces should be
//implemented for a desired backend. e.g(InfluxDB)
func (a *Adapter) WriteData(ma MetricAdapter) error {
	return ma.WriteData()
}

//MetricAdapter represents an object that send metrics over a metric database
type MetricAdapter interface {
	AddDataPoint(name string, tags map[string]string, fields map[string]interface{}, t time.Time) error
	WriteData() error
}

//dumpDataPoint prints out the received metrics to be sent
func dumpDataPoint(name string, tags map[string]string, fields map[string]interface{}, timestamp time.Time) {
	//Getting Build URL
	if _, err := GetEnvVar("DEVOPS_TOOL_DEBUG", "1"); err == nil {
		fmt.Printf("MetricName: [%v]\n", name)
		fmt.Printf("\tTags: [%v]\n", tags)
		fmt.Printf("\tFields: [%v]\n", fields)
		fmt.Printf("\tTimestamp: [%v]\n", timestamp)
	}
}
