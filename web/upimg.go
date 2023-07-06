package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const (
	fileUrl = "https://bg.tencentbot.top/image/uploadfile/"
	imgUrl  = "https://bg.tencentbot.top/image/upload?url="
)

// 上传byte数据
func UpImgByte(file []byte) (url string) {
	// 创建一个buffer，用于构造multipart/form-data格式的表单数据
	t := fmt.Sprintf("%d.jpg", time.Now().Unix())
	log.Println("[ipimg]", t)
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件数据
	part, _ := writer.CreateFormFile("file", t)
	_, _ = part.Write(file)

	// 关闭表单数据的写入
	_ = writer.Close()

	// 创建一个POST请求，设置请求头和请求体
	request, _ := http.NewRequest("POST", fileUrl, &requestBody)

	request.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println("[upimg-err]", err)
		return
	}
	defer response.Body.Close()

	// 处理响应
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("[upimg-err]", err)
		return
	}
	return getBodyUrl(body)
}

// 上传file
func UpImgfile(filePath string) (url string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("[upimg-err]", err)
		return
	}
	return UpImgByte(file)
}
func UpImgUrl(imgurl string) (url string) {
	body, err := GetData(imgurl, "")
	if err != nil {
		log.Println("[upimg-err]", err)
		return
	}
	return getBodyUrl(body)
}
func getBodyUrl(body []byte) (url string) {
	r := new(upJson)
	err := json.Unmarshal(body, r)
	if err != nil {
		log.Println("[upimg-err]", err)
		return
	}
	return r.URL
}

type upJson struct {
	URL       string `json:"url"`
	SecretURL string `json:"secret_url"`
	Object    string `json:"object"`
}
