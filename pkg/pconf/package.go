/**
* @file dependencies.go
* @brief package conf main
* @author ligang
* @date 2016-02-06
 */

package pconf

import (
	"andals/gobox/error"
	"andals/gobox/misc"
	"encoding/json"
	"fmt"
	"gpm/pkg/errno"
	"io/ioutil"
)

type packageJson struct {
	Name         string
	Dependencies map[string]string
}

type PackageConf struct {
	name         string
	dependencies map[string]*DependConf
}

func ParsePackageConf(prjHome string) (*PackageConf, *error.Error) {
	confPath := getPackageJsonPath(prjHome)
	if !misc.FileExist(confPath) {
		return nil, error.NewError(errno.E_CONF_PACKAGE_JSON_NOT_EXISTS, "There is no "+PACKAGE_JSON+" in "+prjHome)
	}

	var pkgJson packageJson

	jsonStr, _ := ioutil.ReadFile(confPath)
	err := json.Unmarshal(jsonStr, &pkgJson)
	if nil != err {
		return nil, error.NewError(errno.E_CONF_PACKAGE_JSON_PARSE_ERROR, "Parse "+PACKAGE_JSON+" error")
	}

	return parseByJson(prjHome, &pkgJson)
}

func GetVendorRoot(prjHome string) string {
	return prjHome + "/" + VENDOR
}

func (this *PackageConf) GetName() string {
	return this.name
}

func (this *PackageConf) GetDependencies() map[string]*DependConf {
	return this.dependencies
}

func parseByJson(prjHome string, pkgJson *packageJson) (*PackageConf, *error.Error) {
	pkgConf := new(PackageConf)

	if pkgJson.Name == "" {
		return nil, error.NewError(errno.E_CONF_PACKAGE_NAME_ERROR, "Parse package name error")
	}

	pkgConf.name = pkgJson.Name

	pkgConf.dependencies = make(map[string]*DependConf, len(pkgJson.Dependencies))
	for name, repStr := range pkgJson.Dependencies {
		depConf, err := parseDependConf(prjHome, name, repStr)
		if err != nil {
			fmt.Println("Parse depend conf error:" + err.Error())
			continue
		}
		pkgConf.dependencies[name] = depConf
	}

	return pkgConf, nil
}

func getPackageJsonPath(prjHome string) string {
	return prjHome + "/" + PACKAGE_JSON
}
