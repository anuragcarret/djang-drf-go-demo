package accounts

import (
	"github.com/anuragcarret/djang-drf-go/core/urls"
	"github.com/anuragcarret/djang-drf-go/drf/views"
	"github.com/anuragcarret/djang-drf-go/orm/db"
)

// RegisterRoutes registers the URL patterns for the accounts app
func RegisterRoutes(database *db.DB) *urls.Router {
	r := urls.NewRouter()

	viewSet := NewAccountViewSet(database)

	// Register ModelHandler for the viewset
	r.Post("/register/", views.ModelHandler(viewSet), "account-register")
	r.Get("/", views.ModelHandler(viewSet), "account-list")
	r.Get("/{id}", views.ModelHandler(viewSet), "account-detail")

	return r
}
