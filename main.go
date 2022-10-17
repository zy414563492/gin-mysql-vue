package main

// 导入gin包
import (
	"fmt"
	"gin-mysql-vue/model"
	"gin-mysql-vue/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 结构体定义
type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

// 控制器函数
func doLogin(c *gin.Context) {
	// 获取post请求参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 通过请求上下文对象Context, 直接往客户端返回一个字符串
	c.String(200, "username=%s,password=%s", username, password)
}

// 控制器函数
func mechanism(c *gin.Context) {
	// 获取get请求参数
	operation := c.Query("operation")
	wash := c.Query("wash")
	fmt.Println("operation:", operation)
	fmt.Println("wash:", wash)

	// 通过请求上下文对象Context, 直接往客户端返回一个字符串
	c.String(200, "operation=%s,wash=%s", operation, wash)
}

// 控制器函数
func getUser(c *gin.Context) {
	// 获取url参数id
	id := c.Param("id")
	fmt.Println("id:", id)

	// 通过请求上下文对象Context, 直接往客户端返回一个字符串
	c.String(200, "id=%s", id)
}

// 控制器函数
func testStruct(c *gin.Context) {
	// 初始化user struct
	u := User{}
	// 通过ShouldBind函数，将请求参数绑定到struct对象， 处理json请求代码是一样的。
	// 如果是post请求则根据Content-Type判断，接收的是json数据，还是普通的http请求参数
	err := c.ShouldBind(&u)
	if err != nil {
		// 绑定成功， 打印请求参数
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

// 控制器函数
// func dbTest(c *gin.Context) {
// 	// 获取post请求参数
// 	name := c.PostForm("name")
// 	email := c.PostForm("email")

// 	db := InitDB()

// 	//创建新用户
// 	newUser := User{
// 		Name:  name,
// 		Email: email,
// 	}
// 	db.Create(&newUser)

// 	//返回结果
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "注册成功",
// 	})
// }

// 控制器函数
func createUser(c *gin.Context) {
	// 获取post请求参数
	name := c.PostForm("name")
	email := c.PostForm("email")
	// role := c.PostForm("role")

	fmt.Println("name:", name)
	fmt.Println("email:", email)

	// 创建新用户
	newUser := model.User{
		Name:  name,
		Email: email,
		Role:  3,
	}
	// fmt.Println("[111]")
	model.CreateUser(&newUser)

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
}

// 全局变量
// var db *gorm.DB

// 入口函数
func main() {

	// 引用数据库
	model.InitDB()

	// 引入路由组件
	routes.InitRouter()

	// 设置release模式
	// gin.SetMode(gin.ReleaseMode)
	// 设置debug模式(不设置也默认为Debug)
	// gin.SetMode(gin.DebugMode)

	// 初始化一个http服务对象
	// r := gin.Default()

	// 路由定义post请求, url路径为：/user/login, 绑定doLogin控制器函数
	// r.POST("/user/login", doLogin)

	// 路由定义get请求
	// r.GET("/mechanism", mechanism)

	// 路由定义获取URL路径参数
	// r.GET("/user/:id", getUser)

	// 路由定义Test
	// r.GET("/test-get-struct", testStruct)
	// r.POST("/test-post-struct", testStruct)

	// //DB test
	// r.POST("/test-db", dbTest)
	// r.POST("/user/create", createUser)

	// 设置一个get请求的路由，url为/ping, 处理函数（或者叫控制器函数）是一个闭包函数。
	// r.GET("/ping", func(c *gin.Context) {
	// 	// 通过请求上下文对象Context, 直接往客户端返回一个json
	// 	// c.JSON(200, gin.H{
	// 	// 	"message": "pong",
	// 	// })
	// 	c.String(http.StatusOK, "Hello Zhuzhu")
	// })

	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	// r.Run(":9999")
}

// func InitDB() *gorm.DB {
// 	//前提是你要先在本机用Navicat创建一个名为go_db的数据库
// 	host := "localhost"
// 	port := "3307"
// 	database := "hisview"
// 	username := "hisview"
// 	password := "hisview"
// 	charset := "utf8"
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
// 		username,
// 		password,
// 		host,
// 		port,
// 		database,
// 		charset)
// 	//这里 gorm.Open()函数与之前版本的不一样，大家注意查看官方最新gorm版本的用法
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Error to Db connection, err: " + err.Error())
// 	}
// 	//这个是gorm自动创建数据表的函数。它会自动在数据库中创建一个名为users的数据表
// 	_ = db.AutoMigrate(&User{})
// 	return db
// }
