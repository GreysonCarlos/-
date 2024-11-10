package service

import (
	"errors"
	"github.com/GreysonCarlos/Topic-web/repository"
	"sync"
)

// 页面信息包含话题、回帖列表
type PageInfo struct {
	Topic		*repository.Topic
	PostList	[]*repository.Post	
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topidId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
			topicId: topidId,
		}
}

type QueryPageInfoFlow struct {
	topicId		int64
	pageInfo	*PageInfo

	topic		*repository.Topic
	posts		[]*repository.Post
}

// 参数校验
func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

// 检查topicId
func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

// 
func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	// 并行处理信息，一个goroutine获取topic信息，一个goroutine获取posts信息
	wg.Add(2)
	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)	// 根据ID获取topic
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryPostByParentId(f.topicId)
		f.posts = posts
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
			Topic: f.topic, 
			PostList: f.posts,
		}
	return nil
}