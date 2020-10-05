package entity

import "time"

type QnA struct {
	ID        int
	Question  string
	Answer    string
	CreatedAt time.Time
}

func NewQnA(id int, question, answer string) *QnA {
	return &QnA{
		ID:        id,
		Question:  question,
		Answer:    answer,
		CreatedAt: time.Now(),
	}
}
