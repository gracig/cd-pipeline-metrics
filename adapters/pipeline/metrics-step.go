package pipeline

//SendDoOperationElapsedTime sends the time an operation tool and its success
import (
	"fmt"
	"time"

	"github.com/gracig/cd-runtime-collector/lib"
)

//SendDoOperationElapsedTime send step elapsed time and statues to metric database
func (o *Operations) SendDoOperationElapsedTime(stepID string) (err error) {

	var (
		stageName     string
		stepName      string
		stepStart     int64
		stepStartTime time.Time
		stepEnd       int64
		elapsedTime   float64
		stepStatus    string
		stepError     string
	)
	//Creates a Metric Adapter object
	adapter := new(lib.Adapter)

	if stageName, err = lib.GetEnvVar(stepID+"_STAGE_NAME", ".+"); err != nil {
		return
	}
	if stepName, err = lib.GetEnvVar(stepID+"_STEP_NAME", ".+"); err != nil {
		return
	}
	if stepStatus, err = lib.GetEnvVar(stepID+"_STEP_STATUS", ".+"); err != nil {
		return
	}
	if stepError, err = lib.GetEnvVar(stepID+"_STEP_ERROR", ".*"); err != nil {
		return
	}
	if stepStart, err = lib.GetEnvVarAsInt64(stepID + "_STEP_START"); err != nil {
		return
	}
	if stepEnd, err = lib.GetEnvVarAsInt64(stepID + "_STEP_END"); err != nil {
		return
	}
	stepStartTime = time.Unix(0, stepStart*int64(time.Millisecond))
	elapsedTime = float64(stepEnd - stepStart)

	tags := o.BaseTag()
	tags["Stage"] = stageName
	tags["Step"] = stepName
	tags["ID"] = stepID

	fields := o.BaseField()
	fields["Duration"] = elapsedTime
	fields["Status"] = stepStatus
	fields["ErrorMessage"] = stepError
	fields["JenkinsBuild"] = o.Cfg.BuildNumber
	fields["Date"] = fmt.Sprintf("%v", stepStartTime)

	if err = adapter.AddNewDataPoint(o.Ma, pipelineMetricName, tags, fields, stepStartTime); err != nil {
		return
	}

	//Send metric values do metric database
	err = adapter.WriteData(o.Ma)

	return
}
