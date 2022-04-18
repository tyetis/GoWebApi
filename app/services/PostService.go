package services

import "github.com/tyetis/goapiexample/app/data/entity"

type IPostService interface {
	Create()
	Update()
	Delete()
	FindOne() []entity.Post
}

type postService struct {
}

var PostService IPostService = &postService{}

func (service *postService) Create() {

}

func (service *postService) Update() {

}

func (service *postService) Delete() {

}

func (service *postService) FindOne() []entity.Post {
	return []entity.Post{
		{Id: 1, Name: "new post", CreateDate: ""},
	}
}
