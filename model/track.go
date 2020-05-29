package model

import "time"

// TrackInfo
type TrackInfo struct {
	TrackCommon TrackCommon `json:"common"`
	Properties  Properties  `json:"properties"`
}

// Properties
type Properties interface{}

// TrackCommon 预置属性
type TrackCommon struct {
	Udid               string `json:"udid"`
	UserId             uint64 `json:"userId,string"`
	DistinctId         uint64 `json:"distinctId"`
	AppId              uint64 `json:"appId,string"`
	Platform           int32  `json:"platform"`
	Time               int64  `json:"time"`
	IosSdkVersions     string `json:"iosSdkVersions"`
	AndroidSdkVersions string `json:"androidSdkVersions"`
	ScreenHeight       int32  `json:"screenHeight"`
	ScreenWidth        int32  `json:"screenWidth"`
	Manufacturer       string `json:"manufacturer"`
	Network            int32  `json:"network"`
	Os                 int32  `json:"os"`
	OsVersion          string `json:"osVersion"`
	Ip                 string `json:"ip"`
	Country            string `json:"country"`
	Province           string `json:"province"`
	City               string `json:"city"`
	Carrier            string `json:"carrier"`
	UtmSource          string `json:"utmSource"`
	UtmMedia           string `json:"utmMedia"`
	UtmCampaign        string `json:"utmCampaign"`
	UtmContent         string `json:"utmContent"`
	UtmTerm            string `json:"utmTerm"`
	AppVersion         string `json:"appVersion"`
}

type EventCommon struct {
	Event        Event  `json:"event"`
	EventTime    int64  `json:"eventTime"`
	FeedConfigId string `json:"feedConfigId"`
	RequestId    string `json:"requestId"`
	AbtestIds    string `json:"abtestIds"`
}

func GenEventCommon(e Event) EventCommon {
	return EventCommon{
		Event:        e,
		EventTime:    time.Now().UnixNano() / 1000000,
		FeedConfigId: "",
		RequestId:    "",
		AbtestIds:    "",
	}
}

// Event 埋点事件
type Event string

const (
	EventRegister       Event = "register"
	EventLogin          Event = "login"
	EventFinishVideo    Event = "finishVideo"
	EventLike           Event = "like"
	EventDislike        Event = "dislike"
	EventComment        Event = "comment"
	EventLikeComment    Event = "likeComment"
	EventDislikeComment Event = "dislikeComment"
	EventFollow         Event = "follow"
	EventShare          Event = "share"
)

// ContentType 内容类型
type ContentType int32

const (
	ContentTypeVideo  ContentType = 1
	ContentTypeShortArticle  ContentType = 2
	ContentTypeLongArticle  ContentType = 3
)

// TrackRegister 注册 有注册结果时触发
type TrackRegister struct {
	EventCommon
	Methord    int32  `json:"methord"`    // 注册方式:  1:手机号   2:微信
	IsSuccess  int32  `json:"isSuccess"`  // 是否成功:  0失败；1 成功
	FailReason string `json:"failReason"` // 失败原因
}

// TrackLogin 登录 有登陆结果时触发
type TrackLogin struct {
	EventCommon
	Methord    int32  `json:"methord"`    // 注册方式:  1:手机号   2:微信
	IsSuccess  int32  `json:"isSuccess"`  // 是否成功:  0失败；1 成功
	FailReason string `json:"failReason"` // 失败原因
}

// TrackLike 点赞/收藏
type TrackLike struct {
	EventCommon
	ContentId   uint64 `json:"contentId,string"`
	ContentType int32  `json:"contentType"`
	VideoTime int32  `json:"videoTime"`
	Duration  int32  `json:"duration"`
	IsFinish  int32  `json:"isFinish"`
}

// TrackFinishVideo 观看视频
type TrackFinishVideo struct {
	EventCommon
	ContentId   uint64 `json:"contentId,string"`
	ContentType int32  `json:"contentType"`
}

// TrackDislike 取消点赞/收藏
type TrackDislike struct {
	EventCommon
	ContentId   uint64 `json:"contentId,string"`
	ContentType int32  `json:"contentType"`
}

// TrackFollow 关注
type TrackFollow struct {
	EventCommon
	ContentId   uint64 `json:"contentId,string"`
	ContentType string `json:"contentType"`
	AuthorId    uint64 `json:"authorId,string"`
}

// TrackComment 评论
type TrackComment struct {
	EventCommon
	ContentId   uint64 `json:"contentId,string"`
	ContentType int32  `json:"contentType"`
}

// TrackLikeComment 评论点赞
type TrackLikeComment struct {
	EventCommon
	ContentId   uint64 `json:"contentId,string"`
	ContentType int32  `json:"contentType"`
	CommentId   uint64 `josn:"commentId,string"`
}

// TrackDislikeComment 取消评论点赞
type TrackDislikeComment struct {
	EventCommon
	ContentId   uint64 `json:"contentId,string"`
	ContentType int32  `json:"contentType"`
	CommentId   uint64 `josn:"commentId,string"`
}
