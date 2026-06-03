package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	domainerror "github.com/renatofagalde/module-error"
	"github.com/renatofagalde/module-error/httperror"
	h2o "github.com/renatofagalde/module-header2object"
)

// InjectHeaders extrai os headers do tenant e popula o gin.Context.
//
// Por padrão exige X-Company-ID, X-Site-ID e X-User-ID.
// Use h2o.SkipSite() para tornar X-Site-ID opcional em rotas que
// operam apenas por company (ex: /pes/*).
//
// Backward compatible: InjectHeaders() continua funcionando idêntico.
func InjectHeaders(opts ...h2o.Option) gin.HandlerFunc {
	cfg := h2o.ResolveOptions(opts...)

	return func(c *gin.Context) {
		companyID := c.GetHeader(h2o.HeaderCompanyID)
		siteID := c.GetHeader(h2o.HeaderSiteID)
		userID := c.GetHeader(h2o.HeaderUserID)

		if companyID == "" || userID == "" {
			httperror.WriteError(c, domainerror.ErrInvalidInput)
			c.Abort()
			return
		}
		if !cfg.SkipSite && siteID == "" {
			httperror.WriteError(c, domainerror.ErrInvalidInput)
			c.Abort()
			return
		}

		correlationID := c.GetHeader(h2o.HeaderCorrelationID)
		if correlationID == "" {
			generated, _ := uuid.NewV7()
			correlationID = generated.String()
		}

		c.Set(h2o.ContextKeyCompanyID, companyID)
		c.Set(h2o.ContextKeySiteID, siteID)
		c.Set(h2o.ContextKeyUserID, userID)
		c.Set(h2o.ContextKeyCorrelationID, correlationID)

		c.Next()
	}
}

// FromGinContext extrai o RequestContext do gin.Context.
//
// Por padrão exige que todos os campos estejam preenchidos.
// Use h2o.SkipSite() para aceitar SiteID vazio.
//
// Backward compatible: FromGinContext(c) continua funcionando idêntico.
func FromGinContext(c *gin.Context, opts ...h2o.Option) (h2o.RequestContext, bool) {
	companyID, ok1 := c.Get(h2o.ContextKeyCompanyID)
	siteID, ok2 := c.Get(h2o.ContextKeySiteID)
	userID, ok3 := c.Get(h2o.ContextKeyUserID)

	if !ok1 || !ok2 || !ok3 {
		return h2o.RequestContext{}, false
	}

	correlationID := ""
	if v, ok := c.Get(h2o.ContextKeyCorrelationID); ok {
		correlationID = v.(string)
	}

	ctx := h2o.RequestContext{
		CompanyID:     companyID.(string),
		SiteID:        siteID.(string),
		UserID:        userID.(string),
		CorrelationID: correlationID,
	}

	return ctx, ctx.IsValid(opts...)
}
