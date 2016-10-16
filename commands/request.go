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

func ShowRequestDetails(client *client.SingularityClient, requestId string) {
	tasks, err := client.GetActiveTasksFor(requestId)
	if err != nil {
		fmt.Printf("Could not load request from singularity: %#v", err)
		panic(err)
	}

	ui.RenderActiveTasksTable(tasks)
}

func PauseRequest(client *client.SingularityClient, requestId string) {
	request, err := client.PauseRequest(requestId)
	if err != nil {
		fmt.Printf("Could not pause %v: %#v", requestId, err)
		panic(err)
	}

	ui.RenderRequest(*request)
}

func UnPauseRequest(client *client.SingularityClient, requestId string) {
	request, err := client.UnPauseRequest(requestId)
	if err != nil {
		fmt.Printf("Could not unpause %v: %#v", requestId, err)
		panic(err)
	}

	ui.RenderRequest(*request)
}
