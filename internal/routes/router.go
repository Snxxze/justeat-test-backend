package routes

import (
	"justeat/internal/config"
	"justeat/internal/handler"
	"justeat/internal/middleware"
	"justeat/internal/repository"
	"justeat/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register จะทำการเชื่อมทุกอย่างเข้า gin router
func Register(r *gin.Engine, db *gorm.DB, cfg *config.Config) {

	// สร้าง Repository ชั้นล่าง -> ส่งให้ Service -> ส่งให้ Handler
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, cfg.JWTSecret)

	// สร้าง API
	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	// protected routes ที่ต้องแนบ token
	auth := api.Group("/")
	auth.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
	{
		auth.GET("/me", userHandler.Me)
	}
}