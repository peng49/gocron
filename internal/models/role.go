package models

import "github.com/go-xorm/xorm"

type Role struct {
	Id          int    `json:"id" xorm:"pk autoincr notnull "`
	Name        string `json:"name" xorm:"varchar(32) notnull unique"` // 用户名
	Permissions string `json:"permissions" xorm:"text null"`
	BaseModel   `json:"-" xorm:"-"`
}

func (role *Role) Create() (int64, error) {
	return Db.Insert(role)
}

func (role *Role) Update() (int64, error) {
	return Db.Cols(`name`).Where("id = ?", role.Id).Update(role)
}

func (role *Role) Total(params CommonMap) (int64, error) {
	session := Db.Alias("p")
	role.parseWhere(session, params)
	return session.Count(role)
}

func (role *Role) List(params CommonMap) ([]Role, error) {
	session := Db.Alias("p")
	roles := make([]Role, 0)
	role.parsePageAndPageSize(params)
	role.parseWhere(session, params)

	err := session.Limit(role.PageSize, role.pageLimitOffset()).Find(&roles)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// 解析where
func (role *Role) parseWhere(session *xorm.Session, params CommonMap) {
	if len(params) == 0 {
		return
	}
	id, ok := params["Id"]
	if ok && id.(int) > 0 {
		session.And("r.id = ?", id)
	}
	name, ok := params["Name"]
	if ok && name.(string) != "" {
		session.And("r.name LIKE ?", "%"+name.(string)+"%")
	}
}

func (role *Role) All() ([]Role, interface{}) {
	roles := make([]Role, 0)
	err := Db.Find(&roles)
	return roles, err
}
