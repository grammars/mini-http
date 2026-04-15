@echo off

REM 创建临时 go.mod 文件，只包含命令行版本需要的依赖
echo module mini-http > go.mod.temp
echo. >> go.mod.temp
echo go 1.20 >> go.mod.temp
echo. >> go.mod.temp
echo require ( >> go.mod.temp
echo 	gopkg.in/yaml.v3 v3.0.1 >> go.mod.temp
echo ) >> go.mod.temp

REM 备份原始 go.mod
copy go.mod go.mod.backup > nul

REM 使用临时 go.mod
copy go.mod.temp go.mod > nul

REM 创建输出目录
if not exist "bin" mkdir "bin"

REM 构建命令行版本
echo 构建命令行版本...
go mod tidy
go build -o bin\mini-http.exe ./cmd/cli

REM 恢复原始 go.mod
copy go.mod.backup go.mod > nul
del go.mod.temp go.mod.backup

if exist "bin\mini-http.exe" (
    echo 命令行版本构建成功！
) else (
    echo 命令行版本构建失败。
)

pause
