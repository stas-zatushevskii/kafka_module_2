package domain

// BlockWords describes a command for blocking word for a specific user
type BlockWords struct {
	// fixme: maybe UserID dont needed
	UserID     int64    `json:"user_id"`
	BlockWords []string `json:"block_words"`
}

// BlockUsers describes a command for blocking user for a specific user
type BlockUsers struct {
	// fixme: maybe UserID dont needed
	UserID       int64   `json:"user_id"`
	BlockUserIDs []int64 `json:"block_user_ids"`
}

// Command describes a universal command for updating a user's message-filter settings.
type Command struct {
	BlockWords *BlockWords `json:"block_words"`
	BlockUsers *BlockUsers `json:"block_users"`
}

// UserFilters describes KTable for block-command-group
type UserFilters struct {
	// fixme: maybe UserID dont needed
	UserID         int64    `json:"user_id"` // unique key
	BlockedUserIDs []int64  `json:"blocked_user_ids"`
	BlockedWords   []string `json:"blocked_words"`
}
