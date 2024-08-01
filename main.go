package main

import(
	"github.com/gin-gonic/gin"
)

func list(c *gin.Context)  {
	c.JSON(200 , "list")
	print("list")
}

func add(c *gin.Context)  {
	c.JSON(200 , "list")
	print("list")
}

func main() {
	route := gin.Default()

	route.GET("/list",list)
	route.POST("/add",add)

	route.Run(":2010")
}