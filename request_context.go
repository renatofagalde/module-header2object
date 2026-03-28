package header2object

const (
	HeaderCompanyID = "X-Company-ID"
	HeaderSiteID    = "X-Site-ID"
	HeaderUserID    = "X-User-ID"
)

const (
	ContextKeyCompanyID = "h2o_company_id"
	ContextKeySiteID    = "h2o_site_id"
	ContextKeyUserID    = "h2o_user_id"
)

// RequestContext holds the tenant and user identifiers extracted from
// HTTP headers after validation by the authorizer.
type RequestContext struct {
	CompanyID string
	SiteID    string
	UserID    string
}

func (r RequestContext) IsValid() bool {
	return r.CompanyID != "" && r.SiteID != "" && r.UserID != ""
}
