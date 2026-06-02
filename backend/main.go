package main

// net/http 標準ライブラリで、HTTPステータスコードを使うために読み込む。
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ハンドラ関数。 Contextがリクエストとレスポンスの情報を持つオブジェクト。

func getting(c *gin.Context) {
	// c.JSON() はレスポンスをJSON形式で返す意味で、第一引数にHTTPステータスコード、
	// 第二引数が中身。
	// gin.H は map[string]interface{}のショートカットで、JSONのキーと値を簡単に書ける。
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PATCH"})
}

func head(c *gin.Context) {
	// c.Status() は body を返さずに、ステータスコードだけを設定します。
	c.Status(http.StatusOK)
}

func options(c *gin.Context) {
	c.Status(http.StatusOK)
}

func main() {
	// ルータを作成しています。 Default では Logger と Recovery 2つのミドルウェアが組み込まれたルーターを返す
	router := gin.Default()

	// ルーティングを定義しています。 （例： GET /someGetが呼ばれたら、ハンドラ関数のgetting()が呼ばれて、JSON形式の{"method" : "GET"}が返される）
	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someeEad", head)
	router.OPTIONS("/someOptions", options)

	router.Run(":8080")
}
