package log

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func RecordAccess(r *http.Request, pattern string, logFolder string) {
	// 确保日志目录存在
	if err := os.MkdirAll(logFolder, 0755); err != nil {
		fmt.Printf("Error creating log directory: %v\n", err)
		return
	}

	// 构建日志文件名
	logFileName := filepath.Join(logFolder, time.Now().Format("2006-01-02")+".log")

	// 替换日志模板中的变量
	logContent := pattern
	logContent = strings.Replace(logContent, "${IP}", getClientIP(r), -1)
	logContent = strings.Replace(logContent, "${METHOD}", r.Method, -1)
	logContent = strings.Replace(logContent, "${URL}", r.URL.Path, -1)

	// 添加时间戳
	logContent = time.Now().Format("2006-01-02 15:04:05") + " " + logContent

	// 写入日志文件
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(logContent + "\n"); err != nil {
		fmt.Printf("Error writing to log file: %v\n", err)
	}
}

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return ip
	}
	return r.RemoteAddr
}
