# app-header2object

Go module that extracts HTTP headers validated by app-cam-authorizer and
injects them as typed properties into usecase inputs across any microservice.

## Headers

| Header       | Context key    | Field in RequestContext |
|--------------|----------------|-------------------------|
| X-Company-ID | h2o_company_id | CompanyID               |
| X-Site-ID    | h2o_site_id    | SiteID                  |
| X-User-ID    | h2o_user_id    | UserID                  |

These headers are validated by the authorizer before reaching the handler.
This module only extracts and types them.

## Installation

    go get github.com/renatofagalde/app-header2object

## Usage

Register the middleware on protected route groups:

    protected := r.Group("/cc")
    protected.Use(middleware.InjectHeaders())
    protected.POST("/movements/", movementHandler.Create)

Implement HeaderSetter on the usecase input:

    import h2o "github.com/renatofagalde/app-header2object"

    type CreateMovementInput struct {
        Amount    decimal.Decimal
        Nature    string
        CompanyID string
        SiteID    string
        UserID    string
    }

    func (i *CreateMovementInput) SetRequestContext(ctx h2o.RequestContext) {
        i.CompanyID = ctx.CompanyID
        i.SiteID    = ctx.SiteID
        i.UserID    = ctx.UserID
    }

Use Enrich in the handler after binding the request body:

    rCtx, ok := middleware.FromGinContext(c)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
        return
    }

    input := extractor.Enrich(ToCreateMovementInput(&dto), rCtx)
    result, err := h.service.Create(c.Request.Context(), input)

## Structure

    app-header2object/
    ├── request_context.go
    ├── middleware/
    │   └── gin_middleware.go
    └── extractor/
        └── extractor.go

## Related

- app-cam: authorizer that validates the headers upstream
- module-error: domain error mapping
