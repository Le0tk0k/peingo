package db

import (
	"database/sql"
	"errors"
	"github.com/Le0tk0k/peingo/domain/entity"
	"github.com/Le0tk0k/peingo/domain/repository"
	"github.com/jmoiron/sqlx"
)

type QnARepository struct {
	conn *sqlx.DB
}

var qnaNotFoundError = errors.New("qna not found")

// NewQnARepository はQnARepositoryのポインタを返す
func NewQnARepository(conn *sqlx.DB) repository.QnARepository {
	return &QnARepository{conn: conn}
}

// FindByID は指定されたIDを持つq&aをDBから取得する
func (r *QnARepository) FindByID(id int) (*entity.QnA, error) {
	var dto qnaDTO
	err := r.conn.Get(&dto, "SELECT id, question, answer FROM qnas WHERE id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, qnaNotFoundError
		}
		return nil, err
	}
	return r.dtoToQnA(dto), nil
}

// FindQnAs はDBから回答済みのq&aを取得する
func (r *QnARepository) FindQnAs() ([]*entity.QnA, error) {
	var dtos []*qnaDTO
	err := r.conn.Select(&dtos, "SELECT id, question, answer FROM qnas WHERE answer IS NOT NULL ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	return r.dtosToQnA(dtos), nil
}

// StoreQuestion は質問を新規保存する
func (r *QnARepository) StoreQuestion(qna *entity.QnA) error {
	tx, err := r.conn.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO qnas (question) VALUES (?)", qna.Question)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// StoreAnswer は指定されたIDの質問の回答を新規保存する
func (r *QnARepository) StoreAnswer(qna *entity.QnA) error {
	tx, err := r.conn.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE qnas SET answer = ? WHERE id = ?", qna.Answer, qna.ID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r *QnARepository) dtoToQnA(dto qnaDTO) *entity.QnA {
	return &entity.QnA{
		ID:       dto.ID,
		Question: dto.Question,
		Answer:   dto.Answer,
	}
}

func (r *QnARepository) dtosToQnA(dtos []*qnaDTO) []*entity.QnA {
	qnas := make([]*entity.QnA, len(dtos))

	for i, dto := range dtos {
		qnas[i] = &entity.QnA{
			ID:       dto.ID,
			Question: dto.Question,
			Answer:   dto.Answer,
		}
	}
	return qnas
}

type qnaDTO struct {
	ID       int     `db:"id"`
	Question string  `db:"question"`
	Answer   *string `db:"answer"`
}
