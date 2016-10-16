package models

type SingularityScaleRequest struct {
	SkipHealthchecks bool   `json:"skipHealthchecks,omitempty"`
	DurationMillis   int64  `json:"durationMillis,omitempty"`
	Message          string `json:"message,omitempty"`
	ActionId         string `json:"actionId,omitempty"`
	Instances        int    `json:"instances,omitempty"`
}
