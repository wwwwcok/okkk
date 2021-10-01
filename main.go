package main

import (
	"fmt"
	"gg/day5_project/config"
	"gg/day5_project/controller"
	"gg/day5_project/datasource"
	"gg/day5_project/service"
	"time"

	//"unicode/utf8"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := newapp()
	configuration(app)
	mvcHandle(app)
	config := config.InitConfig()
	addr := ":" + config.Port
	fmt.Println(addr)
	app.Run(iris.Addr(addr))

}

func newapp() *iris.Application {
	app := iris.New()
	app.HandleDir("/static", "./static") //把访问到/static地址的映射到/static
	app.HandleDir("/manage/static", "./static")
	app.HandleDir("/img", "./static/img")
	app.RegisterView(iris.HTML("./static/", ".html")) //注册当前目录的static下的.html视图文件
	app.Get("/", func(ctx iris.Context) {             //默认加载时访问/时打开视图页面
		ctx.View("index.html")
	})
	return app
}

func configuration(app *iris.Application) {
	app.Configure(iris.WithConfiguration(iris.Configuration{ //配置字符编码
		Charset: "UTF-8",
	}))

	//---------错误处理
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) { //出现错误时以JSON格式返回错误
		ctx.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    "not found4",
			"data":   iris.Map{},
		})
	})

}

//---------MVC架构处理
func mvcHandle(app *iris.Application) {
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})
	engine := datasource.NewMysqlEngine()
	//fmt.Println("DDDDDDDDDDDD")
	//管理员模块功能
	adminService := service.NewAdminservice(engine)
	admin := mvc.New(app.Party("/admin"))
	admin.Register( //接下来如果有实例化controller,就把注册的字段的值赋值给它的对应字段值
		adminService,
		sessManager.Start,
	)
	admin.Handle(new(controller.AdminController))

}
