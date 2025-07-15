package handler

import (
	"justeat/internal/service"

	"net/http"

	"github.com/gin-gonic/gin" // framework ใช้สร้าง API
)

// UserHandler ที่รวมฟังก์ชันรับ request จาก user
type UserHandler struct {
	Service *service.UserService
	Secret string
}

// เชื่อมกับ service
func NewUserHandler(s *service.UserService, secret string) *UserHandler {
	return &UserHandler{Service: s, Secret: secret}
}

// RegisterInput ข้อมูลที่เราคาดหวังจะได้รับจาก user
// binding:"required" คือ ต้องใส่
type RegisterInput struct {
	Name		string	`json:"name" binding:"required"`
	Email		string	`json:"email" binding:"required,email"`
	Password	string	`json:"password" binding:"required,min=6"`
}

// Register คือ POST(ส่งข้อมูลใหม่ไปยังเซิฟเวอร์)
func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterInput		// var(variable declaration)  input(variable name)  Regis..(variable type)

	// ผูกข้อมูลที่ user ส่งมา
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลให้ service เพื่อทำการสมัคร
	// และ Service.Register จะส่งค่า user กลับมาที่ user
	// "รับ input → ส่งให้ service → ตอบกลับผลลัพธ์"
	user, err := h.Service.Register(input.Name, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// จะไม่ส่ง password กลับไป
	c.JSON(http.StatusOK, gin.H{
		"id":	user.ID,
		"name":	user.Name,
		"email":user.Email,
	})

}

// Login
type LoginInput struct {
	Email	string	`json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เรียก service
	token, err := h.Service.Login(input.Email, input.Password, h.Secret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


// struct ใหม่ เพื่อควบคุบการตอบกลับ 
type UserResponse struct {
	Name 	string	`json:"name"`
	Email	string	`json:"email"`
}

// userID จาก token
func (h *UserHandler) Me(c *gin.Context) {
	//ดึง userID จาก context ที่ใส่ไว้ตอน middleware
	userIDRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userIDRaw.(uint)

	user, err := h.Service.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบผู้ใช์"})
		return
	}

	// ตอบกลับ
	c.JSON(http.StatusOK, UserResponse{
		Name:	user.Name,
		Email: user.Email,
	})
}