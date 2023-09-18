package di

import (
	"github.com/akshayur0404/go-aws-s3/pkg/config"
	db "github.com/akshayur0404/go-aws-s3/pkg/db"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) {
	wire.Build(db.ConnectDB)
}
