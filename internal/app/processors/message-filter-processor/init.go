package message_filter_processor

import (
	"kafka_module_2/internal/app/constants"
	"kafka_module_2/internal/app/domain"
	builder "kafka_module_2/internal/pkg/goka-process"
	process "kafka_module_2/internal/pkg/goka-process"
	codec "kafka_module_2/internal/pkg/json-codec"

	"github.com/lovoo/goka"
)

func New() (*builder.Processor, error) {
	return process.NewProcessorBuilder().
		WithGroup(constants.MessageFilterGroup).
		WithInput(constants.TopicMessages, codec.JsonCodec[domain.Message]{}, MessageFilterProcessor).
		WithOutput(constants.TopicFilteredMessages, codec.JsonCodec[domain.Message]{}).
		WithBrokers(constants.Brokers).
		WithLookup(goka.GroupTable(constants.BlockCommandGroup), codec.JsonCodec[domain.UserFilters]{}).
		Build()
}
