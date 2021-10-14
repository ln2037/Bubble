package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {

}

func main() {
	router := gin.Default()
	//模板文件的静态文件的位置
	router.Static("/static", "static")
	//模板文件的位置
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	//v1Group := router.Group("/v")
	//{
	//	//待办事项
	//
	//}

	router.Run(":9999")
}
