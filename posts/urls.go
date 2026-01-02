package posts

import (
	"net/http"

	"github.com/anuragcarret/djang-drf-go/core/urls"
	"github.com/anuragcarret/djang-drf-go/drf/authentication"
	"github.com/anuragcarret/djang-drf-go/drf/views"
	"github.com/anuragcarret/djang-drf-go/orm/db"
)

// Simple token store for demo
type DemoTokenStore struct {
	tokens map[string]interface{}
}

func NewDemoTokenStore() *DemoTokenStore {
	return &DemoTokenStore{
		tokens: make(map[string]interface{}),
	}
}

func (s *DemoTokenStore) ValidateToken(token string) (interface{}, error) {
	user, exists := s.tokens[token]
	if !exists {
		return nil, authentication.ErrInvalidToken
	}
	return user, nil
}

func (s *DemoTokenStore) CreateToken(user interface{}) (string, error) {
	return "demo-token-456", nil
}

func (s *DemoTokenStore) RevokeToken(token string) error {
	delete(s.tokens, token)
	return nil
}

// RegisterRoutes sets up all post endpoints
func RegisterRoutes(database *db.DB) *urls.Router {
	r := urls.NewRouter()

	// Authentication middleware
	tokenStore := NewDemoTokenStore()
	authMiddleware := authentication.AuthenticationMiddleware([]authentication.Authenticator{
		authentication.NewTokenAuthentication(tokenStore),
	})

	// POST endpoints with different features

	// 1. Basic list with pagination & filtering
	// GET /posts/?page=1&page_size=10&title__contains=test
	r.Get("/posts/", wrapHandler(NewPostListView(database)), "post-list")

	// 2. Create post (authenticated)
	// POST /posts/
	r.Post("/posts/", wrapWithAuth(NewPostCreateView(database), authMiddleware), "post-create")

	// 3. Detail view (get/update/delete) - authenticated
	detailView := NewPostDetailView(database)
	r.Get("/posts/{id:int}/", wrapWithAuth(detailView, authMiddleware), "post-detail")
	r.Methods([]string{"PUT"}, "/posts/{id:int}/", wrapWithAuth(detailView, authMiddleware), "post-update")
	r.Methods([]string{"PATCH"}, "/posts/{id:int}/", wrapWithAuth(detailView, authMiddleware), "post-partial-update")
	r.Methods([]string{"DELETE"}, "/posts/{id:int}/", wrapWithAuth(detailView, authMiddleware), "post-delete")

	// 4. Throttled endpoint (rate limited)
	// GET /posts/throttled/ - max 10 req/min for anonymous, 100/min for users
	r.Get("/posts/throttled/", wrapHandler(NewThrottledPostView(database)), "post-throttled")

	// 5. Published posts with search
	// GET /posts/published/?search=keyword&ordering=-created_at
	r.Get("/posts/published/", wrapHandler(NewPublishedPostsView(database)), "post-published")

	return r
}

// Helper to wrap view into http.Handler
func wrapHandler(view interface{}) http.Handler {
	return views.Handler(view)
}

// Helper to wrap view with authentication middleware
func wrapWithAuth(view interface{}, authMiddleware func(http.Handler) http.Handler) http.Handler {
	handler := views.Handler(view)
	return authMiddleware(handler)
}
