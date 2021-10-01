package service

import (
	"fmt"
	"gg/day5_project/model"

	"github.com/go-xorm/xorm"
)

type AdminService interface {
	//通过管理员用户名+密码 获取管理员实体 如果查询到，返回管理员实体，并返回true
	GetByAdminNmaeAndPassword(username, password string) (model.Admin, bool)
}

//管理员服务实现的结构体
type adminService struct {
	engine *xorm.Engine
}

var admin model.Admin

func (ac *adminService) GetByAdminNmaeAndPassword(username, password string) (model.Admin, bool) {
	ac.engine.Where("admin_name=? and pwd=?", username, password).Get(&admin)
	fmt.Println("YYYYYYYYYYYYYYYYYYY")
	return admin, admin.AdminID == 1
}

//调用此函数得到一个接口类型的变量，此变量有查询用户+密码是否正确的方法
func NewAdminservice(db *xorm.Engine) AdminService {
	fmt.Println("DDDDDDDDDDDD")
	return &adminService{
		engine: db,
	}
}
