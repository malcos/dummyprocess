package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

const (
	codeBadUsage = 255
	codeNoError  = 0
)

func main() {

	// declare arguments
	durationArgDef := addArgument("duration", "0ms", "time to wait before terminating the process, supported suffixes are ms,s,m,h,d")
	exitCodeArgDef := addArgument("exit_code", "0", "the code to return when the process ends sleeping")
	signalCodeArgDef := addArgument("signal_code", "65", "the code to return when a signal is captured by the process")

	// help requested? correct number of arguments provided?
	basicOsArgsCheck()

	// read arguments from CLI or fail and go back to prompt with proper exit codes
	sleepDuration := getDurationOrExit(durationArgDef)
	exitCode := getIntegerOrExit(exitCodeArgDef)
	signalCode := getIntegerOrExit(signalCodeArgDef)

	// dispatch the OS signal listener which will terminate the process on any signal
	go osSignalListener(signalCode)

	// happy path: wait and exit
	time.Sleep(sleepDuration)
	os.Exit(exitCode)

}

//
// goroutine for OS signal processor
//

func osSignalListener(errorCode int) {
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	s := <-c
	fmt.Println(s)
	os.Exit(errorCode)
}
