package router

import (
	"github.com/Le0tk0k/peingo/usecase"
	"github.com/Le0tk0k/peingo/web/handler"
	"github.com/labstack/echo"
)

func Router(u usecase.QnAUseCase) *echo.Echo {
	e := echo.New()

	qnaHandler := handler.NewQnAHandler(u)

	e.GET("/qnas", qnaHandler.GetQnAs)
	e.GET("/qnas/:id", qnaHandler.GetQnA)
	e.POST("/qnas", qnaHandler.CreateQuestion)
	e.PUT("/qnas/:id", qnaHandler.CreateAnswer)

	return e
}
