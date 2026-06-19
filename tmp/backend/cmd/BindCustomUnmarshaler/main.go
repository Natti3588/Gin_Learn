package main

import (
	"encoding"
	"strings"

	"github.com/gin-gonic/gin"
)

type Birthday string

func (b *Birthday) UnmarshalText(text []byte) error {
	*b = Birthday(strings.Replace(string(text), "-", "/", -1))
	return nil
}

var _ encoding.TextUnmarshaler = (*Birthday)(nil)

func main() {
	router := gin.Default()
	var request struct {
		Birthday         Birthday   `form:"birthday,parser=encoding.TextUnmarshaler"`
		Birthdays        []Birthday `form:"birthdays,parser=encoding.TextUnmarshaler" collection_format:"csv"`
		BirthdaysDefault []Birthday `form:"birthdaysDef,default=2020-09-01;2020-09-02,parser=encoding.TextUnmarshaler" collection_format:"csv"`
	}
	router.GET("/test", func(c *gin.Context) {
		_ = c.BindQuery(&request)
		c.JSON(200, request)
	})
	router.Run(":8088")
}
