package framework

import "net/http"

type nginx struct {
	Application       *application
	MaxAllowedRequest int
	RateLimiter       map[string]int
}

func NewNginxServer() *nginx {
	return &nginx{
		Application:       &application{},
		MaxAllowedRequest: 2,
		RateLimiter:       make(map[string]int),
	}
}

func (n *nginx) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return http.StatusForbidden, "Not Allowed"
	}
	return n.Application.HandleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.RateLimiter[url] == 0 {
		n.RateLimiter[url] = 1
	}
	if n.RateLimiter[url] > n.MaxAllowedRequest {
		return false
	}
	n.RateLimiter[url] = n.RateLimiter[url] + 1
	return true
}