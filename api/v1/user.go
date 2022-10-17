package v1

// 导入gin包
import (
	"fmt"
	"gin-mysql-vue/model"
	"gin-mysql-vue/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 控制器函数
func DoLogin(c *gin.Context) {
	// 获取post请求参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 通过请求上下文对象Context, 直接往客户端返回一个字符串
	c.String(200, "username=%s,password=%s", username, password)
}

// 控制器函数
func Mechanism(c *gin.Context) {
	// 获取get请求参数
	operation := c.Query("operation")
	wash := c.Query("wash")
	fmt.Println("operation:", operation)
	fmt.Println("wash:", wash)

	// 通过请求上下文对象Context, 直接往客户端返回一个字符串
	c.String(200, "operation=%s,wash=%s", operation, wash)
}

// 控制器函数
// func GetUser(c *gin.Context) {
// 	// 获取url参数id
// 	id := c.Param("id")
// 	fmt.Println("id:", id)

// 	// 通过请求上下文对象Context, 直接往客户端返回一个字符串
// 	c.String(200, "id=%s", id)
// }

// 控制器函数
func AddUser(c *gin.Context) {
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

func GetUser(c *gin.Context) {

	id_str := c.Param("id")

	id, err := strconv.Atoi(id_str)
	if err != nil {
		// c.String(errmsg.ERROR, "error")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, _ := model.GetUser(id)

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)

	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
