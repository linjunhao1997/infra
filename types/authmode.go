package types

const (
	// 系统内部认证
	AuthModeInternal = "internal"
	// 企业sso认证
	AuthModeSSO = "sso"
	// Oauth2第三方认证
	AuthModelOauth2 = "oauth2"
)

var AuthMode = &AuthModeType{
	Internal: AuthModeInternal,
	SSO:      AuthModeSSO,
	Oauth2:   AuthModelOauth2,
}

type AuthModeType struct {
	Internal string
	SSO      string
	Oauth2   string
}
