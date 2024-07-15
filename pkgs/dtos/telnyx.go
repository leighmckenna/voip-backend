package dtos

import "time"

type IncommingMessage struct {
	Data struct {
		EventType  string    `json:"event_type"`
		ID         string    `json:"id"`
		OccurredAt time.Time `json:"occurred_at"`
		Payload    struct {
			AutoresponseType any    `json:"autoresponse_type"`
			Cc               []any  `json:"cc"`
			CompletedAt      any    `json:"completed_at"`
			Cost             any    `json:"cost"`
			Direction        string `json:"direction"`
			Encoding         string `json:"encoding"`
			Errors           []any  `json:"errors"`
			From             struct {
				Carrier     string `json:"carrier"`
				LineType    string `json:"line_type"`
				PhoneNumber string `json:"phone_number"`
			} `json:"from"`
			ID                 string    `json:"id"`
			IsSpam             bool      `json:"is_spam"`
			Media              []any     `json:"media"`
			MessagingProfileID string    `json:"messaging_profile_id"`
			OrganizationID     string    `json:"organization_id"`
			Parts              int       `json:"parts"`
			ReceivedAt         time.Time `json:"received_at"`
			RecordType         string    `json:"record_type"`
			SentAt             any       `json:"sent_at"`
			Subject            string    `json:"subject"`
			Tags               []any     `json:"tags"`
			Text               string    `json:"text"`
			To                 []struct {
				Carrier     string `json:"carrier"`
				LineType    string `json:"line_type"`
				PhoneNumber string `json:"phone_number"`
				Status      string `json:"status"`
			} `json:"to"`
			Type               string `json:"type"`
			ValidUntil         any    `json:"valid_until"`
			WebhookFailoverURL string `json:"webhook_failover_url"`
			WebhookURL         string `json:"webhook_url"`
		} `json:"payload"`
		RecordType string `json:"record_type"`
	} `json:"data"`
	Meta struct {
		Attempt     int    `json:"attempt"`
		DeliveredTo string `json:"delivered_to"`
	} `json:"meta"`
}
