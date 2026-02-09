package server

import (
	"net/http"
	"net/http/httptest"
)

func startTestServer(hub *Hub) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ServeWS(hub, w, r)
	}))
}

func wsURL(server *httptest.Server, path string) string {
	return "ws" + server.URL[len("http"):] + path
}
