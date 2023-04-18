package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandleRoot)
	// server.Handle("/api", HandleHome)
	server.Handle("/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Liten()

}
