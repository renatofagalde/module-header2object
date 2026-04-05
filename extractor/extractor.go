package extractor

import (
	"errors"

	"github.com/gin-gonic/gin"
	h2o "github.com/renatofagalde/module-header2object"
	"github.com/renatofagalde/module-header2object/middleware"
)

// HeaderSetter é implementado por todos os DTOs de rotas protegidas.
// Após extractor.Bind, o DTO terá CompanyID, SiteID, UserID e CorrelationID preenchidos.
type HeaderSetter interface {
	SetRequestContext(ctx h2o.RequestContext)
}

// Bind faz o bind do JSON e injeta o tenant context no DTO em uma única chamada.
// Substitui c.ShouldBindJSON em todos os handlers de rotas protegidas.
func Bind[T HeaderSetter](c *gin.Context, target T) error {
	if err := c.ShouldBindJSON(target); err != nil {
		return err
	}

	rCtx, ok := middleware.FromGinContext(c)
	if !ok {
		return errors.New("missing tenant headers")
	}

	target.SetRequestContext(rCtx)
	return nil
}
