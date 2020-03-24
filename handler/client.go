package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/zlzz-rec/zlyq-go-sdk/auth"
	"github.com/zlzz-rec/zlyq-go-sdk/util"

	"github.com/buger/jsonparser"
)

// Client client
type Client struct {
	AppKey    string
	AppSecret string
	AppId     uint64
	Address   string
}

// HttpMethod POST-GET-PUT-DELETE
func (c *Client) HttpMethod(method, address string, apiUrl string, params map[string]string, body []byte) (string, error) {

	var req *http.Request
	var err error

	// 组装url
	url, err := util.Join(address, apiUrl)
	if err != nil {
		return "", fmt.Errorf("util.Join failed:%s, address=%s, apiUrl=%s", err.Error(), address, apiUrl)
	}
	urlStr := url.String()

	header, urlParams := BuildHttpHeader(c, params)
	if urlParams != "" {
		urlStr = urlStr + "?" + urlParams
	}

	if body == nil || len(body) == 0 {
		req, err = http.NewRequest(method, urlStr, nil)
	} else {
		req, err = http.NewRequest(method, urlStr, strings.NewReader(string(body)))
	}
	fmt.Println(urlStr)
	if header != nil {
		req.Header = header
	}

	if err != nil {
		return "", fmt.Errorf("http.NewRequest failed:%s", err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("client.Do failed:%s", err.Error())
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll failed:%s", err.Error())
	}

	if err := ParseResponse(string(respBody)); err != nil {
		return string(respBody), fmt.Errorf("ParseResponse failed:%s", err.Error())
	}
	return string(respBody), nil
}

// ParseResponse 解析回包
func ParseResponse(resp string) error {
	code, err := jsonparser.GetInt([]byte(resp), "code")
	if err != nil {
		return err
	}
	if code != 0 {
		return errors.New("code != 0 ")
	}
	return nil
}

// BuildHttpHeader HttpHeader
func BuildHttpHeader(client *Client, params map[string]string) (http.Header, string) {
	header := http.Header{}
	header.Add("Content-Type", "application/json")
	sign, urlParam := auth.AddSign(params, client.AppId, client.AppSecret)
	header.Add("X-Sign", sign)
	header.Add("X-App-Token", auth.AddAppToken(client.AppKey, client.AppSecret))
	return header, urlParam
}
