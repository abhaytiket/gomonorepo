package main

import (
	"net/http"

	"github.com/abhaytiket/gomonorepo/lib/msg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(msg.AddMessage("hello app1")))
	})
	http.ListenAndServe(":3000", r)
}
