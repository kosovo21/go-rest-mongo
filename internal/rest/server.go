package rest

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/kosovo21/go-rest-mongo/internal/data"
)

type Server struct {
	userDao data.DAO
}

func New(userDao data.DAO) *Server {
	server := &Server{
		userDao: userDao,
	}

	ws := new(restful.WebService).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	register(ws, server)
	restful.Add(ws)

	return server
}

func (Server) ListenAndServe() error {
	return http.ListenAndServe(":8080", restful.DefaultContainer)
}
