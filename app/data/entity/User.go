package entity

import (
	"time"
)

type User struct {
	Id         int
	Name       string
	Email      string
	Password   string
	CreateDate time.Time
	Orders     []Order
}
