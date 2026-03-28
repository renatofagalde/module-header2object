package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	h2o "github.com/renatofagalde/app-header2object"
)

// InjectHeaders reads X-Company-ID, X-Site-ID and X-User-ID from the request
// headers and stores them in the gin.Context for downstream use.
//
// Apply only on routes that go through the Lambda Authorizer.
// Do not use on public routes such as POST /tokens.
//
// Usage:
//
//	protected := r.Group("/cc")
//	protected.Use(middleware.InjectHeaders())
//	protected.POST("/movements/", movementHandler.Create)
func InjectHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		companyID := c.GetHeader(h2o.HeaderCompanyID)
		siteID := c.GetHeader(h2o.HeaderSiteID)
		userID := c.GetHeader(h2o.HeaderUserID)

		if companyID == "" || siteID == "" || userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing required tenant headers",
			})
			return
		}

		c.Set(h2o.ContextKeyCompanyID, companyID)
		c.Set(h2o.ContextKeySiteID, siteID)
		c.Set(h2o.ContextKeyUserID, userID)

		c.Next()
	}
}

// FromGinContext retrieves the RequestContext stored by InjectHeaders.
// Returns false if the middleware was not applied or any value is missing.
//
// Usage:
//
//	rCtx, ok := middleware.FromGinContext(c)
//	if !ok {
//	    c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
//	    return
//	}
func FromGinContext(c *gin.Context) (h2o.RequestContext, bool) {
	companyID, ok1 := c.Get(h2o.ContextKeyCompanyID)
	siteID, ok2 := c.Get(h2o.ContextKeySiteID)
	userID, ok3 := c.Get(h2o.ContextKeyUserID)

	if !ok1 || !ok2 || !ok3 {
		return h2o.RequestContext{}, false
	}

	ctx := h2o.RequestContext{
		CompanyID: companyID.(string),
		SiteID:    siteID.(string),
		UserID:    userID.(string),
	}

	return ctx, ctx.IsValid()
}
