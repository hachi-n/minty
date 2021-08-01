package mail

import (
	"bytes"
	"io"
	"net/mail"
)

type Mail struct {
	source *bytes.Buffer
	msg    *mail.Message
}

func NewMail(b []byte) *Mail {
	return &Mail{source: bytes.NewBuffer(b)}
}

func (p *Mail) Body() ([]byte, error) {
	msg, err := p.parse()
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(msg.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (p *Mail) parse() (*mail.Message, error) {
	if p.msg != nil {
		return p.msg, nil
	}
	msg, err := mail.ReadMessage(p.source)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
