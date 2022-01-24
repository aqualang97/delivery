package main

import (
	"database/sql"
	"delivery/internal/auth"
	"delivery/internal/models"
	db "delivery/internal/repositories/database"
	//connection "delivery/internal/repositories/database/connection"
	"time"

	open "delivery/internal/repositories/database/connection"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

const AccessSecret = "access_secret"
const RefreshSecret = "refresh_secret"

const AccessTokenLifetimeMinutes = 10
const RefreshTokenLifetimeMinutes = 60

func dbTXBegin(conn *sql.DB) (*sql.Tx, error) {
	TX, err := conn.Begin()
	return TX, err

}

func dbOpen() (*sql.DB, error) {

	dbConn, err := sql.Open(
		"mysql",
		"oboznyi:123123@tcp(127.0.0.1:3306)/oboznyi_db",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	//var id int
	//var name string
	//rows, err := db.Query("SELECT id, login FROM users WHERE id = ?", 1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for rows.Next() {
	//	err := rows.Scan(&id, &name)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(id, name)
	//
	//}
	return dbConn, err
}

type HandlerProvider struct {
	UserRepository             *db.UserDBRepository
	UserAccessTokenRepository  *db.UserAccessTokenRepository
	UserRefreshTokenRepository *db.UserRefreshTokenRepository
}

func main() {
	//реализуем флоу логина юзера
	// юзер дает логин пароль
	// получаем ответ верный илинет юзе
	conn, err := open.OpenMyDB()
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	TX, err := dbTXBegin(conn)
	if err != nil {
		return
	}
	handlerProvider := HandlerProvider{
		UserRepository:             db.NewUserRepo(conn, TX),
		UserAccessTokenRepository:  db.NewAccessTokenRepo(conn, TX),
		UserRefreshTokenRepository: db.NewRefreshTokenRepo(conn, TX),
	}
	http.HandleFunc("/login", handlerProvider.Login) //умеем обрабатывать логин с помощью ф-ции логин
	http.HandleFunc("/profile", handlerProvider.Profile)
	http.HandleFunc("/refresh", handlerProvider.Refresh)
	http.HandleFunc("/registration", handlerProvider.Registration)

	log.Fatal(http.ListenAndServe(":8080", nil)) //слушаем порт 8080 для входящих запросов

}

func (hp *HandlerProvider) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(models.LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil { //берем тело запроса декодим и декодим в тело запроса
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := hp.UserRepository.GetUserByEmail(req.Email)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusBadRequest)
			return
		}
		println(user.ID)

		if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		accessString, err := repositories.GenerateToken(user.ID, AccessTokenLifetimeMinutes, AccessSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		refreshString, err := repositories.GenerateToken(user.ID, RefreshTokenLifetimeMinutes, RefreshSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nowTime := time.Now()
		accessExpiredAt := nowTime.Add(time.Duration(AccessTokenLifetimeMinutes) * time.Minute)
		refreshExpiredAt := nowTime.Add(time.Duration(RefreshTokenLifetimeMinutes) * time.Minute)
		respAccess := models.UserAccessToken{
			UserID:      user.ID,
			AccessToken: accessString,
			ExpiredAt:   &accessExpiredAt,
			Expired:     "false",
		}
		respRefresh := models.UserRefreshToken{
			UserID:       user.ID,
			RefreshToken: refreshString,
			ExpiredAt:    &refreshExpiredAt,
			Expired:      "false",
		}

		//err = hp.UserRepository.InsertAccessAndRefreshTokens(&resp)
		//err = hp.UserAccessTokenRepository.UpdateOldAndInsertNewAccessToken()
		err = hp.UserAccessTokenRepository.InsertAccessToken(respAccess)
		if err != nil {
			log.Fatal(err)
		}
		err = hp.UserRefreshTokenRepository.InsertRefreshToken(respRefresh)
		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respAccess.AccessToken)
		json.NewEncoder(w).Encode(respRefresh.RefreshToken)

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
		tokenString := repositories.GetTokenFromBearerString(r.Header.Get("Authorization"))
		claims, err := repositories.ValidateToken(tokenString, AccessSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		accessToken, err := hp.UserAccessTokenRepository.GetByAccessToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		println(accessToken.AccessToken)

		if accessToken.Expired != "false" {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		//user, err := authRepo.NewUserRepository().GetUserByID(claims.ID)
		//if err != nil {
		//	http.Error(w, "invalid token", http.StatusUnauthorized)
		//	return
		//}

		user, err := hp.UserRepository.GetUserById(claims.ID)

		resp := models.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Login,
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
		////println(req.Token)
		////tokenString := GetTokenFromBearerString(r.Header.Get("Authorization"))
		//
		////access возможно не нужно проверять?
		////смысл их проверять, если они не предназначены для рефреша,
		////а через AccessTokenLifetimeMinutes они станут невалидными
		///*
		//	accessTokenString := req.AccessToken
		//	claims, err := ValidateToken(accessTokenString, AccessSecret)
		//	if err != nil {
		//		http.Error(w, err.Error(), http.StatusUnauthorized)
		//		return
		//	}
		//
		//	user, err := NewUserRepository().GetUserByID(claims.ID)
		//	if err != nil {
		//		http.Error(w, "invalid token", http.StatusUnauthorized)
		//		return
		//	}
		//*
		refreshTokenString := req.RefreshToken
		accessTokenString := req.AccessToken
		//claims, err := repositories.ValidateToken(refreshTokenString, RefreshSecret)
		_, err := repositories.ValidateToken(refreshTokenString, RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userToken, err := hp.UserRefreshTokenRepository.GetByRefreshToken(refreshTokenString)
		//userToken.AccessToken = accessTokenString
		println(userToken.UserID, userToken.Expired, userToken.RefreshToken)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		if userToken.Expired != "false" {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

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

		newAccessTokenString, err := repositories.GenerateToken(userToken.UserID, AccessTokenLifetimeMinutes, AccessSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newRefreshTokenString, err := repositories.GenerateToken(userToken.UserID, RefreshTokenLifetimeMinutes, RefreshSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nowTime := time.Now()
		accessExpiredAt := nowTime.Add(time.Duration(AccessTokenLifetimeMinutes) * time.Minute)
		refreshExpiredAt := nowTime.Add(time.Duration(RefreshTokenLifetimeMinutes) * time.Minute)

		respAccess := models.UserAccessToken{
			AccessToken: newAccessTokenString,
			UserID:      userToken.UserID,
			ExpiredAt:   &accessExpiredAt,
			Expired:     "false",
		}

		respRefresh := models.UserRefreshToken{
			RefreshToken: newRefreshTokenString,
			UserID:       userToken.UserID,
			ExpiredAt:    &refreshExpiredAt,
			Expired:      "false",
		}

		err = hp.UserAccessTokenRepository.UpdateOldAndInsertNewAccessToken(accessTokenString, respAccess)
		err = hp.UserRefreshTokenRepository.UpdateOldAndInsertNewRefreshToken(refreshTokenString, respRefresh)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respAccess.AccessToken)
		json.NewEncoder(w).Encode(respRefresh.RefreshToken)
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

		accessString, err := repositories.GenerateToken(user.ID, AccessTokenLifetimeMinutes, AccessSecret)
		println(user.Email, user.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nowTime := time.Now()
		accessExpiredAt := nowTime.Add(time.Duration(AccessTokenLifetimeMinutes) * time.Minute)
		refreshExpiredAt := nowTime.Add(time.Duration(RefreshTokenLifetimeMinutes) * time.Minute)

		refreshString, err := repositories.GenerateToken(user.ID, RefreshTokenLifetimeMinutes, RefreshSecret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		respAccess := models.UserAccessToken{
			UserID:      user.ID,
			AccessToken: accessString,
			ExpiredAt:   &accessExpiredAt,
			Expired:     "false",
		}
		respRefresh := models.UserRefreshToken{
			UserID:       user.ID,
			RefreshToken: refreshString,
			ExpiredAt:    &refreshExpiredAt,
			Expired:      "false",
		}
		//resp := models.UserToken{
		//	UserID:           user.ID,
		//	AccessToken:      tokenString,
		//	RefreshToken:     refreshString,
		//	AccessExpiredAt:  time.Now().Add(time.Duration(AccessTokenLifetimeMinutes) * time.Minute),
		//	RefreshExpiredAt: time.Now().Add(time.Duration(RefreshTokenLifetimeMinutes) * time.Minute),
		//	Expired:          "false",
		//}

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
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
