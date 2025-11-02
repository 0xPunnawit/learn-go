package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/golang-jwt/jwt/v5"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

var db *sqlx.DB

const jwtSecret = "0xPunnawit"

func main() {
	// โหลดค่า .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// อ่านค่าจาก .env
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// สร้าง connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	// เชื่อมต่อฐานข้อมูล
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect database: %v", err)
	}
	fmt.Println("✅ Database connected successfully!")

	app := fiber.New()

	//  ใช้ middleware ตรวจ JWT ก่อนเข้าถึง /hello
	app.Use("/hello", RequireAuth)
	app.Get("/hello", Hello)

	app.Post("/signup", Signup)
	app.Post("/login", Login)
	app.Post("/hello", Hello)

	app.Listen(":8000")

}

func Signup(c *fiber.Ctx) error {
	request := SignupRequest{}
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// ตรวจสอบค่าว่าง
	if request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "username and password are required")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`

	var id int
	err = db.QueryRow(query, request.Username, string(passwordHash)).Scan(&id)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user_id": id,
	})
}

func Login(c *fiber.Ctx) error {

	request := LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// ตรวจสอบค่าว่าง
	if request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "username and password are required")
	}

	var user User
	query := "SELECT id, username, password FROM users WHERE username=$1"
	err = db.Get(&user, query, request.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid password")
	}

	claims := jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(user.Id),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"jwtToken": tokenString,
	})
}

func RequireAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid Authorization format")
	}

	tokenStr := parts[1]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid or expired token")
	}

	// ✅ เก็บ user id ไว้ใน context
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Locals("user", claims["iss"])
	}

	return c.Next()
}

func Hello(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Hello user %v 👋 You are authorized!", user),
	})
}

type User struct {
	Id       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ==================== Fiber ====================
func FiberTest() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	// Middleware
	app.Use("/hello", func(c *fiber.Ctx) error {
		c.Locals("name", "Punnawit")
		fmt.Println("befor")
		c.Next()
		fmt.Println("after")
		return nil
	})

	app.Use(requestid.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

	// GET
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Locals("name")
		fmt.Println("/hello")
		return c.SendString(fmt.Sprintf("GET: Hello %v !!", name))
	})

	// POST
	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.SendString("POST: Hello World !!")
	})

	// Parameters Optional
	app.Get("/hello/:name/:surname", func(c *fiber.Ctx) error {
		name := c.Params("name")
		surname := c.Params("surname")
		return c.SendString("Name: " + name + ", surname: " + surname)
	})

	// ParamsInt
	app.Get("/hello/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("ID: %v", id))
	})

	// Query
	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		surename := c.Query("surename")
		return c.SendString("\nname: " + name + "\nsurename: " + surename)
	})

	// Query
	app.Get("/query2", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})

	// Wildcards
	app.Get("/wildcards/*", func(c *fiber.Ctx) error {
		wildcard := c.Params("*")
		return c.SendString(wildcard)
	})

	// Static file
	app.Static("/", "./wwwroot", fiber.Static{
		Index:         "index.html",
		CacheDuration: time.Second * 10,
	})

	// NewError
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})

	// Group
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("version", "v1")
		return c.Next()
	})
	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("V1: Hello World !!")
	})

	v2 := app.Group("/v2")
	v2.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("V2: Hello World !!")
	})

	// Mount
	userApp := fiber.New()
	userApp.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})

	app.Mount("/user", userApp)

	// Server
	app.Server().MaxConnsPerIP = 1
	app.Get("/server", func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 30)
		return c.SendString("server")
	})

	// Enviroment
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"BaseUrl":     c.BaseURL(),
			"HostName":    c.Hostname(),
			"IP":          c.IP(),
			"IPs":         c.IPs(),
			"OriginalUrl": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Method":      c.Method(),
			"Subdomains":  c.Subdomains(),
		})
	})

	// Body
	app.Post("/body", func(c *fiber.Ctx) error {
		// fmt.Printf("IsJson: %v\n", c.Is("json"))
		// fmt.Println(string(c.Body()))

		person := Person{}
		err := c.BodyParser(&person)
		if err != nil {
			return err
		}
		fmt.Println(person)
		return nil
	})

	app.Listen(":8000")
}

type Person struct {
	Id   int    `json:id`
	Name string `json:name`
}
