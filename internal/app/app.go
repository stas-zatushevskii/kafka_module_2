package app

import (
	"context"
	"fmt"
	c "kafka_module_2/internal/app/config"
	blockCommandProcessor "kafka_module_2/internal/app/processors/block-command-processor"
	messageFilterProcessor "kafka_module_2/internal/app/processors/message-filter-processor"
	process "kafka_module_2/internal/pkg/goka-process"
	"kafka_module_2/internal/pkg/graceful"
	"log/slog"
)

type App struct {
	blockCommandProcessor  *process.Processor
	filterMessageProcessor *process.Processor
}

// New initializes the application by loading configuration,
// creating OS signal, Goka processors,
// and returning the assembled App instance.
func New(logger *slog.Logger) (*App, error) {

	// load config
	config, err := c.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	// create block-command processor
	BlockCommandProcessor, err := blockCommandProcessor.New(config, logger)
	if err != nil {
		return nil, fmt.Errorf("error creating block command processor: %w", err)
	}

	// create message-filter processor
	MessageFilterProcessor, err := messageFilterProcessor.New(config, logger)
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
