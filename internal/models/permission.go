package models

type Permission struct {
	Id       int    `json:"id" xorm:"pk autoincr notnull "`
	Name     string `json:"name" xorm:"varchar(64) notnull unique"` //权限名称
	HttpPath string `json:"http_path" xorm:"text"`                  //路径
}
