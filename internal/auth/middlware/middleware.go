package middlware

import (
	"delivery/internal/auth/services"
	"delivery/internal/controllers"
	"log"
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
		switch r.Method {
		case "GET":
			tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))

			_, err := services.ValidateToken(tokenString, m.controller.ConfigController.Config.AccessSecret)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				log.Println(err)
				return
			}
			claims, _ := services.Claims(tokenString, m.controller.ConfigController.Config.AccessSecret)
			exist, _ := m.controller.Auth.UserAccessTokenRepository.IsExistAccess(claims.ID) //expired="false" учтен при селекте существования
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
			next.ServeHTTP(w, r)
		case "POST":
			//req := new(models.UserRequestPairTokens)
			//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			//	http.Error(w, err.Error(), http.StatusBadRequest)
			//	return
			//}
			//accessTokenString := services.GetTokenFromBearerString(req.AccessToken)
			//refreshTokenString := services.GetTokenFromBearerString(req.RefreshToken)
			//log.Println("accessTokenString", accessTokenString)
			//log.Println("refreshTokenString", refreshTokenString)
			//tokenString := ""
			//if accessTokenString == "" {
			//	tokenString = refreshTokenString
			//	_, err := services.ValidateToken(tokenString, m.controller.ConfigController.Config.RefreshSecret)
			//	if err != nil {
			//		http.Error(w, err.Error(), http.StatusUnauthorized)
			//		log.Println(err)
			//		return
			//	}
			//	claims, _ := services.Claims(tokenString, m.controller.ConfigController.Config.RefreshSecret)
			//	exist, _ := m.controller.Auth.UserAccessTokenRepository.IsExistAccess(claims.ID) //expired="false" учтен при селекте существования
			//	if !exist {
			//		http.Error(w, "invalid token", http.StatusUnauthorized)
			//		return
			//	}
			//	tokenHash, _ := m.controller.Auth.UserRefreshTokenRepository.GetRefreshTokenByUserID(claims.ID)
			//	equal := services.CompareHashTokenDBAndRequest(tokenHash, tokenString)
			//	log.Println(tokenHash)
			//	log.Println(tokenString)
			//	if !equal {
			//		http.Error(w, "invalid token", http.StatusUnauthorized)
			//		return
			//	}
			//
			//	next.ServeHTTP(w, r)
			//} else if refreshTokenString == "" {
			//	tokenString = accessTokenString
			//	_, err := services.ValidateToken(tokenString, m.controller.ConfigController.Config.AccessSecret)
			//	if err != nil {
			//		http.Error(w, err.Error(), http.StatusUnauthorized)
			//		log.Println(err)
			//		return
			//	}
			//	claims, _ := services.Claims(tokenString, m.controller.ConfigController.Config.AccessSecret)
			//	exist, _ := m.controller.Auth.UserAccessTokenRepository.IsExistAccess(claims.ID) //expired="false" учтен при селекте существования
			//	if !exist {
			//		http.Error(w, "invalid token", http.StatusUnauthorized)
			//		return
			//	}
			//	tokenHash, _ := m.controller.Auth.UserAccessTokenRepository.GetAccessTokenByUserID(claims.ID)
			//	equal := services.CompareHashTokenDBAndRequest(tokenHash, tokenString)
			//	if !equal {
			//		http.Error(w, "invalid token", http.StatusUnauthorized)
			//		return
			//	}
			//	resp := req
			//	json.NewEncoder(w).Encode(resp)
			//	next.ServeHTTP(w, r)
			//} else {
			//	return
			//}
			//log.Println("MW pass")

		default:

		}

	})
}
func (m Middleware) CORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token"
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		log.Println("pass")
		next.ServeHTTP(w, r)

	}
}
