package user_app

import (
	"github.com/anuragcarret/djang-drf-go/framework/admin"
	"github.com/anuragcarret/djang-drf-go/framework/core"
)

type UserAppConfig struct{}

func (c *UserAppConfig) GetName() string {
	return "UserApp"
}

func (c *UserAppConfig) Ready() error {
	// Register with Admin
	admin.Register(&User{}, &admin.ModelAdmin{
		ListDisplay:  []string{"ID", "Username", "Email"},
		SearchFields: []string{"Username", "Email"},
	})
	return nil
}

func init() {
	core.RegisterApp(&UserAppConfig{})
}
