package models

type RolePermission struct {
	Id   int    `json:"id" xorm:"pk autoincr notnull "`
	Name string `json:"name" xorm:"varchar(32) notnull unique"` // 用户名
}
