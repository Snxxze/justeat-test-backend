package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware จะตรวจสอบ JWT token ใน header ของทุก request
func JWTAuthMiddleware(secret string) gin.HandlerFunc{
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")		// ดึง token จาก header

		// ตรวจสอบ header มี bearer มั้ย  ถ้าไม่มี return กลับทันที
		if !strings.HasPrefix(authHeader, "Bearer "){
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "ไม่พบ token ใน header (ต้องใช้ Bearer token)",
			})
			return 
		}
		
		// แยก token ออกจาก header โดย trim
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// ตรวจสอบ token มาจากระบบเราหรือไม่ โดยใช้ secret key ที่เราต้องใน .env
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			// คืนค่า secret ที่ใช้ตรวจสอบ
			return []byte(secret), nil
		})

		// ถ้า token หมดอายุ, ปลอม หรือไม่ valid
		if err != nil || !token.Valid{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token ไม่ถูกต้องหรือหมดอายุ",})
			return 
		}

		// ดึง userID จาก claims ใน token
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["sub"].(float64))		// sub คือ subject ที่กำหนดตอน generate token (internal/utils/jwt.go)
		c.Set("userID", userID)		// แนบ userID ไปกับ context เพื่อให้ handler เอาไปใช้ต่อ
		c.Next()
	}

}