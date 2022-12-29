package main

//
// defining and declaring arguments and their correct order
//

type argumentDefinition struct {
	index        int    // the index to read in os.Args
	defaultValue string // string representation of the default value to use when not enough arguments were provided or its value is "-"
	name         string // name of the argument in the "usage" page
	description  string // description of the argument for the "usage" page
}

var argDefIndex = 0
var argDefs = []*argumentDefinition{}

func addArgument(name, defaultValue, description string) *argumentDefinition {

	argDefIndex++
	argDef := &argumentDefinition{
		index:        argDefIndex,
		defaultValue: defaultValue,
		name:         name,
		description:  description,
	}

	argDefs = append(argDefs, argDef)
	return argDef

}
