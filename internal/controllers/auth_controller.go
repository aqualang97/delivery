package controllers

import (
	"delivery/internal/auth/services"
	"delivery/internal/models"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(models.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil { //берем тело запроса декодим и декодим в тело запроса
			http.Error(w, err.Error(), http.StatusBadRequest)
			//	logger.Error("/login, LoginRequest\n", err)

			return
		}

		user, err := a.UserRepository.GetUserByEmail(req.Email)
		fmt.Println(user)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusBadRequest)
			return
		}
		println(user.PasswordHash)
		println(req.Password)

		if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		_ = a.UserAccessTokenRepository.ExpiredAccessToken(user.ID)
		_ = a.UserRefreshTokenRepository.ExpiredRefreshToken(user.ID)
		//fmt.Println(hp.Config.AccessLifetimeMinutes)
		cfg := a.ConfigController.Config
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
		err = a.UserAccessTokenRepository.InsertAccessToken(respAccess)
		if err != nil {
			log.Fatal(err)
		}
		err = a.UserRefreshTokenRepository.InsertRefreshToken(respRefresh)
		if err != nil {
			log.Fatal(err)
		}
		resp := models.UserResponsePairTokens{
			UserID:       user.ID,
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		w.WriteHeader(http.StatusOK)
		//json.NewEncoder(w).Encode(accessString)
		//json.NewEncoder(w).Encode(refreshString)

		data, _ := json.Marshal(resp)
		w.Write(data)

	}
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))
		req := new(models.UserRequestPairTokens)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tokenString := services.GetTokenFromBearerString(req.AccessToken)
		_, err := services.ValidateToken(tokenString, a.ConfigController.Config.AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			log.Println(err)
			log.Println(err)
			log.Println(err)
			return
		}
		cfg := a.ConfigController.Config
		claims, err := services.Claims(tokenString, a.ConfigController.Config.AccessSecret)
		if err != nil {
			cfg.Logger.Error("Auth err", err)
			return
		}
		exist, _ := a.UserAccessTokenRepository.IsExistAccess(claims.ID) //expired="false" учтен при селекте существования
		if !exist {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		tokenHash, _ := a.UserAccessTokenRepository.GetAccessTokenByUserID(claims.ID)
		equal := services.CompareHashTokenDBAndRequest(tokenHash, tokenString)
		if !equal {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		//user, err := hp.UserRepository.GetUserById(claims.ID)
		err = a.UserRefreshTokenRepository.ExpiredRefreshToken(claims.ID)
		if err != nil {
			cfg.Logger.Error("Auth err", err)
			return
		}
		err = a.UserAccessTokenRepository.ExpiredAccessToken(claims.ID)
		if err != nil {
			cfg.Logger.Error("Auth err", err)
			return
		}
		fmt.Println("logout", claims.ID)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Successful Logout")
	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

func (a AuthController) Registration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		req := new(models.RegistrationRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		exist, err := a.UserRepository.IsExistUserByEmail(req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if exist {

			http.Error(w, "This email is already taken.", http.StatusConflict)
			//data, _ := json.Marshal("This email is already taken.")
			//w.WriteHeader(http.StatusConflict)
			//w.Write(data)
			return
		}

		exist, err = a.UserRepository.IsExistUserByLogin(req.Login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		if exist {
			http.Error(w, "This user name is already taken.", http.StatusConflict)
			return
		}
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		_, err = a.UserRepository.CreateUser(req, string(passwordHash))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := a.UserRepository.GetUserByEmail(req.Email)
		cfg := a.ConfigController.Config
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

		err = a.UserAccessTokenRepository.InsertAccessToken(respAccess)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = a.UserRefreshTokenRepository.InsertRefreshToken(respRefresh)
		resp := models.UserResponsePairTokens{
			UserID:       user.ID,
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}

		//Если оставляем юзера залогиненым поссле регистрации, то даем ему пару токенов
		//user, err := authRepo.NewUserRepository().GetUserByEmail(req.Email)
		//if err != nil {
		//	http.Error(w, "invalid credentials", http.StatusUnauthorized)
		//	return
		//}
		data, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusOK)
		w.Write(data)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
