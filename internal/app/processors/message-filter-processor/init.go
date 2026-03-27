package message_filter_processor

import (
	"kafka_module_2/internal/app/config"
	"kafka_module_2/internal/app/domain"
	builder "kafka_module_2/internal/pkg/goka-process"
	process "kafka_module_2/internal/pkg/goka-process"
	codec "kafka_module_2/internal/pkg/json-codec"
	"log/slog"

	"github.com/lovoo/goka"
)

func New(config *config.Config, logger *slog.Logger) (*builder.Processor, error) {
	return process.NewProcessorBuilder().
		WithGroup(config.MessageFilterGroup()).
		WithInput(config.TopicMessages(), codec.JsonCodec[domain.Message]{}, MessageFilterProcessor(config, logger)).
		WithOutput(config.TopicFilteredMessages(), codec.JsonCodec[domain.Message]{}).
		WithBrokers(config.KafkaBrokers()).
		WithLookup(goka.GroupTable(config.BlockCommandGroup()), codec.JsonCodec[domain.UserFilters]{}).
		Build()
}
