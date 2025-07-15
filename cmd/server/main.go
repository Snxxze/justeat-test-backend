package main

import (
	"justeat/internal/config"
	"justeat/internal/database"
	"justeat/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	// โหลด config
	cfg := config.Load()

	// เชื่อม DB
	db := database.Connect(cfg)

	// สร้าง router 
	r := gin.Default()

	// api/register
	routes.Register(r, db, cfg)

	// run
	log.Println("[SERVER] เริ่มต้นที่พอร์ต : ", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("ไม่สามารถรันเซิฟเวอร์ได้: ", err)
	}
}