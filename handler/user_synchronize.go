package handler

import (
	"encoding/json"
	"fmt"

	"github.com/zlzz-rec/zlyq-go-sdk/util"
	"github.com/zlzz-rec/zlyq-go-sdk/model"
)

// UserInfoSynchronize 用户信息同步，调用次接口后会在中量系统中注册成功用户并得到用户唯一标识
func (client *Client) UserInfoSynchronize(req *model.ReqUserInfoSynchronize) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", AppApi, util.UserInfoSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}

// UserFollowSynchronize 用户关注数据同步 followMap不超过100个key
func (client *Client) UserFollowSynchronize(userId uint64, followMap map[uint64]int64) error {
	followStrMap := make(map[string]int64)
	for k, v := range followMap {
		followStrMap[util.Uint2Str(k)] = v
	}
	bodyParams := map[string]interface{}{
		"userId":    util.Uint2Str(userId),
		"followMap": followStrMap,
	}
	body, err := json.Marshal(bodyParams)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", AppApi, util.UserFollowSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}
