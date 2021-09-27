package service

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	logger "github.com/sirupsen/logrus"
	"topcoder.com/space-traffic-control/internal/app/db_mgr/models"
	respmodels "topcoder.com/space-traffic-control/internal/pkg/models"
)

type DB interface {
	CreateUser(user *models.User) error
	GetUser(creds respmodels.Credentials)(*models.User, error)
	CreateStation(station *models.Station, dock *models.Dock)(string, string, error)
	GetAllStations()([]models.Station, error)
	GetAllDocks()([]models.Dock, error)
	RegisterShip(ship *models.Ship)error
	GetAllShips()([]models.Ship, error)
}



func SetConnection()(*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "password",
		Database: "postgres",
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return db, nil
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.User)(nil),
		(*models.Station)(nil),
		(*models.Dock)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

type ReadWrite struct {
	db *pg.DB
}

func NewDatabaseReadWrite(db *pg.DB) ReadWrite {
	return ReadWrite{db: db}
}

func (rw *ReadWrite) CreateUser(user *models.User) (respmodels.GenericResponse, error) {
	resp := respmodels.GenericResponse{}
	_, err := rw.db.Model(user).Insert()
	if err != nil {
		logger.Errorf("error creating user: %v", err)
		return resp, err
	}

	resp.ID = user.ID.String()
	return resp, err
}

func (rw *ReadWrite) GetUser(creds respmodels.Credentials)(*models.User, error) {
	u := new(models.User)
	err := rw.db.Model(u).
		Where("username = ?", creds.Username).
		Where("password = ?", creds.Password).Select()

	if err != nil {
		logger.Errorf("error fetching user details: %v", err)
		return u, err
	}

	return u, nil
}

func (rw *ReadWrite) CreateStation(station *models.Station, docks []models.Dock)(respmodels.Station, error) {
	var resp respmodels.Station
	_, err := rw.db.Model(station).Insert()
	if err != nil {
		logger.Errorf("error registering station: %v", err)
		return resp, err
	}

	var docksList []respmodels.Dock
	for _, dock := range docks {
		dock.StationID = station.ID
		_, err = rw.db.Model(&dock).Insert()
		if err != nil {
			logger.Errorf("error inserting dock info: %v", err)
			return resp, err
		}

		dockResp := respmodels.Dock{
			ID:              dock.ID.String(),
			NumDockingPorts: dock.NumDockingPorts,
		}

		docksList = append(docksList, dockResp)
	}

	resp.ID = station.ID.String()
	resp.Capacity = station.Capacity
	resp.Docks = docksList

	return resp, nil
}

func (rw *ReadWrite) GetAllStations()([]models.Station, error){
	var stations []models.Station
	err := rw.db.Model(&stations).Select()
	if err != nil {
		logger.Errorf("error fetching list of stations: %v", err)
		return stations, err
	}
	return stations, nil
}

func (rw *ReadWrite) GetAllDocks()([]models.Dock, error){
	var docks []models.Dock
	err := rw.db.Model(&docks).Select()
	if err != nil {
		logger.Errorf("error fetching list of docks: %v", err)
		return docks, err
	}
	return docks, nil
}

func (rw *ReadWrite) CreateShip(ship *models.Ship)error{
	_, err := rw.db.Model(ship).Insert()
	if err != nil {
		logger.Errorf("error inserting ship info: %v", err)
		return  err
	}
	return nil
}

func (rw *ReadWrite) GetAllShips()([]models.Ship, error) {
	var ships []models.Ship
	err := rw.db.Model(&ships).Select()
	if err != nil {
		logger.Errorf("error fetching list of ship: %v", err)
		return ships, err
	}
	return ships, nil
}