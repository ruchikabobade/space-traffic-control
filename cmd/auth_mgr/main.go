package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	client3 "topcoder.com/space-traffic-control/internal/app/auth_mgr/client"
	"topcoder.com/space-traffic-control/internal/app/auth_mgr/handler"
	"topcoder.com/space-traffic-control/internal/app/auth_mgr/service"
)

func main() {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	client := &http.Client{}
	dbServiceClient := client3.NewDBServiceClient(client)
	authSvc := service.NewAuthSvc(dbServiceClient)
	authHandler := handler.NewAuthMgr(authSvc)

	userRoutes := router.Group("/user")
	userRoutes.POST("/signup", authHandler.SignUp)

	authRoutes := router.Group("/auth")
	authRoutes.POST("/login", authHandler.Token)


	err := router.Run(":8082")
	if err != nil {
		logger.Errorf("error in starting auth application %v", err)
		os.Exit(1)
	}
	logger.Infof("Auth Server Started. Listening on port %v", "8082")
}
