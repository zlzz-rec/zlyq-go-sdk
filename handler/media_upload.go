package handler

import (
	"fmt"
  "encoding/json"

	"github.com/zlzz-rec/zlyq-go-sdk/util"
	"github.com/zlzz-rec/zlyq-go-sdk/model"
)

// UploadImage 上传图片
func (client *Client) UploadImage(req *model.ReqImageUpload) error {

  formParam := map[string]string{
    "source": util.Uint2Str(uint64(req.Source)),
    "description": req.Description,
    "userId":util.Uint2Str(uint64(req.UserId)),
    "thirdId":req.ThirdId,
    "thirdExtra": req.ThirdExtra,
  }
  fileMap := map[string][]byte{
    "image":req.Image,
  }

  resp, err := client.HttpMultiForm(client.Address, util.ImageUploadSynchronizeApi, nil,
  	formParam , fileMap)

	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}


// UploadVideo 上传视频
func (client *Client) UploadVideo(req *model.ReqVideoUpload) error {
  formParam := map[string]string{
    "title":req.Title,
    "content":req.Content,
    "orgFileName":req.OrgFileName,
    "os":util.Uint2Str(uint64(req.Os)),
    "source": util.Uint2Str(uint64(req.Source)),
    "userId":util.Uint2Str(uint64(req.UserId)),
    "thirdId":req.ThirdId,
    "thirdExtra": req.ThirdExtra,
  }
  fileMap := map[string][]byte{
    "image":req.Image,
    "video":req.Video,
  }

  resp, err := client.HttpMultiForm(client.Address, util.VideoUploadSynchronizeApi, nil,
  	formParam , fileMap)

	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}

// UploadArticle 上传图文
func (client *Client) UploadArticle(req *model.ReqArticleUpload) error {

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	resp, err := client.HttpMethod("POST", client.Address, util.ArticleUploadSynchronizeApi, nil, body)
	fmt.Println(resp)
	if err != nil {
		return err
	}

	return nil
}
