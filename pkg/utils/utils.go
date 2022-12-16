package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func (resp *Response) JSONResponse(res http.ResponseWriter, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	responseJson, err := json.Marshal(resp)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{"message": "error handling the response."}`))
		return
	}

	res.Write(responseJson)
}

func HttpClient() *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 100
	transport.MaxIdleConnsPerHost = 100
	transport.MaxConnsPerHost = 100
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	return client
}
