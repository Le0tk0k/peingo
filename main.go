package main

import (
	"log"

	"github.com/Le0tk0k/peingo/db"
	"github.com/Le0tk0k/peingo/usecase"
	"github.com/Le0tk0k/peingo/web/router"
	"github.com/joho/godotenv"
)

func readEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := readEnv(); err != nil {
		log.Fatal(err)
	}

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
