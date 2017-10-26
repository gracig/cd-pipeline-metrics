package jenkins

//WorkflowAPIAnswer is the answer when accessing the wfapi from a pipeline job.
//URL sample: http://jenkins/job/TIM/job/devops-tim-rede-inge-cdr-gprs_ericsson/job/master/20/wfapi/
type WorkflowAPIAnswer struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Artifacts struct {
			Href string `json:"href"`
		} `json:"artifacts"`
	} `json:"_links"`
	ID                  string `json:"id"`
	Name                string `json:"name"`
	Status              string `json:"status"`
	StartTimeMillis     int64  `json:"startTimeMillis"`
	EndTimeMillis       int    `json:"endTimeMillis"`
	DurationMillis      int    `json:"durationMillis"`
	QueueDurationMillis int    `json:"queueDurationMillis"`
	PauseDurationMillis int    `json:"pauseDurationMillis"`
	Stages              []struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		ID                  string `json:"id"`
		Name                string `json:"name"`
		ExecNode            string `json:"execNode"`
		Status              string `json:"status"`
		StartTimeMillis     int64  `json:"startTimeMillis"`
		DurationMillis      int    `json:"durationMillis"`
		PauseDurationMillis int    `json:"pauseDurationMillis"`
		Error               struct {
			Message string `json:"message"`
			Type    string `json:"Type"`
		} `json:"error"`
	} `json:"stages"`
}

//WorkflowRun represents a Running instance of a job
//Example url: http://jenkins/job/TIM/job/devops-tim-rede-inge-cdr-gprs_ericsson/job/master/20/api/json?pretty=true
type WorkflowRun struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class  string `json:"_class,omitempty"`
		Causes []struct {
			Class            string `json:"_class"`
			ShortDescription string `json:"shortDescription"`
			UserID           string `json:"userId"`
			UserName         string `json:"userName"`
		} `json:"causes,omitempty"`
		BuildsByBranchName struct {
			Master struct {
				Class       string      `json:"_class"`
				BuildNumber int         `json:"buildNumber"`
				BuildResult interface{} `json:"buildResult"`
				Marked      struct {
					SHA1   string `json:"SHA1"`
					Branch []struct {
						SHA1 string `json:"SHA1"`
						Name string `json:"name"`
					} `json:"branch"`
				} `json:"marked"`
				Revision struct {
					SHA1   string `json:"SHA1"`
					Branch []struct {
						SHA1 string `json:"SHA1"`
						Name string `json:"name"`
					} `json:"branch"`
				} `json:"revision"`
			} `json:"master"`
		} `json:"buildsByBranchName,omitempty"`
		LastBuiltRevision struct {
			SHA1   string `json:"SHA1"`
			Branch []struct {
				SHA1 string `json:"SHA1"`
				Name string `json:"name"`
			} `json:"branch"`
		} `json:"lastBuiltRevision,omitempty"`
		RemoteUrls []string `json:"remoteUrls,omitempty"`
		ScmName    string   `json:"scmName,omitempty"`
		FailCount  int      `json:"failCount,omitempty"`
		SkipCount  int      `json:"skipCount,omitempty"`
		TotalCount int      `json:"totalCount,omitempty"`
		URLName    string   `json:"urlName,omitempty"`
	} `json:"actions"`
	Artifacts []struct {
		DisplayPath  string `json:"displayPath"`
		FileName     string `json:"fileName"`
		RelativePath string `json:"relativePath"`
	} `json:"artifacts"`
	Building          bool          `json:"building"`
	Description       interface{}   `json:"description"`
	DisplayName       string        `json:"displayName"`
	Duration          int           `json:"duration"`
	EstimatedDuration int           `json:"estimatedDuration"`
	Executor          interface{}   `json:"executor"`
	FullDisplayName   string        `json:"fullDisplayName"`
	ID                string        `json:"id"`
	KeepLog           bool          `json:"keepLog"`
	Number            int           `json:"number"`
	QueueID           int           `json:"queueId"`
	Result            string        `json:"result"`
	Timestamp         int64         `json:"timestamp"`
	URL               string        `json:"url"`
	ChangeSets        []interface{} `json:"changeSets"`
	NextBuild         struct {
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"nextBuild"`
	PreviousBuild struct {
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"previousBuild"`
}

