package user_app

import (
	"github.com/anuragcarret/djang-drf-go/framework/drf"
)

type UserSerializer struct {
	drf.BaseSerializer
}

// In a real implementation, we would use reflection to automate this mapping
// For the demo, we use the BaseSerializer's generic Marshal/Unmarshal.
