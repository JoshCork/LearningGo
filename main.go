package main

import (
	"fmt"

	"github.com/joshcork/LearningGo/samplePackages/commandLineFlags"
	"github.com/joshcork/LearningGo/samplePackages/datetimeManipulations"
	"github.com/joshcork/LearningGo/samplePackages/stringutil"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))

	commandLineFlags.CommandLineFlagsExamples()
	datetimeManipulation.DatetimeManipulationExamples()
	datetimeManipulation.ZoneTime()
	datetimeManipulation.TimeDiff()
	datetimeManipulation.DateAddition()

}
