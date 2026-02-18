package transport

import (
	"net/http"

	"github.com/Lucky123-cloud/Domain_Driven_Design/recommendation/internal/recommendation"
	"github.com/gorilla/mux"
)

func NewMux(recHandler recommendation.Handler) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/recommendation", recHandler.GetRecommendation).Methods(http.MethodGet)
	return m
}
