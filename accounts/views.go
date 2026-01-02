package accounts

import (
	"github.com/anuragcarret/djang-drf-go/contrib/auth"
	"github.com/anuragcarret/djang-drf-go/drf/authentication"
	"github.com/anuragcarret/djang-drf-go/drf/views"
	"github.com/anuragcarret/djang-drf-go/orm/db"
)

// Simple in-memory token store for demo
type DemoTokenStore struct {
	tokens map[string]interface{} // token -> user
}

func NewDemoTokenStore() *DemoTokenStore {
	demoUser := &auth.User{
		Username: "demo_user",
		Email:    "demo@example.com",
		IsStaff:  false,
	}
	demoUser.ID = 1

	return &DemoTokenStore{
		tokens: map[string]interface{}{
			"demo-token-123": demoUser,
		},
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
	return "demo-token-123", nil
}

func (s *DemoTokenStore) RevokeToken(token string) error {
	delete(s.tokens, token)
	return nil
}

// Simple demo views using existing AccountViewSet
type AccountViewSet struct {
	views.ModelViewSet[*Account]
}

func NewAccountViewSet(database *db.DB) *AccountViewSet {
	return &AccountViewSet{
		ModelViewSet: views.ModelViewSet[*Account]{
			BaseAPIView: views.BaseAPIView{
				PermissionClasses: []views.Permission{
					&views.AllowAny{},
				},
			},
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

// AuthenticatedAccountView - requires authentication
type AuthenticatedAccountView struct {
	views.ModelViewSet[*Account]
}

func NewAuthenticatedAccountView(database *db.DB) *AuthenticatedAccountView {
	return &AuthenticatedAccountView{
		ModelViewSet: views.ModelViewSet[*Account]{
			BaseAPIView: views.BaseAPIView{
				PermissionClasses: []views.Permission{
					&views.IsAuthenticated{},
				},
			},
			DB: database,
		},
	}
}
