package database

import (
	user_persistence "github.com/dukk308/beetool.dev-go-starter/internal/modules/user/infrastructure/persistence"
)

var Models = []interface{}{
	user_persistence.SQLUser{},
}
