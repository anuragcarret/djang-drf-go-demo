package accounts

import (
	"github.com/anuragcarret/djang-drf-go/contrib/auth"
	"github.com/anuragcarret/djang-drf-go/orm/models"
)

// Account model holds profile information for a User
type Account struct {
	auth.User
	Bio       string    `drf:"bio;null;blank"`
	Avatar    string    `drf:"avatar;null;blank"`
	Location  string    `drf:"location;max_length=100;null"`
	Followers []Account `drf:"m2m=account_follows;to=follower_id;from=following_id;null"`
}

func (a *Account) TableName() string { return "accounts" }

func init() {
	models.RegisterModel("accounts", &Account{})
}
