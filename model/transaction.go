package model

type WxAppPostRawBody struct {
	OrderId       string `json:"orderId"`
	Body          string `json:"body,string"`
	TradeType     string `totalFee:"totalFee"`
	TotalFee      uint64 `totalFee:"tradeType"`
	SpbillCreatIp string `totalFee:"spbillCreatIp"`
	NotifyUrl     string `totalFee:"notifyUrl"`
}
