// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// )

// func main() {
// 	creds, err := loadCredentials("credentials.csv")
// 	if err != nil {
// 		fmt.Println("Error loading credentials:", err)
// 		return
// 	}

// 	// Create a new AWS session using the loaded credentials
// 	awsSession, err := session.NewSession(&aws.Config{
// 		Region:      aws.String("us-east-1"),
// 		Credentials: creds,
// 	})
// 	if err != nil {
// 		fmt.Println("Error creating AWS session:", err)
// 		print("check")
// 		return
// 	}

// 	// Create an S3 client

// 	s3Client := s3.New(awsSession)

// 	// Set up the event notification configuration
// 	// Set up the event notification configuration
// 	notificationConfig := &s3.PutBucketNotificationConfigurationInput{
// 		Bucket: aws.String("myaws"),
// 		NotificationConfiguration: &s3.NotificationConfiguration{
// 			LambdaFunctionConfigurations: []*s3.LambdaFunctionConfiguration{
// 				{
// 					Events: aws.StringSlice([]string{
// 						"s3:ObjectCreated:*",
// 						"s3:ObjectRemoved:*",
// 					}),
// 					LambdaFunctionArn: aws.String("{your arn}"),
// 				},
// 			},
// 		},
// 	}

// 	// Set up the event notification using s3Client
// 	_, err = s3Client.PutBucketNotificationConfiguration(notificationConfig)
// 	if err != nil {
// 		log.Fatalf("unable to set up event notification, %v", err)
// 	}

// 	fmt.Println("S3 event notification configured successfully.")

// 	// List all buckets
// 	result, err := s3Client.ListBuckets(nil)
// 	if err != nil {
// 		fmt.Println("Error listing S3 buckets:", err)
// 		return
// 	}

// 	fmt.Println("S3 Buckets:")
// 	for _, bucket := range result.Buckets {
// 		fmt.Println(*bucket.Name)
// 	}

// }

// func loadCredentials(filePath string) (*credentials.Credentials, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(records) < 2 {
// 		return nil, fmt.Errorf("CSV file must have at least two rows for AccessKeyID and SecretAccessKey")
// 	}

// 	accessKeyID := records[1][0]
// 	secretAccessKey := records[1][1]

// 	return credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""), nil
// }