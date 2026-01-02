package accounts

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/anuragcarret/djang-drf-go/drf/views"
)

// TestDirectCreate tests calling Create directly
func TestDirectCreate(t *testing.T) {
	t.Log("Testing AccountViewSet.Create() directly...")

	// Create viewset with nil DB (will fail at DB operation, but we'll see if Create is called)
	viewSet := &AccountViewSet{}

	// Create mock context
	reqBody := map[string]string{
		"username": "testuser",
		"email":    "test@example.com",
		"password": "secure",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	ctx := &views.Context{
		Request:        req,
		ResponseWriter: w,
		Query:          req.URL.Query(),
		Data:           make(map[string]interface{}),
	}
	ctx.ParseRequest()

	// Call Create directly
	resp := viewSet.Create(ctx)

	t.Logf("Direct Create() returned status: %d", resp.Status)
	t.Logf("Response: %+v", resp)

	// We expect it to fail with DB error, but NOT with 405
	if resp.Status == 405 {
		t.Error("Create() should not return 405 Method Not Allowed")
	}
}
