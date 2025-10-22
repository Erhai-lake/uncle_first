package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DBConfig 数据库配置
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// InitDB 初始化数据库连接
func InitDB(config DBConfig) *sql.DB {
	// 构建DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ 数据库连接失败: ", err)
	}

	// 测试连接
	if err := db.Ping(); err != nil {
		log.Fatal("❌ 数据库ping失败: ", err)
	}

	// 设置连接池参数
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("✅ MySQL数据库连接成功")
	return db
}

// DefaultConfig 返回默认数据库配置
func DefaultConfig() DBConfig {
	return DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "uncle_first",
		Password: "uncle_first",
		Database: "uncle_first",
	}
}
