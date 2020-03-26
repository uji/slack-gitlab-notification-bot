package domain

// Message is notify message
type Message struct {
	to   string
	body string
}

// User is user to notify
type User struct {
	slackID  string
	gitlabID string
}

// GetSlackID get slack ID from User
func (u *User) GetSlackID() string {
	return u.slackID
}
