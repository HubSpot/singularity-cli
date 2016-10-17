package client

import (
	"git.hubteam.com/zklapow/singularity-cli/models"
	"fmt"
	"path/filepath"
)

const (
	api_sandbox_browse = "/api/sandbox/%v/browse"
)

func (c *SingularityClient) BrowseSandbox(taskId, path string) (*models.SingularitySandbox, error) {
	res := &models.SingularitySandbox{}
	url := fmt.Sprintf(api_sandbox_browse, taskId)

	if path != "" {
		url += fmt.Sprintf("?path=%v", filepath.Join(taskId, path))
	}

	err := c.getJson(res, url)
	return res, err
}
