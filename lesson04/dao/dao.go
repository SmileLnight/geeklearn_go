package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

/**
 * @Description 定义全局的数据库对象
 **/
var (
	MyDB *gorm.DB
)


/**
 * @Description 初始化mysql
 * @Param
 * @return err
 **/
func InitMysql()(err error)  {

	dsn := "root:@(192.168.197.129:3306)/dbgolang?charset=utf8mb4&parseTime=True&loc=Local"
	MyDB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect msyql failed, err: %v \n", err)
		return
	}
	// 测试是否能够连通
	return MyDB.DB().Ping()
}

func Close(){
	MyDB.Close()
}