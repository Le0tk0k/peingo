package handler

import (
	"github.com/Le0tk0k/peingo/domain/entity"
	"github.com/Le0tk0k/peingo/usecase"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

type QnaHandler struct {
	qnaUseCase usecase.QnAUseCase
}

func NewQnAHandler(u usecase.QnAUseCase) *QnaHandler {
	return &QnaHandler{qnaUseCase: u}
}

func (h *QnaHandler) GetQnAs(c echo.Context) error {
	qnas, err := h.qnaUseCase.GetQnAs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, &qnasRes{
		QnAs: toQnAJSON(qnas),
	})
}

func (h *QnaHandler) GetQnA(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	qna, err := h.qnaUseCase.GetQnA(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, &qnaRes{
		ID:        qna.ID,
		Question:  qna.Question,
		Answer:    *qna.Answer,
		CreatedAt: qna.CreatedAt,
	})
}

func (h *QnaHandler) CreateQuestion(c echo.Context) error {
	q := entity.QnA{}
	c.Bind(&q)
	err := h.qnaUseCase.CreateQuestion(&q)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, q)
}

func (h *QnaHandler) CreateAnswer(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	a := entity.QnA{}
	a.ID = id
	c.Bind(&a)
	err := h.qnaUseCase.CreateAnswer(&a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, a)
}

func toQnAJSON(qnas []*entity.QnA) []*qnaRes {
	qnasres := make([]*qnaRes, len(qnas))

	for i, qna := range qnas {
		qnasres[i] = &qnaRes{
			ID:        qna.ID,
			Question:  qna.Question,
			Answer:    *qna.Answer,
			CreatedAt: qna.CreatedAt,
		}
	}
	return qnasres
}

type qnaRes struct {
	ID        int       `json:"id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
}

type qnasRes struct {
	QnAs []*qnaRes `json:"qnas"`
}

type AddQuestionsReq struct {
	Body string `json:"body"`
}

type AddAnswerReq struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}
