package notifiers

import (
	"github.com/hachi-n/minty/functions/internal/config"
	line2 "github.com/hachi-n/minty/functions/internal/models/notifiers/line"
	slack2 "github.com/hachi-n/minty/functions/internal/models/notifiers/slack"
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
		d = slack2.NewSlack(slackChannel, messages)
	case "line":
		d = line2.NewLine(messages)
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


