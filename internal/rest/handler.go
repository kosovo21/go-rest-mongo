package rest

import (
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

var (
	TOKEN_EXPIRATION_TIME = time.Duration(1) * time.Hour
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
		return
	}

	// TODO: decrypt password
	if password != user.Password {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TOKEN_EXPIRATION_TIME).Unix(),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secret123"))
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.WriteAsJson(signedToken)
}

func (server *Server) Booking(req *restful.Request, resp *restful.Response) {
	resp.WriteHeader(http.StatusAccepted)
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}
