package domain

import "time"

type Agent struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"created_at"`
}
