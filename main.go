package main

import (
	"fmt"
	// "log"
	"net/http"
	// "os"

	// "glifo/driver"
	// ph "glifo/handler/http"

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
	// db, err := driver.ConnectSQL("./glifo.db")

	// log a fatal error... dãh
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println(err)
	// 	os.Exit(-1)
	// }

	r := chi.NewRouter()
	
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	// sessionHandler := ph.NewSessionHandlerWith(db)

	// r.Route("/", func(router chi.Router) {
	// 	router.Mount("/sessions", sessionRouter(sessionHandler))
	// })

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// Get port config
	// port, exists := os.LookupEnv("PORT")

	conf := config.New()
	print(conf.Port)

	fmt.Println("Server listen at :3000")
	http.ListenAndServe(":3000", r)
}

// A completely separate router for posts routes
// func sessionRouter(sessionHandler *ph.Post) http.Handler {
// 	r := chi.NewRouter()
// 	// r.Get("/", sessionHandler.Fetch)
// 	r.Post("/", sessionHandler.Create)
// 	// r.Get("/{id:[0-9]+}", sessionHandler.GetByID)
// 	// r.Put("/{id:[0-9]+}", sessionHandler.Update)
// 	// r.Delete("/{id:[0-9]+}", sessionHandler.Delete)

// 	return r
// }
