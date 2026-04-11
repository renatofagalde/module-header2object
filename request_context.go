package header2object

const (
	HeaderCompanyID     = "X-Company-ID"
	HeaderSiteID        = "X-Site-ID"
	HeaderUserID        = "X-User-ID"
	HeaderCorrelationID = "X-Correlation-ID"
)

const (
    ContextKeyCompanyID     = "company_id"
    ContextKeySiteID        = "site_id"
    ContextKeyUserID        = "user_id"
    ContextKeyCorrelationID = "correlation_id"
)

type RequestContext struct {
	CompanyID     string
	SiteID        string
	UserID        string
	CorrelationID string
}

func (r RequestContext) IsValid() bool {
	return r.CompanyID != "" && r.SiteID != "" && r.UserID != ""
}
