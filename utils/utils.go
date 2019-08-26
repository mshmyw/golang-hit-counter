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

func CalculateSVGSizes(count int) map[string]int {
	// Calculate the size of the green half based off the length of count
	text := strconv.Itoa(count)
	sizes := map[string]int{
		"width":    80,
		"recWidth": 50,
		"textX":    55,
	}
	if len(text) > 5 {
		sizes["width"] += 6 * (len(text) - 5)
		sizes["recWidth"] += 6 * (len(text) - 5)
		sizes["textX"] += 3 * (len(text) - 5)
	}

	return sizes
}

func GetSVG(count int, width int, recWidth int, textX int, url string) string {
	// """ Put the count in the pre-defined svg and return it """
	return fmt.Sprintf(`<?xml version="1.0"?>
	<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="20">
	<rect width="30" height="20" fill="#555"/>
	<rect x="30" width="%d" height="20" fill="#4c1"/>
	<rect rx="3" width="80" height="20" fill="transparent"/>
		<g fill="#fff" text-anchor="middle"
			font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11">
				<text x="15" y="14">hits</text>
				<text x="%d" y="14">%d</text>
		</g>
	<!-- This count is for the url: %s -->
	</svg>`, width, recWidth, textX, count, url)
}
