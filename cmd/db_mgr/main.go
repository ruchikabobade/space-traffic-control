package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"os"
	"topcoder.com/space-traffic-control/internal/app/db_mgr/handler"
	"topcoder.com/space-traffic-control/internal/app/db_mgr/service"
)

func main(){
	router := gin.New()
	router.HandleMethodNotAllowed = true

	db, err := service.SetConnection()
	if err != nil {
		logger.Errorf("database connection failed. error [%v]. shutting down application", err)
		os.Exit(1)
	}

	dbSvc := service.NewDatabaseReadWrite(db)
	dbHandler := handler.NewDBHandler(dbSvc)

	userRoutes := router.Group("/user")
	userRoutes.POST("/", dbHandler.CreateUser)
	userRoutes.POST("/authenticate", dbHandler.AuthenticateUser)

	stationRoutes := router.Group("/station")
	stationRoutes.POST("/", dbHandler.CreateStation)
	stationRoutes.GET("/", dbHandler.GetAllStations)

	shipRoutes := router.Group("/ship")
	shipRoutes.POST("/", dbHandler.CreateShip)
	shipRoutes.GET("/", dbHandler.GetAllShips)

	err = router.Run(":8081")
	if err != nil {
		logger.Errorf("error in starting db application %v", err)
		os.Exit(1)
	}
	logger.Infof("DB Server Started. Listening on port %v", "8081")
}

