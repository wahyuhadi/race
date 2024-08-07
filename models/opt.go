package models

import "net/http"

type Opt struct {
	File         string
	Mutation     int
	Jsonb        bool
	IgnoreFields string
	Proxy        string
	Version      string
	URL          string
	Round        int
	TotalReq     int
	Duration     int
}

type Connection struct {
	Request  *http.Request
	Response *http.Response
}
