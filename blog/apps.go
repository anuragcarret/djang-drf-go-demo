package blog

import (
	"github.com/anuragcarret/djang-drf-go/core/apps"
	"github.com/anuragcarret/djang-drf-go/orm/models"
)

type Category struct {
	models.Model
	Name string `drf:"name;max_length=50;unique"`
}

func (c *Category) TableName() string { return "blog_categories" }

type BlogApp struct{}

func (a *BlogApp) AppConfig() *apps.AppConfig {
	return &apps.AppConfig{
		Name:  "blog",
		Label: "blog",
	}
}

func (a *BlogApp) Ready() error {
	return nil
}

func init() {
	apps.Apps.Register(&BlogApp{})
	models.RegisterModel("blog", &Category{})
}
