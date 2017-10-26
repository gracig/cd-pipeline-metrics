package pipeline

import (
	"log"
	"strings"
	"time"

	"github.com/gracig/cd-runtime-collector/lib"
)

//Config represents Build parameters inside a running Jeknis Pipeline Job
type Config struct {
	//BuildURL is the URL Status page of this JOB in Jenkins
	BuildURL string
	//BuildTimestamp is the time the this job have started running
	BuildTimestamp time.Time
	//BuildNumber is the number of the running job
	BuildNumber string
	//JobName is the name of the JOB
	JobName string
	//Customer is the name of the project owner
	Customer string
	//ERROR_MESSAGE : The last error on the pipeline if any
	ErrorMessage string
	//MetricBaseKeyLabels
	MetricBaseKeyLabels []string
	//MetricBaseKeyValues
	MetricBaseKeyValues []string
	//MetricBaseFieldLabels
	MetricBaseFieldLabels []string
	//MetricBaseFieldValues
	MetricBaseFieldValues []string
	//MetricBaseFieldTypes
	MetricBaseFieldTypes []string
}

//NewConfigFromEnv creates a new JenkinsConfig from environment variables.
//Expected Environment Variables:
// BUILD_URL BUILD_TIMESTAMP BUILD_NUMBER and JOB_NAME
func NewConfigFromEnv() (cfg Config, err error) {

	//Getting Build URL
	if cfg.BuildURL, err = lib.GetEnvVar("BUILD_URL", "http://.*:\\d*"); err != nil {
		return
	}

	//Getting the metric timestamp from BUILD_TIMESTAMP
	var timeAsString string
	if timeAsString, err = lib.GetEnvVar("BUILD_TIMESTAMP", ".+"); err != nil {
		return
	}
	if cfg.BuildTimestamp, err = time.Parse(time.RFC3339, timeAsString); err != nil {
		log.Println("Could not parse timestamp from BUILD_TIMESTAMP", err)
		return
	}

	//Getting Build Number
	if cfg.BuildNumber, err = lib.GetEnvVar("BUILD_NUMBER", ".+"); err != nil {
		return
	}

	//Getting Job Name
	if cfg.JobName, err = lib.GetEnvVar("JOB_NAME", ".+"); err != nil {
		return
	}

	//Getting Customer
	if cfg.Customer, err = lib.GetEnvVar("CUSTOMER", ".*"); err != nil {
		return
	}

	//Error Message
	if cfg.ErrorMessage, err = lib.GetEnvVar("ERROR_MESSAGE", ".*"); err != nil {
		return
	}

	//Extra keys
	var (
		keyLabels   string
		keyValues   string
		fieldLabels string
		fieldValues string
		fieldTypes  string
	)

	if keyLabels, err = lib.GetEnvVar("METRIC_BASE_KEY_LABELS", ".+"); err != nil {
		return
	}
	cfg.MetricBaseKeyLabels = strings.Split(keyLabels, ",")

	if keyValues, err = lib.GetEnvVar("METRIC_BASE_KEY_VALUES", ".+"); err != nil {
		return
	}
	cfg.MetricBaseKeyValues = strings.Split(keyValues, ",")

	if fieldLabels, err = lib.GetEnvVar("METRIC_BASE_FIELD_LABELS", ".+"); err != nil {
		return
	}
	cfg.MetricBaseFieldLabels = strings.Split(fieldLabels, ",")

	if fieldValues, err = lib.GetEnvVar("METRIC_BASE_FIELD_VALUES", ".+"); err != nil {
		return
	}
	cfg.MetricBaseFieldValues = strings.Split(fieldValues, ",")

	if fieldTypes, err = lib.GetEnvVar("METRIC_BASE_FIELD_TYPES", ".+"); err != nil {
		return
	}
	cfg.MetricBaseFieldTypes = strings.Split(fieldTypes, ",")

	return
}
