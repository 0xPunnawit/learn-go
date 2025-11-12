package main

import (
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New()

	app.Post("/open-account", accountController.OpenAccount)
	app.Post("/deposit-fund", accountController.DepositFund)
	app.Post("/withdraw-fund", accountController.WithdrawFund)
	app.Post("/close-account", accountController.CloseAccount)

	app.Listen(":8080")
}

// func main() {

// 	server := []string{"localhost:9092"}

// 	producer, err := sarama.NewSyncProducer(server, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer producer.Close()

// 	msg := sarama.ProducerMessage{
// 		Topic: "inghello",
// 		Value: sarama.StringEncoder("hello world"),
// 	}

// 	p, o, err := producer.SendMessage(&msg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("partition=%v, offset=%v", p, o)

// }
