package line

import (
	"encoding/json"
	"github.com/hachi-n/minty/functions/internal/apis"
	"github.com/hachi-n/minty/functions/internal/config"
)

type Line struct {
	apiUrl string
	Messages string
}

func NewLine(messages []byte) *Line {
	return &Line{
		apiUrl: config.SettingConfig.Option[config.LINE_API_URL_ENV],
		Messages: string(messages),
	}
}

func (d *Line) Notify() ([]byte, error) {
	jsonByte, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return apis.JsonPost(d.apiUrl, jsonByte)
}
