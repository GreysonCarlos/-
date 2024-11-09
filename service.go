package main

import (
	"github.com/GreysonCarlos/Topic-web/repository"
	"github.com/GreysonCarlos/Topic-web/controller"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init("../data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
			topicId := c.Param("Id")
			data := controller.QueryPageInfo(topicId)
			c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}