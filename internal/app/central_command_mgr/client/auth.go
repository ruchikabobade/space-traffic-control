package client

import (
	"fmt"
	"net/http"
	"topcoder.com/space-traffic-control/internal/pkg/utils"
)

const (
	authPath = "/auth/authorize"
)

var (
	AuthServiceURL = utils.GetEnvOrDefault("AUTH_SERVICE_URL", "")
)

type AuthServiceClient interface {

}


type AuthService struct {
	client *http.Client
}

func NewAuthServiceClient(client *http.Client) AuthService {
	return AuthService{
		client: client,
	}
}

func(auth *AuthService) Authorize() error{
	url := fmt.Sprintf("%s%s", DBServiceURL, authPath)
	headers := map[string][]string{}
	headers["Authorization"] = []string{}


	headersNorm := utils.GetNormalizedValues(headers)
	_, err := utils.ExecuteHttpRequest(auth.client, url, nil, http.MethodPost, headersNorm, nil)
	if err != nil {

	}

	if err != nil {

	}

	return nil
}