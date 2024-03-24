package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/eric3et/go_tutorial_1/internal/handlers"
	"github.com/go-chi/chi/v5"
)

var chiLambda *chiadapter.ChiLambda

func main() {

	// Check if the environment variable indicating cloud environment is set
	cloudEnv := os.Getenv("CLOUD_ENV")

	r := chi.NewRouter()

	handlers.Handler(r)

	if cloudEnv != "" {
		fmt.Println("Running Lambda")
		chiLambda = chiadapter.New(r)
		lambda.StartWithOptions(LambdaHandler, lambda.WithContext(context.Background()))
	} else {
		fmt.Println("Running Local Server :3000")
		http.ListenAndServe(":3000", r)
	}
}

func LambdaHandler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// log request
	fmt.Println("EVENT")
	fmt.Println(event)

	// handle request
	response, err := chiLambda.ProxyWithContext(ctx, event)

	// log response
	fmt.Println("RESPONSE")
	fmt.Println(response)

	// return response
	return response, err
}
