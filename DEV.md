.vscode/
  - launch.json
  - tasks.json

launch.json

```json 
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Run Server",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "main.go",
      "cwd": "${workspaceFolder}",
      "env": {
        "GOWORK": "off"
      },
      "args": []
    }
  ]
}
```

tasks.json

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "build-cli",
      "type": "shell",
      "command": "go",
      "args": [
        "build",
        "-o",
        "mini-http.exe",
        "main.go"
      ],
      "group": "build",
      "problemMatcher": ["$go"]
    },
    {
      "label": "tidy",
      "type": "shell",
      "command": "go",
      "args": ["mod", "tidy"],
      "group": "build",
      "problemMatcher": ["$go"]
    }
  ]
}
```