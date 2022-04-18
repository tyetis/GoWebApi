package services

import (
	"github.com/tyetis/goapiexample/app/data/entity"
	"gorm.io/gorm"
)

type IOrderService interface {
	Create()
	Update()
	Delete()
	FindOne(query *entity.Order) entity.Order
	GetAll(query *entity.Order) []entity.Order
}

type orderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) IOrderService {
	return &orderService{
		db: db,
	}
}

func (service *orderService) Create() {

}

func (service *orderService) Update() {

}

func (service *orderService) Delete() {

}

func (service *orderService) FindOne(query *entity.Order) entity.Order {
	var order entity.Order
	service.db.First(&order)
	return order
}

func (service *orderService) GetAll(query *entity.Order) []entity.Order {
	var orders []entity.Order
	service.db.Find(&orders)
	return orders
}
