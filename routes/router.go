package routes

import (
	v1 "gin-mysql-vue/api/v1"
	// "github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 设置release模式
	// gin.SetMode(gin.ReleaseMode)
	// 设置debug模式(不设置也默认为Debug)
	// gin.SetMode(gin.DebugMode)

	// 初始化一个http服务对象
	r := gin.Default()

	// nil 为不计算，避免性能消耗，上线应当设置
	// _ = r.SetTrustedProxies(nil)

	router := r.Group("api/v1")
	{
		// 路由定义post请求, url路径为：/user/login, 绑定doLogin控制器函数
		router.POST("/user/login", v1.DoLogin)

		// 路由定义get请求
		router.GET("/mechanism", v1.Mechanism)

		// 创建user
		router.POST("/user/add", v1.AddUser)

		// 根据id获取user
		router.GET("/user/:id", v1.GetUser)

		// 获取user清单
		router.GET("users", v1.GetUsers)
	}

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	// r.Run(":9999")

	// _ = r.Run(utils.HttpPort)

}
