package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
	"github.com/pu4mane/go-docker-k8s-demo/pkg/version"
)

var isReady atomic.Bool

func initReady() {
	log.Printf("Readyz probe is negative by default...")
	time.Sleep(10 * time.Second)
	isReady.Store(true)
	log.Printf("Readyz probe is positive.")
}

func Router(buildTime, commit, release string) *mux.Router {
	initReady()

	router := mux.NewRouter()
	router.HandleFunc("/home", homeHendler(buildTime, commit, release)).Methods(http.MethodGet)
	router.HandleFunc("/healthz", healthzHendler)
	router.HandleFunc("/readyz", readyzHendler)

	return router
}

func homeHendler(buildTime, commit, release string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		info := version.Info{
			BuildTime: buildTime,
			Commit:    commit,
			Release:   release,
		}

		body, err := json.Marshal(info)
		if err != nil {
			log.Printf("Could not encode info data: %v", err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func healthzHendler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readyzHendler(w http.ResponseWriter, r *http.Request) {
	if !isReady.Load() {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
}
