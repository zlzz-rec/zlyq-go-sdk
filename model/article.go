package model

// MediaType 媒资类型
type MediaType uint8

const (
	MediaTypeVideo   MediaType = iota + 1 // 视频
	MediaTypeGallery                      // 图集
	MediaTypeArticle                      // 图文
	MediaTypeLink                         // 外链
	MediaTypeImage   MediaType = 10       // 图片
)

// ReqArticleSynchronize ArticleSynchronize request
type ReqArticleSynchronize struct {
	Articles []ArticleSynchronize `json:"articles"`
}

// ArticleSynchronize 图文同步
type ArticleSynchronize struct {
	Title      string    `json:"title"`
	UserId     uint64    `json:"userId,string"`
	Content    string    `json:"content"`
	H5Content  string    `json:"h5content"`
	Os         OsType    `json:"os"`
	DeviceId   string    `json:"deviceId"`
	Ip         string    `json:"ip"`
	Longitude  float32   `json:"longitude"`
	Latitude   float32   `json:"latitude"`
	Source     Source    `json:"source"`
	ThirdId    string    `json:"thirdId"`
	ThirdExtra string    `json:"thirdExtra"`
	CreatedAt  int64     `json:"createdAt"`
	MediaType  MediaType `json:"mediaType"`
}

// ReqArticleUpload 上传图文请求包
type ReqArticleUpload struct {
	CoverIds   []string  `json:"coverIds"`
	GalleryIds []string  `json:"galleryIds"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Os         OsType    `json:"os"`
	Source     Source    `json:"source"`
	MediaType  MediaType `json:"mediaType"`
	H5Content  string    `json:"h5content"`
	UserId     uint64    `json:"userId,string"`
	ThirdId    string    `josn:"thirdId"`
	ThirdExtra string    `json:"thirdExtra"`
}
