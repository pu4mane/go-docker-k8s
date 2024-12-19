package httpserver

import (
	"net/http"

	"github.com/pu4mane/go-docker-k8s-demo/pkg/config"
)

type Server struct {
	http.Server
}

func NewServer(cfg *config.Config, buildTime, commit, release string) *Server {
	srv := &Server{}
	srv.Addr = ":" + cfg.Http.Port
	router := Router(buildTime, commit, release)
	srv.Handler = router
	return srv
}
