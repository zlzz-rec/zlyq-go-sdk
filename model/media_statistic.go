package model

type MediaLike struct {
	MediaId   uint64 `json:"mediaId,string"`
	UserId    uint64 `json:"userId,string"`
	CreatedAt int64  `json:"createdAt"`
}

// ReqMediaLikeSync ReqMediaLike request
type ReqMediaLikeSync struct {
	MediaLikes []MediaLike `json:"mediaLikes"`
}

type MediaFavorite struct {
	MediaId   uint64 `json:"mediaId,string"`
	UserId    uint64 `json:"userId,string"`
	CreatedAt int64  `json:"createdAt"`
}

// ReqMediaFavoriteSync ReqMediaFavorite request
type ReqMediaFavoriteSync struct {
	MediaFavorites []MediaFavorite `json:"mediaFavorites"`
}

type Comment struct {
	MediaId   uint64 `json:"mediaId,string"`
	UserId    uint64 `json:"userId,string"`
	CreatedAt int64  `json:"createdAt"`
	Content   string `json:"content"`
	ThirdId   string `josn:"thirdId,string"`
}

// ReqComment ReqComment request
type ReqCommentSync struct {
	Comments []Comment `json:"comments"`
}

type CommentLike struct {
	MediaId   uint64 `json:"mediaId,string"`
	UserId    uint64 `json:"userId,string"`
	CommentId uint64 `json:"commentId,string"`
	CreatedAt int64  `json:"createdAt"`
}

type ReqCommentLikeSync struct {
	CommentLikes []CommentLike `json:"commentLikes"`
}
