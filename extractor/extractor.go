package extractor

import (
	h2o "github.com/renatofagalde/app-header2object"
)

// HeaderSetter is the constraint that any usecase input must implement
// to receive tenant and user context injection.
//
// Implement this method on your input struct:
//
//	func (i *CreateMovementInput) SetRequestContext(ctx h2o.RequestContext) {
//	    i.CompanyID = ctx.CompanyID
//	    i.SiteID    = ctx.SiteID
//	    i.UserID    = ctx.UserID
//	}
type HeaderSetter interface {
	SetRequestContext(ctx h2o.RequestContext)
}

// Enrich injects the RequestContext into any input that implements HeaderSetter.
// Returns the same pointer to allow inline use.
//
// Usage:
//
//	rCtx, ok := middleware.FromGinContext(c)
//	if !ok { ... }
//
//	input := extractor.Enrich(&CreateMovementInput{
//	    Amount: amount,
//	    Nature: movement.NatureDebit,
//	}, rCtx)
//
//	result, err := h.service.Create(c.Request.Context(), input)
func Enrich[T HeaderSetter](input T, ctx h2o.RequestContext) T {
	input.SetRequestContext(ctx)
	return input
}
