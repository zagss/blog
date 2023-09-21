package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zagss/blog/models"
	"github.com/zagss/blog/pkg/e"
	"github.com/zagss/blog/pkg/setting"
	"github.com/zagss/blog/pkg/util"
	"net/http"
	"strconv"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	// 筛选条件
	maps := make(map[string]interface{})
	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = strconv.Atoi(arg)
		maps["state"] = state
	}

	// 返回值
	data := make(map[string]interface{})
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": "",
	})
}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {

}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {

}
