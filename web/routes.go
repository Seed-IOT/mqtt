package web

import (
	"net/http"
)

func (srv *server) routes() http.Handler {
	//Declare web routing table at here.

	srv.router.POST("protocol/create", srv.CreateProtocol)

	return srv.router
}
