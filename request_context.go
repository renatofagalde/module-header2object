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

// ValidationOpts agrega flags de configuração das validações de RequestContext.
type ValidationOpts struct {
	SkipSite bool
}

// Option configura uma operação de validação ou extração de headers.
// Implementada como functional option (idiomático em Go).
type Option func(*ValidationOpts)

// SkipSite torna a validação do X-Site-ID opcional. Usado em rotas
// que operam apenas no escopo de uma company (ex: /pes/*).
func SkipSite() Option {
	return func(o *ValidationOpts) { o.SkipSite = true }
}

// ResolveOptions aplica todas as opções e retorna a config resultante.
// Exposto para uso pelos pacotes middleware e extractor.
func ResolveOptions(opts ...Option) ValidationOpts {
	cfg := ValidationOpts{}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}

// IsValid valida que os headers obrigatórios estão preenchidos.
//
// Por padrão exige CompanyID, SiteID e UserID.
// Use SkipSite() para pular a validação de SiteID.
//
// Backward compatible: r.IsValid() continua funcionando idêntico.
func (r RequestContext) IsValid(opts ...Option) bool {
	cfg := ResolveOptions(opts...)

	if r.CompanyID == "" || r.UserID == "" {
		return false
	}
	if !cfg.SkipSite && r.SiteID == "" {
		return false
	}
	return true
}
