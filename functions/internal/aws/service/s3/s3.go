package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/hachi-n/minty/functions/internal/aws/sess"
	"github.com/hachi-n/minty/functions/internal/config"
)

type S3Service struct {
	Bucket     string
	Key        string
	downloader *s3manager.Downloader
}

func NewS3Service() *S3Service {
	return &S3Service{
		Bucket:     config.SettingConfig.S3.BucketName,
		Key:        config.SettingConfig.S3.KeyString(),
		downloader: s3manager.NewDownloader(sess.NewSession()),
	}
}

func (svc *S3Service) Download() ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(svc.Bucket),
		Key:    aws.String(svc.Key),
	}

	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := svc.downloader.Download(buf, input)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
