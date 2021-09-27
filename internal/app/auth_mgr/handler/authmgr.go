package handler

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"topcoder.com/space-traffic-control/internal/app/auth_mgr/service"
	"topcoder.com/space-traffic-control/internal/pkg/models"
)

type AuthMgr struct {
	authSvc service.Authorization
}

func NewAuthMgr(authSvc service.Authorization) AuthMgr{
	return AuthMgr{authSvc: authSvc}
}

func (authMgr *AuthMgr) Token(c *gin.Context) {
	input := models.Credentials{}

	err := c.BindJSON(&input)
	if err != nil {
		logger.Errorf("error in request body %v. BAD REQUEST", err)
		c.JSON(http.StatusBadRequest, "err in request body")
		return
	}

	token, err := authMgr.authSvc.GenerateToken(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, token)
}

func (authMgr *AuthMgr) SignUp(c *gin.Context) {
	input := models.User{}

	err := c.BindJSON(&input)
	if err != nil {
		logger.Errorf("error in request body %v. BAD REQUEST", err)
		c.JSON(http.StatusBadRequest, "err in request body")
		return
	}

	err = authMgr.authSvc.SignUp(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, "")
}


