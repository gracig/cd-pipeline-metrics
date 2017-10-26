package sonar

import "github.com/gracig/cd-runtime-collector/lib"

//Config represents parameters to access or manages a Sonar Qube API
type Config struct {
	//APIPWD is the token password of the Sonar Qube API
	APIPWD string
	//ApiUser is the user of the Sonar Qube API
	APIUser string
	//ApiURL is the url where the Sonar Qube API is hosted
	APIURL string
	//APIProjectKey is the id of the project to be queried
	APIProjectKey string
	//APIMetrics is the parameters to get metrics from
	APIMetrics string
}

//NewConfigFromEnv creates a new Sonar QubeConfig from environment variables.
//Expected environment variables:
// Sonar Qube_API_TOKEN Sonar Qube_API_USER Sonar Qube_URL
func NewConfigFromEnv() (cfg Config, err error) {
	if cfg.APIPWD, err = lib.GetEnvVar("SONAR_PWD", ".+"); err != nil {
		return
	}
	if cfg.APIUser, err = lib.GetEnvVar("SONAR_USER", ".+"); err != nil {
		return
	}
	if cfg.APIURL, err = lib.GetEnvVar("SONAR_URL", "http://.*:\\d*"); err != nil {
		return
	}
	if cfg.APIProjectKey, err = lib.GetEnvVar("SONAR_PROJECT_KEY", ".+"); err != nil {
		return
	}
	if cfg.APIMetrics, err = lib.GetEnvVar("SONAR_METRICS", ".+"); err != nil {
		return
	}
	return
}
