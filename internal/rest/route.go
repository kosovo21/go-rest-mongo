package rest

import (
	"github.com/emicklei/go-restful/v3"
)

func register(ws *restful.WebService, server *Server) {

	ws.Route(
		ws.GET("/gettoken").
			To(server.GetToken))

}
