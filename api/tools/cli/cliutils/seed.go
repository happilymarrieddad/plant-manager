package cliutils

import (
	"fmt"
	"plant-manager/internal/models"
	"plant-manager/internal/permissions"
	"plant-manager/internal/types"

	"encoding/json"
	"io/ioutil"

	"xorm.io/xorm"
)

type UserWrapper struct {
	*types.User
	IsCustomerAdmin bool `json:"isCustomerAdmin"`
}

type CustomerWrapper struct {
	*types.Customer
	Users []*UserWrapper `json:"users"`
}

func SeedData(db *xorm.Engine) (err error) {
	customersModel := models.NewCustomer(db)
	usersModel := models.NewUser(db)
	permissionModel := models.NewPermission(db)

	var customers []*CustomerWrapper
	file, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		return
	}
	if err = json.Unmarshal([]byte(file), &customers); err != nil {
		return
	}

	for _, c := range customers {
		cust := c.Customer
		fmt.Println("Attempting to create customer " + cust.Name)
		if err = customersModel.Create(cust); err != nil {
			return
		}

		for _, u := range c.Users {
			u.CustomerID = cust.ID
			u.Password, err = types.EncryptPassword(u.Password)
			if err != nil {
				return
			}

			newUser := u.User
			newUser.CustomerID = cust.ID

			fmt.Println("Attempting to create user " + newUser.Email)
			if err = usersModel.Create(newUser); err != nil {
				return
			}

			// Only if user is customer admin
			if u.IsCustomerAdmin {
				fmt.Printf("Assigning system user permission to user %s \n", newUser.Email)
				if err = permissionModel.Create(&types.Permission{
					Name:       string(permissions.PermissionSystemUser),
					UserID:     newUser.ID,
					CustomerID: cust.ID,
				}); err != nil {
					return
				}

				for _, p := range permissions.AllPermissions {
					fmt.Printf("Assigning permission %s to user %s \n", string(p), newUser.Email)
					if err = permissionModel.Create(&types.Permission{
						Name:       string(p),
						UserID:     newUser.ID,
						CustomerID: cust.ID,
					}); err != nil {
						return
					}
				}
			}
		}
	}

	return
}
