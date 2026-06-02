package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {

		// c.DefaultQuery()はクリエ―パラメータの値を返します。 キーが存在しない場合は、第二引数の値を返す
		firstname := c.DefaultQuery("firstname", "Guest")
		// c.Query()は、クエリパラメータの値を返します。 キーが存在しない場合は、空文字を返す
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":8080")
}
