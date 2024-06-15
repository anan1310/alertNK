package http_util

import (
	"alarm_collector/global"
	"alarm_collector/pkg/utils/common"
	"bytes"
	"io"
	"time"

	"encoding/json"
	"fmt"
	"net/http"
)

// DoJsonRequest 发送JSON请求
func DoJsonRequest(url string, data interface{}) (string, error) {
	if common.IsEmptyStr(url) {
		return "", fmt.Errorf("URL cannot be empty")
	}
	if data == nil {
		return "", fmt.Errorf("data cannot be empty")
	}
	buf, _ := json.Marshal(data)
	reader := bytes.NewReader(buf)
	request, _ := http.NewRequest("POST", url, reader)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, doErr := client.Do(request)
	if doErr != nil {
		return "", doErr
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {

		return "", readErr
	}
	return string(body), nil
}

func Post(url string, bodyReader *bytes.Reader) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodPost, url, bodyReader)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		global.Logger.Sugar().Error("请求建立失败: ", err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		global.Logger.Sugar().Error("请求发送失败: ", err)
		return nil, err
	}

	return resp, nil

}

func Get(url string) (*http.Response, error) {

	client := http.Client{
		Timeout: 1 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		global.Logger.Sugar().Error("请求建立失败: ", err)
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		global.Logger.Sugar().Error("请求发送失败: ", err)
		return nil, err
	}

	return resp, nil
}
