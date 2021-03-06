package handler

import (
	"encoding/json"
	"fmt"

	"github.com/zlzz-rec/zlyq-go-sdk/model"
	"github.com/zlzz-rec/zlyq-go-sdk/util"
)

// TrackSynchronize 上传用户历史埋点数据
func (client *Client) TrackSynchronize(trackInfo model.TrackInfo) error {
	body, err := json.Marshal(trackInfo)
	if err != nil {
		return err
	}

	resp, err := client.HttpMethod("POST", client.Address, util.TrackSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}
