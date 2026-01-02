package posts

import (
	"github.com/anuragcarret/djang-drf-go/drf/filters"
	"github.com/anuragcarret/djang-drf-go/drf/pagination"
	"github.com/anuragcarret/djang-drf-go/drf/throttling"
	"github.com/anuragcarret/djang-drf-go/drf/views"
	"github.com/anuragcarret/djang-drf-go/orm/db"
)

// PostListView - List posts with pagination and filtering
type PostListView struct {
	views.ListAPIView[*Post]
	Paginator      pagination.Paginator
	FilterBackend  *filters.DjangoFilterBackend
	SearchFilter   *filters.SearchFilter
	OrderingFilter *filters.OrderingFilter
}

func NewPostListView(database *db.DB) *PostListView {
	return &PostListView{
		ListAPIView: views.ListAPIView[*Post]{
			GenericAPIView: views.GenericAPIView[*Post]{
				BaseAPIView: views.BaseAPIView{
					PermissionClasses: []views.Permission{
						&views.AllowAny{}, // Public listing
					},
				},
				DB: database,
			},
		},
		Paginator:      pagination.NewPageNumberPagination(),
		FilterBackend:  filters.NewDjangoFilterBackend([]string{"title", "author_id", "published"}),
		SearchFilter:   filters.NewSearchFilter([]string{"title", "content"}),
		OrderingFilter: filters.NewOrderingFilter([]string{"title", "created_at", "updated_at"}),
	}
}

// PostCreateView - Create new post (requires authentication)
type PostCreateView struct {
	views.CreateAPIView[*Post]
}

func NewPostCreateView(database *db.DB) *PostCreateView {
	return &PostCreateView{
		CreateAPIView: views.CreateAPIView[*Post]{
			GenericAPIView: views.GenericAPIView[*Post]{
				BaseAPIView: views.BaseAPIView{
					PermissionClasses: []views.Permission{
						&views.IsAuthenticated{}, // Must be logged in to create
					},
				},
				DB: database,
			},
		},
	}
}

// PostDetailView - Get/Update/Delete specific post
type PostDetailView struct {
	views.RetrieveUpdateDestroyAPIView[*Post]
}

func NewPostDetailView(database *db.DB) *PostDetailView {
	return &PostDetailView{
		RetrieveUpdateDestroyAPIView: views.RetrieveUpdateDestroyAPIView[*Post]{
			GenericAPIView: views.GenericAPIView[*Post]{
				BaseAPIView: views.BaseAPIView{
					PermissionClasses: []views.Permission{
						&views.IsAuthenticated{}, // Auth required for modify/delete
					},
				},
				DB: database,
			},
		},
	}
}

// ThrottledPostView - Rate-limited post listing
type ThrottledPostView struct {
	views.ListAPIView[*Post]
	Throttles []throttling.Throttle
}

func NewThrottledPostView(database *db.DB) *ThrottledPostView {
	storage := throttling.NewInMemoryThrottleStorage()

	return &ThrottledPostView{
		ListAPIView: views.ListAPIView[*Post]{
			GenericAPIView: views.GenericAPIView[*Post]{
				BaseAPIView: views.BaseAPIView{
					PermissionClasses: []views.Permission{
						&views.AllowAny{},
					},
				},
				DB: database,
			},
		},
		Throttles: []throttling.Throttle{
			throttling.NewAnonRateThrottle("10/minute", storage),  // 10 req/min for anonymous
			throttling.NewUserRateThrottle("100/minute", storage), // 100 req/min for authenticated
		},
	}
}

func (v *ThrottledPostView) Get(c *views.Context) views.Response {
	// Check throttles before processing
	allowed, wait := throttling.CheckThrottles(c.Request, v, v.Throttles)
	if !allowed {
		c.ResponseWriter.Header().Set("Retry-After", wait.String())
		return views.Response{
			Status: 429,
			Data:   map[string]string{"error": "Rate limit exceeded. Try again later."},
		}
	}

	// Proceed with normal list operation
	return v.ListModelMixin.List(c)
}

// PublishedPostsView - Only show published posts with search
type PublishedPostsView struct {
	views.ListAPIView[*Post]
	SearchFilter *filters.SearchFilter
}

func NewPublishedPostsView(database *db.DB) *PublishedPostsView {
	return &PublishedPostsView{
		ListAPIView: views.ListAPIView[*Post]{
			GenericAPIView: views.GenericAPIView[*Post]{
				BaseAPIView: views.BaseAPIView{
					PermissionClasses: []views.Permission{
						&views.AllowAny{},
					},
				},
				DB: database,
			},
		},
		SearchFilter: filters.NewSearchFilter([]string{"title", "content", "author"}),
	}
}

// Override List to filter published posts
func (v *PublishedPostsView) Get(c *views.Context) views.Response {
	// This would filter for published=true in a real implementation
	// For now, return all posts
	return v.ListModelMixin.List(c)
}
