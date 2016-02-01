/**
* @file base.go
* @brief commands base
* @author ligang
* @version 1.0
* @date 2016-02-01
 */

package commands

import (
	"andals/gobox/error"
	"gpm/pkg/errno"
)

const (
	COMMAND_NAME_INSTALL = "install"
)

type ICommand interface {
	Run(prjHome string) *error.Error
}

func GetCommand(name string) (ICommand, *error.Error) {
	switch name {
	case COMMAND_NAME_INSTALL:
		return new(installCommand), nil
	default:
		return nil, error.NewError(errno.E_COMMAND_INVALID_NAME, "invalid command name")
	}
}
