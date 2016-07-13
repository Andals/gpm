/**
* @file base.go
* @brief commands base
* @author ligang
* @version 1.0
* @date 2016-02-01
 */

package commands

import (
	"andals/gobox/exception"
	"andals/gobox/log"
	"gpm/pkg/errno"
)

const (
	COMMAND_NAME_INSTALL = "install"
)

type ICommand interface {
	Run(prjHome string, logger log.ILogger) *exception.Exception
}

func GetCommand(name string) (ICommand, *exception.Exception) {
	switch name {
	case COMMAND_NAME_INSTALL:
		return new(installCommand), nil
	default:
		return nil, exception.New(errno.E_COMMAND_INVALID_NAME, "invalid command name")
	}
}
