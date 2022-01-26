package middlware

import (
	repositories "delivery/internal/auth"
	handProv "delivery/internal/auth/handle_provide"
	"net/http"
)

const AccessSecret = "access_secret"
const RefreshSecret = "refresh_secret"

const AccessTokenLifetimeMinutes = 20
const RefreshTokenLifetimeMinutes = 60

type Middleware struct {
	hp *handProv.HandlerProvider
}

func NewMiddleware(hp *handProv.HandlerProvider) *Middleware {
	return &Middleware{
		hp: hp,
	}
}

func (m Middleware) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := repositories.GetTokenFromBearerString(r.Header.Get("Authorization"))
		_, err := repositories.ValidateToken(tokenString, AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		accessToken, err := m.hp.UserAccessTokenRepository.GetByAccessToken(tokenString)

		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		if accessToken.Expired != "false" {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
