package service

type PermissionStruct struct{}

var PermissionService PermissionStruct

type Permission struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type PermissionItem struct {
	Permission Permission `json:"permission"`
	Urls       []string   `json:"urls"`
}

var Permissions = []PermissionItem{
	{Permission: Permission{Name: "首页", Key: "dashboard"}, Urls: []string{"/dashboard"}},
	{Permission: Permission{Name: "节点管理", Key: "host:read"}, Urls: []string{"/host"}},
	{Permission: Permission{Name: "节点可写", Key: "host:write"}, Urls: []string{"/host/store"}},
	{Permission: Permission{Name: "节点删除", Key: "host:delete"}, Urls: []string{"/host/remove/:id"}},
	{Permission: Permission{Name: "登录日志", Key: "login_log:read"}, Urls: []string{"/system/login-log"}},
	//{Permission: Permission{Name: "操作日志", Key: "operate_log:read"}, Urls: []string{"", ""}},
	//{Permission: Permission{Name: "操作日志删除", Key: "operate_log:delete"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "进程管理", Key: "process:read"}, Urls: []string{"/process"}},
	{Permission: Permission{Name: "进程可写", Key: "process:write"}, Urls: []string{"/process/:id", "/process/store", "/process/disable/:id", "/process/start/:id", "/process/stop/:id", "/process/restart/:id"}},
	//{Permission: Permission{Name: "进程删除", Key: "process:delete"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "项目管理", Key: "project:read"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "项目可写", Key: "project:write"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "项目删除", Key: "project:delete"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "系统设置", Key: "setting:read"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "系统设置可写", Key: "setting:write"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "定时任务", Key: "task:read"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "定时任务可写", Key: "task:write"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "定时任务删除", Key: "task:delete"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "任务日志", Key: "task_log:read"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "任务日志删除", Key: "task_log:delete"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "用户管理", Key: "user:read"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "用户可写", Key: "user:write"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "用户删除", Key: "user:delete"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "角色管理", Key: "role:read"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "角色可写", Key: "role:write"}, Urls: []string{"", ""}},
	{Permission: Permission{Name: "角色删除", Key: "role:delete"}, Urls: []string{"", ""}},
}

// Permissions
/*
[
{
	"permission":Permission
	"urls":[]
}
]
通用地址
/system/setting
/host
/host/all
/project/all
*/

func (PermissionStruct) GetAllPermissions() []Permission {
	var permissions []Permission
	for _, item := range Permissions {
		permissions = append(permissions, item.Permission)
	}
	return permissions
}

func (PermissionStruct) GetUrls(permissions []string) []string {
	var result []string
	for _, item := range Permissions {
		for _, value := range permissions {
			if value == item.Permission.Key {
				result = append(result, item.Urls...)
			}
		}
	}
	return result
}
