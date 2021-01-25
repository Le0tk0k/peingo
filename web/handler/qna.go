package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Le0tk0k/peingo/config"
	"github.com/Le0tk0k/peingo/domain/entity"
	"github.com/Le0tk0k/peingo/usecase"
	"github.com/labstack/echo"
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
		ID:       qna.ID,
		Question: qna.Question,
		Answer:   qna.Answer,
	})
}

func (h *QnaHandler) CreateQuestion(c echo.Context) error {
	q := entity.QnA{}
	c.Bind(&q)
	err := h.qnaUseCase.CreateQuestion(&q)

	if err := notifyToSlack(q.Question); err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

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

func (h *QnaHandler) Admin(c echo.Context) error {
	qnas, err := h.qnaUseCase.GetAllQnAs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, &qnasRes{
		QnAs: toQnAJSON(qnas),
	})
}

func (h *QnaHandler) AdminQnA(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	qna, err := h.qnaUseCase.GetQnA(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, &qnaRes{
		ID:       qna.ID,
		Question: qna.Question,
		Answer:   qna.Answer,
	})
}

func notifyToSlack(question string) error {
	jsonStr := `{"text":"` + question + `"}`
	incomingURL := config.WebHook()

	req, err := http.NewRequest(
		"POST",
		incomingURL,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

func toQnAJSON(qnas []*entity.QnA) []*qnaRes {
	qnasres := make([]*qnaRes, len(qnas))

	for i, qna := range qnas {
		qnasres[i] = &qnaRes{
			ID:       qna.ID,
			Question: qna.Question,
			Answer:   qna.Answer,
		}
	}
	return qnasres
}

type qnaRes struct {
	ID       int     `json:"id"`
	Question string  `json:"question"`
	Answer   *string `json:"answer"`
}

type qnasRes struct {
	QnAs []*qnaRes `json:"qnas"`
}
