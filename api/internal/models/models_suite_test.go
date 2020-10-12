package models_test

import (
	"fmt"
	"testing"

	_ "github.com/lib/pq" //postgres driver for sqlx
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

var (
	db *xorm.Engine
)

func cleanMainDatabaseTables() {
	fmt.Println("cleanMainDatabaseTables")
	_, err := db.Exec("TRUNCATE place_slots CASCADE")
	Expect(err).To(BeNil())
	_, err = db.Exec("TRUNCATE places CASCADE")
	Expect(err).To(BeNil())
	_, err = db.Exec("TRUNCATE customers CASCADE")
	Expect(err).To(BeNil())
}

var _ = BeforeSuite(func() {
	defer GinkgoRecover()
	var err error
	db, err = xorm.NewEngine("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?connect_timeout=180&sslmode=disable",
		"postgres",
		"postgres",
		"localhost",
		4432,
		"plant-manager-test",
	))
	Expect(err).To(BeNil())
	db.SetLogLevel(log.LOG_DEBUG)
	db.ShowSQL(true)

	if db != nil {
		cleanMainDatabaseTables()
	}
})

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}
