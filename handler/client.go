package handler

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
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

// HttpPostForm HttpMultiForm
func (c *Client) HttpMultiForm(address string, apiUrl string,
	params map[string]string, formParam map[string]string, fileMap map[string][]byte) (string, error) {

	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// 读文件
	for k, v := range fileMap {
		// 生成一个临时文件名
		tmpfile := fmt.Sprintf("/tmp/%s.jpg", util.RandUint6())
		fw, err := w.CreateFormFile(k, tmpfile)
		if err != nil {
			return "", fmt.Errorf("w.CreateFormFile failed:%s", err)
		}
		_, err = fw.Write(v)
		if err != nil {
			return "", fmt.Errorf("fw.Write failed:%s", err)
		}
	}

	// 读param
	for k, v := range formParam {
		err := w.WriteField(k, v)
		if err != nil {
			return "", fmt.Errorf("w.WriteField failed:%s", err)
		}
	}

	w.Close()

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
	header.Set("Content-Type", w.FormDataContentType())

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", urlStr, &b)
	if err != nil {
		return "", err
	}

	req.Header = header

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
