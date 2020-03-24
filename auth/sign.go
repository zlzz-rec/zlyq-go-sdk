package auth

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/zlzz-rec/zlyq-go-sdk/util"
)

// AddSign 添加sign信息
func AddSign(urlParams map[string]string, appId uint64, appSecret string) (string, string) {
	// 添加时间参数
	if urlParams == nil {
		urlParams = make(map[string]string)
	}
	urlParams["time"] = util.Uint2Str(uint64(time.Now().UnixNano() / 1000000))
	urlParams["appId"] = util.Uint2Str(appId)

	// 参数排序
	keys := []string{}
	for k := range urlParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 参数拼接处理
	var values []string
	for _, k := range keys {
		values = append(values, strings.Join([]string{k, urlParams[k]}, "="))
	}
	str := strings.Join(values, "&")

	var salt string = appSecret
	signature := fmt.Sprintf("%x", md5.Sum([]byte(salt+"&"+str)))
	return signature, str
}
