package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id       int
	Username string
	Password string
}

func (u *User) TableName() string {
	return "users"
}

func init() {
	orm.RegisterModel(&User{})
}

func GetUserById(id int) *User {
	if id == 0 {
		return nil
	}

	o := orm.NewOrm()
	user := User{}
	user.Id = id
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		return nil
	} else if err == orm.ErrMissPK {
		return nil
	}
	return &user
}
