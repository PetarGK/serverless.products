package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/petargk/serverless.products/products"
)

func main() {
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		fmt.Println("Add Product")
		service := products.Products{}
		product := &products.Product{}
		json.Unmarshal([]byte(request.Body), product)
		product, err := service.Create(product)

		var result string
		var code = 201

		if err != nil {
			result = err.Error()
			code = err.Code
		} else {
			body, _ := json.Marshal(product)
			result = string(body)
		}

		return events.APIGatewayProxyResponse{
			Body:       result,
			StatusCode: code,
		}, nil
	})
}
