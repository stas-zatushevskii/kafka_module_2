package config

import "github.com/lovoo/goka"

func (c *Config) KafkaBrokers() []string {
	result := make([]string, len(c.brokers.brokers))
	copy(result, c.brokers.brokers)
	return result
}

func (c *Config) MessageFilterGroup() goka.Group {
	return c.groups.messageFilterGroup
}

func (c *Config) BlockCommandGroup() goka.Group {
	return c.groups.blockCommandGroup
}

func (c *Config) TopicMessages() goka.Stream {
	return c.topics.topicMessages
}

func (c *Config) TopicFilteredMessages() goka.Stream {
	return c.topics.topicFilteredMessages
}

func (c *Config) TopicBlockedMessages() goka.Stream {
	return c.topics.topicBlockedMessages
}

func (c *Config) TopicBlockedUsers() goka.Stream {
	return c.topics.topicBlockedUsers
}
