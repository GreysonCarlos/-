package repository

import (
	"sync"
)

type Topic struct {
	Id			int64		`json:"id"`
	Title 		string		`json:"title"`
	Content 	string		`json:"content"`
	CreateTime 	int64		`json:"create_time"`
}

type TopicDao struct {
}

var topicDao *TopicDao
var topicOnce sync.Once		// 单例模式

func (Topic) TableName() string {
	return "topic"
}


func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do( 
		func () {
			topicDao = &TopicDao{}
		})
	return topicDao
}

// 根据话题ID查询话题
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
