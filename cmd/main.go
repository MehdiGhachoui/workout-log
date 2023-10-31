package main

import (
	"log"

	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mehdighachoui/workout-log/internal/templates/pages"
	// "github.com/mehdighachoui/workout-log/db"
	// "github.com/mehdighachoui/workout-log/pkg"
)

func main() {
	// _, err := db.CreatePostgresConnection()
	// if err != nil {
	// 	log.Fatal("Database connection error:", err)
	// }
	// log.Println("Database connected")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", templ.Handler(pages.Home()).ServeHTTP)

	// pkg.WorkoutHTTP(app)
	log.Fatal(http.ListenAndServe(":8000", r))
}
