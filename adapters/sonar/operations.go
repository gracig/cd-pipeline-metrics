package sonar

import (
	"fmt"
	"log"

	"github.com/gracig/cd-runtime-collector/lib"
)

//NewOperations returns an Operations object
func NewOperations() (ops *Operations, err error) {
	var cfg Config

	if cfg, err = NewConfigFromEnv(); err != nil {
		log.Printf("Could not get SonarQube config from environment. %v", err)
		return
	}
	ops = &Operations{Cfg: cfg}
	return
}

//Operations enables actions upon jenkins server
type Operations struct {
	Cfg Config
}

//GetBasicWebGetter retrieves the SonarQube Response from Sonar API
func (ops *Operations) GetBasicWebGetter() (response lib.WebGetter) {
	return &lib.BasicAuthWebGetter{User: ops.Cfg.APIUser, Pwd: ops.Cfg.APIPWD}
}

//GetComponentTreeAnswer retrieves the ComponentTreeAnswer Response from Sonar Qube API
func (ops *Operations) GetComponentTreeAnswer(wg lib.WebGetter) (response *ComponentTreeAnswer, err error) {

	//builds the url to query sonar qube api
	url := fmt.Sprintf("%v/api/measures/component_tree?additionalFields=metrics,periods&baseComponentKey=%v&metricKeys=%v",
		ops.Cfg.APIURL,
		ops.Cfg.APIProjectKey,
		ops.Cfg.APIMetrics)
	response = new(ComponentTreeAnswer)
	err = wg.DecodeRestJSON(response, url)
	return
}
