package message_filter_processor

import (
	"kafka_module_2/internal/app/config"
	"kafka_module_2/internal/app/domain"
	"log/slog"
	"slices"
	"strconv"
	"strings"

	"github.com/lovoo/goka"
)

func MessageFilterProcessor(config *config.Config, logger *slog.Logger) func(ctx goka.Context, msg interface{}) {
	return func(ctx goka.Context, msg interface{}) {
		var (
			message domain.Message
			filters domain.UserFilters
			ok      bool
		)

		// get message
		if message, ok = msg.(domain.Message); !ok {
			logger.Error("illegal type: %T", msg)
			return
		}

		// get filters for recipient user
		RecipientID := strconv.FormatInt(message.RecipientID, 10)

		f := ctx.Lookup(goka.GroupTable(config.BlockCommandGroup()), RecipientID)
		if filters, ok = f.(domain.UserFilters); !ok {
			logger.Error("illegal stored value type: %T", f)
			return
		}

		// filter by blocked users
		UserID, err := strconv.ParseInt(ctx.Key(), 10, 64)
		if err != nil {
			logger.Error("illegal stored key: %s", ctx.Key())
			return
		}

		if slices.Contains(filters.BlockedUserIDs, UserID) {
			logger.Info("blocked user id: %d", message.RecipientID)
			return
		}

		// filter by blocked words
		if slices.ContainsFunc(filters.BlockedWords, func(s string) bool { return strings.Contains(message.Message, s) }) {
			logger.Info("blocked word: %s", message.Message)
			message.Message = "***"
		}

		// send message to topic 'messages.filtered'
		ctx.Emit(config.TopicFilteredMessages(),
			ctx.Key(),
			message,
		)
		logger.Info("message sent to userID %d: %v", message.RecipientID, message.Message)
	}
}
