package controller

import (
	"encoding/json"
	"fmt"
	"gg/day5_project/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type AdminController struct {
	//自动为每个请求都绑定上下文
	Ctx iris.Context
	//admin实体
	Service service.AdminService
	//session对象
	Session *sessions.Session
}
type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// func (ac *AdminController) GetInfo() mvc.Result {
// 	fmt.Println("7777777777777777777777666666")
// 	return mvc.Response{
// 		Object: map[string]interface{}{
// 			"status":  "0",
// 			"success": "失败",
// 			"message": "填写用户名密码",
// 		},
// 	}
// }

func (ac *AdminController) PostLogin(context iris.Context) mvc.Result { //mvc自动匹配
	fmt.Println("66666666666666666666666666666666666666666")
	iris.New().Logger().Info("admin login")
	fmt.Println("admin")
	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)
	//信息校验
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "失败",
				"message": "填写用户名密码",
			},
		}
	}
	admin, exist := ac.Service.GetByAdminNmaeAndPassword(adminLogin.UserName, adminLogin.Password)
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  "0",
				"success": "失败",
				"message": "用户名密码错误",
			},
		}
	}
	userByte, _ := json.Marshal(admin) //成功了就把userByte就找ac.Session用set方法写入会话session中的key:"admin"对应的值里
	ac.Session.Set("admin", userByte)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  "1",
			"success": "成功",
			"message": "成功了",
		},
	}

}
