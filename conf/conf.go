/**
* @file conf.go
* @brief global server conf
* @author ligang
* @date 2016-02-06
 */

package conf

import (
	"andals/gobox/log"
	"flag"
	//     "fmt"
)

type serverConf struct {
	logLevel int
	logPath  string
}

var confData serverConf

func init() {
	flag.IntVar(&confData.logLevel, "logLevel", log.LEVEL_INFO, "log level")

	flag.Parse()
}

func GetLogLevel() int {
	return confData.logLevel
}
