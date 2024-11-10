package main

import (
	"log"
	"os"
	"github.com/GreysonCarlos/Topic-web/controller"
	"github.com/GreysonCarlos/Topic-web/repository"
	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	topicId := c.Param("id")
	data := controller.QueryPageInfo(topicId)
	c.JSON(200, data)
}
func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", Handler)
	err := r.Run()
	if err != nil {
		log.Printf("Run failed: %#v", err)
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		log.Printf("Filepath error: %#v", err)
		return err
	}
	return nil
}
// "/home/greyson/ginweb/data/post"