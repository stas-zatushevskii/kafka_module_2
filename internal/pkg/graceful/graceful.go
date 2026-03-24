package graceful

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"
)

type graceful struct {
	processes []process
}

func New(processes ...process) *graceful {
	return &graceful{
		processes: processes,
	}
}

func (gr *graceful) Start(ctx context.Context) error {
	g, gCtx := errgroup.WithContext(ctx)

	for _, proc := range gr.processes {
		if proc.disabled {
			continue
		}

		g.Go(func() error {
			err := proc.starter.Start(gCtx)
			if err != nil {
				log.Println("graceful start error:", err)
			}
			return err
		})

	}

	err := g.Wait()
	if err != nil {
		log.Println("Application stopped with error:", err)

		return err
	}

	return nil
}
