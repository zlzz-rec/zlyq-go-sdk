package util

import (
	"fmt"
	"net/url"
	"path"
	"strconv"
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
