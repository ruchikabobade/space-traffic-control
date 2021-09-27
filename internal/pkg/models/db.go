package models

import "io"

type RequestParams struct {
	Headers map[string][]string
	Query map[string][]string
	Body io.Reader
	Path string
}
