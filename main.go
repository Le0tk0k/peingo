package main

import (
	"log"

	"github.com/Le0tk0k/peingo/db"
	"github.com/Le0tk0k/peingo/usecase"
	"github.com/Le0tk0k/peingo/web/router"
)

func main() {
	conn, err := db.NewDb()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	qnaRepo := db.NewQnARepository(conn)
	qnaUC := usecase.NewQnAUseCase(qnaRepo)
	r := router.Router(qnaUC)

	r.Logger.Fatal(r.Start(":1323"))
}
