package delivery

import (
	"github.com/gin-gonic/gin"
	"profile_service/app"
)

type AssetsDelivery struct {
	log app.Logger
}

func NewAssetsDelivery(log app.Logger) *AssetsDelivery {
	return &AssetsDelivery{
		log: log,
	}
}

func (d *AssetsDelivery) Route(g *gin.RouterGroup) error {
	g.Static("/", "./assets")
	return nil
}
