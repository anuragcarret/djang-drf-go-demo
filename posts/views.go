package posts

import (
	"github.com/anuragcarret/djang-drf-go/drf/views"
	"github.com/anuragcarret/djang-drf-go/orm/db"
	"github.com/anuragcarret/djang-drf-go/orm/queryset"
)

type PostViewSet struct {
	views.ModelViewSet[*Post]
}

func NewPostViewSet(database *db.DB) *PostViewSet {
	return &PostViewSet{
		ModelViewSet: views.ModelViewSet[*Post]{
			BaseAPIView: views.BaseAPIView{
				PermissionClasses: []views.Permission{&views.IsAuthenticated{}},
			},
			DB:    database,
			Depth: 1, // Nest Author and Likes
		},
	}
}

// List overrides to add prefetching
func (v *PostViewSet) List(c *views.Context) views.Response {
	// Use the generic NewQuerySet helper from the queryset package
	qs := queryset.NewQuerySet[*Post](v.DB)
	posts, err := qs.PrefetchRelated("Likes", "Comments").All()
	if err != nil {
		return views.BadRequest(err.Error())
	}
	return views.OK(posts)
}
