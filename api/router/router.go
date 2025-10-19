package router

import (
	"api/controller"
	"api/middleware"
	"net/http"
	"strings"
)

// SetupRouter 设置路由
// func SetupRouter(db any) http.Handler {
func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	// 基础路由
	mux.HandleFunc("GET /api/ping", controller.Ping())

	// 添加中间件
	handler := middleware.CORS(mux)
	handler = middleware.Logger(handler)
	return handler
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
