package usecase

import (
	"github.com/Le0tk0k/peingo/domain/entity"
	"github.com/Le0tk0k/peingo/domain/repository"
)

type QnAUseCase interface {
	GetQnA(id int) (*entity.QnA, error)
	GetQnAs() ([]*entity.QnA, error)
	CreateQuestion(body string) error
	CreateAnswer(id int, body string) error
}

type qnaUseCase struct {
	qnaRepository repository.QnARepository
}

func NewQnAUseCase(r repository.QnARepository) QnAUseCase {
	return &qnaUseCase{qnaRepository: r}
}

func (u *qnaUseCase) GetQnA(id int) (*entity.QnA, error) {
	qna, err := u.qnaRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return qna, nil
}

func (u *qnaUseCase) GetQnAs() ([]*entity.QnA, error) {
	qnas, err := u.qnaRepository.FindQnAs()
	if err != nil {
		return nil, err
	}
	return qnas, err
}

func (u *qnaUseCase) CreateQuestion(body string) error {
	err := u.qnaRepository.StoreQuestion(body)
	if err != nil {
		return err
	}
	return nil
}

func (u *qnaUseCase) CreateAnswer(id int, body string) error {
	err := u.qnaRepository.StoreAnswer(id, body)
	if err != nil {
		return err
	}
	return nil
}
