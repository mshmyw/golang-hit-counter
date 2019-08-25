package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// parse url
func GetURL(URL string) string {
	u, _ := url.Parse(URL)

	return u.Host + u.Path
}

func CheckValidCookie(c *gin.Context, url string) bool {
	cookie, err := c.Request.Cookie(url)
	if err != nil {
		fmt.Println("cookie wrong")
		return false
	}
	int64ExpireTime, err := strconv.ParseInt(cookie.Value, 10, 64)
	if err == nil {
		if int64ExpireTime > time.Now().Unix() {
			return true
		}
	}

	return false
}

// func findPage(url string) Page {
// 	var page Page
// 	db := GetDB()
// 	db.Where("url = ?", url).Find(&page)

// 	return page
// }

// get title from database or by httpClient
// func GetTitle(url string) string {
// 	// get title from database
// 	page := findPage(url)
// 	if page.Id != 0 {
// 		return page.Title
// 	}

// 	// get title by httpClient
// 	doc, err := goquery.NewDocument(url)
// 	if err != nil {
// 		return url
// 	}
// 	title := doc.Find("title").Text()
// 	title = strings.Trim(title, " ")
// 	if len(title) == 0 {
// 		return url
// 	}
// 	return title
// }
