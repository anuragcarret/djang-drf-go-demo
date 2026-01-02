package accounts

import (
	"github.com/anuragcarret/djang-drf-go/core/apps"
)

type AccountsApp struct{}

func (a *AccountsApp) AppConfig() *apps.AppConfig {
	return &apps.AppConfig{
		Name:  "accounts",
		Label: "accounts",
	}
}

func (a *AccountsApp) Ready() error {
	return nil
}

func init() {
	apps.Apps.Register(&AccountsApp{})
}
