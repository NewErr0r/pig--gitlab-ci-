package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/rs/zerolog/log"
)

func RunServer() {
	log.Info().Msg("Running PIG Dummy Service on port 8000")
	r := mux.NewRouter()
	r.HandleFunc("/internal/healthz", healthcheck)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("resources/"))))
	r.Use(loggingMiddleware)
	http.Handle("/", r)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Debug().Msg(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	var response API
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Server", "PIG Dummy Service")
	if r.Method != http.MethodGet {
		response = API{
			Code:    405,
			Message: "Method not allowed",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		response = API{
			Code:    200,
			Message: "healthy",
		}
		w.WriteHeader(http.StatusOK)
	}
	js, err := json.Marshal(response)
	if err != nil {
		log.Debug().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
