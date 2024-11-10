package repository

import (
	"encoding/json"
	"os"
	"sync"
)

type Post struct {
	Id				int64		`json:"id"`
	ParentId		int64		`json:"parent_id"`
	Content		 	string		`json:"content"`
	CreateTime 		int64		`json:"create_time"`	
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

// 添加回帖功能
func (*PostDao) InsertPost(post *Post) error {
	f, err := os.OpenFile("../data/post", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	marshal, _ := json.Marshal(post)
	if _, err = f.WriteString(string(marshal)); err != nil {
		return err
	}

	rwMutex.Lock()
	postlist, ok := postIndexMap[post.ParentId]
	// 插入帖子
	if !ok {
		postIndexMap[post.ParentId] = []*Post{post}
	} else {
		postlist = append(postlist, post)
		postIndexMap[post.ParentId] = postlist
	}
	rwMutex.Unlock()
	return nil
}