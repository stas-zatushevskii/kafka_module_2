package goka_process

import (
	"errors"

	"github.com/lovoo/goka"
)

type Codec interface {
	Encode(any) ([]byte, error)
	Decode([]byte) (any, error)
}

type ProcessorBuilder struct {
	brokers   []string
	groupName goka.Group
	edges     []goka.Edge
}

func NewProcessorBuilder() *ProcessorBuilder {
	return &ProcessorBuilder{
		edges: []goka.Edge{},
	}
}

func (b *ProcessorBuilder) WithInput(topic goka.Stream, codec Codec, f func(ctx goka.Context, msg interface{})) *ProcessorBuilder {
	b.edges = append(b.edges, goka.Input(topic, codec, f))
	return b
}

func (b *ProcessorBuilder) WithOutput(topic goka.Stream, codec Codec) *ProcessorBuilder {
	b.edges = append(b.edges, goka.Output(topic, codec))
	return b
}

func (b *ProcessorBuilder) WithPersist(codec Codec) *ProcessorBuilder {
	b.edges = append(b.edges, goka.Persist(codec))
	return b
}

func (b *ProcessorBuilder) WithGroup(groupName goka.Group) *ProcessorBuilder {
	b.groupName = groupName
	return b
}

func (b *ProcessorBuilder) WithLookup(table goka.Table, codec Codec) *ProcessorBuilder {
	b.edges = append(b.edges, goka.Lookup(table, codec))
	return b
}

func (b *ProcessorBuilder) WithBrokers(brokers []string) *ProcessorBuilder {
	b.brokers = append(b.brokers, brokers...)
	return b
}

func (b *ProcessorBuilder) Build() (*Processor, error) {

	// validate provided values
	if len(b.brokers) == 0 {
		return nil, errors.New("must provide at least one broker")
	}

	if b.groupName == "" {
		return nil, errors.New("must provide group name")
	}

	// create new group
	g := goka.DefineGroup(b.groupName, b.edges...)

	// create new process
	tmConfig := goka.NewTopicManagerConfig()

	// fixme: only for local tests
	tmConfig.Table.Replication = 1
	tmConfig.Stream.Replication = 1

	p, err := goka.NewProcessor(
		b.brokers,
		g,
		goka.WithTopicManagerBuilder(
			goka.TopicManagerBuilderWithTopicManagerConfig(tmConfig),
		),
	)
	if err != nil {
		return nil, err
	}

	return &Processor{p}, nil
}
