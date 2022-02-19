package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adeniyistephen/withkoa/business"
	"github.com/pkg/errors"
)

type convertGroup struct {
	convert business.NewConvert
}

// Handle Service API for converting to naira
func (cg convertGroup) ConvertNaira(w http.ResponseWriter, r *http.Request) {
	var cm business.ConvertMoney

	// Decode the request body into the struct.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cm); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	matches := r.Header.Get("Authorization")
	if matches == "" {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	// Convert to naira
	cvm, err := cg.convert.ConvertToNaira(cm, matches)
	if err != nil {
		log.Panic(errors.Wrap(err, "Could not convert to Naira"))
	}

	// Return the converted amount
	respondWithJSON(w, http.StatusCreated, cvm)
}

// Handle Service API for converting to cedis
func (cg convertGroup) ConvertCedis(w http.ResponseWriter, r *http.Request) {
	var cm business.ConvertMoney

	// Decode the request body into the struct.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cm); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Convert to cedis
	cvm, err := cg.convert.ConvertToCedis(cm)
	if err != nil {
		log.Panic(errors.Wrap(err, "Could not convert to Cedis"))
	}

	// Return the converted amount
	respondWithJSON(w, http.StatusCreated, cvm)
}

// Handle Service API for converting to shilling
func (cg convertGroup) ConvertShilling(w http.ResponseWriter, r *http.Request) {
	var cm business.ConvertMoney

	// Decode the request body into the struct.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cm); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Convert to shillinh
	cvm, err := cg.convert.ConvertToShillings(cm)
	if err != nil {
		log.Panic(errors.Wrap(err, "Could not convert to Shilling"))
	}

	// Return the converted amount
	respondWithJSON(w, http.StatusCreated, cvm)
}

func (cg convertGroup) Token(w http.ResponseWriter, r *http.Request) {
	tkn := cg.convert.Token()
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Auth successful", "token": tkn})
}

// Return an error message as a JSON object
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Return a JSON object
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
