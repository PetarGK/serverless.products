package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/petargk/serverless.products/products"
)

func main() {
	service := products.Products{}

	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Println("Add Product")

		product := Product{}

		json.Unmarshal([]byte(request.Body), product)

		if _, err := service.Create(product); err != nil {
			return events.APIGatewayProxyResponse{ // Error HTTP response
				Body:       err.Error(),
				StatusCode: 500,
			}, nil
		}

		body, _ := json.Marshal(product)
		return events.APIGatewayProxyResponse{ // Success HTTP response
			Body:       string(body),
			StatusCode: 200,
		}, nil
	})
}
