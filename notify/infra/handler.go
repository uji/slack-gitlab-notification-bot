package infra

import (
	"encoding/json"
	"errors"
	"notify/usecase"

	"github.com/aws/aws-lambda-go/events"
)

var (
	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func Handler(request events.APIGatewayProxyRequest, notifyer usecase.Notifier) (events.APIGatewayProxyResponse, error) {
	j := map[string]string{} // TODO: define request struct
	if err := json.Unmarshal([]byte(request.Body), j); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if err := notifier.Notify(gitlabID); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
