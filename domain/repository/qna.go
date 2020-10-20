package repository

import "github.com/Le0tk0k/peingo/domain/entity"

type QnARepository interface {
	FindByID(id int) (*entity.QnA, error)
	FindQnAs() ([]*entity.QnA, error)
	//FindAllQnAs() ([]*entity.QnA, error)
	StoreQuestion(qna *entity.QnA) error
	StoreAnswer(qna *entity.QnA) error
}
