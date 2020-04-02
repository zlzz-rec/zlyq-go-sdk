package model

// ReqUserInfoSynchronize UserInfoSynchronize request
type ReqUserInfoSynchronize struct {
	ThirdId       string `json:"thirdId"` // 三方用户id 必传
	Account       string `json:"account"` // 可展示在客户端的用户ID 用于搜索
	Password      string `json:"password"`
	Nickname      string `json:"nickname"`
	Avatar        string `json:"avatar"`
	Gender        int32  `json:"gender"`
	Birthday      int64  `json:"birthday"`
	Signature     string `json:"signature"`
	Region        string `json:"region"`
	ExtraInfo     string `json:"extraInfo"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
	LoginTime     int64  // 最后一次打开app时间
	IsRobot       int32  `json:"is_robot"`      // 是否是马甲号 0否 1是
	Status        int32  `json:"status"`        // 1正常 2非正常
	AvatarStorage bool   `json:"avatarStorage"` // 头像是否要存储到我们的oss中
	Phone         string `json:"phone"`
	// WxPassport    string `json:"wxPassport"`
	// QqPassport    string `json:"qqPassport"`
	// WbPassport    string `json:"wbPassport"`
	DeviceId   string `json:"udid"` // 必传
	DeviceType string `json:"deviceType"`
}

// ReqUserFollow  UserFollowSynchronize 接口 请求 body 体定义
type ReqUserFollow struct {
	UserId    string           `json:"userId"`
	FollowMap map[string]int64 `json:"followMap"`
}
