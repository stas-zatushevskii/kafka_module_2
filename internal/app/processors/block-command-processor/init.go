package block_command_processor

import (
	"kafka_module_2/internal/app/config"
	"kafka_module_2/internal/app/domain"
	builder "kafka_module_2/internal/pkg/goka-process"
	process "kafka_module_2/internal/pkg/goka-process"
	codec "kafka_module_2/internal/pkg/json-codec"
	"log/slog"
)

func New(config *config.Config, logger *slog.Logger) (*builder.Processor, error) {
	return process.NewProcessorBuilder().
		WithGroup(config.BlockCommandGroup()).
		WithInput(config.TopicBlockedMessages(), codec.JsonCodec[domain.Command]{}, BlockCommandProcessor(logger)).
		WithInput(config.TopicBlockedUsers(), codec.JsonCodec[domain.Command]{}, BlockCommandProcessor(logger)).
		WithPersist(codec.JsonCodec[domain.UserFilters]{}).
		WithBrokers(config.KafkaBrokers()).
		Build()
}
