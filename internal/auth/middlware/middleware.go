package middlware

import (
	"delivery/internal/auth/services"
	"delivery/internal/controllers"
	"net/http"
)

type Middleware struct {
	controller *controllers.Controllers
}

func NewMiddleware(c *controllers.Controllers) *Middleware {
	return &Middleware{
		controller: c,
	}
}

func (m Middleware) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))
		_, err := services.ValidateToken(tokenString, m.controller.ConfigController.Config.AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		//accessTokenHash, err := m.hp.UserAccessTokenRepository.GetByAccessToken(tokenString)
		claims, _ := services.Claims(tokenString, m.controller.ConfigController.Config.AccessSecret)
		exist, _ := m.controller.Auth.UserAccessTokenRepository.IsExistAccess(claims.ID) //expired="false" учтен при селекте существования

		//if err != nil {
		//	http.Error(w, "invalid token", http.StatusUnauthorized)
		//	return
		//}
		if !exist {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		tokenHash, _ := m.controller.Auth.UserAccessTokenRepository.GetAccessTokenByUserID(claims.ID)
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
