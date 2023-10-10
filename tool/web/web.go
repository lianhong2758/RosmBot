package web

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// 请求快捷包装
func Web(client *http.Client, url, method string, setheaders func(*http.Request), body io.Reader) (data []byte, err error) {
	// 提交请求
	var request *http.Request
	request, err = http.NewRequest(method, url, body)
	if err == nil {
		// 增加header选项
		setheaders(request)
		var response *http.Response
		response, err = client.Do(request)
		if err != nil {
			return
		}
		if response.StatusCode != http.StatusOK {
			s := fmt.Sprintf("status code: %d", response.StatusCode)
			err = errors.New(s)
			return
		}
		data, err = io.ReadAll(response.Body)
		response.Body.Close()
	}
	return
}

// RequestDataWith 使用自定义请求头获取数据
func RequestDataWith(client *http.Client, url, method, referer, ua string, body io.Reader) (data []byte, err error) {
	// 提交请求
	var request *http.Request
	request, err = http.NewRequest(method, url, body)
	if err == nil {
		// 增加header选项
		if referer != "" {
			request.Header.Add("Referer", referer)
		}
		if ua != "" {
			request.Header.Add("User-Agent", ua)
		}
		var response *http.Response
		response, err = client.Do(request)
		if err == nil {
			if response.StatusCode != http.StatusOK {
				s := fmt.Sprintf("status code: %d", response.StatusCode)
				err = errors.New(s)
				return
			}
			data, err = io.ReadAll(response.Body)
			response.Body.Close()
		}
	}
	return
}

// GetRealURL 获取跳转后的链接
func GetRealURL(url string) (realurl string, err error) {
	data, err := http.Head(url)
	if err != nil {
		return
	}
	_ = data.Body.Close()
	realurl = data.Request.URL.String()
	return
}

// 快速获取
func GetData(url, ua string) (body []byte, err error) {
	var client = &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}
	if ua != "" {
		req.Header.Add("User-Agent", ua)
	}
	req.Header.Add("Accept", "*/*")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("获取数据失败, Code: " + strconv.Itoa(res.StatusCode))
	}
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return
}

func NewDefaultClient() *http.Client {
	return &http.Client{}
}
