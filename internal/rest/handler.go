package rest

import (
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/sirupsen/logrus"
)

func (server *Server) GetToken(req *restful.Request, resp *restful.Response) {
	logrus.Info("request incoming...")
	username, password, ok := req.Request.BasicAuth()
	if !ok {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := server.userDao.FindUser(username)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user == nil {
		resp.WriteHeader(http.StatusBadRequest)
	}

	if password != user.Password {
		resp.WriteHeader(http.StatusBadRequest)
	}

	resp.WriteHeader(http.StatusOK)
}
