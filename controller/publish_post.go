package controller

import (
	"github.com/GreysonCarlos/Topic-web/service"
	"strconv"
)

func PublishPost(topicIdStr, content string) *PageData{
	topicId, _ := strconv.ParseInt(topicIdStr, 10, 64)
	postId, err := service.PublishPost(topicId, content)
	if err != nil {
		return &PageData{
				Code: -1,
				Msg: err.Error(),
		}
	}
	return &PageData{
			Code: 1,
			Msg: "success",
			Data: map[string]int64{
					"post_id": postId,
			},
	}
}

