package usecase

import (
	"errors"
	"go-grow/model"
	"go-grow/repository"
	"go-grow/util"
)

type UserUsecase interface {
	PostRegisterUser(request *model.RegisterUserRequest) (*model.User, error)
	PostEmailAvailability(request *model.EmailAvailabilityRequest) (bool, error)
}

type userUsecase struct {
	repo repository.BaseRepository
}

func NewUserUsecase(repo repository.BaseRepository) *userUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) PostRegisterUser(request *model.RegisterUserRequest) (*model.User, error) {
	user := model.User{}
	user.Name = request.Name
	
	userInputtedEmail := request.Email
	checkEmailString, err := util.EmailStringRegex(userInputtedEmail)
	if !checkEmailString {
		return &user, err
	}

	checkEmailAvailable, err := u.repo.FindUserByEmail(userInputtedEmail)
	if err != nil {
		return &user, err
	}

	if checkEmailAvailable.ID != 0 {
		return &user, errors.New("user with this email is already exist")
	}

	user.Email = userInputtedEmail

	encrpytedPassword, err := util.AESMethodEncyrpt(request.Password)
	if err != nil {
		return nil, err
	}
	
	user.Password = encrpytedPassword
	user.Role = "user"
	user.IsCoolMember = false

	newUser, err := u.repo.CreateUserToDB(&user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (u *userUsecase) PostEmailAvailability(request *model.EmailAvailabilityRequest) (bool, error) {
	userEmail := request.Email
	
	checkEmailString, err := util.EmailStringRegex(userEmail)
	if !checkEmailString {
		return false, err
	}

	user, err := u.repo.FindUserByEmail(userEmail)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

