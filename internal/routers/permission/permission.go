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

func EditRole(ctx *macaron.Context, role RoleForm) string {
	resp := utils.JsonResponse{}
	return resp.Success("Success", map[string]interface{}{
		"data": "0",
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
