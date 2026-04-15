package main

import (
	"fmt"

	"mini-http/internal/config"
	"mini-http/internal/server"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// 创建服务器
	s := server.NewServer(cfg)

	// 启动服务器
	if err := s.Start(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
