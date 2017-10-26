package jenkins

import (
	"log"

	"github.com/gracig/cd-runtime-collector/lib"
)

//NewOperations returns an Operations object
func NewOperations() (ops *Operations, err error) {
	var cfg Config

	if cfg, err = NewConfigFromEnv(); err != nil {
		log.Printf("Could not get Jenkins config from environment. %v", err)
		return
	}
	ops = &Operations{Cfg: cfg}
	return
}

//Operations enables actions upon jenkins server
type Operations struct {
	Cfg Config
}

//GetBasicWebGetter retrieves the Hudson Response from Jenkins API
func (ops *Operations) GetBasicWebGetter() (response lib.WebGetter) {
	return &lib.BasicAuthWebGetter{User: ops.Cfg.APIUser, Pwd: ops.Cfg.APIToken}
}

//GetHudson retrieves the Hudson Response from Jenkins API
func (ops *Operations) GetHudson(wg lib.WebGetter) (response *Hudson, err error) {
	url := ops.Cfg.APIURL + "/api/json?pretty=true"
	response = new(Hudson)
	err = wg.DecodeRestJSON(response, url)
	return
}

//GetOrganizationFolder retrieves the OrganizationFolder from Jenkins API
func (ops *Operations) GetOrganizationFolder(wg lib.WebGetter, path string) (response *OrganizationFolder, err error) {
	url := path + "/api/json?pretty=true"
	response = new(OrganizationFolder)
	err = wg.DecodeRestJSON(response, url)
	return
}

//GetWorkflowMultiBranchProject retrieves the WorkflowMultiBranchProject from Jenkins API
func (ops *Operations) GetWorkflowMultiBranchProject(wg lib.WebGetter, path string) (response *WorkflowMultiBranchProject, err error) {
	url := path + "/api/json?pretty=true"
	response = new(WorkflowMultiBranchProject)
	err = wg.DecodeRestJSON(response, url)
	return
}

//GetWorkflowJob retrieves the WorkflowJob from Jenkins API
func (ops *Operations) GetWorkflowJob(wg lib.WebGetter, path string) (response *WorkflowJob, err error) {
	url := path + "/api/json?pretty=true"
	response = new(WorkflowJob)
	err = wg.DecodeRestJSON(response, url)
	return
}

//GetWorkflowRun retrieves the WorkflowRun from Jenkins API
func (ops *Operations) GetWorkflowRun(wg lib.WebGetter, path string) (response *WorkflowRun, err error) {
	url := path + "/api/json?pretty=true"
	response = new(WorkflowRun)
	err = wg.DecodeRestJSON(response, url)
	return
}

//GetWorkflowAPIAnswer retrieves the WorkflowAPIAnswer from Jenkins API
func (ops *Operations) GetWorkflowAPIAnswer(wg lib.WebGetter, path string) (response *WorkflowAPIAnswer, err error) {
	url := path + "/wfapi/"
	response = new(WorkflowAPIAnswer)
	err = wg.DecodeRestJSON(response, url)
	return
}
