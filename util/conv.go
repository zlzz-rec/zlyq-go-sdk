package util

const (
	UserFollowSynchronizeApi = "/api/v1/synchronize/userFollow" // 用户关注数据迁移api
	UserInfoSynchronizeApi   = "/api/v1/synchronize/userInfo"   // 用户数据迁移api

	HistorySynchronizeApi = "/trace" // 用户历史埋点数据上报

	VideoSynchronizeApi   = "/api/v1/videoSync"   // 同步视频
	ArticleSynchronizeApi = "/api/v1/articleSync" // 同步图文

	ImageUploadSynchronizeApi   = "/api/v1/imageUploadSync"   // 上传并同步图片
	VideoUploadSynchronizeApi   = "/api/v1/videoUploadSync"   // 上传并同步视频
	ArticleUploadSynchronizeApi = "/api/v1/articleUploadSync" // 上传并同步图文

	MediaLikeSynchronizeApi     = "/api/v1/mediaLikeSync"     // 点赞数据
	MediaFavoriteSynchronizeApi = "/api/v1/mediaFavoriteSync" // 收藏数据
	CommentSynchronizeApi       = "/api/v1/commentSync"       // 评论数据
	CommentLikeSynchronizeApi   = "/api/v1/commentLikeSync"   // 评论点赞数据
)
