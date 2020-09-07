package handler

import (
	"encoding/json"
	"fmt"
	"github.com/zlzz-rec/zlyq-go-sdk/model"

	"github.com/zlzz-rec/zlyq-go-sdk/util"
)

// startWeiXinPay 发起微信支付
func (client *Client) WxUnifiedOrder(req *model.WxAppPostRawBody) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.WxUnifiedOrderApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}
