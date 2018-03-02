package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/petargk/serverless.products/products"
)

func main() {
	service := products.Products{}

	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Println("Add Product")

		Product product;

		json.Unmarshal([]byte(request.Body), product)

		return service.Create(product), nil
	})
}
