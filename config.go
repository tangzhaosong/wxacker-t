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
