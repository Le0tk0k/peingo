package usecase

import (
	"github.com/Le0tk0k/peingo/domain/entity"
	"github.com/Le0tk0k/peingo/domain/repository"
)

type QnAUseCase interface {
	GetQnA(id int) (*entity.QnA, error)
	GetQnAs() ([]*entity.QnA, error)
	GetAllQnAs() ([]*entity.QnA, error)
	CreateQuestion(qna *entity.QnA) error
	CreateAnswer(qna *entity.QnA) error
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

func (u *qnaUseCase) GetAllQnAs() ([]*entity.QnA, error) {
	qnas, err := u.qnaRepository.FindAllQnAs()
	if err != nil {
		return nil, err
	}
	return qnas, err
}

func (u *qnaUseCase) CreateQuestion(qna *entity.QnA) error {
	err := u.qnaRepository.StoreQuestion(qna)
	if err != nil {
		return err
	}
	return nil
}

func (u *qnaUseCase) CreateAnswer(qna *entity.QnA) error {
	err := u.qnaRepository.StoreAnswer(qna)
	if err != nil {
		return err
	}
	return nil
}
