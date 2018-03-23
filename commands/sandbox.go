package commands

import (
	"git.hubteam.com/zklapow/singularity-cli/client"
	"fmt"
	"git.hubteam.com/zklapow/singularity-cli/models"
	"git.hubteam.com/zklapow/singularity-cli/ui"
)

func BrowseSandbox(client *client.SingularityClient, requestId, path string, instance int) {
	task, err := taskForRequest(client, requestId, instance)
	if err != nil {
		fmt.Printf("Could not load tasks for request %v: %#v", requestId, err)
		panic(err)
	}

	if task == nil {
		fmt.Printf("Could not find task with ID %v for request %v", instance, requestId)
		return
	}

	sandbox, err := client.BrowseSandbox(task.Id, path)
	if err != nil {
		fmt.Printf("Could not browse sandbox of task %v: %#v", task.Id, err)
		panic(err)
	}

	ui.RenderSandboxFileList(*sandbox)
}

func CatFile(client *client.SingularityClient, requestId, path string, instance int) {
	task, err := taskForRequest(client, requestId, instance)
	if err != nil {
		fmt.Printf("Could not load tasks for request %v: %#v", requestId, err)
		panic(err)
	}

	if task == nil {
		fmt.Printf("Could not find task with ID %v for request %v", instance, requestId)
		return
	}

	chunk, err := client.GetFileChunk(task.Id, path)
	if err != nil {
		fmt.Printf("Could not get file %#v sandbox of task %v: %#v", path, task.Id, err)
		panic(err)
	}

	fmt.Println(chunk.Data)
	//ui.RenderSandboxFileList(*sandbox)
}

func taskForRequest(client *client.SingularityClient, requestId string, instance int) (*models.SingularityTaskId, error) {
	tasks, err := client.GetActiveTasksFor(requestId)
	if err != nil {
		fmt.Printf("Could not load tasks for request %v: %#v", requestId, err)
		return nil, err
	}

	var task *models.SingularityTaskId
	for i, t := range tasks {
		if t.TaskId.InstanceNo == instance {
			task = &tasks[i].TaskId
		}
	}

	return task, nil
}
