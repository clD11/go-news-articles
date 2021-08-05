package handler

import (
	"encoding/json"
	"ncu-main-recruitment/newsarticles/internal/model"
	"ncu-main-recruitment/newsarticles/internal/storage"

	"net/http"
)

// POST /articles
func PostArticles(db storage.DB, w http.ResponseWriter, r *http.Request) {
	var article model.Articles
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Could not decode request body")
		return
	}
	defer r.Body.Close()

	// store the articles against the user id
	if err := db.Insert(article.UserID, article); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Could not insert article")
		return
	}

	writeResponse(w, http.StatusCreated, article)
}

func writeErrorResponse(w http.ResponseWriter, code int, message string) {
	writeResponse(w, code, map[string]string{"error": message})
}

func writeResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
