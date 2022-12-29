package middleware

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"os"
)

func RespondMessage(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", response)
}

func ResponseMessage(body string, code int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
			IsBase64Encoded: false,
			Headers: map[string]string{
				"Access-Control-Expose-Headers": "x-api-key",
				"Access-Control-Allow-Headers":  "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With",
				"Access-Control-Allow-Methods":  "GET, POST, OPTIONS",
				"Content-Type":                  "application/json",
				"Access-Control-Allow-Origin":   os.Getenv("OriginHeader"),
			},
			Body:       body,
			StatusCode: code,
		},
		nil
}
