package main

import (
	"fmt"
	"plant-manager/internal/api"

	_ "github.com/lib/pq" //postgres driver for sqlx
	"xorm.io/xorm"
)

func main() {
	db, err := xorm.NewEngine("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?connect_timeout=180&sslmode=disable",
		"postgres",
		"postgres",
		"localhost",
		4432,
		"plant-manager",
	))
	if err != nil {
		panic(err)
	}
	db.ShowSQL(true)

	api.Run(db, 50051)
}
