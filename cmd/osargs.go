package main

import (
	"os"
	"strconv"
	"time"
)

//
// basic checks for user requesting help or providing too many arguments
//

func basicOsArgsCheck() {
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" {
			exitWithUsage(codeNoError)
		}
	}
	if len(os.Args)-1 > len(argDefs) {
		exitWithUsage(codeBadUsage)
	}
}

//
// reading arguments from CLI and supporting default values
//

func getDurationOrExit(argDef *argumentDefinition) time.Duration {
	duration, err := time.ParseDuration(getOsArg(argDef))
	exitWithUsageIfError(err)
	return duration
}

func getIntegerOrExit(argDef *argumentDefinition) int {
	integer, err := strconv.Atoi(getOsArg(argDef))
	exitWithUsageIfError(err)
	return integer
}

func getOsArg(argDef *argumentDefinition) string {
	if argDef.index < len(os.Args) {
		osArgValue := os.Args[argDef.index]
		if osArgValue != "-" {
			return osArgValue
		}
	}
	return argDef.defaultValue
}
