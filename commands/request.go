package commands

import (
	"fmt"
	"git.hubteam.com/zklapow/singularity-cli/client"
	"git.hubteam.com/zklapow/singularity-cli/models"
	"git.hubteam.com/zklapow/singularity-cli/ui"
	"strings"
)

func ListAllRequests(client *client.SingularityClient) {
	reqs, err := client.ListAllRequests()
	if err != nil {
		fmt.Printf("Could not load requests from singularity: %#v", err)
		panic(err)
	}

	ui.RenderRequestTable(reqs)
}

func FindRequestsMatching(client *client.SingularityClient, query string) {
	reqs, err := client.ListAllRequests()
	if err != nil {
		fmt.Printf("Could not load requests from singularity: %#v", err)
		panic(err)
	}

	suggested := []models.RequestParent{}
	for _, req := range reqs {
		if strings.Contains(req.Request.Id, query) {
			suggested = append(suggested, req)
		}
	}

	ui.RenderRequestTable(suggested)
}

func indexRequestsById(reqs []models.RequestParent) map[string]models.RequestParent {
	result := map[string]models.RequestParent{}

	for _, req := range reqs {
		result[req.Request.Id] = req
	}

	return result
}
