package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/peacemakr-io/peacemakr-go-sdk/pkg/tools"

	"fmt"
)

// Create structs to hold info about new item
type ItemInfo struct {
	Plot   string `json:"plot" encrypt:"true"`
	Rating string `json:"rating"`
}

type Item struct {
	Year  int      `json:"year"`
	Title string   `json:"title" encrypt:"true"`
	Info  ItemInfo `json:"info"`
}

func getDynamoDBClient() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000")})

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	return svc
}

func createExampleTable(svc *dynamodb.DynamoDB, tableName string) error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("year"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("year"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Created the table %v in us-west-2\n", tableName)
	return nil
}

func main() {

	svc := getDynamoDBClient()

	createExampleTable(svc, "Movies")

	info := ItemInfo{
		Plot:   "Encrypted Nothing happens at all.",
		Rating: "2.0",
	}

	item := Item{
		Year:  2020,
		Title: "Encrypted Text",
		Info:  info,
	}

	// Set up Peacemakr Configs

	config := tools.EncryptorConfig{
		ApiKey:     "YOUR_API_KEY_HERE",
		ClientName: "Sample Encrypt Client",
		Url:        nil,
		Persister:  nil,
		Logger:     nil,
	}

	av, err := dynamodbattribute.EncryptAndMarshalMap(&item, &config)

	log.Println("item to be added", item)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Movies"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added data")
}
