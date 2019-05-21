package web

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"search_center/core/kafka"
)

var db = make(map[string]string)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Get user value
	//r.GET("/center/:router", routerGet)
	r.POST("/center", centerPost)
	//测试kafka
	r.GET("/centerMagic", centerMagic)
	//测试重定向，
	r.GET("/redirect", redirectTest)

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
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

	return r
}

//get
var routerGet = func(c *gin.Context) {
	router := c.Params.ByName("router")
	key := c.DefaultQuery("key", "HaHa!")
	fmt.Println("router: " + router + " , key: " + key)
	//c.JSONP(http.StatusOK,gin.H{"router":router,"key":key})
	c.JSON(http.StatusOK, gin.H{"router": router, "key": key})
}

//Post
var centerPost = func(c *gin.Context) {

	// Read from the Query Param
	key := c.DefaultQuery("key", "haha query!")
	fmt.Println("query: ", key)

	// Read from the Form-data
	key = c.DefaultPostForm("key", "haha form!")
	fmt.Println("form: ", key)

	// Read the Body content
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Continue to use the Body, like Binding it to a struct
	fmt.Println("body: ", string(bodyBytes))
}

var centerMagic = func(c *gin.Context) {
	param := c.DefaultQuery("param", "nil")

	//获取cookie、session、header、body
	fmt.Println("---header/--- \r")
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}

	fmt.Println("---cookies/--- \r")
	for k, v := range c.Request.Cookies() {
		fmt.Println(k, v)
	}

	kafka.Producer(param)

	//重定向
	c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/redirect")
}

var redirectTest = func(c *gin.Context) {

	fmt.Println("---redirect header/--- \r")
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}

	fmt.Println("---cookies/--- \r")
	for k, v := range c.Request.Cookies() {
		fmt.Println(k, v)
	}
	c.String(http.StatusOK, "done")
}
