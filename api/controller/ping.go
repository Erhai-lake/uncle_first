package controller

import (
	"api/utils"
	"net/http"
	"strconv"
	"time"
)

// Ping 测试接口, 返回延迟时间
func Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 只处理GET请求
		if r.Method != http.MethodGet {
			utils.Error(w, http.StatusMethodNotAllowed, "方法不允许")
			return
		}

		start := time.Now()
		time.Sleep(10 * time.Millisecond)
		latency := time.Since(start).Milliseconds()
		latencyStr := strconv.FormatInt(latency, 10) + "ms"

		utils.Success(w, map[string]any{
			"latency": latencyStr,
			"message": "pong",
		})
	}
}
