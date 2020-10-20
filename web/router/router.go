package router

import (
	"github.com/Le0tk0k/peingo/usecase"
	"github.com/Le0tk0k/peingo/web/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router(u usecase.QnAUseCase) *echo.Echo {
	//t := &handler.Template{
	//	Templates: template.Must(template.ParseGlob("templates/*.html")),
	//}

	e := echo.New()

	//e.Renderer = t

	e.Use(middleware.CORS())

	qnaHandler := handler.NewQnAHandler(u)

	e.GET("/qnas", qnaHandler.GetQnAs)
	e.GET("/qnas/:id", qnaHandler.GetQnA)
	e.POST("/qnas", qnaHandler.CreateQuestion)
	e.PUT("/qnas/:id", qnaHandler.CreateAnswer)

	//e.GET("/admin", qnaHandler.Admin)
	//e.GET("admin/:id", qnaHandler.AdminQnA)

	return e
}
