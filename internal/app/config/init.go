package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/lovoo/goka"
)

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	cfg := &Config{
		brokers: brokers{
			brokers: parseArray(os.Getenv("KAFKA_BROKERS")),
		},
		groups: groups{
			messageFilterGroup: goka.Group(os.Getenv("MESSAGE_FILTER_GROUP")),
			blockCommandGroup:  goka.Group(os.Getenv("BLOCK_COMMAND_GROUP")),
		},
		topics: topics{
			topicMessages:         goka.Stream(os.Getenv("TOPIC_MESSAGES")),
			topicFilteredMessages: goka.Stream(os.Getenv("TOPIC_FILTERED_MESSAGES")),
			topicBlockedMessages:  goka.Stream(os.Getenv("TOPIC_BLOCKED_MESSAGES")),
			topicBlockedUsers:     goka.Stream(os.Getenv("TOPIC_BLOCKED_USERS")),
		},
	}

	if err := validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func parseArray(value string) []string {
	if value == "" {
		return nil
	}

	parts := strings.Split(value, ",")
	result := make([]string, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}

	return result
}

func validate(cfg *Config) error {
	var missing []string

	if len(cfg.brokers.brokers) == 0 {
		missing = append(missing, "KAFKA_BROKERS")
	}

	if cfg.groups.messageFilterGroup == "" {
		missing = append(missing, "MESSAGE_FILTER_GROUP")
	}

	if cfg.groups.blockCommandGroup == "" {
		missing = append(missing, "BLOCK_COMMAND_GROUP")
	}

	if cfg.topics.topicMessages == "" {
		missing = append(missing, "TOPIC_MESSAGES")
	}

	if cfg.topics.topicFilteredMessages == "" {
		missing = append(missing, "TOPIC_FILTERED_MESSAGES")
	}

	if cfg.topics.topicBlockedMessages == "" {
		missing = append(missing, "TOPIC_BLOCKED_MESSAGES")
	}

	if cfg.topics.topicBlockedUsers == "" {
		missing = append(missing, "TOPIC_BLOCKED_USERS")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required env vars: %s", strings.Join(missing, ", "))
	}

	return nil
}
