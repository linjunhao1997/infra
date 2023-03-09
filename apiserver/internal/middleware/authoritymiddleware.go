package middleware

import (
	"infra/services/system/systemservice"
	"net/http"
)

type AuthorityMiddleware struct {
	systemService systemservice.SystemService
}

func NewAuthorityMiddleware(service systemservice.SystemService) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		systemService: service,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
