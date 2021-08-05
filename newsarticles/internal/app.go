package internal

import (
	"log"
	"ncu-main-recruitment/newsarticles/internal/handler"
	"ncu-main-recruitment/newsarticles/internal/storage"

	"github.com/gorilla/mux"

	"net/http"
)

type App struct {
	Router *mux.Router
	DB     storage.DB
}

func (a *App) Initialize() {
	a.createDatabase()
	a.routes()
}

func (a *App) PostArticles(w http.ResponseWriter, r *http.Request) {
	handler.PostArticles(a.DB, w, r)
}

func (a *App) routes() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/articles", a.PostArticles).Methods(http.MethodPost)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) createDatabase() {
	a.DB = storage.NewDB()
}
