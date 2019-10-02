package web

import (
	"net/http"
)

func (srv *server) routes() http.Handler {
	//Declare web routing table at here.

	return srv.router
}
