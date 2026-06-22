package domain

import "time"

type TicketStatus string

const (
	StatusOpen       TicketStatus = "open"
	StatusInProgress TicketStatus = "in_progress"
	StatusResolved   TicketStatus = "resolved"
)

type Ticket struct {
	ID          int64        `json:"id"`
	ClientID    int64        `json:"client_id"`
	AgentID     *int64       `json:"agent_id"`
	Subject     string       `json:"subject"`
	Description string       `json:"description"`
	Status      TicketStatus `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	ResolvedAt  *time.Time   `json:"resolved_at"`
}
