/**
* @file dependencies.go
* @brief package conf pkg repository
* @author ligang
* @date 2016-02-06
 */

package pconf

import (
	"andals/gobox/exception"
	"gpm/pkg/errno"
	"strings"
)

type RepositoryConf struct {
	scheme string
	url    string
	tag    string
}

func parseRepositoryConf(repStr string) (*RepositoryConf, *exception.Exception) {
	repConf := new(RepositoryConf)
	repoRune := []rune(repStr)

	schemePos := strings.Index(repStr, SEP_SCHEME)
	if schemePos == -1 {
		return nil, exception.New(errno.E_CONF_REPOSITORY_STR_ERROR, "invalid repository str: don't have scheme pos")
	}
	repConf.scheme = string(repoRune[0:schemePos])

	url := ""
	tagPos := strings.LastIndex(repStr, SEP_TAG)
	if tagPos != -1 {
		repConf.tag = string(repoRune[tagPos+1:])
		url = string(repoRune[schemePos+len(SEP_SCHEME) : tagPos])
	} else {
		url = string(repoRune[schemePos+2:])
	}

	switch repConf.scheme {
	case "git+ssh:":
		repConf.url = url
	case "git+https:":
		repConf.url = "https://" + url
	case "git+http:":
		repConf.url = "http://" + url
	case "git:":
		repConf.url = "git://" + url
	default:
		return nil, exception.New(errno.E_CONF_REPOSITORY_STR_ERROR, "invalid repository str: undefined scheme")
	}

	return repConf, nil
}

func (this *RepositoryConf) GetScheme() string {
	return this.scheme
}

func (this *RepositoryConf) GetUrl() string {
	return this.url
}

func (this *RepositoryConf) GetTag() string {
	return this.tag
}
