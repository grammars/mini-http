param(
    [string]$OutputName = "mini-http.exe"
)

$ErrorActionPreference = "Stop"

Write-Host "=== Building Windows version ===" -ForegroundColor Cyan

# 设置环境变量
$env:GOWORK = "off"
$env:GOOS = "windows"
$env:GOARCH = "amd64"

try {
    # 清理旧的构建文件
    if (Test-Path $OutputName) {
        Remove-Item $OutputName -Force
        Write-Host "Removed existing $OutputName"
    }

    # 执行构建
    Write-Host "Running: go build -o $OutputName main.go"
    go build -o $OutputName main.go

    # 验证构建结果
    if (Test-Path $OutputName) {
        Write-Host "`nBuild succeeded!" -ForegroundColor Green
        $fileInfo = Get-Item $OutputName
        Write-Host "Output: $($fileInfo.FullName)"
        Write-Host "Size: $($fileInfo.Length) bytes"
    } else {
        Write-Host "Build failed: output file not found" -ForegroundColor Red
        exit 1
    }
} catch {
    Write-Host "`nBuild failed with error: $_" -ForegroundColor Red
    exit 1
}