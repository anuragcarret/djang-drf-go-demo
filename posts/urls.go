package posts

import (
	"github.com/anuragcarret/djang-drf-go/core/urls"
	"github.com/anuragcarret/djang-drf-go/drf/views"
	"github.com/anuragcarret/djang-drf-go/orm/db"
)

// RegisterRoutes registers the URL patterns for the posts app
func RegisterRoutes(database *db.DB) *urls.Router {
	r := urls.NewRouter()

	viewSet := NewPostViewSet(database)

	r.Get("/", views.ModelHandler(&viewSet.ModelViewSet), "post-list")
	r.Get("/{id}", views.ModelHandler(&viewSet.ModelViewSet), "post-detail")

	return r
}
