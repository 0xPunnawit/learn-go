package main

import (
	"fmt"
	"goredis/repositories"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db := initDatabase()
	redisClient := initRedis()

	productRepo := repositories.NewProductRepositoryRedis(db, redisClient)
	products, err := productRepo.GetProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(products)
}

func initDatabase() *gorm.DB {
	// ค่าเชื่อมต่อแบบตรง ๆ ไม่ .env
	host := "localhost"
	port := "5888"
	user := "postgres"
	password := "postgres"
	dbname := "postgres"

	// สร้าง DSN (Connection string)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, user, password, dbname, port)

	// เชื่อมต่อด้วย GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	return db

}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
