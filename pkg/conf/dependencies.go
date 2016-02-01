package conf

import (
	"andals/gobox/error"
)

type DependConf struct {
	path       *PathConf
	repository *RepositoryConf
}

func (this *DependConf) GetPath() *PathConf {
	return this.path
}

func (this *DependConf) GetRepository() *RepositoryConf {
	return this.repository
}

func parseDependConf(prjHome string, pkgName string, repStr string) (*DependConf, *error.Error) {
	repository, err := parseRepositoryConf(repStr)
	if err != nil {
		return nil, err
	}

	conf := new(DependConf)

	conf.path = parsePathConf(prjHome, pkgName)
	conf.repository = repository

	return conf, nil
}
