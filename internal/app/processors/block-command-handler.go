package processors

import (
	"kafka_module_2/internal/app/domain"
	"log"

	"github.com/lovoo/goka"
)

func BlockCommandProcessor(ctx goka.Context, msg interface{}) {
	var (
		filters domain.UserFilters
		command domain.Command
		ok      bool
	)

	// get command
	if command, ok = msg.(domain.Command); !ok {
		log.Printf("illegal type: %T", msg)
		return
	}

	// get existed filters
	if value := ctx.Value(); value != nil {
		if filters, ok = value.(domain.UserFilters); !ok {
			log.Printf("illegal stored value type: %T", value)
			return
		}
	}

	// update filters according command type
	switch {
	case command.BlockUsers != nil:
		filters.BlockedUserIDs = append(filters.BlockedUserIDs, command.BlockUsers.BlockUserIDs...)
	case command.BlockWords != nil:
		filters.BlockedWords = append(filters.BlockedWords, command.BlockWords.BlockWords...)
	}

	// set updated filters
	ctx.SetValue(filters)
}
