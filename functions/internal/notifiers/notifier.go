package notifiers

import (
	"github.com/hachi-n/minty/functions/internal/config"
	"github.com/hachi-n/minty/functions/internal/models/notifiers/line"
	"github.com/hachi-n/minty/functions/internal/models/notifiers/slack"
)

type Notifier interface {
	Notify() ([]byte, error)
}

func NewNotifier(messages []byte) Notifier {
	options := config.SettingConfig.Option
	notificationType := options[config.NOTIFICATION_TYPE_ENV]

	var d Notifier
	switch notificationType {
	case "slack":
		slackChannel := options[config.SLACK_CHANNEL_ENV]
		d = slack.NewSlack(slackChannel, messages)
	case "line":
		d = line.NewLine(messages)
	}
	return d
}

func Notify(notifier Notifier) error {
	body, err := notifier.Notify()
	if err != nil {
		return err
	}
	_ = body

	return nil
}


