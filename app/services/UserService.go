package services

import (
	"github.com/tyetis/goapiexample/app/data/entity"
	"github.com/tyetis/goapiexample/app/models"
	"gorm.io/gorm"
)

type IUserService interface {
	Create()
	Update()
	Delete()
	FindOne(query *entity.User) entity.User
	GetAll(filter *entity.User, perPage int, page int) models.PaginationResult
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) IUserService {
	return &userService{
		db: db,
	}
}

func (service *userService) Create() {

}

func (service *userService) Update() {

}

func (service *userService) Delete() {

}

func (service *userService) FindOne(query *entity.User) entity.User {
	return service.GetAll(query, 1, 0).Data.([]entity.User)[0]
}

func (service *userService) GetAll(filter *entity.User, perPage int, page int) models.PaginationResult {
	var users []entity.User
	query := service.db
	if filter.Id != 0 {
		query = query.Where("id = ?", filter.Id)
	}
	if filter.Name != "" {
		query = query.Where("name like ?", "%"+filter.Name+"%")
	}
	if filter.Email != "" {
		query = query.Where("email like ?", "%"+filter.Email+"%")
	}
	result := query.Offset(perPage * page).Limit(perPage).Find(&users)
	return models.PaginationResult{
		Data:        users,
		Total:       int(result.RowsAffected),
		PerPage:     perPage,
		CurrentPage: page,
	}
}
