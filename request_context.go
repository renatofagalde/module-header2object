package header2object

const (
	HeaderCompanyID     = "X-Company-ID"
	HeaderSiteID        = "X-Site-ID"
	HeaderUserID        = "X-User-ID"
	HeaderCorrelationID = "X-Correlation-ID"
)

const (
	ContextKeyCompanyID     = "h2o_company_id"
	ContextKeySiteID        = "h2o_site_id"
	ContextKeyUserID        = "h2o_user_id"
	ContextKeyCorrelationID = "h2o_correlation_id"
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
