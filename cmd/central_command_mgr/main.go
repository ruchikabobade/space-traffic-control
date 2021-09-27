package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
	client3 "topcoder.com/space-traffic-control/internal/app/central_command_mgr/client"
	"topcoder.com/space-traffic-control/internal/app/central_command_mgr/handler"
	"topcoder.com/space-traffic-control/internal/app/central_command_mgr/service"
)

func main() {
	router := gin.New()
	router.HandleMethodNotAllowed = true

	client := &http.Client{}
	dbServiceClient := client3.NewDBServiceClient(client)
	authClient := client3.NewAuthServiceClient(client)
	svc := service.NewCentralCommand(dbServiceClient, authClient)
	handler := handler.NewCentralCommandMgr(svc)

	stationRoutes := router.Group("/centcom")
	stationRoutes.POST("/station/register", handler.Register)
	stationRoutes.GET("/station/all", handler.GetStations)

	err := router.Run(":8083")
	if err != nil {
		logger.Errorf("error in starting central command application %v", err)
		os.Exit(1)
	}
	logger.Infof("Central Command Server Started. Listening on port %v", "8082")
}
