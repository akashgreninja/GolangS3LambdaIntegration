package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, s3Event events.S3Event) error {
	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		fmt.Printf("S3 Event Received - Bucket: %s, Key: %s\n", bucket, key)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
