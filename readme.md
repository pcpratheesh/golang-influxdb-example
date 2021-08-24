# golang-influx-example
This is a sample respository to leanrn how we can perform influx db operations with an api endpoint.

To get all items from db
```golang
    app.Get("/get", func(c *fiber.Ctx) error {
		data, err := influxInstance.GetAllItems()
		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(err.Error())
		}
		return c.JSON(data)
	})
```

To create new entry to the db
```golang
    app.Post("/create", func(c *fiber.Ctx) error {
		err := influxInstance.InsertSample()
		if err != nil {
			return c.Status(fiber.ErrBadGateway.Code).JSON(err.Error())
		}
		return c.JSON("successfully inserted")
	})
```

### Run
```golang
    go run main.go
```


## References
- InfluxDB : https://docs.influxdata.com/influxdb/v1.8/introduction/get-started/
- InfluxDB Golang Client : https://github.com/influxdata/influxdb-client-go
- Go Fiber : https://github.com/gofiber/fiber