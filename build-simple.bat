@echo off

REM 设置 Go 环境变量
set GO111MODULE=on

REM 构建命令行版本
echo 构建命令行版本...
go build -o bin\mini-http.exe ./cmd/cli

REM 构建GUI版本
echo 构建GUI版本...
go build -o bin\mini-http-gui.exe ./cmd/gui

echo 构建完成！
dir bin

pause
