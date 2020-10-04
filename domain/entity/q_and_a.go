package entity

import "time"

type QAndA struct {
	ID        string
	Question  string
	Answer    string
	CreatedAt time.Time
}

func NewQAndA(id, question, answer) *QAndA {
	return &QAndA{
		ID:        id,
		Question:  question,
		Answer:    answer,
		CreatedAt: time.Now(),
	}
}
