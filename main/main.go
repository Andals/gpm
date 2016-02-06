/**
* @file main.go
* @brief andals-gpm main
* @author ligang
* @date 2016-02-01
 */

package main

import (
	"andals/gobox/log"
	logWriter "andals/gobox/log/writer"
	"flag"
	"fmt"
	"gpm/conf"
	"gpm/pkg/commands"
	"os"
)

func main() {
	if flag.NArg() < 1 {
		fmt.Println("Usage: andals-gpm install")
		return
	}

	cmd, err := commands.GetCommand(flag.Arg(0))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	writer, _ := logWriter.NewFileWriter(conf.GetLogPath())
	logger, _ := log.NewSimpleLogger(writer, conf.GetLogLevel())

	prjHome, _ := os.Getwd()
	cmd.Run(prjHome, logger)
}
