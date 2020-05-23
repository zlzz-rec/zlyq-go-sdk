package model

// ReqImageUpload 上传图片
type ReqImageUpload struct {
	Image       []byte `form:"image"`
	Source      Source `form:"source"`
	Description string `form:"description"`
	UserId      uint64 `form:"userId"`
	ThirdId     string `form:"thirdId"`
	ThirdExtra  string `form:"thirdExtra"`
}
