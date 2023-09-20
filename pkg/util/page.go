package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0
	fmt.Println("page: ", c.Query("page"))
	return result
}
