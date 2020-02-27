package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/peacemakr-io/peacemakr-go-sdk/pkg/tools"
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

func main() {

	svc := getDynamoDBClient()

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String("2020"),
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	item := Item{}

	config := tools.EncryptorConfig{
		ApiKey:     "your-api-key-here",
		ClientName: "Sample Decrypt Client",
		Url:        nil,
		Persister:  nil,
		Logger:     nil,
	}

	err = dynamodbattribute.UnmarshalMapAndDecrypt(result.Item, &item, &config)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if item.Title == "" {
		fmt.Println("Could not find 'The Big New Movie' (2015)")
		return
	}

	fmt.Println("Found item:")
	fmt.Println("Year:  ", item.Year)
	fmt.Println("Title: ", item.Title)
	fmt.Println("Plot:  ", item.Info.Plot)
	fmt.Println("Rating:", item.Info.Rating)
}
