package conf

import (
	"path/filepath"
)

type PathConf struct {
	levelDir string
	baseDir  string
}

func parsePathConf(prjHome string, pkgName string) *PathConf {
	conf := new(PathConf)

	conf.levelDir = prjHome + "/" + VENDOR + "/" + filepath.Dir(pkgName)
	conf.baseDir = filepath.Base(pkgName)

	return conf
}

func (this *PathConf) GetLevelDir() string {
	return this.levelDir
}

func (this *PathConf) GetBaseDir() string {
	return this.baseDir
}
