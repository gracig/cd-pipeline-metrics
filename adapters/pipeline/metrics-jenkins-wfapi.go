package pipeline

import (
	"fmt"
	"log"
	"time"

	"github.com/gracig/cd-runtime-collector/adapters/jenkins"
	"github.com/gracig/cd-runtime-collector/lib"
)

//SendJenkinsWFAPIMetrics reads jenkins wfapi and send metric to metric adapter
func (o *Operations) SendJenkinsWFAPIMetrics() (err error) {

	//Variable declaration
	var (
		resp *jenkins.WorkflowAPIAnswer
	)

	//Creates a Metric Adapter object
	adapter := new(lib.Adapter)
	//Load a Jenkins operations
	jops, err := jenkins.NewOperations()
	if err != nil {
		log.Printf("Could not create a new Jenkins Operations. %v", err)
	}

	//Retrieving a web getter to be used
	wg := jops.GetBasicWebGetter()

	//Retrieves the WorkflowApiResponse from Jenkins API
	if resp, err = jops.GetWorkflowAPIAnswer(wg, o.Cfg.BuildURL); err != nil {
		return
	}
	timestamp := time.Unix(0, resp.StartTimeMillis*int64(time.Millisecond))

	//This adds the Job and Build Number
	tags := o.BaseTag()
	fields := o.BaseField()

	fields["Duration"] = resp.DurationMillis
	fields["Date"] = fmt.Sprintf("%v", timestamp)

	//Iterates of stages to get metrics from
	for _, stage := range resp.Stages {

		//Do not count on Declarative: Post Actions
		if stage.Name == "Declarative: Post Actions" {
			continue
		}

		fields[stage.Name+"-Duration"] = stage.DurationMillis
		fields[stage.Name+"-WaitTime"] = stage.PauseDurationMillis
		fields[stage.Name+"-ExecTime"] = stage.DurationMillis - stage.PauseDurationMillis
		fields[stage.Name+"-Status"] = stage.Status

		//Writes the last Stage
		fields["Stage"] = stage.Name
		fields["Status"] = stage.Status

		switch stage.Status {
		case "SUCCESS":
			fields[stage.Name+"-StatusCode"] = 0
			fields["StatusCode"] = 0
		case "IN_PROGRESS":
			fields[stage.Name+"-StatusCode"] = 1
			fields["StatusCode"] = 1
		default:
			fields[stage.Name+"-StatusCode"] = 2
			fields["StatusCode"] = 2
		}
		if stage.Error.Type != "" {
			fields["ErrorType"] = stage.Error.Type
			fields["ErrorMsg"] = stage.Error.Message
		}
		if stage.Status == "FAILED" {
			break
		}
	}

	if err = adapter.AddNewDataPoint(o.Ma, jenkinsMetricName, tags, fields, timestamp); err != nil {
		return
	}

	//Send metric values do metric database
	err = adapter.WriteData(o.Ma)
	return
}
