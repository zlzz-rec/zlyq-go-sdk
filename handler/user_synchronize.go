package handler

import (
	"encoding/json"
	"fmt"

	"github.com/zlzz-rec/zlyq-go-sdk/model"
	"github.com/zlzz-rec/zlyq-go-sdk/util"
)

// UserInfoSynchronize 用户信息同步，调用次接口后会在中量系统中注册成功用户并得到用户唯一标识
func (client *Client) UserInfoSynchronize(req *model.ReqUserInfoSynchronize) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.UserInfoSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}

// UserFollowSynchronize 用户关注数据同步 followMap不超过100个key wiki url: https://wiki.zplatform.cn/pages/server_user_iface.html#%E7%94%A8%E6%88%B7%E5%85%B3%E6%B3%A8%E6%95%B0%E6%8D%AE%E5%AF%BC%E5%85%A5
func (client *Client) UserFollowSynchronize(req *model.ReqUserFollow) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.UserFollowSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}
