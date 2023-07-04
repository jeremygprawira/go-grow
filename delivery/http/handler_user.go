package http

import (
	"go-grow/model"
	"go-grow/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler (userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) RegisterUser(ctx *gin.Context) {
	var userRequest model.RegisterUserRequest

	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42201",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	user, err := h.userUsecase.PostRegisterUser(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40001",
			"responseMessage": "Usecase PostRegisterUser is not working properly: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "User has been registered successfully",
		"userName": user.Name,
		"userEmail": user.Email,
		"userEncryptedPassword": user.Password,
	})
}

func (h *UserHandler) EmailAvailability(ctx *gin.Context) {
	var userRequest model.EmailAvailabilityRequest

	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"responseCode": "42202",
			"responseMessage": "The required field on the body request is empty or invalid.",
		})
		return
	}

	isAvailable, err := h.userUsecase.PostEmailAvailability(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"responseCode": "40002",
			"responseMessage": "Usecase PostRegisterUser is not working properly: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseCode": "20000",
		"responseMessage": "Request has been successfully proceed",
		"isEmailAvailable": isAvailable,
	})
}