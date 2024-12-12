package handler

import (
	"20241212/class/2/domain"
	"20241212/class/2/helper"
	"20241212/class/2/service"
	"github.com/gin-gonic/gin"
	"github.com/mailersend/mailersend-go"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type UserController struct {
	service service.Service
	logger  *zap.Logger
}

func NewUserController(service service.Service, logger *zap.Logger) *UserController {
	return &UserController{service, logger}
}

// Check Email endpoint
// @Summary Check Email
// @Description email must be valid when users want to reset their passwords
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} handler.Response "email is valid"
// @Failure 404 {object} handler.Response "user not found"
// @Router  /users [get]
func (ctrl *UserController) All(c *gin.Context) {
	searchParam := domain.User{Email: c.Query("email")}

	if searchParam.Email == "" {
		BadResponse(c, "invalid parameter", http.StatusBadRequest)
		return
	}

	if _, err := ctrl.service.User.All(searchParam); err != nil {
		BadResponse(c, err.Error(), http.StatusNotFound)
		return
	}

	GoodResponseWithData(c, "email is valid", http.StatusOK, nil)
}

// Registration endpoint
// @Summary Staff Registration
// @Description register staff
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param domain.User body domain.User true " "
// @Success 200 {object} handler.Response "login successfully"
// @Failure 500 {object} handler.Response "server error"
// @Router  /register [post]
func (ctrl *UserController) Registration(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		BadResponse(c, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := ctrl.service.User.Register(&user); err != nil {
		BadResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	to := mailersend.Recipient{Name: user.FullName, Email: user.Email}
	data := struct {
		OTP string
	}{
		OTP: helper.GenerateOTP(),
	}
	messageId, err := ctrl.service.Email.Send(to, "Welcome", "registration", data)
	if err != nil {
		log.Println(err)
	}
	log.Println("messageId:", messageId)
	GoodResponseWithData(c, "user registered", http.StatusCreated, user)
}
