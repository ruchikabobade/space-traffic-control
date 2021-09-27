package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	models2 "topcoder.com/space-traffic-control/internal/pkg/models"
	"topcoder.com/space-traffic-control/internal/pkg/utils"
)

const (
	userPath = "/user"
)
var (
	DBServiceURL = utils.GetEnvOrDefault("DB_SERVICE_URL", "")
)

type DBServiceClient interface {
	CreateUser(user models2.User) error
	Login(ctx context.Context)(models2.User, error)
}

type DBService struct {
	client *http.Client
}

func NewDBServiceClient(client *http.Client) DBService {
	return DBService{
		client: client,
	}
}

func(dbSvc *DBService) CreateUser(user models2.User) error{
	url := fmt.Sprintf("%s%s", DBServiceURL, userPath)
	headers := map[string][]string{}
	headers["Authorization"] = []string{}

	userReqBodyBytes, err := json.Marshal(user)
	if err != nil {

	}

	body := ioutil.NopCloser(bytes.NewBuffer(userReqBodyBytes))

	headersNorm := utils.GetNormalizedValues(headers)
	_, err = utils.ExecuteHttpRequest(dbSvc.client, url, nil, http.MethodPost, headersNorm,body)
	if err != nil {

	}

	if err != nil {

	}

	return nil
}

func(dbSvc *DBService) Login(ctx context.Context)(models2.User, error){
	var user models2.User
	url := fmt.Sprintf("%s%s", DBServiceURL, userPath)
	headers := map[string][]string{}
	headers["Authorization"] = []string{}

	headersNorm := utils.GetNormalizedValues(headers)
	resp, err := utils.ExecuteHttpRequest(dbSvc.client, url, nil, http.MethodPost, headersNorm,nil)
	if err != nil {

	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {

	}

	return user, nil
}



