package accounts

import (
	"time"

	"github.com/anuragcarret/djang-drf-go/contrib/auth"
	"github.com/anuragcarret/djang-drf-go/orm/models"
)

// Account model holds profile information for a User
type Account struct {
	auth.User
	Bio      string `drf:"bio;null;blank"`
	Avatar   string `drf:"avatar;null;blank"`
	Location string `drf:"location;max_length=100;null"`
}

func (a *Account) TableName() string { return "accounts" }

type ComplexData struct {
	models.Model
	Status         string                 `drf:"status;max_length=20;default=active"`
	Score          float64                `drf:"score;type=numeric"`
	Tags           []string               `drf:"tags"`
	Metadata       map[string]interface{} `drf:"metadata"`
	ExternalID     string                 `drf:"external_id;type=uuid;unique"`
	IsProcessed    bool                   `drf:"is_processed;default=false"`
	ActivationDate time.Time              `drf:"activation_date;type=date;null"`
}

func (c *ComplexData) TableName() string { return "complex_data" }

func init() {
	models.RegisterModel("accounts", &Account{})
	models.RegisterModel("accounts", &ComplexData{})
}
