package repository

import (
	"sync"
	"time"
)

type Post struct {
	Id				int64		`json:"id"`
	ParentId		int64		`json:"parent_id"`
	Content		 	string		`json:"content"`
	CreateTime 		time.Time	`json:"create_time"`	
}

func (Post) TableName() string {
	return "post"
}

type PostDao struct {
}

var postDao *PostDao
var postOnce sync.Once

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

// 根据话题ID查询回帖列表
func (*PostDao) QueryPostByParentId(parentId int64) []*Post {
	return postIndexMap[parentId]
}
