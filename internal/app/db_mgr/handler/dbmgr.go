package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"topcoder.com/space-traffic-control/internal/app/db_mgr/models"
	"topcoder.com/space-traffic-control/internal/app/db_mgr/service"
	reqmodels "topcoder.com/space-traffic-control/internal/pkg/models"
)

type DBMgr struct {
	dbSvc service.ReadWrite
}

func NewDBHandler(dbSvc service.ReadWrite) DBMgr{
	return DBMgr{
		dbSvc: dbSvc,
	}
}

func (dbMgr *DBMgr) CreateUser(c *gin.Context) {
	input := reqmodels.User{}

	err := c.BindJSON(&input)
	if err != nil {
		logger.Errorf("error in request body %v. BAD REQUEST", err)
		c.JSON(http.StatusBadRequest, "err in request body")
		return
	}

	user := &models.User{
		ID:       uuid.UUID{},
		Username: input.Username,
		Password: input.Password,
		Role:     input.Role,
	}

	resp, err := dbMgr.dbSvc.CreateUser(user)
	if err != nil {
		logger.Errorf("error creating user %v", err)
		c.JSON(http.StatusInternalServerError, "error creating user")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (dbMgr *DBMgr) AuthenticateUser(c *gin.Context) {
	input := reqmodels.Credentials{}

	err := c.BindJSON(&input)
	if err != nil {
		logger.Errorf("error in request body %v. BAD REQUEST", err)
		c.JSON(http.StatusBadRequest, "err in request body")
		return
	}

	userDB, err := dbMgr.dbSvc.GetUser(input)
	if err != nil {
		logger.Errorf("error fetching user details : %v", err)
		c.JSON(http.StatusInternalServerError, "error fetching user details")
		return
	}

	if userDB.Username == "" {
		logger.Error("no user found for given username")
		c.JSON(http.StatusNotFound, "no user found for given username")
		return
	}

	user := reqmodels.User{
		UserID:   userDB.ID.String(),
		Username: userDB.Username,
		Password: userDB.Password,
		Role:     userDB.Role,
	}

	c.JSON(http.StatusOK, user)
}

func (dbMgr *DBMgr) CreateStation(c *gin.Context) {
	input := reqmodels.Station{}

	err := c.BindJSON(&input)
	if err != nil {
		logger.Errorf("error in request body %v. BAD REQUEST", err)
		c.JSON(http.StatusBadRequest, "err in request body")
		return
	}

	station := &models.Station{
		ID:       uuid.UUID{},
		Capacity: input.Capacity,
	}

	var docks []models.Dock

	for _, d := range input.Docks {
		dock := models.Dock{
			ID:              uuid.UUID{},
			NumDockingPorts: d.NumDockingPorts,
		}

		docks = append(docks, dock)
	}

	resp, err := dbMgr.dbSvc.CreateStation(station, docks)
	if err != nil {
		logger.Errorf("error creating todo %v", err)
		c.JSON(http.StatusInternalServerError, "error creating todo")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (dbMgr *DBMgr) GetAllStations(c *gin.Context) {

	stationsDB, err := dbMgr.dbSvc.GetAllStations()
	if err != nil {
		logger.Errorf("error fetching stations from DB %v", err)
		c.JSON(http.StatusInternalServerError, "error fetching stations from DB")
		return
	}

	docksDB, err := dbMgr.dbSvc.GetAllDocks()
	if err != nil {
		logger.Errorf("error fetching stations from DB %v", err)
		c.JSON(http.StatusInternalServerError, "error fetching stations from DB")
		return
	}
	var stations []reqmodels.Station
	for _, st := range stationsDB {
		var docks []reqmodels.Dock
		for _, d := range docksDB {
			if d.StationID == st.ID {
				dock := reqmodels.Dock{
					ID:              d.ID.String(),
					NumDockingPorts: d.NumDockingPorts,
					Occupied:        d.Occupied,
					Weight:          d.Weight,
				}
				docks = append(docks, dock)
			}
			continue
		}

		station := reqmodels.Station{
			ID:           st.ID.String(),
			Capacity:     st.Capacity,
			UsedCapacity: st.UsedCapacity,
			Docks:        docks,
		}

		stations = append(stations, station)
	}

	c.JSON(http.StatusOK, stations)
}

func (dbMgr *DBMgr) CreateShip(c *gin.Context) {
	input := reqmodels.Ship{}

	err := c.BindJSON(&input)
	if err != nil {
		logger.Errorf("error in request body %v. BAD REQUEST", err)
		c.JSON(http.StatusBadRequest, "err in request body")
		return
	}

	ship := &models.Ship{
		ID:     uuid.UUID{},
		Weight: input.Weight,
	}

	err = dbMgr.dbSvc.CreateShip(ship)
	if err != nil {
		logger.Errorf("error creating todo %v", err)
		c.JSON(http.StatusInternalServerError, "error creating todo")
		return
	}
	c.JSON(http.StatusOK, "")
}

func (dbMgr *DBMgr) GetAllShips(c *gin.Context) {

	shipsDB, err := dbMgr.dbSvc.GetAllShips()
	if err != nil {
		logger.Errorf("error creating todo %v", err)
		c.JSON(http.StatusInternalServerError, "error creating todo")
		return
	}

	var ships []reqmodels.Ship
	for _, sh := range shipsDB {
		ship := reqmodels.Ship{
			ID:     sh.ID.String(),
			Status: sh.Status,
			Weight: sh.Weight,
		}

		ships = append(ships, ship)
	}

	c.JSON(http.StatusOK, ships)
}