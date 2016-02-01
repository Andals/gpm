/**
* @file main.go
* @brief andals-gpm main
* @author ligang
* @date 2016-02-01
 */

package main

import (
	"flag"
	"fmt"
	"gpm/pkg/commands"
	"os"
)

var debugMode bool

func main() {
	flag.BoolVar(&debugMode, "debug", false, "")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: andals-gpm install")
		return
	}

	cmd, err := commands.GetCommand(flag.Arg(0))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	prjHome, _ := os.Getwd()
	cmd.Run(prjHome)
}
