package service_test

import (
	"github.com/ouqiang/gocron/internal/service"
	"regexp"
	"testing"
)

func TestPermissionStruct_GetAllPermissions(t *testing.T) {
	t.Log(service.PermissionService.GetAllPermissions())
}

func TestRouteMatch(t *testing.T) {
	permissions := service.PermissionService.GetAllPermissions()
	var keys []string
	for _, permission := range permissions {
		keys = append(keys, permission.Key)
	}
	urls := service.PermissionService.GetUrls(keys)
	re, _ := regexp.Compile(":\\w+") //url参数占位符替换成正则
	regexpGroup := make([]*regexp.Regexp, 0)
	for _, url := range urls {
		if url == "" {
			continue
		}
		pattern, _ := regexp.Compile(re.ReplaceAllString(url, "[^/]*") + "$")
		regexpGroup = append(regexpGroup, pattern)
	}

	path := "/dashboard/123"
	allow := false
	for _, p := range regexpGroup {
		if p.Match([]byte(path)) {
			allow = true
			break
		}
	}
	t.Log(allow)

	//uri := re.ReplaceAllString("/project/edit/:id", "[^/]*") + "$"
	//t.Log(uri)
	//url := "/project/edit/12123"
	//pattern, _ := regexp.Compile(uri)
	//
	//t.Log(pattern.Match([]byte(url)))
}
