package ui

import (
	"git.hubteam.com/zklapow/singularity-cli/models"
	"github.com/olekukonko/tablewriter"
	"os"
	"github.com/fatih/color"
)

var red func(...interface{}) string = color.New(color.FgHiRed).SprintFunc()
var green func(...interface{}) string = color.New(color.FgGreen).SprintFunc()
var blue func(...interface{}) string = color.New(color.FgBlue).SprintFunc()

func RenderRequestTable(reqs []models.RequestParent) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Request ID", "Status"})
	table.SetBorder(false)
	table.AppendBulk(requestToStringArray(reqs))

	table.Render()
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
