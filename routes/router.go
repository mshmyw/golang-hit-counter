package routes

import (
	"github.com/gin-gonic/gin"
	"go.xstore.local/go-hit-counter/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// v1 := router.Group("/api/v1", middleware.Counter())
	v1 := router.Group("/api/v1")
	{
		hitCounter := new(controllers.HitCounter)
		v1.GET("/hit-counter", hitCounter.Index)
		v1.GET("/count", hitCounter.Count)
		v1.GET("/count/tag.svg", hitCounter.CountTag)
		v1.POST("/hit-counter", hitCounter.Store)
		v1.PUT("/hit-counter/:id", hitCounter.Update)
		v1.DELETE("/hit-counter/:id", hitCounter.Destroy)
		v1.GET("/hit-counter/:id", hitCounter.Show)

	}

	return router

}
