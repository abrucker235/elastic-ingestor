package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))

	signer := v4.NewSigner(sess.Config.Credentials)

	lambda.Start(handler)
}

func handler(ctx context.Context, event events.KinesisEvent) error {
	for _, record := range event.Records {
		//TODO Publish Records From Kinesis to Elastic Search
	}
	return nil
}
