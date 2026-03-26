package app

import (
	"context"
	"fmt"
	"kafka_module_2/internal/app/constants"
	"kafka_module_2/internal/app/domain"
	"kafka_module_2/internal/app/processors"
	process "kafka_module_2/internal/pkg/goka-process"
	"kafka_module_2/internal/pkg/graceful"
	codec "kafka_module_2/internal/pkg/json-codec"

	"github.com/lovoo/goka"
)

type App struct {
	blockCommandProcessor  *process.Processor
	filterMessageProcessor *process.Processor
}

func New() (*App, error) {

	blockCommandProcessor, err := process.NewProcessorBuilder().
		WithGroup(constants.BlockCommandGroup).
		WithInput(constants.TopicBlockedMessages, codec.JsonCodec[domain.Command]{}, processors.BlockCommandProcessor).
		WithInput(constants.TopicBlockedUsers, codec.JsonCodec[domain.Command]{}, processors.BlockCommandProcessor).
		WithPersist(codec.JsonCodec[domain.UserFilters]{}).
		WithBrokers(constants.Brokers...).
		Build()

	if err != nil {
		return nil, fmt.Errorf("create block command processor error: %w", err)
	}

	filterMessageProcessor, err := process.NewProcessorBuilder().
		WithGroup(constants.MessageFilterGroup).
		WithInput(constants.TopicMessages, codec.JsonCodec[domain.Message]{}, processors.MessageFilterProcessor).
		WithOutput(constants.TopicFilteredMessages, codec.JsonCodec[domain.Message]{}).
		WithBrokers(constants.Brokers...).
		WithLookup(goka.GroupTable(constants.BlockCommandGroup), codec.JsonCodec[domain.UserFilters]{}).
		Build()

	if err != nil {
		return nil, fmt.Errorf("create filter message processor error: %w", err)
	}

	return &App{
		blockCommandProcessor:  blockCommandProcessor,
		filterMessageProcessor: filterMessageProcessor,
	}, nil
}

// Start runs the application processes under a graceful shutdown supervisor.
func (app *App) Start() error {
	gr := graceful.New(
		graceful.NewProcess(app.blockCommandProcessor),
		graceful.NewProcess(app.filterMessageProcessor),
	)

	err := gr.Start(context.Background())
	if err != nil {
		return err
	}

	return nil
}
