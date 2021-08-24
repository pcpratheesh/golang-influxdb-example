package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pcpratheesh/golang-influxdb-example/config"
	"github.com/pcpratheesh/golang-influxdb-example/influxdb"
)

func main() {

	cfg, err := config.LoadConfiguration()
	if err != nil {
		panic(err)
	}

	// initiate the db instance and connect
	influxInstance := influxdb.NewInfluxDBInstance(cfg.InfluxInstance)
	err = influxInstance.Connect()
	if err != nil {
		panic(err)
	}
	defer influxInstance.Close()
	log.Println("Successfully connect to influxdb ...")

	// create database if not exists
	influxInstance.Create()

	app := fiber.New()

	// get data
	app.Get("/get", func(c *fiber.Ctx) error {
		data, err := influxInstance.GetAllItems()
		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(err.Error())
		}
		return c.JSON(data)
	})

	// create new data
	app.Post("/create", func(c *fiber.Ctx) error {
		err := influxInstance.InsertSample()
		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(err.Error())
		}
		return c.JSON("successfully inserted")
	})

	log.Println("Start listening...")
	if err := app.Listen(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
