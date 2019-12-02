package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"glifo/driver"
	ph "glifo/handler/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Glifo CRUD operations

// POST /sessions
// GET /sessions
// GET /session/{id}
// PUT /session/{id}
// DELETE /session/{id}

func main() {
	db, err := driver.ConnectSQL("./glifo.db")

	// log a fatal error... dãh
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	sessionHandler := ph.NewPostHandler(db)

	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", sessionRouter(sessionHandler))
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}

// A completely separate router for posts routes
func sessionRouter(sessionHandler *ph.Post) http.Handler {
	r := chi.NewRouter()
	// r.Get("/", sessionHandler.Fetch)
	r.Post("/", sessionHandler.Create)
	// r.Get("/{id:[0-9]+}", sessionHandler.GetByID)
	// r.Put("/{id:[0-9]+}", sessionHandler.Update)
	// r.Delete("/{id:[0-9]+}", sessionHandler.Delete)

	return r
}
