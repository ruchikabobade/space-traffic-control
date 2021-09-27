package transport

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

type LoggerTransport struct {
	tr http.RoundTripper
}

func NewLoggerTransport(tr http.RoundTripper) http.RoundTripper {
	return &LoggerTransport{tr: tr}
}

func (trans *LoggerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Infof("method: [%s], path:[%s]", req.Method, req.URL.String())
	resp, err := trans.tr.RoundTrip(req)
	if resp != nil {
		if resp.StatusCode <399 {
			log.Infof("method: [%s], path: [%s], response status: [%d]", req.Method, req.URL.String(), resp.StatusCode)
		} else {
			log.Errorf("method: [%s], path: [%s], response status: [%d]", req.Method, req.URL.String(), resp.StatusCode)
		}
	}
	return resp, err
}