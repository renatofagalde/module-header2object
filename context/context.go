package context

import "context"

type contextKey string

const (
    companyIDKey     contextKey = "company_id"
    siteIDKey        contextKey = "site_id"
    userIDKey        contextKey = "user_id"
    correlationIDKey contextKey = "correlation_id"
)

func WithCompanyID(ctx context.Context, v string) context.Context {
    return context.WithValue(ctx, companyIDKey, v)
}

func CompanyIDFromContext(ctx context.Context) string {
    v, _ := ctx.Value(companyIDKey).(string)
    return v
}

func WithSiteID(ctx context.Context, v string) context.Context {
    return context.WithValue(ctx, siteIDKey, v)
}

func SiteIDFromContext(ctx context.Context) string {
    v, _ := ctx.Value(siteIDKey).(string)
    return v
}

func WithCorrelationID(ctx context.Context, v string) context.Context {
    return context.WithValue(ctx, correlationIDKey, v)
}

func CorrelationIDFromContext(ctx context.Context) string {
    v, _ := ctx.Value(correlationIDKey).(string)
    return v
}