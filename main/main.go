/**
* @file main.go
* @brief andals-gpm main
* @author ligang
* @date 2016-02-01
 */

package main

import (
	"andals/gobox/misc"
	"andals/gobox/shell"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	PACKAGE_JSON = "package.json"
	VENDOR       = "vendor"

	SEP_SCHEME = "//"
	SEP_TAG    = "#"
)

type PkgConf struct {
	Name         string
	Dependencies map[string]string
}

type repositoryConf struct {
	scheme string
	url    string
	tag    string
}

func main() {
	prjHome, _ := os.Getwd()

	if len(os.Args) < 2 {
		fmt.Println("Usage: andals-gpm install")
		return
	}

	switch os.Args[1] {
	case "install":
		install(prjHome)
	default:
		fmt.Println("Not support")
	}
}

func install(prjHome string) {
	confPath := prjHome + "/" + PACKAGE_JSON

	if !misc.FileExist(confPath) {
		fmt.Println("There is no " + PACKAGE_JSON + " in " + prjHome)
		return
	}

	var confData PkgConf

	confJson, _ := ioutil.ReadFile(confPath)
	err := json.Unmarshal(confJson, &confData)
	if nil != err {
		fmt.Println("Parse " + PACKAGE_JSON + " error")
		return
	}

	for name, repository := range confData.Dependencies {
		fmt.Println("Start install " + name)

		repoConf := parseRepositoryConf(repository)

		pkgLevelDir := prjHome + "/" + VENDOR + "/" + filepath.Dir(name)
		pkgBaseDir := filepath.Base(name)
		pkgPrjHome := pkgLevelDir + "/" + pkgBaseDir

		os.RemoveAll(pkgPrjHome)
		os.MkdirAll(pkgLevelDir, 0755)

		cmd := "cd " + pkgLevelDir + ";"
		cmd += "git clone " + repoConf.url + " " + pkgBaseDir + ";"
		cmd += "cd " + pkgBaseDir + ";"
		if repoConf.tag != "" {
			cmd += "git checkout " + repoConf.tag + " -b " + repoConf.tag + ";"
		}
		cmd += "rm -rf .git; 2>&1"

		shell.RunCmd(cmd)

		install(pkgPrjHome)

		fmt.Println("End install " + name)
	}
}

func parseRepositoryConf(repository string) *repositoryConf {
	conf := new(repositoryConf)
	repoRune := []rune(repository)

	schemePos := strings.Index(repository, SEP_SCHEME)
	conf.scheme = string(repoRune[0:schemePos])

	url := ""
	tagPos := strings.LastIndex(repository, SEP_TAG)
	if tagPos != -1 {
		conf.tag = string(repoRune[tagPos+1:])
		url = string(repoRune[schemePos+len(SEP_SCHEME) : tagPos])
	} else {
		url = string(repoRune[schemePos+2:])
	}

	switch conf.scheme {
	case "git+https:":
		conf.url = "https://" + url
	case "git+http:":
		conf.url = "http://" + url
	case "git:":
		conf.url = "git://" + url
	default:
		conf.url = url
	}

	return conf
}
