package handler

import (
	"encoding/json"
	"fmt"

	"github.com/zlzz-rec/zlyq-go-sdk/util"
	"github.com/zlzz-rec/zlyq-go-sdk/model"
)

// HistorySynchronize 上传用户历史埋点数据
func (client *Client) HistorySynchronize(trackInfo model.TrackInfo) error {
	body, err := json.Marshal(trackInfo)
	if err != nil {
		return err
	}

	resp, err := client.HttpMethod("POST", client.Address, util.HistorySynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}
