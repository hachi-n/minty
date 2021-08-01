package mail

import (
	svc "github.com/hachi-n/minty/functions/internal/aws/service/s3"
)

type Downloader interface {
	Download() ([]byte, error)
}

type Client struct {
	downloader Downloader
}

func NewClient() *Client {
	return &Client{
		downloader: svc.NewS3Service(),
	}
}

func (client *Client) DownloadMail() (*Mail, error) {
	messages, err := client.downloader.Download()
	if err != nil {
		return nil, err
	}
	return NewMail(messages), nil
}
