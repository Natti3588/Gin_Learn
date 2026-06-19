package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定数としてアップロードされるサイズを1MBでセット
const (
	MaxUploadSize = 1 << 20
)

func uploadHandler(c *gin.Context) {
	// http.MaxBytesReader が c.Request.Body をラップします。 許可された容量を超えた場合、リーダーは停止してエラーを返す
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxUploadSize)

	// ParseMultipartForm でフォームを解析
	if err := c.Request.ParseMultipartForm(MaxUploadSize); err != nil {
		// エラーが MaxBytesError の時に実行
		if _, ok := err.(*http.MaxBytesError); ok {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": fmt.Sprintf("file too large (max: %d bytes)", MaxUploadSize),
			})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file form required"})
		return
	}
	defer file.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": "upload successful",
	})
}

func main() {
	router := gin.Default()
	router.POST("/upload", uploadHandler)
	router.Run(":8080")
}
