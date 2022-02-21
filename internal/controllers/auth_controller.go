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
		w.Header().Add("Access-Control-Allow-Origin", "*")
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
		//w.WriteHeader(http.StatusOK)
		//json.NewEncoder(w).Encode(accessString)
		//json.NewEncoder(w).Encode(refreshString)

		data, _ := json.Marshal(resp)
		w.Write(data)

	}
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))
		cfg := a.ConfigController.Config
		claims, _ := services.Claims(tokenString, cfg.AccessSecret)
		//user, err := hp.UserRepository.GetUserById(claims.ID)
		_ = a.UserRefreshTokenRepository.ExpiredRefreshToken(claims.ID)
		_ = a.UserAccessTokenRepository.ExpiredAccessToken(claims.ID)
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

		user, err := a.UserRepository.GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(user.Email) != 0 {
			http.Error(w, "This email is already taken.", http.StatusConflict)
			return
		}

		user, err = a.UserRepository.GetUserByLogin(req.Login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(user.Login) != 0 {
			http.Error(w, "This user name is already taken.", http.StatusConflict)
			return
		}
		passwordHash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

		_, err = a.UserRepository.CreateUser(req, string(passwordHash))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err = a.UserRepository.GetUserByEmail(req.Email)
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
