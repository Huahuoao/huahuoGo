package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	history := r.Group("/history")

	history.POST("/upload", upload)
	history.GET("/list", list)
	history.GET("/clear", clearHistory)
	rate := r.Group("/rate")
	rate.POST("/getRatesBorrow", listRateBorrow)
	rate.POST("/getRatesSave", listRateSave)
	rate.POST("/update", update)
	rate.POST("/result/save", getSaveResult)
	rate.POST("/result/borrow", getBorrowResult)
}
