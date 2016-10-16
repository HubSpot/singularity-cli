package ui

import (
	"fmt"
	"git.hubteam.com/zklapow/singularity-cli/models"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
)

func RenderSandboxFileList(sandbox models.SingularitySandbox) {
	fmt.Println(sandbox.SlaveHostname)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetColumnSeparator("")

	table.AppendBulk(sandboxToStrings(sandbox))

	table.Render()
}

func sandboxToStrings(sandbox models.SingularitySandbox) [][]string {
	result := make([][]string, len(sandbox.Files))
	for i, file := range sandbox.Files {
		result[i] = fileToStrings(file)
	}

	return result
}

func fileToStrings(file models.SingularitySandboxFile) []string {

	name := file.Name
	if strings.HasPrefix(file.Mode, "d") {
		name = bold(name)
	} else if strings.HasSuffix(file.Mode, "x") {
		name = red(name)
	}

	return []string{
		file.Mode,
		strconv.FormatUint(file.Size, 10),
		unixSecToHumanTime(file.Mtime),
		name,
	}
}
