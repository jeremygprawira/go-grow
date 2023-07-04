package usecase

import (
	"go-grow/model"
	"go-grow/repository"
)

type CoolUsecase interface {
	PostCreateCool(request *model.CreateCoolRequest) (*model.Cool, error)
}

type coolUsecase struct {
	repo repository.BaseRepository
}

func NewCoolUsecase(repo repository.BaseRepository) *coolUsecase {
	return &coolUsecase{repo}
}

func (u *coolUsecase) PostCreateCool(request *model.CreateCoolRequest) (*model.Cool, error) {
	var cool *model.Cool
	return cool, nil
}