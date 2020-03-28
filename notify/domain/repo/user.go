package repo

import "notify/domain"

// User is repository of User model
type User interface {
	FindByGitlabID(gitlabID string) (*domain.User, error)
}
