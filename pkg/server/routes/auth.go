package routes

import (
	"github.com/cmelgarejo/go-gql-server/internal/handlers/auth"
	"github.com/danymarita/go-gql-server/internal/orm"
	"github.com/danymarita/go-gql-server/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Auth(cfg *utils.ServerConfig, r *gin.Engine, orm *orm.ORM) error {
	// OAuth handlers
	g := r.Group(cfg.VersionedEndpoint("/auth"))
	g.GET("/:provider", auth.Begin())
	g.GET("/:provider/callback", auth.Callback(cfg, orm))

	return nil
}
