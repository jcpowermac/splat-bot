package util

import (
	"fmt"

	"github.com/slack-go/slack"
)

type SlackClientInterface interface {
	PostEphemeral(channelID string, userID string, options ...slack.MsgOption) (string, error)
	PostMessage(channelID string, options ...slack.MsgOption) (string, string, error)
	OpenConversation(params *slack.OpenConversationParameters) (*slack.Channel, bool, bool, error)
	GetConversationReplies(params *slack.GetConversationRepliesParameters) (msgs []slack.Message, hasMore bool, nextCursor string, err error)
}

type StubInterface struct {
}

func (s *StubInterface) PostEphemeral(channelID string, userID string, options ...slack.MsgOption) (string, error) {
	return "", fmt.Errorf("PostEphemeral")
}

func (s *StubInterface) PostMessage(channelID string, options ...slack.MsgOption) (string, string, error) {
	return "", "", fmt.Errorf("PostMessage")
}

func (s *StubInterface) OpenConversation(params *slack.OpenConversationParameters) (*slack.Channel, bool, bool, error) {
	return nil, false, false, fmt.Errorf("OpenConversation")
}

func (s *StubInterface) GetConversationReplies(params *slack.GetConversationRepliesParameters) (msgs []slack.Message, hasMore bool, nextCursor string, err error) {
	return nil, false, "", fmt.Errorf("GetConversationReplies")
}
