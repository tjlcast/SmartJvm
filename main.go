package main

import (
"fmt"
"strings"
"jvm_pro/classpath"
)

/**
	project args:


 */

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v \nclass:%v \nargs:%v \n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	// className: java/lang/Object
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", className)
	}
	fmt.Printf("class data: %v\n", classData)
}

