package main

import (
	"api/router"
	"api/storage"
	"database/sql"
	"log"
	"net/http"
	"runtime"
	"runtime/debug"
)

func main() {
	// 限制CPU使用(1核)
	CPUCore := 1
	runtime.GOMAXPROCS(CPUCore)
	// 设置内存限制(500MB)
	MemoryLimit := int64(500 * 1024 * 1024)
	debug.SetMemoryLimit(MemoryLimit)
	// 初始化数据库
	config := storage.DefaultConfig()
	db := storage.InitDB(config)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("❌ 数据库关闭失败: ", err)
		}
	}(db)
	// 获取路由处理器
	handler := router.SetupRouter(db)
	Port := "8080"
	log.Printf("API文档在 https://tmlo5zfmuq.apifox.cn/, 访问密码: S2jb4tnO")
	log.Printf("🚀 服务器启动在 http://localhost:%s", Port)
	log.Printf("💾 内存限制: %dMB", MemoryLimit/(1024*1024))
	log.Printf("⚡ CPU核心: %d", CPUCore)

	// 启动HTTP服务器
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("❌ 服务器启动失败: ", err)
	}
}
