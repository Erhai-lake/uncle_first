package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse 统一的JSON响应
func JSONResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

// Success 成功响应
func Success(w http.ResponseWriter, data any) {
	JSONResponse(w, http.StatusOK, map[string]any{
		"success": true,
		"data":    data,
	})
}

// Error 错误响应
func Error(w http.ResponseWriter, status int, message string) {
	JSONResponse(w, status, map[string]any{
		"success": false,
		"error":   message,
	})
}

// GetQueryParam 获取查询参数
func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

// GetQueryParamDefault 获取查询参数, 如果为空返回默认值
func GetQueryParamDefault(r *http.Request, key, defaultValue string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		return defaultValue
	}
	return value
}
