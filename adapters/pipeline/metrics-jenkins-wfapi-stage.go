package pipeline

import (
	"fmt"
	"log"
	"time"

	"github.com/gracig/cd-runtime-collector/adapters/jenkins"
	"github.com/gracig/cd-runtime-collector/lib"
)

//SendJenkinsWFAPIMetrics reads jenkins wfapi and send metric to metric adapter
func (o *Operations) SendJenkinsWFAPIStageMetrics() (err error) {

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

	//Iterates of stages to get metrics from
	for i, stage := range resp.Stages {
		tags["Stage"] = stage.Name
		fields["Sequence"] = fmt.Sprintf("%03d", i)
		fields["Duration"] = stage.DurationMillis
		fields["WaitTime"] = stage.PauseDurationMillis
		fields["ExecTime"] = stage.DurationMillis - stage.PauseDurationMillis
		fields["Status"] = stage.Status
		if stage.Error.Type != "" {
			fields["ErrorType"] = stage.Error.Type
			fields["ErrorMsg"] = stage.Error.Message
		}
		switch stage.Status {
		case "SUCCESS":
			fields["StatusCode"] = 0
		case "IN_PROGRESS":
			fields["StatusCode"] = 1
		default:
			fields["StatusCode"] = 2
		}
		if err = adapter.AddNewDataPoint(o.Ma, jenkinsMetricName+"-STAGE", tags, fields, timestamp); err != nil {
			return
		}
	}

	//Send metric values do metric database
	err = adapter.WriteData(o.Ma)
	return
}
