package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

var MyDB *sql.DB

type Account struct {
	id   int
	name string
	passwd string
}

func InitMysql()(err error)  {

	dsn := "root:@(192.168.197.129:3306)/dbgolang?charset=utf8mb4&parseTime=True&loc=Local"
	MyDB, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect msyql failed, err: %v \n", err)
		return
	}
	// 测试是否能够连通
	return MyDB.Ping()
}

func Close(){
	MyDB.Close()
}

// 将dao层产生的错误包装并往上抛
func QueryOneUser(id int) (*Account, error) {
	var account Account
	err := MyDB.QueryRow(`select id,name,passwd from account where id = ?`, id).Scan(&account.id, &account.name, &account.passwd)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("Error:%w", err)
	}
	if err != nil {
		return nil, err
	}
	return &account, nil
}


func main() {
	InitMysql()
	account, err := QueryOneUser(2021)
	if err != nil {
		log.Printf("QueryOneUser is Error,error:%v", err)
		return
	}
	fmt.Printf("Account: %#v", account)
}
