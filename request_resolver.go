package main

import (
	"io"
	"net/http"
)

func HandleRootRequest(w http.ResponseWriter, r *http.Request) {
	parsedRequest := ParseRequest(r)

	if parsedRequest["service_name"] == "" {
		panic("Service-Name is missing from the request's headers")
	}

	services := DecodeGatewayConfig()
	var service_request serviceRequest

	for _, service := range services.Services {
		if service.Name == parsedRequest["service_name"] {
			service_request = GetServiceRequest(parsedRequest, service)
			break
		}
	}

	if (service_request == serviceRequest{}) {
		panic("The requested service is not found")
	}

	ExecuteServiceRequest(service_request)
}

func ParseRequest(r *http.Request) map[string]string {
	body := r.Body

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		panic("Something went wrong while reading request body")
	}

	parsedRequest := make(map[string]string)

	parsedRequest["method"] = r.Method
	// parsedRequest["headers"] = r.Header
	parsedRequest["host"] = r.Host
	parsedRequest["url"] = r.URL.Path
	parsedRequest["query"] = r.URL.RawQuery
	parsedRequest["body"] = string(bodyBytes)

	for key, value := range r.Header {
		if key == "Service-Name" && len(value) > 0 {
			parsedRequest["service_name"] = value[0]
		}
	}

	return parsedRequest
}
