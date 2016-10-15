package main

import "git.hubteam.com/zklapow/singularity-cli/ui"

func main() {
	client := NewSingularityClient("https://bootstrap.hubteam.com/singularity/v3", map[string]string{"X-HubSpot-User": "zklapow"})
	reqs, err := client.ListAllRequests()
	if err != nil {
		panic(err)
	}

	ui.RenderRequestTable(reqs)
}
