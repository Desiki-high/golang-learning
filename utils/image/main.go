package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// GetLocalImgBase64 获取本地文件
func GetLocalImgBase64(path string) (baseImg string, err error) {
	file, err := os.Open(path)
	if err != nil {
		err = errors.New("获取本地图片失败")
		return
	}
	defer file.Close()
	imgByte, _ := io.ReadAll(file)

	mimeType := http.DetectContentType(imgByte)
	//拼接
	switch mimeType {
	case "image/bmp":
		baseImg = "data:image/bmp;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/webp":
		baseImg = "data:image/webp;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/jpeg":
		baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/png":
		baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	}
	return
}

func main() {
	image, _ := GetLocalImgBase64("util/image/1080p.jpeg")
	fmt.Println(image)
}
