package handler

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//参数含义:数据库用户名、密码、主机ip、连接的数据库、端口号
func dbConn(User, Password, Host, Db string, Port int) *gorm.DB {
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", User, Password, Host, Port, Db)
	db, err := gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		return nil
	}
	//开启连接池
	DB, _ := db.DB()
	DB.SetMaxIdleConns(100)   //最大空闲连接
	DB.SetMaxOpenConns(10000) //最大连接数
	DB.SetConnMaxLifetime(30) //最大生存时间(s)

	return db
}

func GetDb() (conn *gorm.DB) {
	for {
		conn = dbConn("root", "admin123", "127.0.0.1", "mysql", 3306)
		if conn != nil {
			break
		}
		fmt.Println("本次未获取到mysql连接")
	}
	return conn
}
