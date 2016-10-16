package models

type Deploy struct {
	CustomExecutorId                  string                       `json:"customExecutorId"`
	Uris                              []string                     `json:"uris"`
	LoadBalancerDomains               []string                     `json:"loadBalancerDomains"`
	Arguments                         []string                     `json:"arguments"`
	TaskEnv                           map[string]map[string]string `json:"taskEnv"`
	AutoAdvanceDeploySteps            bool                         `json:"autoAdvanceDeploySteps"`
	ServiceBasePath                   string                       `json:"serviceBasePath"`
	CustomExecutorSource              string                       `json:"customExecutorSource"`
	Metadata                          map[string]string            `json:"metadata"`
	HealthcheckMaxRetries             int                          `json:"healthcheckMaxRetries"`
	HealthcheckTimeoutSeconds         int64                        `json:"healthcheckTimeoutSeconds"`
	HealthcheckProtocol               string                       `json:"healthcheckProtocol"`
	TaskLabels                        map[string]string            `json:"taskLabels"`
	HealthcheckPortIndex              int                          `json:"healthcheckPortIndex"`
	HealthcheckMaxTotalTimeoutSeconds int64                        `json:"healthcheckMaxTotalTimeoutSeconds"`
	LoadBalancerServiceIdOverride     string                       `json:"loadBalancerServiceIdOverride"`
	Labels                            map[string]string            `json:"labels"`
	HealthcheckUri                    string                       `json:"healthcheckUri"`
	User                              string                       `json:"user"`
	RequestId                         string                       `json:"requestId"`
	Command                           string                       `json:"command"`
	Timestamp                         int64                        `json:"timestamp"`
	Env                               map[string]string            `json:"env"`
}
