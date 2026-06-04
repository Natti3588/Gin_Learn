package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 8 << 20 は8MBです。 multipart処理時にメモリに保持する上限です。
	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(c *gin.Context) {

		// フォームの file フィールドからアップロードファイルを取得します。
		// 失敗した場合は 400 bad Request を JSONで返して終了
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 受け取ったファイル名をサーバーのログに出力
		log.Println(file.Filename)

		// 保存先のパスを組み立てる。
		dst := filepath.Join("./files/", filepath.Base(file.Filename))

		// ファイルを ./files/ ディレクトリに保存
		c.SaveUploadedFile(file, dst)

		// 成功メッセージを返す
		c.String(http.StatusOK, fmt.Sprintf("`%s` uploaded!", file.Filename))
	})

	router.Run(":8080")
}
