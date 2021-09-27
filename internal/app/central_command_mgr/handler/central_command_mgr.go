package handler

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"topcoder.com/space-traffic-control/internal/app/central_command_mgr/service"
	"topcoder.com/space-traffic-control/internal/pkg/models"
)

type CentralCommandMgr struct {
	svc service.CentralCommand
}

func NewCentralCommandMgr(svc service.CentralCommand) CentralCommandMgr{
	return CentralCommandMgr{
		svc: svc,
	}
}

func (ccm *CentralCommandMgr) Register(c *gin.Context) {
	input := models.Station{}

	err := c.BindJSON(&input)
	if err != nil {
		logger.Errorf("error in request body %v. BAD REQUEST", err)
		c.JSON(http.StatusBadRequest, "err in request body")
		return
	}

	resp, err := ccm.svc.RegisterStation(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (ccm *CentralCommandMgr) GetStations(c *gin.Context) {
	stations, err := ccm.svc.GetAllStations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, stations)
}