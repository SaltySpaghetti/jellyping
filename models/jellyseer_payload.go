package models

type JellyseerPayload struct {
	NotificationType string  `json:"notification_type"`
	Event            string  `json:"event"`
	Subject          string  `json:"subject"`
	Message          string  `json:"message"`
	Image            string  `json:"image"`
	Media            Media   `json:"media"`
	Request          Request `json:"request"`
	Issue            any     `json:"issue"`
	Comment          any     `json:"comment"`
	Extra            []any   `json:"extra"`
}

type Media struct {
	MediaType string `json:"media_type"`
	TmdbID    string `json:"tmdbId"`
	TvdbID    string `json:"tvdbId"`
	Status    string `json:"status"`
	Status4k  string `json:"status4k"`
}

type Request struct {
	RequestID                   string `json:"request_id"`
	RequestedByEmail            string `json:"requestedBy_email"`
	RequestedByUsername         string `json:"requestedBy_username"`
	RequestedByAvatar           string `json:"requestedBy_avatar"`
	RequestedBySettingsDiscord  string `json:"requestedBy_settings_discordId"`
	RequestedBySettingsTelegram string `json:"requestedBy_settings_telegramChatId"`
}
