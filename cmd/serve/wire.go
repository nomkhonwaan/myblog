//+build wireinject

package serve

import (
	"github.com/google/wire"
	"github.com/nomkhonwaan/myblog/pkg/facebook"
	"github.com/nomkhonwaan/myblog/pkg/mongo"
	"github.com/nomkhonwaan/myblog/pkg/storage"
)

func initFacebookHandler(db mongo.Database) *facebook.Handler {
	wire.Build(facebook.NewHandler, storage.NewFileRepository)
	return &facebook.Handler{}
}
