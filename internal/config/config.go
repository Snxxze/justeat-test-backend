package config

import (
	"os"		//	ใช้เพื่ออ่านค่าจากระบบ
	"log"		
	"github.com/joho/godotenv"		// โหลดค่าจากไฟล์
)

// เก็บค่าของ config ทั้งระบบ
type Config struct {
	Port	string
	DBPath	string
	JWTSecret	string
}

// โหลดค่าจาก .env หรือระบบ
func Load() *Config {
	// ลองโหลด .env
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == ""{
		dbPath = "justeat.db"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "supersecretkey"
	}

	// แสดงค่า config
	log.Println("[CONFIG] Loaded Port = ", port, " DB_Path = ", dbPath)

	// คืนค่าให้ส่วนอื่นๆ ใช้
	return &Config{
		Port: port,
		DBPath: dbPath,
		JWTSecret: jwtSecret,
	}
}