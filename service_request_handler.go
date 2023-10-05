package main

import (
	"net/http"
	"strconv"
)

type serviceRequest struct {
	Name      string
	Endpoint  string
	Host      string
	Method    string
	SecretKey string
}

func (s serviceRequest) GetURL() string {
	return s.Host + s.Endpoint
}

func GetServiceRequest(request map[string]string, service service) serviceRequest {
	var service_request serviceRequest

	service_request.Name = service.Name
	service_request.Host = service.Base_url + ":" + strconv.Itoa(service.Port)
	service_request.SecretKey = service.SecretKey

	if request["url"] != "" {
		for _, endpoint := range service.Endpoints.GET {
			if endpoint == request["url"] {
				service_request.Endpoint = endpoint
				service_request.Method = "GET"
				break
			}
		}
		if service_request.Endpoint == "" {
			panic("Endpoint not found in the gateway config")
		}
	}

	return service_request
}

func ExecuteServiceRequest(service_request serviceRequest) {
	request, err := http.NewRequest(service_request.Method)
}
