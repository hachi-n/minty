package slack

import (
	"encoding/json"
	"github.com/hachi-n/minty/functions/internal/apis"
	"github.com/hachi-n/minty/functions/internal/config"
)

type Slack struct {
	apiUrl  string
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func NewSlack(channel string, messages []byte) *Slack {
	return &Slack{
		Channel: channel,
		Text:    string(messages),
		apiUrl: config.SettingConfig.Option[config.SLACK_API_URL_ENV],
	}
}

func (d *Slack) Notify() ([]byte, error) {
	jsonByte, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return apis.JsonPost(d.apiUrl, jsonByte)
}
