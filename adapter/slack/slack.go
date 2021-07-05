package slack

import (
	"github.com/slack-go/slack"
)

type SlackAdapter interface {
	SendMessage(title, username, content string) error
}

type SlackAdapterImpl struct {
	token   string
	channel string
	bot     *slack.Client
}

func NewSlackAdapter(token string, channel string) SlackAdapter {
	return &SlackAdapterImpl{
		token:   token,
		channel: channel,
		bot:     slack.New(token),
	}
}

func (s SlackAdapterImpl) SendMessage(title, username, content string) error {
	_, _, _, err := s.bot.SendMessage(s.channel, slack.MsgOptionAttachments(slack.Attachment{
		Title: title,
		Text:  content,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "[From]",
				Value: username,
			},
		},
	}))
	if err != nil {
		return err
	}

	return nil
}
