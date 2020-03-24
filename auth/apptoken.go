package auth

import (
	"crypto/md5"
	"fmt"
	"time"
)

// AddAppToken 添加apptoken信息
func AddAppToken(appKey, appSecret string) string {
	timeStr := time.Now().Format("2006-01-02 15:04")
	str := appKey + appSecret + timeStr
	signature := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return signature
}
