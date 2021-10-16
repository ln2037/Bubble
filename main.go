package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

type Todo struct {
	ID int	`json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

var (
	DB *gorm.DB
)

func (Todo)TableName()string {
	return "todo"
}

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
	DB.AutoMigrate(Todo{})
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
	v1Group := router.Group("/v1")
	{
		//添加
		v1Group.POST("/todo", func(context *gin.Context) {
			//1.获取前端传送的参数
			var todo Todo
			//context.BindJSON(&todo)
			//context.BindJSON()
			context.ShouldBind(&todo)
			fmt.Println(todo)
			//2.保存到数据库中
			if err := DB.Debug().Create(&todo).Error; err != nil {
				context.JSON(http.StatusOK,  gin.H{
					"error":err.Error(),
				})
			} else {
				context.JSON(http.StatusOK, todo)
			}
			//3.返回响应

		})
		//查看所有事项
		v1Group.GET("/todo", func(context *gin.Context) {
			//查找数据库中所有的项目
			var todoList []Todo
			err := DB.Debug().Find(&todoList).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error":err,
				})
			}
			//返回结果
			context.JSON(http.StatusOK, todoList)
		})
		//修改事项
		v1Group.PUT("/todo/:id", func(context *gin.Context) {
			//获取id
			id, ok := context.Params.Get("id")
			if ok == false {
				context.JSON(http.StatusOK, gin.H{
					"error":"无效的id",
				})
				return
			}
			var todo Todo
			//根据id查询具体信息,若查询不到直接返回
			if err := DB.Where("id = ?", id).First(&todo).Error; err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error": err,
				})
				return
			}
			//获取前端传送的todo信息
			context.BindJSON(&todo)
			//重新存储到数据库中
			//返回前端结果
			err := DB.Debug().Save(todo).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error":err,
				})
				return
			}
			context.JSON(http.StatusOK, todo)
		})
		//删除事项
		v1Group.DELETE("/todo/:id", func(context *gin.Context) {
			//获取url中的参数
			id, ok := context.Params.Get("id")
			if !ok {
				context.JSON(http.StatusOK, gin.H{
					"error":"id不存在",
				})
			}
			//判断id是否存在
			var todo Todo
			err := DB.Debug().Where("id = ? ", id).First(&todo).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"error":err,
				})
				return
			}
			//删除相关的数据
			err = DB.Debug().Delete(todo).Error
			if err != nil {
				context.JSON(http.StatusOK, gin.H{"error":err})
				return
			}
			context.JSON(http.StatusOK, todo)
		})


	}

	router.Run(":9999")
}
