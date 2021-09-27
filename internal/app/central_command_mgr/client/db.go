package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	models "topcoder.com/space-traffic-control/internal/pkg/models"
	"topcoder.com/space-traffic-control/internal/pkg/utils"
)

const (
	stationPath = "/station"
	shipPath = "/ship"
)

var (
	DBServiceURL = utils.GetEnvOrDefault("DB_SERVICE_URL", "")
)


type DBServiceClient interface {
	CreateStation(station models.Station) error
	GetAllStation()([]models.Station, error)
	CreateShip(ship models.Ship) error
	GetAllShip()([]models.Ship, error)
}

type DBService struct {
	client *http.Client
}

func NewDBServiceClient(client *http.Client) DBService {
	return DBService{
		client: client,
	}
}


func(dbSvc *DBService) CreateStation(station models.Station)(models.Station, error){
	var stationResp models.Station
	url := fmt.Sprintf("%s%s", DBServiceURL, stationPath)
	headers := map[string][]string{}
	headers["Authorization"] = []string{}

	stationReqBodyBytes, err := json.Marshal(station)
	if err != nil {

	}

	body := ioutil.NopCloser(bytes.NewBuffer(stationReqBodyBytes))

	headersNorm := utils.GetNormalizedValues(headers)
	resp, err := utils.ExecuteHttpRequest(dbSvc.client, url, nil, http.MethodPost, headersNorm,body)
	if err != nil {

	}
	err = json.NewDecoder(resp.Body).Decode(&stationResp)
	if err != nil {

	}

	return stationResp, nil
}

func(dbSvc *DBService) GetAllStation()([]models.Station, error){
	var stations []models.Station
	url := fmt.Sprintf("%s%s", DBServiceURL, stationPath)
	headers := map[string][]string{}
	headers["Authorization"] = []string{}

	headersNorm := utils.GetNormalizedValues(headers)
	resp, err := utils.ExecuteHttpRequest(dbSvc.client, url, nil, http.MethodPost, headersNorm,nil)
	if err != nil {

	}

	err = json.NewDecoder(resp.Body).Decode(&stations)
	if err != nil {

	}

	return stations, nil
}


func(dbSvc *DBService) CreateShip(ship models.Ship) error{
	url := fmt.Sprintf("%s%s", DBServiceURL, shipPath)
	headers := map[string][]string{}
	headers["Authorization"] = []string{}

	stationReqBodyBytes, err := json.Marshal(ship)
	if err != nil {

	}

	body := ioutil.NopCloser(bytes.NewBuffer(stationReqBodyBytes))

	headersNorm := utils.GetNormalizedValues(headers)
	_, err = utils.ExecuteHttpRequest(dbSvc.client, url, nil, http.MethodGet, headersNorm,body)
	if err != nil {

	}

	if err != nil {

	}

	return nil
}

func(dbSvc *DBService) GetAllShip()([]models.Ship, error){
	var stations []models.Ship
	url := fmt.Sprintf("%s%s", DBServiceURL, shipPath)
	headers := map[string][]string{}
	headers["Authorization"] = []string{}

	headersNorm := utils.GetNormalizedValues(headers)
	resp, err := utils.ExecuteHttpRequest(dbSvc.client, url, nil, http.MethodGet, headersNorm,nil)
	if err != nil {

	}

	err = json.NewDecoder(resp.Body).Decode(&stations)
	if err != nil {

	}

	return stations, nil
}

