package model

import (
	"fmt"

	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// Database 一个API服务可能需要同时访问多个数据库实例
// 使用Database便于对多个数据库实例进行连接管理
type Database struct {
	Self	*gorm.DB
	Docker	*gorm.DB
}

// DB 数据库连接对象
var DB *Database

func openDB(userName, pwd, addr, dbName string) *gorm.DB {
	connCfg := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4",
						userName, pwd, addr, dbName)
	db, err := gorm.Open(mysql.Open(connCfg), &gorm.Config{})
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", dbName)
	}

	return db
}

// InitSelfDB 初始化Self实例连接
func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("self_db.user_name"),
				viper.GetString("self_db.pwd"),
				viper.GetString("self_db.addr"),
				viper.GetString("self_db.db_name"))
}

// InitDockerDB 初始化Docker实例连接
func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db:user_name"),
				viper.GetString("docker_db.pwd"),
				viper.GetString("docker_db.addr"),
				viper.GetString("docker_db.db_name"))
}

// Init 初始化数据库连接
func (db *Database) Init() {
	DB = &Database {
		Self:	InitSelfDB(),
		Docker:	InitDockerDB(),
	}
}

// Close 关闭数据库连接
func (db *Database) Close() {
	DB.Self.Close()
	DB.Docker.Close()
}