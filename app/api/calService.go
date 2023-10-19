package api

import (
	"com.huahuo/app/app/dao/model"
	"com.huahuo/app/dao"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"time"
)

func upload(ctx *gin.Context) {
	data := ctx.Query("data")
	result := ctx.Query("result")
	historyTime := time.Now().Format("2006-01-02 15:04:05")
	history := model.CalHistory{Data: data, Result: result, Time: historyTime}
	dao.DB.Create(&history)
	ctx.JSON(200, "保存历史成功")
}
func list(ctx *gin.Context) {
	var historys []model.CalHistory // 改为切片类型
	dao.DB.Find(&historys)          // 传入切片对象的指针
	ctx.JSON(200, historys)
}
func clearHistory(ctx *gin.Context) {
	dao.DB.Exec("truncate table cal_history;")
	ctx.JSON(200, "清空历史记录成功！")
}
func listRateBorrow(ctx *gin.Context) {
	var rates []model.CalRate                  // 改为切片类型
	dao.DB.Where("type = ?", "1").Find(&rates) // 传入切片对象的指针
	ctx.JSON(200, rates)
}

func listRateSave(ctx *gin.Context) {
	var rates []model.CalRate                  // 改为切片类型
	dao.DB.Where("type = ?", "0").Find(&rates) // 传入切片对象的指针
	ctx.JSON(200, rates)
}

func update(ctx *gin.Context) {
	var rates []model.CalRate
	err := ctx.ShouldBindJSON(&rates)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	for _, rate := range rates {
		err := dao.DB.Model(&model.CalRate{}).Where("id = ?", rate.ID).Update("rate", rate.Rate).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rate"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Rates updated successfully"})
}

func getSaveResult(ctx *gin.Context) {
	var rates []model.CalRate                  // 改为切片类型
	dao.DB.Where("type = ?", "0").Find(&rates) // 传入切片对象的指针
	num := ctx.Query("num")
	principal := ctx.Query("principal")
	result := calculateAmount(num, principal, rates)
	ctx.JSON(http.StatusOK, result)
}
func getBorrowResult(ctx *gin.Context) {
	var rates []model.CalRate                  // 改为切片类型
	dao.DB.Where("type = ?", "1").Find(&rates) // 传入切片对象的指针
	num := ctx.Query("num")
	principal := ctx.Query("principal")
	result := calculateAmount(num, principal, rates)
	ctx.JSON(http.StatusOK, result)
}

func calculateAmount(numStr string, principalStr string, rates []model.CalRate) float64 {
	num, _ := strconv.Atoi(numStr)
	principal, _ := strconv.ParseFloat(principalStr, 64)
	prevNum := 0
	var rate float64
	for _, r := range rates {
		if int(r.Num) >= num {
			rate, _ = strconv.ParseFloat(r.Rate, 64)
			break
		}
		prevNum = int(r.Num)
	}
	rate /= 100.0 // 将利率转换为小数
	months := (num - prevNum) * 12
	finalAmount := principal * math.Pow(1+rate/12, float64(months))
	return math.Floor(finalAmount*100+0.5) / 100
}
