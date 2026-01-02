package posts

import (
	"demo/accounts"

	"github.com/anuragcarret/djang-drf-go/orm/models"
)

// Post model represents a social media post
type Post struct {
	models.Model
	Title    string             `drf:"title;max_length=100;null"`
	AuthorID uint64             `drf:"author_id;index"` // Relates to auth.User ID
	Content  string             `drf:"content;max_length=280"`
	Likes    []accounts.Account `drf:"m2m=post_likes;to=account_id;from=post_id"`
	Comments []Comment          `drf:"relation=comments.post_id"`
}

func (p *Post) TableName() string { return "posts" }

// Comment model represents a comment on a post
type Comment struct {
	models.Model
	PostID   uint64 `drf:"post_id;index"`
	AuthorID uint64 `drf:"author_id;index"`
	Text     string `drf:"text;max_length=500"`
}

func (c *Comment) TableName() string { return "comments" }

func init() {
	models.RegisterModel("posts", &Post{})
	models.RegisterModel("posts", &Comment{})
}
