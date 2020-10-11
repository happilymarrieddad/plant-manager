// Package types defines all of the data entities this API works with.
// These may be used within repos, within API routes, etc...
package types

import (
	validator "gopkg.in/go-playground/validator.v9"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

var (
	validCop *validator.Validate = validator.New()
	logger                       = log15.New(log15.Ctx{"pkg": "types"})
)

// ValidatingEntity is an entity that can check itself for general validity errors
type ValidatingEntity interface {
	Validate() error
}

func Validate(entity interface{}) error {
	ve, ok := entity.(ValidatingEntity)
	if ok {
		return ve.Validate()
	}
	return validCop.Struct(entity)
}
