package app

import (
	"context"
	"fmt"
	blockCommandProcessor "kafka_module_2/internal/app/processors/block-command-processor"
	messageFilterProcessor "kafka_module_2/internal/app/processors/message-filter-processor"
	process "kafka_module_2/internal/pkg/goka-process"
	"kafka_module_2/internal/pkg/graceful"
)

type App struct {
	blockCommandProcessor  *process.Processor
	filterMessageProcessor *process.Processor
}

func New() (*App, error) {

	BlockCommandProcessor, err := blockCommandProcessor.New()
	if err != nil {
		return nil, fmt.Errorf("error creating block command processor: %w", err)
	}

	MessageFilterProcessor, err := messageFilterProcessor.New()
	if err != nil {
		return nil, fmt.Errorf("error creating message filter processor: %w", err)
	}

	return &App{
		blockCommandProcessor:  BlockCommandProcessor,
		filterMessageProcessor: MessageFilterProcessor,
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
