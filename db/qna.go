package db

import (
	"github.com/Le0tk0k/peingo/domain/entity"
	"github.com/Le0tk0k/peingo/domain/repository"
	"github.com/jmoiron/sqlx"
)

type QnARepository struct {
	conn *sqlx.DB
}

func NewQnARepository(conn *sqlx.DB) repository.QnARepository {
	return &QnARepository{conn: conn}
}

func (r *QnARepository) FindByID(id int) (*entity.QnA, error) {
	qna := entity.QnA{}
	err := r.conn.Get(&qna, "SELECT * FROM qnas WHEHE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &qna, nil
}

func (r *QnARepository) FindQnAs() ([]*entity.QnA, error) {
	var qnas []*entity.QnA
	err := r.conn.Select(&qnas, "SELECT * FROM qnas ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	return qnas, nil

}

func (r *QnARepository) StoreQuestion(body string) error {
	tx, err := r.conn.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO qandas (question) VALUES (?)", body)
	if err != nil {
		return err
	}
	_ = tx.Commit()
	return nil
}

func (r *QnARepository) StoreAnswer(id int, body string) error {
	tx, err := r.conn.Beginx()
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE qandas SET answer = ? WHERE id = ?", body, id)
	if err != nil {
		return err
	}
	_ = tx.Commit()
	return nil
}
