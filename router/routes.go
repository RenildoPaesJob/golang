package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func inicializeRoutes(router *gin.Engine){

	v1 := router.Group("/api/v1")

	{
		// SHOW
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":"GET Opening!",
			})
		})

		// CREATE
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":"POST Opening!",
			})
		})

		// SHOW
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":"DELETE Opening!",
			})
		})

		// SHOW
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":"PUT Opening!",
			})
		})

		// SHOW
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":"List Openings!",
			})
		})

	}
}