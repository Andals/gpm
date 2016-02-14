/**
* @file install.go
* @brief install command
* @author ligang
* @date 2016-02-01
 */

package commands

import (
	"andals/gobox/color"
	"andals/gobox/error"
	"andals/gobox/log"
	"andals/gobox/misc"
	"andals/gobox/shell"
	"gpm/pkg/pconf"
	"os"
)

type installCommand struct {
}

func (this *installCommand) Run(prjHome string, logger log.ILogger) *error.Error {
	pkgConf, err := pconf.ParsePackageConf(prjHome)
	if err != nil {
		logger.Debug([]byte(err.GetMsg() + "\n"))

		return err
	}

	logger.Debug([]byte("Remove vender root\n"))
	os.RemoveAll(pconf.GetVendorRoot(prjHome))

	dependencies := pkgConf.GetDependencies()
	for name, depConf := range dependencies {
		goSrcRoot := os.Getenv("GOPATH") + "/src"
		if misc.DirExist(goSrcRoot + "/" + name) {
			logger.Warning(color.Red([]byte("There is " + name + " in go src root\n")))
		}

		logger.Info([]byte("Start install " + name + "\n"))

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

		result := shell.RunCmd(cmd)
		logger.Debug([]byte(result.Output))

		command := new(installCommand)
		command.Run(pkgPrjHome, logger)

		logger.Info([]byte("End install " + name + "\n"))
	}

	return nil
}
