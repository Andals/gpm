package conf

import (
	"andals/gobox/error"
	"gpm/pkg/errno"
	"strings"
)

type RepositoryConf struct {
	scheme string
	url    string
	tag    string
}

func parseRepositoryConf(repStr string) (*RepositoryConf, *error.Error) {
	repConf := new(RepositoryConf)
	repoRune := []rune(repStr)

	schemePos := strings.Index(repStr, SEP_SCHEME)
	if schemePos == -1 {
		return nil, error.NewError(errno.E_CONF_REPOSITORY_STR_ERROR, "invalid repository str: don't have scheme pos")
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
		return nil, error.NewError(errno.E_CONF_REPOSITORY_STR_ERROR, "invalid repository str: undefined scheme")
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
