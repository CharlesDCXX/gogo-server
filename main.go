package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test(http://localhost/ping)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value(http://localhost/user/donghai)
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
	// AsciiJSON
	r.GET("/someJson", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"name": "donghai< b",
		//	"age":  18,
		//}
		//也可以用下main方是给data当赋值
		// gin.H{} 是一个用于创建 map[string]interface{} 类型的快捷方式。它常用于在 Gin 路由处理程序中传递数据给 HTML 模板或 JSON 响应
		data := gin.H{"name": "donghai", "age": 18}

		// 使用 ASCII 字符集返回 JSON 响应
		//c.AsciiJSON(http.StatusOK, data)
		//// 使用 UTF-8 字符集返回 JSON 响应
		c.JSON(http.StatusOK, data)
	})
	// HTML渲染
	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLFiles("templates/index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title": "webHTML"})
	})
	return r
}

func main() {
	db["donghai"] = "董海donghai"
	db["xiaomi"] = "xiaoming"
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080

	r.Run(":8080")

}
