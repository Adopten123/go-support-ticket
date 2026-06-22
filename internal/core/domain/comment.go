package domain

import "time"

type AuthorType string

const (
	AuthorClient AuthorType = "client"
	AuthorAgent  AuthorType = "agent"
)

type Comment struct {
	ID         int64      `json:"id"`
	TicketID   int64      `json:"ticket_id"`
	AuthorID   int64      `json:"author_id"`
	AuthorType AuthorType `json:"author_type"`
	Text       string     `json:"text"`
	CreatedAt  time.Time  `json:"created_at"`
}
