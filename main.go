package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Can not load environment variables", err))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		c.Response().Header.Set("Content-Type", "application/json")
		// c.Request().Header.Add("Content-Type", "application/json")
		return c.JSON(fiber.Map{
			"ok":      true,
			"message": "get you msg",
		})
	})

	app.Use(compress.New())

	go func() {
		if err := app.Listen(port); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // > Block the main thread until interrupted
	app.Shutdown()
	fmt.Println("💀 > Server shutting down...")
}
