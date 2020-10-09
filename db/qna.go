package db

import (
	"database/sql"
	"errors"
	"time"

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
	err := r.conn.Get(&dto, "SELECT * FROM qnas WHEHE id = ?", id)
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
	err := r.conn.Select(&dtos, "SELECT * FROM qnas WHERE answer IS NOT NULL ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	return r.dtosToQnA(dtos), nil
}

// StoreQuestion は質問を新規保存する
func (r *QnARepository) StoreQuestion(body string) error {
	tx, err := r.conn.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO qandas (question) VALUES (?)", body)
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
func (r *QnARepository) StoreAnswer(id int, body string) error {
	tx, err := r.conn.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE qandas SET answer = ? WHERE id = ?", body, id)
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
		ID:        dto.ID,
		Question:  dto.Question,
		Answer:    dto.Answer,
		CreatedAt: dto.CreatedAt,
	}
}

func (r *QnARepository) dtosToQnA(dtos []*qnaDTO) []*entity.QnA {
	qnas := make([]*entity.QnA, len(dtos))

	for i, dto := range dtos {
		qnas[i] = &entity.QnA{
			ID:        dto.ID,
			Question:  dto.Question,
			Answer:    dto.Answer,
			CreatedAt: dto.CreatedAt,
		}
	}
	return qnas
}

type qnaDTO struct {
	ID        int       `db:"id"`
	Question  string    `db:"question"`
	Answer    *string   `db:"string"`
	CreatedAt time.Time `db:"created_at"`
}
