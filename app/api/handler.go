package api

import (
	"log"

	"github.com/gorilla/mux"

	"github.com/adeniyistephen/withkoa/business"
)

func Handle(log *log.Logger, key []byte) *mux.Router {
	r := mux.NewRouter()

	// Register passenger endpoints.
	cg := convertGroup{
		convert: business.New(log, key),
	}

	// Register convertion endpoints.
	r.HandleFunc("/auth", cg.Token).Methods("GET")
	r.HandleFunc("/naira", cg.ConvertNaira).Methods("POST")
	r.HandleFunc("/cedis", cg.ConvertCedis).Methods("POST")
	r.HandleFunc("/shillings", cg.ConvertShilling).Methods("POST")

	return r
}
