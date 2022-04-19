package services

import (
	"time"

	"github.com/tyetis/goapiexample/app/data/entity"
	"github.com/tyetis/goapiexample/app/models"
	"gorm.io/gorm"
)

type IUserService interface {
	Create(user *entity.User) models.BaseResult
	Update(user *entity.User) models.BaseResult
	Delete(id int) models.BaseResult
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

func (service *userService) Create(model *entity.User) models.BaseResult {
	var newuser = entity.User{
		Name:       model.Name,
		Email:      model.Email,
		Createdate: time.Now(),
	}
	result := service.db.Create(&newuser)
	return ReturnResult(result)
}

func (service *userService) Update(user *entity.User) models.BaseResult {
	result := service.db.Model(&entity.User{Id: user.Id}).Updates(&entity.User{
		Name:  user.Name,
		Email: user.Email,
	})
	return ReturnResult(result)
}

func (service *userService) Delete(id int) models.BaseResult {
	result := service.db.Delete(&entity.User{Id: id})
	return ReturnResult(result)
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

func ReturnResult(result *gorm.DB) models.BaseResult {
	if result.Error == nil {
		if result.RowsAffected == 0 {
			return models.BaseResult{Result: false, Message: "User not found"}
		}
		return models.BaseResult{Result: true}
	}
	return models.BaseResult{
		Result:  false,
		Message: result.Error.Error(),
	}
}
