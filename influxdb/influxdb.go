package influxdb

import (
	"fmt"
	"log"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/pcpratheesh/golang-influxdb-example/config"
)

type Item struct {
	Name  string
	Value int
	time  time.Time
}
type Instance struct {
	Host, DB, User, Password, Port, Token string
	Client                                client.Client
}

func NewInfluxDBInstance(conf config.InfluxInstance) *Instance {
	return &Instance{
		Host:     conf.Host,
		Port:     conf.Port,
		DB:       conf.DB,
		Password: conf.Password,
		User:     conf.User,
		Token:    conf.Token,
	}
}

/**
 * Initiate connection
 */
func (i *Instance) Connect() error {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     fmt.Sprintf("http://%s:%s", i.Host, i.Port),
		Username: i.User,
		Password: i.Password,
	})
	if err != nil {
		return err
	}
	_, _, err = c.Ping(10)
	if err != nil {
		return err
	}

	i.Client = c

	return nil
}

/**
 * Close the client connection
 */
func (i *Instance) Close() {
	i.Client.Close()
}

// Create influx database
func (i *Instance) Create() error {
	// Workaround, since daocloud influxdb haven't privision an instance
	// create the db instance here
	q := client.Query{
		Command:  fmt.Sprintf("create database %s", i.DB),
		Database: i.DB,
	}

	// ignore the error of existing database
	_, err := i.Client.Query(q)
	return err
}

// insert sample data
func (i *Instance) InsertSample() error {
	// Create a new point batch
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  i.DB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}

	item := map[string]interface{}{
		"name":  "Item name 2 ",
		"value": 10,
	}
	// Create a point and add to batch
	tags := map[string]string{}
	fields := item

	pt, err := client.NewPoint("items", tags, fields, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	// Write the batch
	if err := i.Client.Write(bp); err != nil {
		return err
	}

	return nil
}

func (i *Instance) GetAllItems() (res []client.Result, err error) {
	q := client.Query{
		Command:  "select * from items;",
		Database: i.DB,
	}

	if response, err := i.Client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
