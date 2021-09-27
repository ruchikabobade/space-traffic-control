package service

import (
	logger "github.com/sirupsen/logrus"
	"topcoder.com/space-traffic-control/internal/app/central_command_mgr/client"
	"topcoder.com/space-traffic-control/internal/pkg/models"
)

type CentralCommandSvc interface {

}

type CentralCommand struct {
	dbClient client.DBService
	authClient client.AuthServiceClient
}

func NewCentralCommand(serviceClient client.DBService, authClient client.AuthServiceClient) CentralCommand{
	return CentralCommand{
		dbClient: serviceClient,
		authClient: authClient,
	}
}

func (cen *CentralCommand) RegisterStation(station models.Station) (models.Station, error) {
	resp, err := cen.dbClient.CreateStation(station)
	if err != nil {
		logger.Errorf("error registering station with command center")
		return resp, err
	}
	return resp, nil
}

func (cen *CentralCommand) GetAllStations()([]models.Station, error) {
	stations, err := cen.dbClient.GetAllStation()
	if err != nil {

	}
	return stations, nil
}

func (cen *CentralCommand) RegisterShip(ship models.Ship) error {
	err := cen.dbClient.CreateShip(ship)
	if err != nil {

	}
	return nil
}

func (cen *CentralCommand) GetAllShips()([]models.Ship, error) {
	ships, err := cen.dbClient.GetAllShip()
	if err != nil {

	}
	return ships, nil
}