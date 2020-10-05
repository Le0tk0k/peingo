package entity

import "time"

type QnA struct {
	ID        string
	Question  string
	Answer    string
	CreatedAt time.Time
}

func NewQnA(id, question, answer) *QnA {
	return &QnA{
		ID:        id,
		Question:  question,
		Answer:    answer,
		CreatedAt: time.Now(),
	}
}
