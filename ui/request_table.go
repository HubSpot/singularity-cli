package ui

import (
	"git.hubteam.com/zklapow/singularity-cli/models"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func RenderRequest(req models.RequestParent) {
	RenderRequestTable([]models.RequestParent{req})
}

func RenderRequestTable(reqs []models.RequestParent) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Request ID", "State"})
	table.SetBorder(false)
	table.AppendBulk(requestToStringArray(reqs))

	table.Render()
}

func RenderActiveTasksTable(tasks []models.SingularityTaskIdHistory) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Instance #", "Host", "Started At", "State"})
	table.SetBorder(false)

	table.AppendBulk(tasksToStringArray(tasks))

	table.Render()
}

func tasksToStringArray(tasks []models.SingularityTaskIdHistory) [][]string {
	result := make([][]string, len(tasks))

	for i, task := range tasks {
		result[i] = taskToStrings(task)
	}

	return result
}

func taskToStrings(task models.SingularityTaskIdHistory) []string {
	var state string
	switch task.LastTaskState {
	case "TASK_RUNNING":
		state = green(task.LastTaskState)
	case "TASK_KILLED":
		state = red(task.LastTaskState)
	case "TASK_LAUNCHED":
		state = blue(task.LastTaskState)
	case "TASK_CLEANING":
		state = yellow(task.LastTaskState)
	case "TASK_STARTING":
		state = yellow(task.LastTaskState)
	default:
		state = task.LastTaskState
	}

	return []string{
		strconv.Itoa(task.TaskId.InstanceNo),
		task.TaskId.Host,
		unixMsToHumanTime(task.TaskId.StartedAt),
		state,
	}
}

func requestToStringArray(reqs []models.RequestParent) [][]string {
	result := make([][]string, len(reqs))

	for i, req := range reqs {
		result[i] = requestToStrings(req)
	}

	return result
}

func requestToStrings(req models.RequestParent) []string {
	var state string
	switch req.State {
	case "ACTIVE":
		state = green(req.State)
	case "PAUSED":
		state = blue(req.State)
	case "SYSTEM_COOLDOWN":
		state = red(req.State)
	default:
		state = req.State
	}

	return []string{req.Request.Id, state}
}
