package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 8 << 20 は8MBです。 multipart処理時にメモリに保持する上限です。
	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(c *gin.Context) {

		// c.MultipartForm() でフォーム全体を取得
		// 失敗すれば 400 でJSONを返す
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// filesという名前で複数送られてきたファイルをリストで取得
		files := form.File["files"]

		// 取得したファイルを一つずつ回して./files/に格納する処理
		for _, file := range files {
			dst := filepath.Join("./files/", filepath.Base(file.Filename))
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	router.Run(":8080")
}
