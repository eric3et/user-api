package tools

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var dynamo *dynamodb.DynamoDB

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Age      int    `json:"age,omitempty"`
	Token    string `json:"token,omitempty"`
}

const TABLE_NAME = "users"

func init() {
	dynamo = connectDynamo()
}

func connectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})))
}

func CreateTable() {
	_, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(TABLE_NAME),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func DBPutUser(user User) {
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(user.Id)),
			},
			"UserName": {
				S: aws.String(user.Username),
			},
			"Age": {
				N: aws.String(strconv.Itoa(user.Age)),
			},
			"Token": {
				S: aws.String(user.Token),
			},
		},
		TableName: aws.String(TABLE_NAME),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func DBUpdateUser(user User) {
	_, err := dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#U": aws.String("Username"),
			"#A": aws.String("Age"),
			"#T": aws.String("Token"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Username": {
				S: aws.String(user.Username),
			},
			":Age": {
				S: aws.String(strconv.Itoa(user.Age)),
			},
			":Token": {
				S: aws.String(user.Token),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(user.Id)),
			},
		},
		TableName:        aws.String(TABLE_NAME),
		UpdateExpression: aws.String("SET #U = :Username, #A = :Age, #T = :Token"),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}
}

func DBDeleteUser(id int) {
	_, err := dynamo.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String(TABLE_NAME),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
	}

}

func DBGetUser(id int) (user *User, err error) {
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: aws.String(TABLE_NAME),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
		return nil, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func DBListUser() ([]*User, error) {
	// Define the scan input parameters
	input := &dynamodb.ScanInput{
		TableName: aws.String(TABLE_NAME),
	}

	// Perform the scan operation
	result, err := dynamo.Scan(input)
	if err != nil {
		return nil, err
	}

	// Initialize a slice to store the users
	var users []*User

	// Unmarshal the scanned items into User structs
	for _, item := range result.Items {
		var user User
		err := dynamodbattribute.UnmarshalMap(item, &user)
		if err != nil {
			return nil, err
		}
		// Append a copy of the user to the users slice
		users = append(users, &user)
	}

	return users, nil
}
