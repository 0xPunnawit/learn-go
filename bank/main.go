package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {

	initTimeZone()
	db := initDatabase()
	fmt.Println("✅ Connected to PostgreSQL successfully!")

	// customerRepository := repository.NewCustomerRepositoryDB(db)
	// _ = customerRepository

	// customers, err := customerRepository.GetAll()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)

	// customer, err := customerRepository.GetById(2001)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer)

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)

	// customerRepositoryMock := repository.NewCustomerRepositoryMock()
	// _ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepositoryDB := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accountService)

	// customers, err := customerService.GetCustomer(2001)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	port := ":8000"
	logs.Info("Banking service started", zap.String("port", port))

	http.ListenAndServe(":8000", router)

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	// โหลด environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// สร้าง connection string จาก .env
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, dbPass, dbName,
	)

	// ใช้ sqlx.Connect() เพื่อเชื่อม PostgreSQL
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatal("❌ Cannot connect to database:", err)
	}

	db.SetConnMaxIdleTime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
