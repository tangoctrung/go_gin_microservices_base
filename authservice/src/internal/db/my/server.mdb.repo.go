package my

import (
	"context"

	"github.com/google/wire"
	"github.com/minh1611/go_structure/authservice/src/internal/db/user"
	"gorm.io/gorm"
)

var CDBRepoSet = wire.NewSet(wire.Struct(new(ServerCDBRepo), "*"), InterfaceProvider)

type DBInterfaces []interface{}

func InterfaceProvider() DBInterfaces {
	return DBInterfaces{
		user.User{},
	}
}

type ServerCDBRepo struct {
	Db         *gorm.DB
	Context    context.Context
	Interfaces DBInterfaces
	//Logger     *logrus.Logger
}
