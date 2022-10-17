package model

import (
	"fmt"
	// "os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	err error
)

// func InitDb() *gorm.DB {
func InitDB() {
	// var db *gorm.DB
	// var err error

	// 在本机用Navicat创建一个名为mecha-glancer的数据库
	host := "localhost"
	port := "3306"
	database := "mecha-glancer"
	username := "glancer"
	password := "glancer"
	charset := "utf8mb4"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	// dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	"glancer",
	// 	"glancer",
	// 	"localhost",
	// 	"3306",
	// 	"mecha-glancer",
	// )

	//这里 gorm.Open()函数与之前版本的不一样，大家注意查看官方最新gorm版本的用法
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Error to Db connection, err: " + err.Error())
		// fmt.Println("连接数据库失败，请检查参数：", err)
		// os.Exit(1)
	}
	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	_ = db.AutoMigrate(&User{})
	// _ = db.AutoMigrate(&User{}, &Article{}, &Category{}, Profile{}, Comment{})

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	DB = db
	// return db
}

func GetDB() *gorm.DB {
	return DB
}
