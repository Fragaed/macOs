package handler

import (
	"Fragaed/internal/service"
	"Fragaed/static"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	log.Print("Роутер запущен")
	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Post("/", h.Create)
			r.Post("/list", h.List)
			r.Put("/{id}", h.Get)
			r.Post("/{id}", h.Update)
			r.Delete("/{id}", h.Delete)
		})

	})
	r.Get("/swagger", static.SwaggerUI)
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
	})
	return r
}
