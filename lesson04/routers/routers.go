package routers

import (
	"github.com/gin-gonic/gin"
	"lesson04/controller"
	"github.com/mattn/go-colorable"
)

func SetupRouter() *gin.Engine{
	// 启用gin的日志输出带颜色
	gin.ForceConsoleColor()
	// 替换默认Writer（关键步骤）
	gin.DefaultWriter = colorable.NewColorableStdout()
	// 定义Gin默认路由
	r := gin.Default()

	// 告诉Gin框架模板文件引用的静态文件去哪里找
	r.Static("/static","static")

	// 告诉Gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")

	// 访问待办事项首页
	r.GET("/", controller.IndexHandler)


	// v1  待办事项
	// 前端页面填写代办事项，点击提交就会发送请求到这里
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateMenu)

		// 查看【查看所有】
		v1Group.GET("/todo", controller.GetAllMenus)

		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateMenu)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteMenu)

	}
	return r
}
