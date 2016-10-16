package ui

import (
	"git.hubteam.com/zklapow/singularity-cli/models"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"time"
)

var red func(...interface{}) string = color.New(color.FgHiRed).SprintFunc()
var green func(...interface{}) string = color.New(color.FgGreen).SprintFunc()
var blue func(...interface{}) string = color.New(color.FgBlue).SprintFunc()

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
		break
	default:
		state = task.LastTaskState
		break
	}

	humanReadableStartTime, err := time.Unix(task.TaskId.StartedAt/1000, 0).MarshalText()
	if err != nil {
		humanReadableStartTime = []byte("UNKNOWN")
	}

	return []string{strconv.Itoa(task.TaskId.InstanceNo), task.TaskId.Host, string(humanReadableStartTime), state}
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
		break
	case "PAUSED":
		state = blue(req.State)
		break
	case "SYSTEM_COOLDOWN":
		state = red(req.State)
		break
	default:
		state = req.State
		break
	}

	return []string{req.Request.Id, state}
}
