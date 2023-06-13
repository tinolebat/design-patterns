package main

//Real Object
type Application struct {
}

func (a *Application) handleRequest(url string, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User created"
	}
	return 404, "Not found"
}
