package constants

import "github.com/lovoo/goka"

var (
	TopicMessages         goka.Stream = "messages"
	TopicFilteredMessages goka.Stream = "messages.filtered"
	TopicBlockedMessages  goka.Stream = "messages.blocked"
	TopicBlockedUsers     goka.Stream = "users.blocked"
)
