package sonar

//ComponentTreeAnswer represents an ComponentTree response from the sonar qube API
//Example url:	url := fmt.Sprintf("%v/api/measures/component_tree?additionalFields=metrics,periods&baseComponentKey=%v&metricKeys=%v",
//		SONAR_URL,
//		SONAR_PROJECT_KEY,
//		SONAR_METRICS)
type ComponentTreeAnswer struct {
	Paging struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		Total     int `json:"total"`
	} `json:"paging"`
	BaseComponent struct {
		ID        string `json:"id"`
		Key       string `json:"key"`
		Name      string `json:"name"`
		Qualifier string `json:"qualifier"`
		Measures  []struct {
			Metric  string `json:"metric"`
			Value   string `json:"value"`
			Periods []struct {
				Index int    `json:"index"`
				Value string `json:"value"`
			} `json:"periods,omitempty"`
		} `json:"measures"`
	} `json:"baseComponent"`
	Components []struct {
		ID        string `json:"id"`
		Key       string `json:"key"`
		Name      string `json:"name"`
		Qualifier string `json:"qualifier"`
		Path      string `json:"path"`
		Language  string `json:"language,omitempty"`
		Measures  []struct {
			Metric string `json:"metric"`
			Value  string `json:"value"`
		} `json:"measures"`
	} `json:"components"`
	Metrics []struct {
		Key                   string `json:"key"`
		Name                  string `json:"name"`
		Description           string `json:"description"`
		Domain                string `json:"domain"`
		Type                  string `json:"type"`
		HigherValuesAreBetter bool   `json:"higherValuesAreBetter,omitempty"`
		Qualitative           bool   `json:"qualitative"`
		Hidden                bool   `json:"hidden"`
		Custom                bool   `json:"custom"`
		DecimalScale          int    `json:"decimalScale,omitempty"`
		BestValue             string `json:"bestValue,omitempty"`
		WorstValue            string `json:"worstValue,omitempty"`
	} `json:"metrics"`
	Periods []struct {
		Index     int    `json:"index"`
		Mode      string `json:"mode"`
		Date      string `json:"date"`
		Parameter string `json:"parameter,omitempty"`
	} `json:"periods"`
}