//WorkflowJob represents a pipeline job
//Example url: http://jenkins/job/TIM/job/devops-tim-rede-inge-cdr-gprs_ericsson/job/master/api/json?pretty=true
type WorkflowJob struct {
	Class             string      `json:"_class"`
	Description       interface{} `json:"description"`
	DisplayName       string      `json:"displayName"`
	DisplayNameOrNull interface{} `json:"displayNameOrNull"`
	Name              string      `json:"name"`
	URL               string      `json:"url"`
	Buildable         bool        `json:"buildable"`
	Builds            []struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"builds"`
	Color      string `json:"color"`
	FirstBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"firstBuild"`
	HealthReport []struct {
		Description   string `json:"description"`
		IconClassName string `json:"iconClassName"`
		IconURL       string `json:"iconUrl"`
		Score         int    `json:"score"`
	} `json:"healthReport"`
	InQueue          bool `json:"inQueue"`
	KeepDependencies bool `json:"keepDependencies"`
	LastBuild        struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastBuild"`
	LastCompletedBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastCompletedBuild"`
	LastFailedBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastFailedBuild"`
	LastStableBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastStableBuild"`
	LastSuccessfulBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastSuccessfulBuild"`
	LastUnstableBuild     interface{} `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastUnsuccessfulBuild"`
	NextBuildNumber int `json:"nextBuildNumber"`
	Property        []struct {
		Class  string `json:"_class"`
		Branch struct {
		} `json:"branch"`
	} `json:"property"`
	QueueItem       interface{} `json:"queueItem"`
	ConcurrentBuild bool        `json:"concurrentBuild"`
}

//WorkflowMultiBranchProject represents a Workflow Multi Branch Project.
//Example url: http://jenkins/job/TIM/job/devops-tim-rede-inge-cdr-gprs_ericsson/api/json?pretty=true
type WorkflowMultiBranchProject struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class string `json:"_class,omitempty"`
	} `json:"actions"`
	Description       interface{} `json:"description"`
	DisplayName       string      `json:"displayName"`
	DisplayNameOrNull interface{} `json:"displayNameOrNull"`
	Name              string      `json:"name"`
	URL               string      `json:"url"`
	HealthReport      []struct {
		Description   string `json:"description"`
		IconClassName string `json:"iconClassName"`
		IconURL       string `json:"iconUrl"`
		Score         int    `json:"score"`
	} `json:"healthReport"`
	Jobs []struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
		Color string `json:"color"`
	} `json:"jobs"`
	Views []struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"views"`
	PrimaryView struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"primaryView"`
}

//OrganizationFolder Represents a Organization Folder class that is linked against a git repository
//Example url: http://jenkins/job/TIM/api/json?pretty=true
type OrganizationFolder struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class string `json:"_class,omitempty"`
	} `json:"actions"`
	Description       string `json:"description"`
	DisplayName       string `json:"displayName"`
	DisplayNameOrNull string `json:"displayNameOrNull"`
	Name              string `json:"name"`
	URL               string `json:"url"`
	HealthReport      []struct {
		Description   string `json:"description"`
		IconClassName string `json:"iconClassName"`
		IconURL       string `json:"iconUrl"`
		Score         int    `json:"score"`
	} `json:"healthReport"`
	Jobs []struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"jobs"`
	PrimaryView struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"primaryView"`
	Views []struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"views"`
}

//Hudson represents the Jenkins root folder
//Example url: http://jenkins/api/json?pretty=true
type Hudson struct {
	Class          string `json:"_class"`
	AssignedLabels []struct {
	} `json:"assignedLabels"`
	Mode            string      `json:"mode"`
	NodeDescription string      `json:"nodeDescription"`
	NodeName        string      `json:"nodeName"`
	NumExecutors    int         `json:"numExecutors"`
	Description     interface{} `json:"description"`
	Jobs            []struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"jobs"`
	OverallLoad struct {
	} `json:"overallLoad"`
	PrimaryView struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"primaryView"`
	QuietingDown   bool `json:"quietingDown"`
	SlaveAgentPort int  `json:"slaveAgentPort"`
	UnlabeledLoad  struct {
		Class string `json:"_class"`
	} `json:"unlabeledLoad"`
	UseCrumbs   bool `json:"useCrumbs"`
	UseSecurity bool `json:"useSecurity"`
	Views       []struct {
		Class string `json:"_class"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"views"`
}
