package accounts

import (
	"net/http"

	"github.com/anuragcarret/djang-drf-go/core/urls"
	"github.com/anuragcarret/djang-drf-go/drf/authentication"
	"github.com/anuragcarret/djang-drf-go/drf/views"
	"github.com/anuragcarret/djang-drf-go/orm/db"
)

// RegisterRoutes registers the URL patterns for the accounts app
func RegisterRoutes(database *db.DB) *urls.Router {
	r := urls.NewRouter()

	// Create authentication middleware
	tokenStore := NewDemoTokenStore()
	authMiddleware := authentication.AuthenticationMiddleware([]authentication.Authenticator{
		authentication.NewTokenAuthentication(tokenStore),
	})

	// Public endpoints (no auth required)
	accountViewSet := NewAccountViewSet(database)
	r.Get("/", views.ModelHandler(accountViewSet), "account-list")
	r.Post("/", views.ModelHandler(accountViewSet), "account-create")

	// Authenticated endpoints
	authViewSet := NewAuthenticatedAccountView(database)
	r.Get("/me/", wrapWithAuth(authViewSet, authMiddleware), "account-me")

	return r
}

// Helper to wrap view with authentication middleware
func wrapWithAuth(view interface{}, authMiddleware func(http.Handler) http.Handler) http.Handler {
	handler := views.ModelHandler(view.(views.APIViewSet))
	return authMiddleware(handler)
}
