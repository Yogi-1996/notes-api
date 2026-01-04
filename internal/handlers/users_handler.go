package handlers

import (
	"net/http"

	"github.com/Yogi-1996/notes-backend/internal/models"
	"github.com/Yogi-1996/notes-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserServiceInterface
}

func NewUserHandler(u services.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: u,
	}
}

func (u *UserHandler) UserRegister(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	addeduser, err := u.service.AddUser(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user registered sucessfull",
		"data": gin.H{
			"Email":      addeduser.Email,
			"created at": addeduser.CreatedAt,
		},
	})

}

func (u *UserHandler) UserLogin(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := u.service.VerifyUser(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user Login sucessfull",
		"token":   token,
	})

}
