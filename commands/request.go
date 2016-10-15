package commands

import (
	"git.hubteam.com/zklapow/singularity-cli/client"
	"git.hubteam.com/zklapow/singularity-cli/ui"
	"fmt"
)

func ListAllRequests(client *client.SingularityClient) {
	reqs, err := client.ListAllRequests()
	if err != nil {
		fmt.Printf("Could not load requests from singularity: %#v", err)
		panic(err)
	}

	ui.RenderRequestTable(reqs)
}
