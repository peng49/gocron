package permission

import (
	"github.com/ouqiang/gocron/internal/models"
	"github.com/ouqiang/gocron/internal/modules/logger"
	"github.com/ouqiang/gocron/internal/modules/utils"
	"github.com/ouqiang/gocron/internal/routers/base"
	"github.com/ouqiang/gocron/internal/service"
	"gopkg.in/macaron.v1"
)

type RoleForm struct {
	Id          int
	Slug        string `binding:"Required"`
	Name        string `binding:"Required"`
	Permissions string
}

func Roles(ctx *macaron.Context) string {
	role := new(models.Role)
	queryParams := parseQueryParams(ctx)
	total, err := role.Total(queryParams)
	if err != nil {
		logger.Error(err)
	}
	roles, err := role.List(queryParams)
	if err != nil {
		logger.Error(err)
	}

	resp := utils.JsonResponse{}
	return resp.Success("Success", map[string]interface{}{
		"data":        roles,
		"total":       total,
		"permissions": service.PermissionService.GetAllPermissions(),
	})
}

func AllRole(ctx *macaron.Context) string {
	role := new(models.Role)
	roles, _ := role.AllList()
	resp := utils.JsonResponse{}
	return resp.Success("Success", roles)
}

func EditRole(ctx *macaron.Context, form RoleForm) string {
	role := models.Role{}
	role.Id = form.Id
	role.Slug = form.Slug
	role.Name = form.Name
	role.Permissions = form.Permissions
	if form.Id > 0 {
		role.Update()
	} else {
		role.Create()
	}
	resp := utils.JsonResponse{}
	return resp.Success("Success", map[string]interface{}{
		"data": role,
	})
}

// 解析查询参数
func parseQueryParams(ctx *macaron.Context) models.CommonMap {
	var params = models.CommonMap{}
	params["Id"] = ctx.QueryInt("id")
	params["Name"] = ctx.QueryTrim("name")
	base.ParsePageAndPageSize(ctx, params)

	return params
}
