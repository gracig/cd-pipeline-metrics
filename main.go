package main

import (
	"flag"
	"log"
	"os"

	"github.com/gracig/cd-runtime-collector/adapters/pipeline"
	"github.com/gracig/cd-runtime-collector/lib"
)

func main() {

	//Flag declaratios
	version := flag.Bool("version", false, "Prints the version of this tool")
	jkwfapi := flag.Bool("jkwfapi", false, "Send Jenkins Metrics By Job")
	jkstage := flag.Bool("jkstage", false, "Send Jenkins Metrics By Stage")
	jkwfrun := flag.Bool("jkwfrun", false, "Send Jenkins Workflow Run data to Metric Database")
	sonar := flag.Bool("sonar", false, "Send Sonar Quality data to Metric Database")
	step := flag.Bool("step", false, "Send Elapsed time from current step to Metric Database")
	stepID := flag.String("stepID", "", "The step identificator")
	flag.Parse()

	if *version {
		lib.PrintVersion()
	}

	//variable declaration
	var (
		ops *pipeline.Operations //ops contains pipeline operation functions
		err error                //stores any errors
	)

	if *jkwfapi || *jkwfrun || *sonar || *step || *jkstage {
		if ops, err = pipeline.NewOperations(); err != nil {
			log.Printf("Could not create a new Pipeline Operations. %v", err)
			os.Exit(1)
		}
	}

	//Send Workflow pipeline stage times to metric database
	if *jkwfapi {
		if err = ops.SendJenkinsWFAPIMetrics(); err != nil {
			log.Printf("Could not send Jenkins Workflow API data to metric database. %v", err)
			os.Exit(1)
		}
	}

	//Send Workflow pipeline stage times to metric database
	if *jkstage {
		if err = ops.SendJenkinsWFAPIStageMetrics(); err != nil {
			log.Printf("Could not send Jenkins Workflow API data to metric database. %v", err)
			os.Exit(1)
		}
	}

	//Send Workflow Run times to metric database
	if *jkwfrun {
		if err = ops.SendJenkinsWorkflowRunMetrics(); err != nil {
			log.Printf("Could not send Jenkins Workflow Run data to metric database. %v", err)
			os.Exit(1)
		}
	}

	//Send sonar quality variables to Metric database
	if *sonar {
		if err = ops.SendSonarJobMetrics(); err != nil {
			log.Printf("Could not send Sonar Quality data to metric database. %v", err)
			os.Exit(1)
		}
	}

	//Send Step metrics to Metric Database
	if *step {
		if err = ops.SendDoOperationElapsedTime(*stepID); err != nil {
			log.Printf("Could not send Step Elapsed Time to Metric Database. %v", err)
			os.Exit(1)
		}
	}

}
