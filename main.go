package webui

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"
)

//go:embed static
var Dist embed.FS

// 静态文件处理
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// 下载
func Download(url string, ws string) {

	// 发送请求
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(ws)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

}

// 创建目录
func CreateDirectory(directoryPath string) {
	err := os.Mkdir(directoryPath, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

}
