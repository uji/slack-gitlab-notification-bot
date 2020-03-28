package usecase

import (
	"notify/domain"
	"testing"
)

type userRepositoryMock struct{}

func (u userRepositoryMock) FindByGitlabID(gitlabID string) (*domain.User, error) {
	return &domain.User{}, nil
}

type slackClientMock struct{}

func (c slackClientMock) Notify(channelID string, text string) error {
	return nil
}

func TestNotify(t *testing.T) {
	userRepo := userRepositoryMock{}
	slackClientMock := slackClientMock{}
	n := NewNotifier(userRepo, slackClientMock)
	if err := n.Notify("text"); err != nil {
		t.Fatalf("get Error")
	}
}
