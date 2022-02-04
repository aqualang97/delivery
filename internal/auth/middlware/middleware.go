package middlware

import (
	handProv "delivery/internal/auth/handle_provide"
	"delivery/internal/auth/services"
	"net/http"
)

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
		tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))
		_, err := services.ValidateToken(tokenString, m.hp.Config.AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		//accessTokenHash, err := m.hp.UserAccessTokenRepository.GetByAccessToken(tokenString)
		claims, _ := services.Claims(tokenString, m.hp.Config.AccessSecret)
		exist, _ := m.hp.UserAccessTokenRepository.IsExistAccess(claims.ID) //expired="false" учтен при селекте существования

		//if err != nil {
		//	http.Error(w, "invalid token", http.StatusUnauthorized)
		//	return
		//}
		if !exist {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		tokenHash, _ := m.hp.UserAccessTokenRepository.GetAccessTokenByUserID(claims.ID)
		equal := services.CompareHashTokenDBAndRequest(tokenHash, tokenString)
		if !equal {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		//if accessToken.Expired != "false" {
		//	http.Error(w, "invalid token", http.StatusUnauthorized)
		//	return
		//}
		next.ServeHTTP(w, r)
	})
}
