package utils

import (
	"encoding/json"
	"net/http"
)

// ParseJSON 解析JSON请求体
func ParseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
