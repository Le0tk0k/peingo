package entity

import "time"

type QnA struct {
	ID        int
	Question  string
	Answer    *string
	CreatedAt time.Time
}

func NewQnA(id int, question string) *QnA {
	return &QnA{
		ID:        id,
		Question:  question,
		CreatedAt: time.Now(),
	}
}
