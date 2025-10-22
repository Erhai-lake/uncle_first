package router

import (
	"api/controller"
	"api/middleware"
	"api/models"
	"database/sql"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// SetupRouter 设置路由
func SetupRouter(db *sql.DB) http.Handler {
	// 初始化仓库和控制器
	userRepo := models.NewUserRepository(db)
	userController := controller.NewUserController(userRepo)

	r := chi.NewRouter()

	// 添加中间件
	r.Use(middleware.Logger)
	r.Use(middleware.CORS)

	// 基础路由
	r.Get("/api/ping", controller.Ping())

	// 用户路由
	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", userController.GetUsers())
		r.Post("/", userController.CreateUser())
		r.Get("/{id}", userController.GetUser())
		// 可以继续添加更新和删除路由
		// r.Put("/{id}", userController.UpdateUser())
		// r.Delete("/{id}", userController.DeleteUser())
	})
	return r
}

// GetPathParam 路径参数提取辅助函数
func GetPathParam(r *http.Request, param string) string {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	// 在路径中查找参数位置
	for i, part := range parts {
		if part == param {
			if i+1 < len(parts) {
				return parts[i+1]
			}
		}
	}
	return ""
}
