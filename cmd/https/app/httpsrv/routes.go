package httpsrv

import "net/http"

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleRoot())
}

func (s *server) handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
