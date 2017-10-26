package pipeline

import (
	"log"

	"github.com/gracig/cd-runtime-collector/lib"
	"github.com/gracig/cd-runtime-collector/adapters/sonar"
)

//SendSonarJobMetrics reads job details from Sonar API and send to metric database
func (o *Operations) SendSonarJobMetrics() (err error) {

	//Variable declaration
	var (
		resp *sonar.ComponentTreeAnswer
	)

	//Creates a Metric Adapter object
	adapter := new(lib.Adapter)

	//Load a Sonar Operations
	sops, err := sonar.NewOperations()
	if err != nil {
		log.Printf("Could not create a new Sonar Operations. %v", err)
	}

	//Retrieving a web getter to be used
	wg := sops.GetBasicWebGetter()

	//Retrieves the WorkflowApiResponse from Jenkins API
	if resp, err = sops.GetComponentTreeAnswer(wg); err != nil {
		return
	}

	//Iterates over the reponse Components to get metrics from
	for _, comp := range resp.Components {
		//Creates base tag and add more information
		tags := o.BaseTag()
		tags["ID"] = comp.ID
		tags["Key"] = comp.Key
		tags["Name"] = comp.Name
		tags["Type"] = comp.Qualifier
		tags["Lang"] = comp.Language
		tags["Path"] = comp.Path

		//Creates the fields map and populate it from measure.Metrics
		fields := make(map[string]interface{})
		for _, measure := range comp.Measures {
			fields[measure.Metric] = lib.Text2Float(measure.Value)
		}

		//Adds the data to the metric adapter

		if err = adapter.AddNewDataPoint(o.Ma, qualityMetricName, tags, fields, o.Cfg.BuildTimestamp); err != nil {
			return
		}
	}

	//Create metrics of the Base Component
	//Creates base tag and add more information
	tags := o.BaseTag()
	tags["ID"] = resp.BaseComponent.ID
	tags["Key"] = resp.BaseComponent.Key
	tags["Name"] = resp.BaseComponent.Name
	tags["Type"] = resp.BaseComponent.Qualifier

	//Creates the fields map and populate it from measure.Metrics
	fields := make(map[string]interface{})
	for _, measure := range resp.BaseComponent.Measures {
		fields[measure.Metric] = lib.Text2Float(measure.Value)
	}

	//Adds the data to the metric adapter
	if err = adapter.AddNewDataPoint(o.Ma, qualityMetricName, tags, fields, o.Cfg.BuildTimestamp); err != nil {
		return
	}

	//Write metrics to database
	err = adapter.WriteData(o.Ma)

	return
}
