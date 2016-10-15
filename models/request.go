package models

type RequestParent struct {
	State         string `json:"state"`
	PendingDeploy Deploy `json:"pendingDeploy"`
	ActiveDeploy  Deploy `json:"activeDeploy"`
	Request       Request `json:"request"`
}

type Request struct {
	Id                     string `json:"id"`
	Owners                 []string `json:"owners"`
	AllowedSlaveAttributes map[string]string `json:"allowedSlaveAttributes"`
	RackSensitive          bool `json:"rackSensitive"`
	Group                  string `json:"group"`
	BounceAfterScale       bool `json:"bounceAfterScale"`
	RackAffinity           []string `json:"rackAffinity"`
}
