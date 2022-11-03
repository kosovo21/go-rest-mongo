package server

import (
	"io"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/sirupsen/logrus"
)

func RunServer() error {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	restful.Add(ws)
	logrus.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
