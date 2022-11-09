package rest

import (
	"net/http"
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/golang-jwt/jwt"
)

func register(ws *restful.WebService, server *Server) {

	ws.Route(
		ws.GET("/gettoken").
			To(server.GetToken))

	ws.Route(
		ws.POST("/booking").
			Filter(jwtAuth).
			To(server.Booking))

}

func jwtAuth(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	authHeader := req.HeaderParameter("Authorization")

	if !validJWT(authHeader) {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func validJWT(authHeader string) bool {
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return false
	}

	jwtToken := strings.Split(authHeader, " ")
	if len(jwtToken) < 2 {
		return false
	}
	parts := strings.Split(jwtToken[1], ".")
	err := jwt.SigningMethodHS512.Verify(strings.Join(parts[0:2], "."), parts[2], []byte("secret123"))
	return err != nil
}
