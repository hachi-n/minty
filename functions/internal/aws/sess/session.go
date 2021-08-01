package sess

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hachi-n/minty/functions/internal/config"
)
var sess *session.Session

func NewSession() *session.Session {
	if sess != nil {
		return sess
	}
	awsConfig := config.SettingConfig.Aws
	s := session.Must(
		session.NewSession(
			&aws.Config{
				Credentials: credentials.NewStaticCredentials(awsConfig.AccessKey, awsConfig.SecretKey, ""),
				Region:      aws.String(awsConfig.Region),
			},
		),
	)

	sess = s
	return sess
}
