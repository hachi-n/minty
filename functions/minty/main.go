package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hachi-n/minty/functions/internal/config"
	"github.com/hachi-n/minty/functions/internal/handlers"
	"os"
	"time"
)

type MyEvent struct {
	Records []struct {
		EventSource  string `json:"eventSource"`
		EventVersion string `json:"eventVersion"`
		Ses          struct {
			Mail struct {
				CommonHeaders struct {
					Date       string   `json:"date"`
					From       []string `json:"from"`
					MessageID  string   `json:"messageId"`
					ReturnPath string   `json:"returnPath"`
					Subject    string   `json:"subject"`
					To         []string `json:"to"`
				} `json:"commonHeaders"`
				Destination []string `json:"destination"`
				Headers     []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"headers"`
				HeadersTruncated bool      `json:"headersTruncated"`
				MessageID        string    `json:"messageId"`
				Source           string    `json:"source"`
				Timestamp        time.Time `json:"timestamp"`
			} `json:"mail"`
			Receipt struct {
				Action struct {
					FunctionArn    string `json:"functionArn"`
					InvocationType string `json:"invocationType"`
					Type           string `json:"type"`
				} `json:"action"`
				DkimVerdict struct {
					Status string `json:"status"`
				} `json:"dkimVerdict"`
				ProcessingTimeMillis int      `json:"processingTimeMillis"`
				Recipients           []string `json:"recipients"`
				SpamVerdict          struct {
					Status string `json:"status"`
				} `json:"spamVerdict"`
				SpfVerdict struct {
					Status string `json:"status"`
				} `json:"spfVerdict"`
				Timestamp    time.Time `json:"timestamp"`
				VirusVerdict struct {
					Status string `json:"status"`
				} `json:"virusVerdict"`
			} `json:"receipt"`
		} `json:"ses"`
	} `json:"Records"`
}

func initializeConfig(event MyEvent) error {
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

func handler(ctx context.Context, event MyEvent) error {
	err := initializeConfig(event)
	if err != nil {
		return err
	}
	return handlers.Apply()
}

func main() {
	lambda.Start(handler)
}
