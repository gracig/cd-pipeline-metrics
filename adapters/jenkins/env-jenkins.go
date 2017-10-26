package jenkins

import "github.com/gracig/cd-runtime-collector/lib"

//Config represents parameters to access or manages a Jenkins API
type Config struct {
	//ApiToken is the token password of the Jenkins API
	APIToken string
	//ApiUser is the user of the Jenkins API
	APIUser string
	//ApiURL is the url where the Jenkins API is hosted
	APIURL string
}

//NewConfigFromEnv creates a new JenkinsConfig from environment variables.
//Expected environment variables:
// JENKINS_API_TOKEN JENKINS_API_USER JENKINS_URL
func NewConfigFromEnv() (cfg Config, err error) {
	if cfg.APIToken, err = lib.GetEnvVar("JENKINS_API_TOKEN", ".+"); err != nil {
		return
	}
	if cfg.APIUser, err = lib.GetEnvVar("JENKINS_API_USER", ".+"); err != nil {
		return
	}
	if cfg.APIURL, err = lib.GetEnvVar("JENKINS_URL", "http://.*:\\d*"); err != nil {
		return
	}

	return
}
