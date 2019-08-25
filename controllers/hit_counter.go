package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.xstore.local/go-hit-counter/database"
	"go.xstore.local/go-hit-counter/models"
	"go.xstore.local/go-hit-counter/utils"
)

type HitCounter struct {
	Basic
}

func (a *HitCounter) Count(c *gin.Context) {
	// Get the url out of a request either passed as a query parameter or taken from the referrer. Remove any query
	referer := c.DefaultQuery("url", c.Request.Referer())
	url := utils.GetURL(referer)
	if len(url) == 0 {
		a.JsonFail(c, http.StatusBadRequest, "not found url")
		return
	}

	valid_cookie := utils.CheckValidCookie(c, url)
	if valid_cookie == false {
		database.AddView(url)

		duration := time.Minute * 5
		c.SetCookie(
			url, 
			strconv.FormatInt(time.Now().Add(time.Minute*5).Unix(), 10),
			int(duration.Seconds()),
			"/", 
			"localhost", 
			false, 
			true,
		)
	}
	fmt.Println(c.Request.Cookie(url))

	hitCounter := database.GetCount(url)
	a.JsonSuccess(c, http.StatusOK, gin.H{"data": *hitCounter.Count})
}

func (a *HitCounter) Index(c *gin.Context) {
	var counters []models.HitCounter

	database.DB.Select("url, count, memo").Order("count").Find(&counters)

	a.JsonSuccess(c, http.StatusOK, gin.H{"data": counters})
}

func (a *HitCounter) Store(c *gin.Context) {
	var request CreateRequest

	if err := c.ShouldBind(&request); err == nil {
		var count int
		database.DB.Model(&models.HitCounter{}).Where("url = ?", request.URL).Count(&count)

		if count < 0 {
			a.JsonFail(c, http.StatusBadRequest, "hit count wrong")
			return
		}

		password := []byte(request.Memo)
		md5Ctx := md5.New()
		md5Ctx.Write(password)
		cipherStr := md5Ctx.Sum(nil)
		hitCount := models.HitCounter{
			URL:   request.URL,
			Count: &request.Count,
			Memo:  hex.EncodeToString(cipherStr),
		}

		if err := database.DB.Create(&hitCount).Error; err != nil {
			a.JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		a.JsonSuccess(c, http.StatusCreated, gin.H{"message": "创建成功"})
	} else {
		a.JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

func (a *HitCounter) Update(c *gin.Context) {
	var request UpdateRequest

	if err := c.ShouldBind(&request); err == nil {
		var hitCounter models.HitCounter
		if database.DB.First(&hitCounter, c.Param("url")).Error != nil {
			a.JsonFail(c, http.StatusNotFound, "数据不存在")
			return
		}

		hitCounter.Count = &request.Count

		if err := database.DB.Save(&hitCounter).Error; err != nil {
			a.JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		a.JsonSuccess(c, http.StatusCreated, gin.H{})
	} else {
		a.JsonFail(c, http.StatusBadRequest, err.Error())
	}
}

func (a *HitCounter) Show(c *gin.Context) {
	var hitCounter models.HitCounter

	if database.DB.Select("url, count, memo, created_at, updated_at").First(&hitCounter, c.Param("url")).Error != nil {
		a.JsonFail(c, http.StatusNotFound, "数据不存在")
		return
	}

	a.JsonSuccess(c, http.StatusCreated, gin.H{"data": hitCounter})
}

func (a *HitCounter) Destroy(c *gin.Context) {
	var hitCounter models.HitCounter

	if database.DB.First(&hitCounter, c.Param("url")).Error != nil {
		a.JsonFail(c, http.StatusNotFound, "数据不存在")
		return
	}

	if err := database.DB.Unscoped().Delete(&hitCounter).Error; err != nil {
		a.JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}

	a.JsonSuccess(c, http.StatusCreated, gin.H{})

}

type CountRequest struct {
	URL string `form:"url" json:"url"`
}

type UpdateRequest struct {
	Count int `form:"count" json:"count" binding:"required"`
	Memo  int `form:"memo" json:"memo"`
}

type CreateRequest struct {
	URL   string `form:"url" json:"url" binding:"required"`
	Count int    `form:"count" json:"count"` // 允许0值
	Memo  string `form:"memo" json:"memo"`
}
