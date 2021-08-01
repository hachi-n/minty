package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hachi-n/minty/functions/internal/aws/lambda/events"
	"github.com/hachi-n/minty/functions/internal/config"
	"github.com/hachi-n/minty/functions/internal/handlers"
	"os"
)

func initializeConfig(event events.AmazonSesEvent) error {
	awsConfig := &config.AwsConfig{
		AccessKey: os.Getenv(config.ACCESS_KEY_ENV),
		SecretKey: os.Getenv(config.SECRET_KEY_ENV),
		Region:    os.Getenv(config.REGION_ENV),
	}

	s3Config := &config.S3Config{
		BucketName: os.Getenv(config.S3_BUCKET_NAME_ENV),
		Prefix:     os.Getenv(config.S3_KEY_PREFIX_ENV),
		Key:        event.Records[0].Ses.Mail.MessageID,
	}

	options := make(config.Option)
	notificationType := os.Getenv(config.NOTIFICATION_TYPE_ENV)
	switch notificationType {
	case "slack":
		options[config.SLACK_CHANNEL_ENV] = os.Getenv(config.SLACK_CHANNEL_ENV)
		options[config.SLACK_API_URL_ENV] = os.Getenv(config.SLACK_API_URL_ENV)
	case "line":
		//TODO
		// LINE Setting.
	}

	options[config.NOTIFICATION_TYPE_ENV] = notificationType

	if len(options) == 0 || awsConfig == nil || s3Config == nil {
		return fmt.Errorf("ENV ERROR!!!!!!!!!!!")
	}

	config.SettingConfig = &config.Config{
		Aws:    awsConfig,
		S3:     s3Config,
		Option: options,
	}
	return nil
}

func handler(ctx context.Context, event events.AmazonSesEvent) error {
	err := initializeConfig(event)
	if err != nil {
		return err
	}
	return handlers.Apply()
}

func main() {
	lambda.Start(handler)
}
