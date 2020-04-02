package main

import (
	"fmt"
	"github.com/zlzz-rec/zlyq-go-sdk/util"

	"github.com/zlzz-rec/zlyq-go-sdk/handler"
	"github.com/zlzz-rec/zlyq-go-sdk/model"
)

// sync_user_example
func sync_user_example(client *handler.Client) {

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

// sync_user_follow_example
func sync_user_follow_example(client *handler.Client) {

	// 用户关注数据导入
	var userId uint64 = 457741243892858880
	followMap := map[string]int64{
		"457751121931763712": 0,
	}
	req := &model.ReqUserFollow{
		UserId:    util.Uint2Str(userId),
		FollowMap: followMap,
	}

	if err := client.UserFollowSynchronize(req); err != nil {
		fmt.Println(err)
	}
}

// sync_history_example
func sync_history_example(client *handler.Client) {

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
	if err := client.HistorySynchronize(req); err != nil {
		fmt.Println(err)
	}
}

func main() {

	// 同步用户数据example
	//appClient := &handler.Client{
	//	AppId:     450007472627785728,
	//	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	//	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	//	Address:     "http://testappapi.zplatform.cn",
	//}
	//sync_user_example(appClient)

	// 同步用户关注 example
	appClient := &handler.Client{
		AppId:     450007472627785728,
		AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
		AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
		Address:   "http://testappapi.zplatform.cn",
	}
	sync_user_follow_example(appClient)

	// 同步历史交互数据example
	//trackClient := &handler.Client{
	//	AppId:     450007472627785728,
	//	AppKey:    "bb4ddb451bdd80af204d9f464fbf07df",
	//	AppSecret: "2d4964bbafde4bf415f9e5b81c4556b3",
	//	Address:   "http://testtrackapi.zplatform.cn",
	//}
	//sync_history_example(trackClient)

}
