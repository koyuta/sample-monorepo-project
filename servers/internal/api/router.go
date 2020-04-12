package application

import (
	"net/http"

	"servers/interfaces/handler"
	"servers/interfaces/handler/api/v1"
	"servers/interfaces/middleware"
	"servers/pkg/rds"
	"servers/pkg/router"
	"servers/repository"
	"servers/usecase"

	"github.com/go-chi/chi"
)

// BuildRouter builds a http router.
func BuildRouter(r rds.RDS) http.Handler {
	var chiContext = router.NewChiRouter()

	// repositories
	var (
		usersRepository = repository.NewUsers(r)
	)

	// usecases
	var (
		usersUsecase = usecase.NewUsers(usersRepository)
	)

	// handlers
	var (
		users = api.NewUsers(chiContext, usersUsecase)
		root  = handler.Root{}
	)

	var router = chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Use(middleware.Common)
		r.Use(middleware.Recoverer)
		r.Get("/", users.Get)
		r.Post("/", users.Post)
	})
	router.Get("/healthz", root.Healthz)

	return router
}
