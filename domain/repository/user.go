package repository

import "github.com/Le0tk0k/peingo/domain/entity"

type UserRepositry interface {
	FindByID(id int) (*entity.QnA, error)
	FindQnAs() (*[]entity.QnA, error)
	StoreQuestion(body string) error
	StoreAnswer(id int, body string) error
}
