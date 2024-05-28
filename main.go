package main

import (
	"chat/conn"
	"chat/handlers"
	"chat/middleware"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	db := conn.ConnectToDB()
	defer db.Close()

	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()

	ws := http.HandlerFunc(handlers.SocketHandler)
	serveFiles := http.HandlerFunc(handlers.HandleFiles)
	redgister:= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handlers.Redgister(db, w, r) })
    login:= http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handlers.Login(db, w, r) })
    validate:= http.HandlerFunc(handlers.Verify)


    mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.Handle("/ws", middleware.HeaderMiddleware(ws))
	mux.Handle("/", middleware.HeaderMiddleware(serveFiles))
	mux.Handle("/api/redgister", middleware.HeaderMiddleware(redgister))
	mux.Handle("/api/login", middleware.HeaderMiddleware(login))
	mux.Handle("/api/validate", middleware.HeaderMiddleware(validate))

	log.Print("running on 3000")

	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
