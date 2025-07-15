package database

import (
	"justeat/internal/config"	// ใช้ดึง config DB
	"justeat/internal/models"	// ดึงโครงสร้างที่ใช้ในตารางฐานข้อมูล หรือ entity แหละ

	"gorm.io/driver/sqlite"		// ใช้เชื่อมต่อ sqlite
	"gorm.io/gorm"
)

// ฟังก์ชัน Connect 
func Connect(cfg *config.Config) *gorm.DB {
	// สร้างการเชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		panic("ไม่สามารถเชื่อมต่อฐานข้อมูลได้: " + err.Error())
	}

	// gorm สร้างตารางตาม models ที่สร้างไว้
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("ไม่สามารถสร้างตารางในฐานข้อมูลได้: " + err.Error())
	}

	return db
}