package main

import (
	"fmt"
	"io/ioutil"

	"github.com/zlzz-rec/zlyq-go-sdk/handler"
	"github.com/zlzz-rec/zlyq-go-sdk/model"
)

// syncUserExample
func syncUserExample(client *handler.Client) {

	// 用户基本信息导入
	req := &model.ReqUserInfoSynchronize{
		ThirdId:   "1",
		Nickname:  "昵称1",
		DeviceId:  "1",
		CreatedAt: 1440000000002,
	}
	if err := client.UserInfoSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// syncUserFollowExample
func syncUserFollowExample(client *handler.Client) {

	// 用户关注数据导入
	var userId uint64 = 457741243892858880
	followMap := map[uint64]int64{
		457751121931763712: 0,
	}
	if err := client.UserFollowSynchronize(userId, followMap); err != nil {
		fmt.Println(err)
	}
}

// syncTrackExample
func syncTrackExample(client *handler.Client) {

	// 用户基本信息导入
	trackCommon := model.TrackCommon{
		Udid:       "JKAFJ-LAKJS-JAKSJ-IWNSK",
		UserId:     450007472627785728,
		DistinctId: 450007472627785728,
		AppId:      450007472627785728,
	}

	// 用户的一个埋点事件
	like := model.TrackLike{
		EventCommon: model.GenEventCommon(model.EventLike),
		ContentId:   1291872937198273,
		ContentType: int32(model.ContentTypeVideo),
	}

	// 用户的另一个埋点事件
	finishVideo := model.TrackFinishVideo{
		EventCommon: model.GenEventCommon(model.EventFinishVideo),
		ContentId:   1291872937198273,
		ContentType: int32(model.ContentTypeVideo),
	}
	properties := []model.Properties{like, finishVideo}

	req := model.TrackInfo{
		TrackCommon: trackCommon,
		Properties:  properties,
	}

	// 同步用户历史数据
	if err := client.TrackSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// syncVideoExample 视频同步
func syncVideoExample(client *handler.Client) {

	req := &model.ReqVideoSynchronize{}
	r := model.VideoSynchronize{
		Title:       "title test",
		UserId:      457378584727416832,
		Content:     "content test",
		UploadType:  model.UploadTypeUnknown,
		OrgFileName: "1.mp4",
		Os:          model.OsUnknown,
		DeviceId:    "",
		Ip:          "",
		Longitude:   0,
		Latitude:    0,
		AudioId:     0,
		Source:      model.SourceAdminUpload,
		ThirdId:     "12345677",
		ThirdExtra:  "",
		CreatedAt:   1585711691000,
	}
	req.Videos = append(req.Videos, r)

	if err := client.VideoSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// syncArticleExample 图文同步
func syncArticleExample(client *handler.Client) {

	req := &model.ReqArticleSynchronize{}
	r := model.ArticleSynchronize{
		Title:      "title test",
		UserId:     457378584727416832,
		Content:    "content test",
		H5Content:  "",
		Os:         model.OsUnknown,
		DeviceId:   "",
		Ip:         "",
		Longitude:  0,
		Latitude:   0,
		Source:     model.SourceAdminUpload,
		ThirdId:    "12345677",
		ThirdExtra: "",
		CreatedAt:  1585711691000,
		MediaType:  model.MediaTypeGallery,
	}
	req.Articles = append(req.Articles, r)

	if err := client.ArticleSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// syncMediaLikeExample 点赞同步
func syncMediaLikeExample(client *handler.Client) {

	req := &model.ReqMediaLikeSync{}
	r := model.MediaLike{
		MediaId:   467366407865606144,
		UserId:    457751121931763712,
		CreatedAt: 1585711691000,
	}
	req.MediaLikes = append(req.MediaLikes, r)

	if err := client.MediaLikeSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// syncMediaFavoriteExample 收藏同步
func syncMediaFavoriteExample(client *handler.Client) {

	req := &model.ReqMediaFavoriteSync{}
	r := model.MediaFavorite{
		MediaId:   467366407865606144,
		UserId:    457751121931763712,
		CreatedAt: 1585711691000,
	}
	req.MediaFavorites = append(req.MediaFavorites, r)

	if err := client.MediaFavoriteSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// syncCommentExample 评论同步
func syncCommentExample(client *handler.Client) {

	req := &model.ReqCommentSync{}
	r := model.Comment{
		Content:   "comment content",
		MediaId:   467366407865606144,
		UserId:    457751121931763712,
		CreatedAt: 1585711691000,
		ThirdId:   "",
	}
	req.Comments = append(req.Comments, r)

	if err := client.CommentSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// syncCommentLikeExample 点赞评论同步
func syncCommentLikeExample(client *handler.Client) {

	req := &model.ReqCommentLikeSync{}
	r := model.CommentLike{
		CommentId: 467428084384505856,
		MediaId:   467366407865606144,
		UserId:    457751121931763712,
		CreatedAt: 1585711691000,
	}
	req.CommentLikes = append(req.CommentLikes, r)

	if err := client.CommentLikeSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// uploadArticleExample 上传图文
func uploadArticleExample(client *handler.Client) {

	req := &model.ReqArticleUpload{
		CoverIds:   []string{},
		GalleryIds: []string{"455522923617492992"},
		Title:      "article title",
		Content:    "article content",
		Os:         model.OsUnknown,
		Source:     model.SourceUserUpload,
		MediaType:  model.MediaTypeGallery,
		H5Content:  "",
		UserId:     457751121931763712,
		ThirdId:    "",
		ThirdExtra: "",
	}

	if err := client.UploadArticle(req); err != nil {
		fmt.Println(err)
	}
}

// uploadImageExample 上传图片
func uploadImageExample(client *handler.Client) error {

	// 测试时，请将test.jpg替换成本地图片路径
	f, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		fmt.Println("read fail", err)
		return err
	}

	req := &model.ReqImageUpload{
		Image:       f,
		Source:      model.SourceUserUpload,
		Description: "description",
		UserId:      457751121931763712,
		ThirdId:     "",
		ThirdExtra:  "",
	}

	if err := client.UploadImage(req); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// uploadVideoExample 上传视频
func uploadVideoExample(client *handler.Client) error {

	// 测试时，请将test.jpg, test.mp4替换成本地图片、视频路径
	i, err := ioutil.ReadFile("test.jpg")
	if err != nil {
		fmt.Println("read fail", err)
		return err
	}

	v, err := ioutil.ReadFile("test.mp4")
	if err != nil {
		fmt.Println("read fail", err)
		return err
	}

	req := &model.ReqVideoUpload{
		Video:       v,
		Image:       i,
		Title:       "video title",
		Content:     "video contetn",
		OrgFileName: "video file name",
		Os:          model.OsUnknown,
		Source:      model.SourceUserUpload,
		UserId:      457751121931763712,
		ThirdId:     "",
		ThirdExtra:  "",
	}

	if err := client.UploadVideo(req); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {

	// 同步用户数据example
	appClient := &handler.Client{
		AppId:     450007472627785728,
		AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
		AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
		Address:   "http://testappapi.zplatform.cn",
	}
	syncUserExample(appClient)

	// // 同步历史交互数据example
	// trackClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testtrackapi.zplatform.cn",
	// }
	// syncTrackExample(trackClient)

	// 同步视频example
	// videoClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// syncVideoExample(videoClient)

	// 同步图文example
	// articleClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// syncArticleExample(articleClient)

	// 同步点赞example
	// mediaLikeClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// syncMediaLikeExample(mediaLikeClient)

	// 同步收藏example
	// mediaFavoriteClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// syncMediaFavoriteExample(mediaFavoriteClient)

	// 同步评论example
	// commentClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// syncCommentExample(commentClient)

	// 同步点赞评论example
	// commentLikeClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// syncCommentLikeExample(commentLikeClient)

	// 上传图文example
	// articleClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// uploadArticleExample(articleClient)

	// 上传图片example
	// imageClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// uploadImageExample(imageClient)

	// 上传视频example
	// videoClient := &handler.Client{
	// 	AppId:     450007472627785728,
	// 	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	// 	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	// 	Address:   "http://testappapi.zplatform.cn",
	// }
	// uploadVideoExample(videoClient)

}
