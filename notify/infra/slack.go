package infra

import "github.com/slack-go/slack"

// SlackClient is post message client of slack
type SlackClient struct {
	client *slack.Client
}

// NewSlackClient is constructor of slackClient
func NewSlackClient(botToken string) *SlackClient {
	return &SlackClient{slack.New(botToken)}
}

// Notify post message to slack channel
func (c *SlackClient) Notify(channelID string, text string) error {
	_, _, err := c.client.PostMessage(channelID, slack.MsgOptionText(text, false))
	return err
}
