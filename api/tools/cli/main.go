package main

import (
	"fmt"
	"plant-manager/tools/cli/cliutils"

	_ "github.com/lib/pq" //postgres driver for sqlx
	"xorm.io/xorm"
)

func truncateTable(db *xorm.Engine, name string) error {
	fmt.Println("Truncating table " + name)
	_, err := db.Exec(name)
	return err
}

func main() {

	db, err := xorm.NewEngine("postgres", "postgres://postgres:postgres@localhost:4432/plant-manager?connect_timeout=180&sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = truncateTable(db, "TRUNCATE permissions"); err != nil {
		panic(err)
	}
	if err = truncateTable(db, "TRUNCATE users"); err != nil {
		panic(err)
	}
	if err = truncateTable(db, "TRUNCATE customers"); err != nil {
		panic(err)
	}

	if err = cliutils.SeedData(db); err != nil {
		panic(err)
	}

	fmt.Println("Success!")
}
