package admin_config

import (
	"demo/accounts"
	"demo/posts"

	framework_admin "github.com/anuragcarret/djang-drf-go/admin"
	"github.com/anuragcarret/djang-drf-go/contrib/auth"
	"github.com/anuragcarret/djang-drf-go/orm/queryset"
)

func init() {
	// Register models with the admin site
	framework_admin.Register[*auth.User](framework_admin.DefaultSite, &framework_admin.ModelAdmin{
		ListDisplay: []string{"Username", "Email", "IsStaff"},
	})

	framework_admin.Register[*accounts.Account](framework_admin.DefaultSite, &framework_admin.ModelAdmin{
		ListDisplay: []string{"Username", "DateJoined"},
	})

	framework_admin.Register[*posts.Post](framework_admin.DefaultSite, &framework_admin.ModelAdmin{
		ListDisplay: []string{"Content", "Created"},
		Actions: []framework_admin.AdminAction{
			{
				Name:  "log_to_console",
				Label: "Log to Console",
				Handler: func(qs *queryset.QuerySet[*posts.Post], ids []uint64) (string, error) {
					println("LOGGING SELECTED POSTS:", len(ids))
					for _, id := range ids {
						println("  - Post ID:", id)
					}
					return "Logged to console", nil
				},
			},
		},
	})

	framework_admin.Register[*posts.Comment](framework_admin.DefaultSite, &framework_admin.ModelAdmin{
		ListDisplay: []string{"Text", "Created"},
	})
}
