@echo off

REM 设置Go环境变量
set GO111MODULE=on

REM 安装依赖
echo 安装依赖...
go mod tidy

REM 创建输出目录
if not exist "bin" mkdir "bin"

REM 构建命令行版本
echo 构建命令行版本...
go build -o bin\mini-http.exe ./cmd/cli

REM 构建GUI版本
echo 构建GUI版本...
go build -o bin\mini-http-gui.exe ./cmd/gui

echo 构建完成！可执行文件位于 bin 目录中。
pause
