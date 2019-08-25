package middleware

// func Counter() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if cookie, err := c.Request.Cookie("counter"); err == nil {
// 			value := cookie.Value
// 			if len(value) == 0 {
// 				cookie.Value = "0"
// 			} else {
// 				if v, err := strconv.Atoi(value); err == nil {
// 					i := v + 1
// 					cookie.Value = fmt.Sprintf("%d", i)
// 				}
// 			}
// 			http.SetCookie(c.Writer, &http.Cookie{
// 				Name:    url,
// 				Value:   strconv.FormatInt(time.Now().Add(time.Minute*5).Unix(), 10),
// 				Expires: time.Unix(60*5, 0),
// 			})
// 			//before request
// 			c.Next()
// 			//after request
// 			fmt.Println("before middleware " + cookie.Value)
// 		}
// 	}
// }
