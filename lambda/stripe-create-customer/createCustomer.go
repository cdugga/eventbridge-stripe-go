package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
)

type Item struct {
	CustomerID   string
}


func CreateCustomerHandler(ctx context.Context, request events.APIGatewayProxyRequest)  (string, error){

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Update item in table Movies
	tableName := os.Getenv("CUSTOMER_TABLE_NAME")
	fmt.Println("Tablename:", tableName)

	item := Item{
		CustomerID: "Some customer ID",
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		fmt.Println("Got error marshalling new item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("marshalled:", av)

	input := &dynamodb.PutItemInput{
		Item: av,
		//ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName: aws.String(tableName),
	}

	result, err := svc.PutItem(input)
	if err != nil {
		fmt.Println("Some error:", err.Error())

	}
	fmt.Println("Successfully updated")

	return fmt.Sprintf("Hello %s!", result ), nil
}


func main() {
	lambda.Start(CreateCustomerHandler)
}

