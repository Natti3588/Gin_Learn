package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// :name のようにコロンで始まる部分は「動的パラメータ」と呼んで、URLのその位置に来た値を変数として取り出せる。
	router.GET("/user/:name", func(c *gin.Context) {
		// c.Param() で、:name と書いた部分の実際の値を取得する。
		name := c.Param("name")
		// c.String() は文字列としてレスポンスを返すメソッドです。
		// 第二引数に、フォーマットが使える。
		c.String(http.StatusOK, "Hello %s", name)
	})

	// *action は / を含めて残り全部にマッチします。
	// （例： GET /user/Natti/send -> /send が取得される） -
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
