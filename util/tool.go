package util

import (
	"fmt"
	"math/rand"
	"net/url"
	"path"
	"strconv"
	"time"
)

// Uint2Str 无符号整型到字符串
func Uint2Str(n uint64) string {
	return strconv.FormatUint(n, 10)
}

func Join(basePath string, paths ...string) (*url.URL, error) {

	u, err := url.Parse(basePath)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %s", basePath)
	}

	p2 := append([]string{u.Path}, paths...)
	result := path.Join(p2...)

	u.Path = result

	return u, nil
}

// RandUint6 生成首位非0的6位随机码
func RandUint6() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := fmt.Sprintf("%05v", rnd.Int31n(100000))
	prefix := rnd.Int31n(9) + 1
	result = fmt.Sprintf("%01v", prefix) + result
	return result
}
