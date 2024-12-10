package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	service userService
}

func NewUserHandler(service userService) *userHandler {
	return &userHandler{service: service}
}

func (u *userHandler) RegisterHandler(c *gin.Context) {
	//  nhan user tu gia tri cua
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = u.service.registerService(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (u *userHandler) LoginHandler(c *gin.Context) {
	var login Login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	info, token, err := u.service.Login(c, login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": info, "token": token})
}
func (u *userHandler) GetUseByIdHandler(c *gin.Context) {
	//  lay id tu token wwa va chuyen sang uuid de su ly
	uuid := c.MustGet("uuid")

	user, err := u.service.GetUserByIdService(c, uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
func (u *userHandler) UpdateByIdHandler(c *gin.Context) {
	var user UpdateUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuid := c.MustGet("uuid")
	err = u.service.UpdateUserByIdService(c, uuid, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "status": "success"})
}
func (u *userHandler) ChangePasswordHandler(c *gin.Context) {
	var password ChangePassword
	err := c.ShouldBindJSON(&password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uuid := c.MustGet("uuid")
	err = u.service.ChangePassword(c, uuid, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
