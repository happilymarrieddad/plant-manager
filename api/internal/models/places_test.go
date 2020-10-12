package models_test

import (
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "plant-manager/internal/models"
	"plant-manager/internal/types"
)

var _ = Describe("PlacesModel", func() {

	var placesModel Places
	var pl *types.Place
	var pl2 *types.Place
	var cust *types.Customer

	setup := func() {
		custModel := NewCustomer(db)

		cust = &types.Customer{
			Name: "Some Customer",
		}
		err := custModel.Create(cust)
		Expect(err).To(BeNil())

		placesModel = NewPlace(db)

		pl = &types.Place{
			Name:       "Greenhouse 1",
			Rows:       4,
			Columns:    4,
			CustomerID: cust.ID,
		}

		pl2 = &types.Place{
			Name:       "Greenhouse 2",
			Rows:       4,
			Columns:    5,
			CustomerID: cust.ID,
		}
	}

	create := func() {
		err := placesModel.Create(pl)
		Expect(err).To(BeNil())
		err = placesModel.Create(pl2)
		Expect(err).To(BeNil())
	}

	BeforeEach(func() {
		cleanMainDatabaseTables()
		setup()
	})

	Context("Create", func() {
		It("should fail to create a place when an invalid place is passed in", func() {
			err := placesModel.Create(nil)
			Expect(err).NotTo(BeNil())
		})

		It("should fail to create a place when an invalid place is passed in 2", func() {
			err := placesModel.Create(&types.Place{})
			Expect(err).NotTo(BeNil())
		})

		It("should fail to create a place because the customer_id does NOT exist", func() {
			pl.CustomerID = 2
			err := placesModel.Create(pl)
			Expect(err).NotTo(BeNil())
		})

		It("should successfully create a place", create)
	})

	Context("Find", func() {
		BeforeEach(create)

		It("should find place", func() {
			ps, err := placesModel.Find(cust.ID, 100, 0)
			Expect(err).To(BeNil())
			Expect(ps).NotTo(BeEmpty())
			Expect(len(ps)).To(Equal(2))
			spew.Dump(ps)
		})
	})

})
