package ui

import (
	"github.com/fatih/color"
	"time"
)

var red func(...interface{}) string = color.New(color.FgHiRed).SprintFunc()
var green func(...interface{}) string = color.New(color.FgGreen).SprintFunc()
var blue func(...interface{}) string = color.New(color.FgBlue).SprintFunc()
var yellow func(...interface{}) string = color.New(color.FgYellow).SprintFunc()
var bold func(...interface{}) string = color.New(color.Bold).SprintFunc()

func unixMsToHumanTime(timestamp int64) string {
	return unixSecToHumanTime(timestamp / 1000)
}

func unixSecToHumanTime(timestamp int64) string {
	humanReadableTime, err := time.Unix(timestamp, 0).MarshalText()
	if err != nil {
		return "UNKNOWN"
	}

	return string(humanReadableTime)
}
