package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	fileUrl = "http://bg.tencentbot.top/image/uploadfile/"
	imgUrl  = "http://bg.tencentbot.top/image/upload?url="
)

// 上传byte数据
func UpImgByte(file []byte) (url string, con image.Config) {
	// 创建一个buffer，用于构造multipart/form-data格式的表单数据
	t := fmt.Sprintf("%d.jpg", time.Now().Unix())
	log.Infoln("[web](upimg)FileName: ", t)
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
		log.Errorln("[upimg]", err)
		return
	}
	defer response.Body.Close()

	// 处理响应
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorln("[upimg]", err)
		return
	}
	con, _ = BytesToConfig(file)
	return getBodyUrl(body), con
}

// 上传file
func UpImgfile(filePath string) (url string, con image.Config) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Errorln("[upimg]", err)
		return
	}
	return UpImgByte(file)
}
func UpImgUrl(imgurl string) (url string) {
	body, err := GetData(imgUrl+imgurl, "")
	if err != nil {
		log.Errorln("[upimg]", err)
		return
	}
	return getBodyUrl(body)
}
func getBodyUrl(body []byte) (url string) {
	r := new(upJson)
	err := json.Unmarshal(body, r)
	if err != nil {
		log.Errorln("[upimg]", err)
		return
	}
	return r.URL
}

type upJson struct {
	URL       string `json:"url"`
	SecretURL string `json:"secret_url"`
	Object    string `json:"object"`
}

// 将 []byte 类型的图片数据转换为 image.Image 类型
func BytesToImage(imgBytes []byte) (image.Image, error) {
	img, err := jpeg.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}
	return img, nil
}

// 返回本地文件的img对象
func FileToImage(path string) (image.Image, error) {
	fileT, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fileT.Close()
	fileExtT := strings.ToLower(filepath.Ext(path))
	var imgT image.Image
	switch fileExtT {
	case ".png":
		imgT, err = png.Decode(fileT)
	case ".gif":
		imgT, err = gif.Decode(fileT)
	default:
		imgT, err = jpeg.Decode(fileT) //默认按jpg格式处理
	}
	return imgT, err
}

// 获取 []byte图片的config
func BytesToConfig(imgBytes []byte) (image.Config, error) {
	imgConfig, _, err := image.DecodeConfig(bytes.NewReader(imgBytes))
	if err != nil {
		return image.Config{}, err
	}
	return imgConfig, nil
}

// 返回本地文件的config
func FileToConfig(path string) (image.Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return image.Config{}, err
	}
	return BytesToConfig(file)
}

// 返回本地文件的config
func URLToConfig(url string) (image.Config, error) {
	file, err := GetData(url, "")
	if err != nil {
		return image.Config{}, err
	}
	return BytesToConfig(file)
}
