package products

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
)

// Products service
type Products struct{}

// Product struct
type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

// List products struct
type List struct {
	Products []Product `json:"products"`
}

var ddb *dynamodb.DynamoDB

func init() {
	region := os.Getenv("AWS_REGION")
	if session, err := session.NewSession(&aws.Config{ // Use aws sdk to connect to dynamoDB
		Region: &region,
	}); err != nil {
		fmt.Println(fmt.Sprintf("Failed to connect to AWS: %s", err.Error()))
	} else {
		ddb = dynamodb.New(session) // Create DynamoDB client
	}
}

// Create product
func (products Products) Create(product Product) (Product, error) {
	var (
		id        = uuid.Must(uuid.NewV4(), nil).String()
		tableName = aws.String(os.Getenv("PRODUCTS_TABLE_NAME"))
	)

	// validate product
	if product.Name == "" {
		return product, errors.New("Name can't be empty")
	}

	if product.Description == "" {
		return product, errors.New("Description can't be empty")
	}

	// persist product
	product.ID = id
	product.CreatedAt = time.Now().String()

	item, _ := dynamodbattribute.MarshalMap(product)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: tableName,
	}

	_, err := ddb.PutItem(input)

	return product, err
}

// Update product
func (products Products) Update(product Product) (Product, error) {

	return product, nil
}

// Delete product
func (products Products) Delete(id string) (bool, error) {

	return true, nil
}

// GetAll products
//func (products Products) GetAll() (List, error) {
//
//	return nil, nil
//}
