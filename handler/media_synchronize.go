package handler

import (
	"encoding/json"
	"fmt"

	"github.com/zlzz-rec/zlyq-go-sdk/model"
	"github.com/zlzz-rec/zlyq-go-sdk/util"
)

// VideoSynchronize 同步视频信息
func (client *Client) VideoSynchronize(req *model.ReqVideoSynchronize) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.VideoSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}

// ArticleSynchronize 同步图文信息
func (client *Client) ArticleSynchronize(req *model.ReqArticleSynchronize) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.ArticleSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}

// MediaLikeSynchronize 同步点赞信息
func (client *Client) MediaLikeSynchronize(req *model.ReqMediaLikeSync) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.MediaLikeSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}

// MediaFavoriteSynchronize 同步收藏信息
func (client *Client) MediaFavoriteSynchronize(req *model.ReqMediaFavoriteSync) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.MediaFavoriteSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}

// CommentSynchronize 同步评论信息
func (client *Client) CommentSynchronize(req *model.ReqCommentSync) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.CommentSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}

// CommentLikeSynchronize 同步评论点赞信息
func (client *Client) CommentLikeSynchronize(req *model.ReqCommentLikeSync) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.CommentLikeSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}
