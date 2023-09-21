package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zagss/blog/pkg/setting"
	v1 "github.com/zagss/blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiV1.GET("/tags", v1.GetTags)
		// 新建标签
		apiV1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiV1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
