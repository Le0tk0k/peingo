package entity

import "time"

type QnA struct {
	ID        int
	Question  string
	Answer    *string
	CreatedAt time.Time
}
