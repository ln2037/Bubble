package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
	ID int	`json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

var (
	DB *gorm.DB
)

func initMysql() (err error){
	connStr := "root:1234@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open("mysql", connStr)
	return
}

func main() {
	//创建数据库连接
	if err := initMysql(); err != nil {
		panic(err)
	}
	defer DB.Close()
	//把结构体与数据库中的表进行绑定
	DB.AutoMigrate()
	//加载页面
	router := gin.Default()
	//模板文件的静态文件的位置
	router.Static("/static", "static")
	//模板文件的位置
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	//访问请求
	v1Group := router.Group("/v")
	{
		//添加
		v1Group.POST("/todo", func(context *gin.Context) {
			//1.获取前端传送的参数
			//2.保存到数据库中

		})
		//查看所有事项

		//修改事项

		//删除事项


	}

	router.Run(":9999")
}
