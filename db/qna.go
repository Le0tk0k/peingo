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
	qna := entity.QnA{}
	err := r.conn.Get(&qna, "SELECT * FROM qnas WHEHE id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, qnaNotFoundError
		}
		return nil, err
	}
	return &qna, nil
}

// FindQnAs はDBから回答済みのq&aを取得する
func (r *QnARepository) FindQnAs() ([]*entity.QnA, error) {
	var qnas []*entity.QnA
	err := r.conn.Select(&qnas, "SELECT * FROM qnas WHERE answer IS NOT NULL ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	return qnas, nil
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
