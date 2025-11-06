package main

import (
	"fmt"
	"goredis/handlers"
	"goredis/repositories"
	"goredis/services"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogServiceRedis(productRepo, redisClient)
	productHandler := handlers.NewCatalogHandler(productService)

	app := fiber.New()

	app.Get("/products", productHandler.GetProducts)

	app.Listen(":8000")
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
