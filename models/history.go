package models

type SingularityTaskIdHistory struct {
	TaskId        SingularityTaskId `json:"taskId"`
	RunId         string            `json:"runId"`
	UpdatedAt     int64             `json:"updatedAt"`
	LastTaskState string            `json:"lastTaskState"`
}

type SingularityTaskId struct {
	RequestId       string `json:"requestId"`
	Host            string `json:"host"`
	DeployId        string `json:"deployId"`
	SanitizedHost   string `json:"sanitizedHost"`
	RackId          string `json:"rackId"`
	SanitizedRackId string `json:"sanitizedRackId"`
	InstanceNo      int    `json:"instanceNo"`
	StartedAt       int64  `json:"startedAt"`
	Id              string `json:"id"`
}
