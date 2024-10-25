package utils

import (
	"fmt"
	_ "github.com/tangzhaosong/wxacker-t/config"
	"io/ioutil"
	"net/http"
)

// 列出文件的处理函数
func listFiles(w http.ResponseWriter, r *http.Request) {
	// 获取当前目录
	dirPath := r.URL.Path
	if dirPath == "/" {
		dirPath = "C:\\"
	} else {
		dirPath = "C:\\" + dirPath[1:] // 去掉前面的斜杠
	}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		http.Error(w, "无法读取目录", http.StatusInternalServerError)
		return
	}

	// 提取文件信息
	var fileInfos []FileInfo
	for _, file := range files {
		// 构建文件或目录的链接
		if file.IsDir() {
			// 如果是目录，加上斜杠
			fileInfos = append(fileInfos, FileInfo{Name: file.Name(), Path: r.URL.Path + file.Name() + "/"})
		} else {
			// 如果是文件
			fileInfos = append(fileInfos, FileInfo{Name: file.Name(), Path: r.URL.Path + file.Name()})
		}
	}

	// 页面数据
	pageData := PageData{
		Title: fmt.Sprintf("文件列表 - %s", dirPath),
		Files: fileInfos,
	}

	// 解析和执行模板
	tmpl, err := template.New("fileList").Parse(fileListTemplate)
	if err != nil {
		http.Error(w, "模板解析错误", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, "模板执行错误", http.StatusInternalServerError)
	}
}
