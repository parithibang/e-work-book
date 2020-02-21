package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	prefix := beego.AppConfig.String("db_prefix")
	orm.RegisterModelWithPrefix(prefix, new(Pods))
	orm.RegisterModelWithPrefix(prefix, new(Roles))
	orm.RegisterModelWithPrefix(prefix, new(Teams))
	orm.RegisterModelWithPrefix(prefix, new(Users))
	orm.RegisterModelWithPrefix(prefix, new(Units))
	orm.RegisterModelWithPrefix(prefix, new(Projects))
	orm.RegisterModelWithPrefix(prefix, new(UsersRoles))
	orm.RegisterModelWithPrefix(prefix, new(Permissions))
	orm.RegisterModelWithPrefix(prefix, new(PodsProjects))
	orm.RegisterModelWithPrefix(prefix, new(RolesPermissions))
	orm.RegisterModelWithPrefix(prefix, new(UsersProjects))
}
