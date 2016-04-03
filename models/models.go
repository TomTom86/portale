package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//AuthUser struct for login manager
type AuthUser struct {
	ID            int    `orm:"auto"`
	First         string `orm:"size(20)"`
	Last          string `orm:"size(20)"`
	Email         string `orm:"unique"`
	Password      string `orm:"size(60)"`
	IsApproved    bool
	IDkey         string    `orm:"size(20)"`
	RegDate       time.Time `orm:"auto_now_add;type(datetime)"`
	LastLoginDate time.Time `orm:"auto_now_add;type(datetime)"`
	LastEditDate  time.Time `orm:"auto_now_add;type(datetime)"`
	ResetKey      string    `orm:"size(20)"`
	BlockControll int
	Group         int
	Note          string   `orm:"size(100)"`
	AuthApp       *AuthApp `orm:"rel(one)"`
}

//AuthApp struct for app manager
type AuthApp struct {
	ID        int
	Automezzi bool
	Servizi   bool
	AuthUser  *AuthUser `orm:"reverse(one)"`
}



//*************FINE DB AUTOMEZZI****************
func init() {
	//Login & App Manager DB
	orm.RegisterModel(new(AuthUser), new(AuthApp))
}
