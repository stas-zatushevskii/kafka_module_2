package block_command_processor

import (
	"kafka_module_2/internal/app/constants"
	"kafka_module_2/internal/app/domain"
	builder "kafka_module_2/internal/pkg/goka-process"
	process "kafka_module_2/internal/pkg/goka-process"
	codec "kafka_module_2/internal/pkg/json-codec"
)

func New() (*builder.Processor, error) {
	return process.NewProcessorBuilder().
		WithGroup(constants.BlockCommandGroup).
		WithInput(constants.TopicBlockedMessages, codec.JsonCodec[domain.Command]{}, BlockCommandProcessor).
		WithInput(constants.TopicBlockedUsers, codec.JsonCodec[domain.Command]{}, BlockCommandProcessor).
		WithPersist(codec.JsonCodec[domain.UserFilters]{}).
		WithBrokers(constants.Brokers).
		Build()
}
