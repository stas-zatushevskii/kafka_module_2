package processors

import (
	"kafka_module_2/internal/app/constants"
	"kafka_module_2/internal/app/domain"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/lovoo/goka"
)

func MessageFilterProcessor(ctx goka.Context, msg interface{}) {
	var (
		message domain.Message
		filters domain.UserFilters
		ok      bool
	)

	// get message
	if message, ok = msg.(domain.Message); !ok {
		log.Printf("illegal type: %T", msg)
		return
	}

	// get filters for recipient user
	RecipientID := strconv.FormatInt(message.RecipientID, 10)

	f := ctx.Lookup(goka.GroupTable(constants.BlockCommandGroup), RecipientID)
	if filters, ok = f.(domain.UserFilters); !ok {
		log.Printf("illegal stored value type: %T", f)
		return
	}

	// filter by blocked users
	UserID, err := strconv.ParseInt(ctx.Key(), 10, 64)
	if err != nil {
		log.Printf("illegal stored key: %s", ctx.Key())
	}

	if slices.Contains(filters.BlockedUserIDs, UserID) {
		log.Printf("blocked user id: %d", message.RecipientID)
		return
	}

	// filter by blocked words
	if slices.ContainsFunc(filters.BlockedWords, func(s string) bool { return strings.Contains(message.Message, s) }) {
		log.Printf("blocked word: %s", message.Message)
		message.Message = "***"
	}

	// send message to topic 'messages.filtered'
	ctx.Emit(constants.TopicFilteredMessages,
		ctx.Key(),
		message,
	)
	log.Printf("message sent to userID %d: %v", message.RecipientID, message.Message)
}
