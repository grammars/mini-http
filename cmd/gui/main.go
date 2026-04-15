package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mini-http/internal/config"
	"mini-http/internal/server"
)

var (
	srv        *server.Server
	httpServer *http.Server
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// 创建服务器实例
	srv = server.NewServer(cfg)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Port)
	httpServer = &http.Server{
		Addr:    addr,
		Handler: nil,
	}

	// 注册处理函数
	http.HandleFunc("/", srv.HandleRequest)

	// 启动服务器
	go func() {
		fmt.Printf("Server starting on port %d...\n", cfg.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 优雅关闭服务器
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown error: %v\n", err)
	}

	fmt.Println("Server exited")
}
