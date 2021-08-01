package main

import (
	"lesson04/dao"
	"lesson04/models"
	"lesson04/routers"
)



func main() {

	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	// 延迟关闭数据库
	defer dao.Close()

	// 模型关闭 自动迁移【把结构体和数据表进行对应】
	dao.MyDB.AutoMigrate(&models.Menu{})

	//注册路由
	r := routers.SetupRouter()

	//9000端口启动
	r.Run(":9000")

}