package TTVClient

import (
	"time"

	"github.com/TomAvecLeVdsl/go-ttv-pubsub/pkg/Topic"
)

type IncomingMessage struct {
	Type string `json:"type"`
	Data struct {
		Message string `json:"message"`
		Topic   string `json:"topic"`
	} `json:"data"`
	Nonce string `json:"nonce"`
	Error string `json:"error"`
}

type OutgoingMessage struct {
	Type  string `json:"type,omitempty"`
	Nonce string `json:"nonce,omitempty"`
	Data  struct {
		Topics    []Topic.Topic `json:"topics,omitempty"`
		AuthToken string        `json:"auth_token,omitempty"`
	} `json:"data,omitempty"`
}

type ModerationActionMsg struct {
	Data struct {
		Type             string   `json:"type"`
		ModerationAction string   `json:"moderation_action"`
		Args             []string `json:"args"`
		CreatedBy        string   `json:"created_by"`
		CreatedByUserID  string   `json:"created_by_user_id"`
		MsgID            string   `json:"msg_id"`
		TargetUserID     string   `json:"target_user_id"`
		TargetUserLogin  string   `json:"target_user_login"`
	} `json:"data"`
}

type PointsMsg struct {
	Type string `json:"type"`
	Data struct {
		Timestamp  time.Time `json:"timestamp"`
		Redemption struct {
			ID   string `json:"id"`
			User struct {
				ID          string `json:"id"`
				Login       string `json:"login"`
				DisplayName string `json:"display_name"`
			} `json:"user"`
			ChannelID  string    `json:"channel_id"`
			RedeemedAt time.Time `json:"redeemed_at"`
			Reward     struct {
				ID                  string      `json:"id"`
				ChannelID           string      `json:"channel_id"`
				Title               string      `json:"title"`
				Prompt              string      `json:"prompt"`
				Cost                int         `json:"cost"`
				IsUserInputRequired bool        `json:"is_user_input_required"`
				IsSubOnly           bool        `json:"is_sub_only"`
				Image               interface{} `json:"image"`
				DefaultImage        struct {
					URL1X string `json:"url_1x"`
					URL2X string `json:"url_2x"`
					URL4X string `json:"url_4x"`
				} `json:"default_image"`
				BackgroundColor string `json:"background_color"`
				IsEnabled       bool   `json:"is_enabled"`
				IsPaused        bool   `json:"is_paused"`
				IsInStock       bool   `json:"is_in_stock"`
				MaxPerStream    struct {
					IsEnabled    bool `json:"is_enabled"`
					MaxPerStream int  `json:"max_per_stream"`
				} `json:"max_per_stream"`
				ShouldRedemptionsSkipRequestQueue bool      `json:"should_redemptions_skip_request_queue"`
				TemplateID                        string    `json:"template_id"`
				UpdatedForIndicatorAt             time.Time `json:"updated_for_indicator_at"`
			} `json:"reward"`
			Status string `json:"status"`
		} `json:"redemption"`
	} `json:"data"`
}

type WhisperMsg struct {
	Type string `json:"type"`
	Data struct {
		ID       string `json:"id"`
		LastRead int    `json:"last_read"`
		Archived bool   `json:"archived"`
		Muted    bool   `json:"muted"`
		SpamInfo struct {
			Likelihood        string `json:"likelihood"`
			LastMarkedNotSpam int    `json:"last_marked_not_spam"`
		} `json:"spam_info"`
		WhitelistedUntil string `json:"whitelisted_until"`
	} `json:"data"`
	DataObject struct {
		ID       string `json:"id"`
		LastRead int    `json:"last_read"`
		Archived bool   `json:"archived"`
		Muted    bool   `json:"muted"`
		SpamInfo struct {
			Likelihood        string `json:"likelihood"`
			LastMarkedNotSpam int    `json:"last_marked_not_spam"`
		} `json:"spam_info"`
		WhitelistedUntil string `json:"whitelisted_until"`
	} `json:"data_object"`
}

type CommerceMsg struct {
	UserName        string `json:"user_name"`
	DisplayName     string `json:"display_name"`
	ChannelName     string `json:"channel_name"`
	UserID          string `json:"user_id"`
	ChannelID       string `json:"channel_id"`
	Time            string `json:"time"`
	ItemImageURL    string `json:"item_image_url"`
	ItemDescription string `json:"item_description"`
	SupportsChannel bool   `json:"supports_channel"`
	PurchaseMessage struct {
		Message string `json:"message"`
		Emotes  []struct {
			Start int `json:"start"`
			End   int `json:"end"`
			ID    int `json:"id"`
		} `json:"emotes"`
	} `json:"purchase_message"`
}

type SubscriptionMsg struct {
	UserName    string `json:"user_name"`
	DisplayName string `json:"display_name"`
	ChannelName string `json:"channel_name"`
	UserID      string `json:"user_id"`
	ChannelID   string `json:"channel_id"`
	Time        string `json:"time"`
	SubPlan     string `json:"sub_plan"`
	SubPlanName string `json:"sub_plan_name"`
	Months      int    `json:"months"`
	Context     string `json:"context"`
	SubMessage  struct {
		Message string      `json:"message"`
		Emotes  interface{} `json:"emotes"`
	} `json:"sub_message"`
	RecipientID          string `json:"recipient_id"`
	RecipientUserName    string `json:"recipient_user_name"`
	RecipientDisplayName string `json:"recipient_display_name"`
}

type BitsBadgeMsg struct {
	UserID      string    `json:"user_id"`
	UserName    string    `json:"user_name"`
	ChannelID   string    `json:"channel_id"`
	ChannelName string    `json:"channel_name"`
	BadgeTier   int       `json:"badge_tier"`
	ChatMessage string    `json:"chat_message"`
	Time        time.Time `json:"time"`
}

type BitsMsg struct {
	Data struct {
		UserName         string    `json:"user_name"`
		ChannelName      string    `json:"channel_name"`
		UserID           string    `json:"user_id"`
		ChannelID        string    `json:"channel_id"`
		Time             time.Time `json:"time"`
		ChatMessage      string    `json:"chat_message"`
		BitsUsed         int       `json:"bits_used"`
		TotalBitsUsed    int       `json:"total_bits_used"`
		Context          string    `json:"context"`
		BadgeEntitlement struct {
			NewVersion      int `json:"new_version"`
			PreviousVersion int `json:"previous_version"`
		} `json:"badge_entitlement"`
	} `json:"data"`
	Version     string `json:"version"`
	MessageType string `json:"message_type"`
	MessageID   string `json:"message_id"`
	IsAnonymous bool   `json:"is_anonymous"`
}
