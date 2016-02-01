/**
* @file install.go
* @brief install command
* @author ligang
* @date 2016-02-01
 */

package commands

import (
	"andals/gobox/error"
	"andals/gobox/shell"
	"fmt"
	"gpm/pkg/conf"
	"os"
)

type installCommand struct {
}

func (this *installCommand) Run(prjHome string) *error.Error {
	pkgConf, err := conf.ParsePackageConf(prjHome)
	if err != nil {
		return err
	}

	os.RemoveAll(conf.GetVendorRoot(prjHome))

	dependencies := pkgConf.GetDependencies()
	for name, depConf := range dependencies {
		fmt.Println("Start install " + name)

		path := depConf.GetPath()
		repConf := depConf.GetRepository()

		pkgLevelDir := path.GetLevelDir()
		pkgBaseDir := path.GetBaseDir()
		pkgPrjHome := pkgLevelDir + "/" + pkgBaseDir

		os.MkdirAll(pkgLevelDir, 0755)

		cmd := "cd " + pkgLevelDir + ";"
		cmd += "git clone " + repConf.GetUrl() + " " + pkgBaseDir + ";"
		cmd += "cd " + pkgBaseDir + ";"

		tag := repConf.GetTag()
		if tag != "" {
			cmd += "git checkout " + tag + " -b " + tag + ";"
		}
		cmd += "rm -rf .git; 2>&1"

		shell.RunCmd(cmd)

		command := new(installCommand)
		command.Run(pkgPrjHome)

		fmt.Println("End install " + name)
	}

	return nil
}
