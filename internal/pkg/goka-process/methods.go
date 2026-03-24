package processors

import (
	"context"

	"github.com/lovoo/goka"
)

type Processor struct {
	processor *goka.Processor
}

func (p *Processor) Start(ctx context.Context) error {
	return p.processor.Run(ctx)
}
