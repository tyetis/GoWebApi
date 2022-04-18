package entity

type Order struct {
	Id        int
	UserId    int `gorm:"column:UserId"`
	ProductId int `gorm:"column:ProductId"`
	User      User
	Product   Product
}
