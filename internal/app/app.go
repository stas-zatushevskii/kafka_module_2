package app

import "github.com/lovoo/goka"

type App struct{}

var (
	brokers = []string{"localhost:9092"}

	topicMessages         goka.Stream = "messages"
	topicFilteredMessages goka.Stream = "messages.filtered"
	topicBlockedMessages  goka.Stream = "messages.blocked"
	topicBlockedUsers     goka.Stream = "users.blocked"

	messageFilterGroup goka.Group = "message-filter-group"
	blockCommandGroup  goka.Group = "block-command-group"
)

func New() (*App, error) {

	// 1 create use cases

	// 2 create processors with default builder

	// 3 add use cases to processors

	// 4 run all process in graceful

	return &App{}, nil
}
