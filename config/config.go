package config

type FileInfo struct {
	Name string
	Path string
}

// 文件列表页面数据结构
type PageData struct {
	Title string
	Files []FileInfo
}

// 文件列表页面的 HTML 模板
const fileListTemplate = `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>文件列表</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <ul>
        {{range .Files}}
        <li>
            <a href="{{.Path}}">{{.Name}}</a>
        </li>
        {{end}}
    </ul>
    <a href="/">返回根目录</a>
</body>
</html>
`
