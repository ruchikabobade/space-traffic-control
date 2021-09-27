package client

import "topcoder.com/space-traffic-control/internal/pkg/client"

type DBServiceClient interface {
	RequestLanding()
	Land()
}

type DBService struct {
	client client.DatabaseClient
}


func (dbSvc *DBService) RequestLanding(){

}


func (dbSvc *DBService) Land(){

}