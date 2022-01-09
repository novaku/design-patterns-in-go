package framework

import "net/http"

type application struct {
}

func (a *application) HandleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return http.StatusOK, "Ok"
	}

	if url == "/create/user" && method == "POST" {
		return http.StatusCreated, "User Created"
	}
	return 404, "Not Ok"
}