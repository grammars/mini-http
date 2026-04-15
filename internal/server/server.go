package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"mini-http/internal/config"
	"mini-http/internal/log"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

func (s *Server) Config() *config.Config {
	return s.config
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.handleRequest)
	addr := fmt.Sprintf(":%d", s.config.Port)
	fmt.Printf("Server starting on port %d...\n", s.config.Port)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) HandleRequest(w http.ResponseWriter, r *http.Request) {
	s.handleRequest(w, r)
}

func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	// 记录访问日志
	log.RecordAccess(r, s.config.Log.Pattern, s.config.Log.Folder)

	// 处理虚拟目录
	for _, vhost := range s.config.Vhost {
		if strings.HasPrefix(r.URL.Path, vhost.Path) {
			// 构建真实文件路径
			relativePath := strings.TrimPrefix(r.URL.Path, vhost.Path)
			filePath := filepath.Join(vhost.Folder, relativePath)
			s.serveFile(w, r, filePath)
			return
		}
	}

	// 处理默认目录
	filePath := filepath.Join("./content", r.URL.Path)
	s.serveFile(w, r, filePath)
}

func (s *Server) serveFile(w http.ResponseWriter, r *http.Request, filePath string) {
	// 防止目录遍历攻击
	if strings.Contains(filePath, "..") {
		http.Error(w, "403 Forbidden", http.StatusForbidden)
		return
	}

	// 提供文件服务
	http.ServeFile(w, r, filePath)
}
