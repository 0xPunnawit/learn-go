package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/api", api)

	app.Listen(":8001")

}

func init() {
	hystrix.ConfigureCommand("api", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 1,
		ErrorPercentThreshold:  100,
		SleepWindow:            15000,
	})

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	go func() {
		fmt.Println("Starting Hystrix stream on :8002 ...")
		if err := http.ListenAndServe(":8002", hystrixStreamHandler); err != nil {
			fmt.Println("Hystrix stream error:", err)
		}
	}()
}

func api(c *fiber.Ctx) error {

	hystrix.Go("api", func() error {

		res, err := http.Get("http://localhost:8000/api")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		msg := string(data)
		fmt.Println(msg)

		return nil
	}, func(err error) error {
		fmt.Println(err)
		return nil
	})

	return nil
}
