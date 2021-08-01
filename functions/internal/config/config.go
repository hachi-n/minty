package config

import "fmt"

var SettingConfig *Config

type Config struct {
	Aws    *AwsConfig
	S3     *S3Config
	Option Option
}

const (
	ACCESS_KEY_ENV = "ACCESS_KEY"
	SECRET_KEY_ENV = "SECRET_KEY"
	REGION_ENV     = "REGION"
)

type AwsConfig struct {
	AccessKey string
	SecretKey string
	Region    string
}

const (
	S3_BUCKET_NAME_ENV = "BUCKET_NAME"
	S3_KEY_PREFIX_ENV  = "KEY_PREFIX"
)

type S3Config struct {
	BucketName string
	Prefix     string
	Key        string
}

func (c *S3Config) KeyString() string {
	return fmt.Sprintf("%s/%s", c.Prefix, c.Key)
}

// slack
const (
	SLACK_CHANNEL_ENV = "SLACK_CHANNEL"
	SLACK_API_URL_ENV = "SLACK_API_URL"
)

// line
const (
	LINE_API_URL_ENV = "LINE_API_URL"
)

const (
	NOTIFICATION_TYPE_ENV = "NOTIFICATION_TYPE"
)

type Option = map[string]string
