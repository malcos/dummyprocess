package main

import (
	"fmt"
	"os"
)

//
// CLI output for errors and usage
//

func exitWithUsageIfError(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		exitWithUsage(codeBadUsage)
	}
}

func exitWithUsage(errorCode int) {
	usage()
	os.Exit(errorCode)
}

func usage() {

	fmt.Print(os.Args[0])
	for _, def := range argDefs {
		fmt.Printf(" <%v>", def.name)
	}
	fmt.Println()

	for _, def := range argDefs {
		fmt.Printf("    <%v> : %v. Default: %v\n", def.name, def.description, def.defaultValue)
	}

	fmt.Println()
	fmt.Println("Specifying \"-\" as a parameter will cause to use its default value.")
	fmt.Println()

}
