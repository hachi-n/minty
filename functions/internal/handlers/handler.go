package handlers

import (
	"github.com/hachi-n/minty/functions/internal/models/mail"
	"github.com/hachi-n/minty/functions/internal/notifiers"
)

func Apply() error {
	m, err := mail.NewClient().DownloadMail()
	if err != nil {
		return err
	}

	//TODO
	// m.Body() should be normalized.
	body, err := m.Body()
	if err != nil {
		return err
	}
	notifier := notifiers.NewNotifier(body)
	return notifiers.Notify(notifier)
}
