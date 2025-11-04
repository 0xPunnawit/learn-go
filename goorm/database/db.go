package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("SQL: %v\n=====================================\n", sql)
}

// DB คือ instance ของ *gorm.DB ที่จะใช้ทั่วทั้งระบบ
var DB *gorm.DB

func ConnectDB() {
	// โหลดค่าในไฟล์ .env เข้ามาใน environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// ดึงค่าตัวแปรจาก environment
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// สร้าง Data Source Name (DSN)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, user, password, dbName, port,
	)

	// เชื่อมต่อฐานข้อมูลผ่าน GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// เก็บไว้ในตัวแปร global
	DB = db
	log.Println("✅ Connected to PostgreSQL successfully!")
}
