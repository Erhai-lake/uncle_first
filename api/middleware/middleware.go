package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 颜色常量
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// CORS 中间件
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Logger 中间件
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// 包装ResponseWriter来捕获状态码
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)
		log.Printf("%s", formatLog(r.Method, r.URL.Path, wrapped.statusCode, duration))
	})
}

// 格式化日志输出
func formatLog(method, path string, status int, latency time.Duration) string {
	// 方法颜色
	var methodColor string
	switch method {
	case "GET":
		methodColor = colorBlue
	case "POST":
		methodColor = colorGreen
	case "PUT":
		methodColor = colorYellow
	case "DELETE":
		methodColor = colorRed
	default:
		methodColor = colorWhite
	}

	// 状态码颜色
	var statusColor string
	switch {
	case status >= 200 && status < 300:
		statusColor = colorGreen
	case status >= 300 && status < 400:
		statusColor = colorCyan
	case status >= 400 && status < 500:
		statusColor = colorYellow
	case status >= 500:
		statusColor = colorRed
	default:
		statusColor = colorWhite
	}

	// 延迟颜色
	var latencyColor string
	switch {
	case latency < 100*time.Millisecond:
		latencyColor = colorGreen
	case latency < 500*time.Millisecond:
		latencyColor = colorYellow
	default:
		latencyColor = colorRed
	}

	methodStr := fmt.Sprintf("%s%-7s%s", methodColor, method, colorReset)
	statusStr := fmt.Sprintf("%s%-3d%s", statusColor, status, colorReset)
	latencyStr := fmt.Sprintf("%s%v%s", latencyColor, latency, colorReset)
	pathStr := fmt.Sprintf("%s%s%s", colorCyan, path, colorReset)

	return fmt.Sprintf("%s | %s | %13s | %s", methodStr, statusStr, latencyStr, pathStr)
}

// 自定义ResponseWriter来捕获状态码
type responseWriter struct {
	http.ResponseWriter
	statusCode  int
	wroteHeader bool
}

func (rw *responseWriter) WriteHeader(code int) {
	if !rw.wroteHeader {
		rw.statusCode = code
		rw.ResponseWriter.WriteHeader(code)
		rw.wroteHeader = true
	}
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(data)
}
