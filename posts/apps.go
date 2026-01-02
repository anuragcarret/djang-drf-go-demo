package posts

import (
	"github.com/anuragcarret/djang-drf-go/core/apps"
)

type PostsApp struct{}

func (a *PostsApp) AppConfig() *apps.AppConfig {
	return &apps.AppConfig{
		Name:  "posts",
		Label: "posts",
	}
}

func (a *PostsApp) Ready() error {
	return nil
}

func init() {
	apps.Apps.Register(&PostsApp{})
}
