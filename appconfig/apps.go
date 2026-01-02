package appconfig

import (
	"github.com/anuragcarret/djang-drf-go/core/apps"
)

type DemoApp struct{}

func (a *DemoApp) AppConfig() *apps.AppConfig {
	return &apps.AppConfig{
		Name:  "demo",
		Label: "demo",
	}
}

func (a *DemoApp) Ready() error {
	return nil
}

func init() {
	apps.Apps.Register(&DemoApp{})
}
