package user_app

import (
	"fmt"

	"github.com/anuragcarret/djang-drf-go/framework/orm"
)

// User represents the system user model.
type User struct {
	orm.BaseModel
	Username string `orm:"size:100;unique" json:"username"`
	Email    string `orm:"size:255;unique" json:"email"`
}

func (u *User) GetTableName() string {
	return "auth_user"
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}
