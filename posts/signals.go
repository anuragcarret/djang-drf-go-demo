package posts

import (
	"fmt"

	"github.com/anuragcarret/djang-drf-go/orm/signals"
)

func init() {
	signals.Register(signals.PostSave, "posts", func(sender interface{}, instance interface{}, kwargs map[string]interface{}) {
		post, ok := instance.(*Post)
		if !ok {
			return
		}

		created := kwargs["created"].(bool)
		if created {
			fmt.Printf("[SIGNAL] New post created by AuthorID %d: %s\n", post.AuthorID, post.Content)
		} else {
			fmt.Printf("[SIGNAL] Post %d updated\n", post.ID)
		}
	})
}
