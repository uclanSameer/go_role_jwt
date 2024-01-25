package config

import (
	"backend_01/ent"
	"context"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var client *ent.Client

func InitDataSource() *ent.Client {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	c, err := ent.Open("postgres", "host=localhost port=5432 user="+username+" dbname=test_01 password="+password+" sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := c.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if client == nil {
		client = c
	}
	return c
}

func SingletonClient() *ent.Client {
	if client != nil {
		return client
	} else {
		c := InitDataSource()
		client = c
		return c
	}
}

func UserClient() *ent.UserClient {
	return SingletonClient().User
}
