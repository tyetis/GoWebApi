package entity

import (
	"time"
)

type User struct {
	Id         int
	Name       string `validate:"required,min=3,max=32"`
	Email      string
	Password   string
	Createdate time.Time
	Orders     []Order
}
