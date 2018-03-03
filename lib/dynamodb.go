package serverless

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Dynamodb struct
type Dynamodb struct{}

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

// Add method
func (db Dynamodb) Add(entity interface{}, tableName string) error {
	item, _ := dynamodbattribute.MarshalMap(entity)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err := ddb.PutItem(input)

	return err
}
