package web

import (
	"github.com/Le0tk0k/peingo/usecase"
	"github.com/Le0tk0k/peingo/web/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router(u usecase.QnAUseCase) *echo.Echo {

	e := echo.New()

	e.Use(middleware.CORS())

	qnaHandler := handler.NewQnAHandler(u)

	e.GET("/qnas", qnaHandler.GetQnAs)
	e.GET("/qnas/:id", qnaHandler.GetQnA)
	e.POST("/qnas", qnaHandler.CreateQuestion)

	admin := e.Group("/admin")
	admin.Use(middleware.BasicAuth(validateUser))
	admin.GET("", qnaHandler.Admin)
	admin.GET("/:id", qnaHandler.AdminQnA)
	admin.PUT("/:id", qnaHandler.CreateAnswer)

	return e
}
