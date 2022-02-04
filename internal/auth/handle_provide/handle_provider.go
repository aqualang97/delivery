package handle_provide

import (
	config "delivery/configs"
	"delivery/internal/auth/services"
	"delivery/internal/models"
	r "delivery/internal/repositories/database"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"

	"time"
)

type HandlerProvider struct {
	UserRepository             *r.UserDBRepository
	UserAccessTokenRepository  *r.UserAccessTokenRepository
	UserRefreshTokenRepository *r.UserRefreshTokenRepository
	Config                     *config.Config
}

func (hp *HandlerProvider) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(models.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil { //берем тело запроса декодим и декодим в тело запроса
			http.Error(w, err.Error(), http.StatusBadRequest)
			//	logger.Error("/login, LoginRequest\n", err)

			return
		}

		user, err := hp.UserRepository.GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusBadRequest)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		_ = hp.UserAccessTokenRepository.ExpiredAccessToken(user.ID)
		_ = hp.UserRefreshTokenRepository.ExpiredRefreshToken(user.ID)

		//fmt.Println(hp.Config.AccessLifetimeMinutes)
		cfg := hp.Config
		accessString, err := services.GenerateToken(user.ID, cfg.AccessLifetimeMinutes, cfg.AccessSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		accessHash, _ := services.GetHashOfToken(accessString)
		refreshString, err := services.GenerateToken(user.ID, cfg.RefreshLifetimeMinutes, cfg.RefreshSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		refreshHash, _ := services.GetHashOfToken(refreshString)

		nowTime := time.Now()

		accessExpiredAt := nowTime.Add(time.Duration(cfg.AccessLifetimeMinutes) * time.Minute)
		refreshExpiredAt := nowTime.Add(time.Duration(cfg.RefreshLifetimeMinutes) * time.Minute)

		respAccess := models.UserAccessToken{
			UserID:      user.ID,
			AccessToken: accessHash,
			ExpiredAt:   &accessExpiredAt,
			Expired:     "false",
		}
		respRefresh := models.UserRefreshToken{
			UserID:       user.ID,
			RefreshToken: refreshHash,
			ExpiredAt:    &refreshExpiredAt,
			Expired:      "false",
		}
		err = hp.UserAccessTokenRepository.InsertAccessToken(respAccess)
		if err != nil {
			log.Fatal(err)
		}
		err = hp.UserRefreshTokenRepository.InsertRefreshToken(respRefresh)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accessString)
		json.NewEncoder(w).Encode(refreshString)
	}
}

/*func Login1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(repositories.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil { //берем тело запроса декодим и декодим в тело запроса
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := authRepo.NewUserRepository().GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "invalid ", http.StatusUnauthorized)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		tokenString, err := repositories.GenerateToken(user.ID, AccessTokenLifetimeMinutes, AccessSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		refreshString, err := repositories.GenerateToken(user.ID, RefreshTokenLifetimeMinutes, RefreshSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := repositories.LoginResponse{
			AccessToken:  tokenString,
			RefreshToken: refreshString,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:

		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}*/

func (hp *HandlerProvider) Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))
		cfg := hp.Config

		claims, _ := services.Claims(tokenString, cfg.AccessSecret)

		user, err := hp.UserRepository.GetUserById(claims.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		resp := models.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Login: user.Login,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}
