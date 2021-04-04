package httpsrv

import "net/http"

type server struct {
	router *http.ServeMux
}

func NewServer() *server {
	return &server{
		router: http.NewServeMux(),
	}
}
