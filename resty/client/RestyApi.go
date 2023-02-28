package client

//client包 用于发送请求调用外部接口

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"gorm.io/datatypes"
)

// RestyGet 使用Resty获取json
// url 接口地址
func RestyGet(url string) datatypes.JSON {
	client := resty.New()
	resp, err := client.R().Get(url)
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Body       :\n", resp)
	json, _ := datatypes.JSON.MarshalJSON(resp.Body())
	return json
}

// RestyPost 使用Resty发送Post请求
// url 接口地址 content发送的结构体(JSON)
func RestyPost(url string, content interface{}) {
	client := resty.New()
	resp, err := client.R().
		SetBody(content).
		Post(url)
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Body       :\n", resp)
}

// RestyUpload 使用Resty上传文件
// url 接口地址  filePath 本地文件绝对路径
func RestyUpload(url string, filePath string) {
	client := resty.New()
	resp, err := client.R().
		SetFile("file", filePath).
		Post(url)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(filePath + "文件读取失败")
		}
	}()
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Body       :\n", resp)
}
