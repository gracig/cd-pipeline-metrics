package pipeline

import (
	"log"
	"time"

	"github.com/gracig/cd-runtime-collector/lib"
	"github.com/gracig/cd-runtime-collector/adapters/jenkins"
)

//SendJenkinsWorkflowRunMetrics reads jenkins wfapi and send metric to metric adapter
func (o *Operations) SendJenkinsWorkflowRunMetrics() (err error) {

	//Variable declaration
	var (
		resp *jenkins.WorkflowRun
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
	if resp, err = jops.GetWorkflowRun(wg, o.Cfg.BuildURL); err != nil {
		return
	}

	//Getting timestmp from response
	timestamp := time.Unix(0, resp.Timestamp*int64(time.Millisecond))

	//Building tags and values
	tags := o.BaseTag()
	fields := map[string]interface{}{
		"Building":          resp.Building,
		"EstimatedDuration": resp.EstimatedDuration,
		"Duration":          resp.Duration,
		"Result":            resp.Result,
	}

	//Add the metric value
	if err = adapter.AddNewDataPoint(o.Ma, jenkinsMetricName, tags, fields, timestamp); err != nil {
		return
	}

	//Send metric values do metric database
	err = adapter.WriteData(o.Ma)

	return
}