func (hp *HandlerProvider) Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(models.UserRequestPairTokens)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cfg := hp.Config

		refreshTokenString := req.RefreshToken

		//accessTokenString := req.AccessToken
		//claims, err := repositories.ValidateToken(refreshTokenString, RefreshSecret)
		_, err := services.ValidateToken(refreshTokenString, cfg.RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, _ := services.Claims(refreshTokenString, cfg.RefreshSecret)
		exist, _ := hp.UserRefreshTokenRepository.IsExistRefresh(claims.ID)
		if !exist {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		tokenHash, _ := hp.UserRefreshTokenRepository.GetRefreshTokenByUserID(claims.ID)
		equal := services.CompareHashTokenDBAndRequest(tokenHash, refreshTokenString)
		if !equal {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		//userToken, err := hp.UserRefreshTokenRepository.GetByRefreshToken(refreshTokenString)
		//userToken.AccessToken = accessTokenString

		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusUnauthorized)
		//}
		//if userToken.Expired != "false" {
		//	http.Error(w, "invalid token", http.StatusUnauthorized)
		//	return
		//}

		//user, err := authRepo.NewUserRepository().GetUserByID(claims.ID)
		//if err != nil {
		//	http.Error(w, "invalid token", http.StatusUnauthorized)
		//	return
		//}

		/*
			err := ValidateTokenToRefresh(accessTokenString, AccessSecret)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			refreshTokenString := req.AccessToken
			err = ValidateTokenToRefresh(refreshTokenString, RefreshSecret)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}*/

		newAccessTokenString, err := services.GenerateToken(claims.ID, cfg.AccessLifetimeMinutes, cfg.AccessSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newRefreshTokenString, err := services.GenerateToken(claims.ID, cfg.RefreshLifetimeMinutes, cfg.RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		accessHash, _ := services.GetHashOfToken(newAccessTokenString)
		refreshHash, _ := services.GetHashOfToken(newRefreshTokenString)

		nowTime := time.Now()
		accessExpiredAt := nowTime.Add(time.Duration(cfg.AccessLifetimeMinutes) * time.Minute)
		refreshExpiredAt := nowTime.Add(time.Duration(cfg.RefreshLifetimeMinutes) * time.Minute)

		respAccess := models.UserAccessToken{
			AccessToken: accessHash,
			UserID:      claims.ID,
			ExpiredAt:   &accessExpiredAt,
			Expired:     "false",
		}

		respRefresh := models.UserRefreshToken{
			RefreshToken: refreshHash,
			UserID:       claims.ID,
			ExpiredAt:    &refreshExpiredAt,
			Expired:      "false",
		}

		err = hp.UserAccessTokenRepository.UpdateOldAndInsertNewAccessToken(claims.ID, respAccess)
		err = hp.UserRefreshTokenRepository.UpdateOldAndInsertNewRefreshToken(claims.ID, respRefresh)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newAccessTokenString)
		json.NewEncoder(w).Encode(newRefreshTokenString)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
func (hp *HandlerProvider) Registration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		req := new(models.RegistrationRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := hp.UserRepository.GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(user.Email) != 0 {
			http.Error(w, "This email is already taken.", http.StatusConflict)
			return
		}

		user, err = hp.UserRepository.GetUserByLogin(req.Login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(user.Login) != 0 {
			http.Error(w, "This user name is already taken.", http.StatusConflict)
			return
		}
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		_, err = hp.UserRepository.CreateUser(req, string(passwordHash))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err = hp.UserRepository.GetUserByEmail(req.Email)
		cfg := hp.Config
		accessString, err := services.GenerateToken(user.ID, cfg.AccessLifetimeMinutes, cfg.AccessSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nowTime := time.Now()
		accessExpiredAt := nowTime.Add(time.Duration(cfg.AccessLifetimeMinutes) * time.Minute)
		refreshExpiredAt := nowTime.Add(time.Duration(cfg.RefreshLifetimeMinutes) * time.Minute)

		refreshString, err := services.GenerateToken(user.ID, cfg.RefreshLifetimeMinutes, cfg.RefreshSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		accessHash, _ := services.GetHashOfToken(accessString)
		refreshHash, _ := services.GetHashOfToken(refreshString)
		respAccess := models.UserAccessToken{
			ID:          user.ID,
			UserID:      user.ID,
			AccessToken: accessHash,
			ExpiredAt:   &accessExpiredAt,
			Expired:     "false",
		}
		respRefresh := models.UserRefreshToken{
			ID:           user.ID,
			UserID:       user.ID,
			RefreshToken: refreshHash,
			ExpiredAt:    &refreshExpiredAt,
			Expired:      "false",
		}

		err = hp.UserAccessTokenRepository.InsertAccessToken(respAccess)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = hp.UserRefreshTokenRepository.InsertRefreshToken(respRefresh)

		//Если оставляем юзера залогиненым поссле регистрации, то даем ему пару токенов
		//user, err := authRepo.NewUserRepository().GetUserByEmail(req.Email)
		//if err != nil {
		//	http.Error(w, "invalid credentials", http.StatusUnauthorized)
		//	return
		//}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respAccess)
		json.NewEncoder(w).Encode(respRefresh)
		json.NewEncoder(w).Encode(accessString)
		json.NewEncoder(w).Encode(refreshString)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
func (hp *HandlerProvider) Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))
		cfg := hp.Config
		claims, _ := services.Claims(tokenString, cfg.AccessSecret)
		//user, err := hp.UserRepository.GetUserById(claims.ID)
		_ = hp.UserRefreshTokenRepository.ExpiredRefreshToken(claims.ID)
		_ = hp.UserAccessTokenRepository.ExpiredAccessToken(claims.ID)
		fmt.Println("logout", claims.ID)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Successful Logout")
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
