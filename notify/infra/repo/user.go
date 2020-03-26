package repo

import "slack-gitlab-notification-bot/notify/domain"

// User is repository of User model
type User struct {
}

// FindByGitlabID is get user from database
func (u *User) FindByGitlabID(gitlabID string) (*domain.User, error) {
	return &domain.User{}, nil
}
