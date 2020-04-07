package model

// UploadType 上传类型
type UploadType uint8

const (
	UploadTypeUnknown UploadType = iota + 1 // 未知
	UploadTypeShot                          // 拍摄
	UploadTypeImport                        // 本地导入
)

// OsType 操作系统
type OsType uint8

const (
	OsIos     OsType = iota + 1 // ios
	OsAndroid                   // Android
	OsUnknown                   // 未知
)

// Source 视频来源
type Source uint8

const (
	SourceUserUpload  Source = iota + 1 // 用户上传
	SourceAdminUpload                   // 后台上传
	SourceSpider                        // 爬虫
)

// ReqVideoSynchronize VideoSynchronize request
type ReqVideoSynchronize struct {
	Videos []VideoSynchronize `json:"videos"`
}

type VideoSynchronize struct {
	Title       string     `json:"title"`
	UserId      uint64     `json:"userId,string"`
	Content     string     `json:"content"`
	UploadType  UploadType `json:"uploadType"`
	OrgFileName string     `json:"orgFileName"`
	Os          OsType     `json:"os"`
	DeviceId    string     `json:"deviceId"`
	Ip          string     `json:"ip"`
	Longitude   float32    `json:"longitude"`
	Latitude    float32    `json:"latitude"`
	AudioId     uint64     `json:"audioId"`
	Source      Source     `json:"source"`
	ThirdId     string     `json:"thirdId"`
	ThirdExtra  string     `json:"thirdExtra"`
	CreatedAt   int64      `json:"createdAt"`
}

// ReqVideoUpload 上传视频请求包
type ReqVideoUpload struct {
	Video       []byte `form:"video"`       // 视频，byte流，必传
	Image       []byte `form:"image"`       // 封图，byte流，必传
	Title       string `form:"title"`       // 视频标题
	Content     string `form:"content"`     // 内容文本
	OrgFileName string `form:"orgFileName"` // 原始文件名
	Os          OsType `form:"os"`          // 操作系统
	Source      Source `form:"source"`      //1用户上传，2后台上传，3爬取
	UserId      uint64 `form:"userId"`
	ThirdId     string `form:"thirdId"`
	ThirdExtra  string `form:"thirdExtra"`
}
