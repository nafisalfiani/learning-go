package main

import (
	"data-center/config"
	"data-center/handler"
	"data-center/repository"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func accessLog(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		log.Printf("%s -> HTTP request sent to %s %s", time.Now().Format("2006/01/02 - 15:04:05"), r.Method, r.URL.Path)
		next(w, r, p)
	}
}

func main() {
	log.Println("Starting server...")

	cfg, err := config.InitEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.InitSql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// init customer repository
	userRepo := repository.InitUser(db)

	// init validator
	validator := validator.New(validator.WithRequiredStructEnabled())

	// init handler
	handler := handler.Init(userRepo, validator)

	// registering routes
	router := httprouter.New()
	router.POST("/register", accessLog(handler.Register))
	router.POST("/login", accessLog(handler.Login))

	log.Println("Server is up")
	log.Fatal(http.ListenAndServe(":8080", router))
}
