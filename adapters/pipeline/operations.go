package pipeline

import (
	"fmt"
	"log"

	"github.com/gracig/cd-runtime-collector/lib"
	"github.com/gracig/cd-runtime-collector/adapters/influxdb"
)

const (
	pipelineMetricName = "PIPELINE"
	qualityMetricName  = "QUALITY"
	jenkinsMetricName  = "JENKINS"
)

//NewOperations returns an Operations object
func NewOperations() (ops *Operations, err error) {

	//Declares the cfg variable
	var cfg Config

	//Loads Pipeline config from environment
	if cfg, err = NewConfigFromEnv(); err != nil {
		log.Printf("Could not get pipeline config from environment. %v", err)
		return
	}

	//Creates a metric adapter. In this case from influxdb
	ma, err := influxdb.NewMetricAdapter()
	if err != nil {
		log.Printf("Could not create a new Metric Adapter. %v", err)
	}

	//Instantiates a new pipeline.Operations object
	ops = &Operations{Cfg: cfg, Ma: ma}
	return
}

//Operations enables actions upon jenkins server
type Operations struct {
	Cfg Config
	Ma  lib.MetricAdapter
}

//BaseTag Creates a base tag for all metric operations
func (o *Operations) BaseTag() map[string]string {
	result := make(map[string]string)
	for i, label := range o.Cfg.MetricBaseKeyLabels {
		if value, err := lib.GetEnvVar(o.Cfg.MetricBaseKeyValues[i], ".+"); err != nil {
			log.Fatalln("Error while build BaseTag: ", err)
		} else {
			result[label] = value
		}
	}
	return result
}

//BaseField Creates a base field for all metric operations
func (o *Operations) BaseField() map[string]interface{} {
	result := make(map[string]interface{})

	for i, label := range o.Cfg.MetricBaseFieldLabels {

		var (
			value interface{}
			err   error
		)

		if value, err = lib.GetEnvVar(o.Cfg.MetricBaseFieldValues[i], ".*"); err != nil {
			log.Fatalln("Error while building BaseField: ", err)
		}
		switch o.Cfg.MetricBaseFieldTypes[i] {
		case "string":
		case "float":
			value = lib.Text2Float(fmt.Sprintf("%v", value))
		default:
			log.Println("Could not determine type for field value! Using string value")
		}

		result[label] = value

	}
	return result
}
