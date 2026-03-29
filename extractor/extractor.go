package extractor

import (
	"errors"

	"github.com/gin-gonic/gin"
	h2o "github.com/renatofagalde/app-header2object"
	"github.com/renatofagalde/app-header2object/middleware"
)

type HeaderSetter interface {
	SetRequestContext(ctx h2o.RequestContext)
}

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
