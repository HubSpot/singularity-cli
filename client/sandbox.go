package client

import (
	"git.hubteam.com/zklapow/singularity-cli/models"
)

const (
	api_sandbox_browse = "/api/sandbox/%v/browse"
)

func (c *SingularityClient) BrowseSandbox(taskId string) (*models.SingularitySandbox, error) {
	res := &models.SingularitySandbox{}
	err := c.getJson(res, api_sandbox_browse, taskId)
	return res, err
}
