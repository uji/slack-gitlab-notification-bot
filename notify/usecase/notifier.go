package usecase

import "notify/domain/repo"

// Notifier is notify from gitlab to slack
type Notifier struct {
	userRepo    repo.User
	slackClient slackClient
}

type slackClient interface {
	Notify(channelID string, text string) error
}

// NewNotifier is constructor of Notifier
func NewNotifier(userRepo repo.User, slackClient slackClient) *Notifier {
	return &Notifier{userRepo, slackClient}
}

// Notify is get slackID from gitlabID and notify to slack
func (n *Notifier) Notify(gitlabID string) error {
	user, err := n.userRepo.FindByGitlabID(gitlabID)
	if err != nil {
		return err
	}
	n.slackClient.Notify(user.GetSlackID(), "")
	return nil
}
