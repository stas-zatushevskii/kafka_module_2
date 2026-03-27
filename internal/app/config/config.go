package config

import "github.com/lovoo/goka"

type Config struct {
	brokers brokers
	groups  groups
	topics  topics
}

type groups struct {
	messageFilterGroup goka.Group
	blockCommandGroup  goka.Group
}

type topics struct {
	topicMessages         goka.Stream
	topicFilteredMessages goka.Stream
	topicBlockedMessages  goka.Stream
	topicBlockedUsers     goka.Stream
}

type brokers struct {
	brokers []string
}
