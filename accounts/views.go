package accounts

import (
	"github.com/anuragcarret/djang-drf-go/drf/views"
	"github.com/anuragcarret/djang-drf-go/orm/db"
)

type AccountViewSet struct {
	views.ModelViewSet[*Account]
}

func NewAccountViewSet(database *db.DB) *AccountViewSet {
	return &AccountViewSet{
		ModelViewSet: views.ModelViewSet[*Account]{
			DB: database,
		},
	}
}

func (v *AccountViewSet) Create(c *views.Context) views.Response {
	var a Account
	if err := c.Bind(&a); err != nil {
		return views.BadRequest(map[string]string{"error": "Failed to bind data: " + err.Error()})
	}

	a.SetPassword(a.Password)

	if err := v.ModelViewSet.PerformCreate(c, &a); err != nil {
		return views.BadRequest(map[string]string{"error": "Failed to create: " + err.Error()})
	}

	return views.Created(&a)
}
