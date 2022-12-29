package search

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"mobile/internal/middleware"
	"net/http"
)

// SearchLocationByCords exported func to find a location using long and lat
func SearchLocationByCords(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// only GET is allowed
	if req.HTTPMethod != "GET" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Methos not allowed\n"),
		}, nil
	}

	// check parameter
	lat, ok := req.PathParameters["lat"]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Invalid Parameter"),
		}, nil
	}

	// check parameter
	long, ok := req.PathParameters["long"]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Invalid Parameter"),
		}, nil
	}

	location, err := dataSearchLocationByCords(lat, long)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("No data found"),
		}, nil
	}

	locationJson, err := json.Marshal(location)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"status": "failed", "msg": "` + err.Error() + `"}`),
		}, nil
	}

	response, err := middleware.ResponseMessage(string(locationJson), http.StatusOK)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"status": "failed", "msg": "` + err.Error() + `"}`),
		}, nil
	}

	return response, nil

}

// NotFound - path not found
func NotFound(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// only GET is allowed
	if req.HTTPMethod != "GET" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Method not allowed...",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Body:       fmt.Sprintf("Nothing to see here"),
	}, nil
}
