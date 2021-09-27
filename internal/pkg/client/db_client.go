package client

import (
	"net/http"
	"topcoder.com/space-traffic-control/internal/pkg/models"
	"topcoder.com/space-traffic-control/internal/pkg/utils"
)

type DatabaseClient interface {
	ExecuteDatabaseOperation(params models.RequestParams, method string,URL string, v interface{})(*http.Response, error)
}

type DatabaseService struct {
	client *http.Client
}

func NewDatabaseClient(client *http.Client) DatabaseClient{
	return &DatabaseService{client: client}
}

func (db *DatabaseService) ExecuteDatabaseOperation(params models.RequestParams, method string, URL string, v interface{})(*http.Response, error){
	headersNorm := utils.GetNormalizedValues(params.Headers)
	QueryNorm := utils.GetNormalizedValues(params.Query)
	response, err := utils.ExecuteHttpRequest(db.client, URL, QueryNorm, method, headersNorm, params.Body)
	if err != nil {

	}
	return response, nil
}