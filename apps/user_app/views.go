package user_app

import (
	"fmt"
	"net/http"

	"github.com/anuragcarret/djang-drf-go/framework/drf"
)

type UserViewSet struct {
	drf.ModelViewSet
}

func (v *UserViewSet) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `[{"id": 1, "username": "admin", "email": "admin@example.com"}]`)
}
