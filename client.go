package grequests

import (
	"net/http"
	"time"
)

// Client requests client
type Client struct {
	Request  *http.Request
	Response *http.Response
}

// Request struct holds request
type Request struct {
	Request *http.Request
}

// Response struct holds response
type Response struct {
	Response *http.Response

	body       []byte
	size       int64
	receivedAt time.Time
}

// NewResponse new response
func NewResponse() *Response {
	return new(Response)
}
