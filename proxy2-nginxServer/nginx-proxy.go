package main

type NginxProxy struct {
	app               *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxProxyServer() *NginxProxy {
	return &NginxProxy{
		app:               &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *NginxProxy) handleRequestAtProxy(url, method string) (int, string) {
	allowed := n.checkRateLimiter(url)
	if !allowed {
		return 203, "Not allowed"
	}
	return n.app.handleRequest(url, method)

}

func (n *NginxProxy) checkRateLimiter(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true

}
