package controllers

import (
	"delivery/internal/auth/services"
	"delivery/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (u UserController) Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))
		cfg := u.ConfigController.Config

		claims, _ := services.Claims(tokenString, cfg.AccessSecret)
		user, err := u.UserRepository.GetUserById(claims.ID)
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

func (u UserController) Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(models.UserRequestPairTokens)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cfg := u.ConfigController.Config

		refreshTokenString := req.RefreshToken

		//accessTokenString := req.AccessToken
		//claims, err := repositories.ValidateToken(refreshTokenString, RefreshSecret)
		_, err := services.ValidateToken(refreshTokenString, cfg.RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		claims, _ := services.Claims(refreshTokenString, cfg.RefreshSecret)
		exist, _ := u.UserRefreshTokenRepository.IsExistRefresh(claims.ID)
		if !exist {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		tokenHash, _ := u.UserRefreshTokenRepository.GetRefreshTokenByUserID(claims.ID)
		equal := services.CompareHashTokenDBAndRequest(tokenHash, refreshTokenString)
		if !equal {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

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

		err = u.UserAccessTokenRepository.UpdateOldAndInsertNewAccessToken(claims.ID, respAccess)
		err = u.UserRefreshTokenRepository.UpdateOldAndInsertNewRefreshToken(claims.ID, respRefresh)

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

func (u UserController) GetProductsInCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Access-Control-Allow-Origin", "*")
		req := new(models.UserRequestPairTokens)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		refreshTokenString := req.AccessToken
		cfg := u.ConfigController.Config
		claims, _ := services.Claims(refreshTokenString, cfg.AccessSecret)
		//user, err := u.UserRepository.GetUserById(claims.ID)
		productsInOrderProd, err := u.OrderProductRepository.GetAllProductsByOrderID(claims.ID)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(productsInOrderProd)
		data, _ := json.Marshal(productsInOrderProd)
		w.Write(data)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (u UserController) AddProductsInCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Access-Control-Allow-Origin", "*")
		//req := new([]models.OrderProducts)
		var req []models.OrderProducts

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//var listOP []models.OrderProducts
		////user, err := u.UserRepository.GetUserById(claims.ID)
		//for i := range req{
		//	orderProduct := models.OrderProducts{
		//		Id:               0,
		//		ProductId:        req[i].ProductId,
		//		ProductName:      req[i].ProductName,
		//		OrderId:          req[i].OrderId,
		//		NumbersOfProduct: req[i].NumbersOfProduct,
		//		PurchasePrice:    req[i].PurchasePrice,
		//		AccessToken:      "",
		//		RefreshToken:     "",
		//		CreatedAt:        nil,
		//		UpdatedAt:        nil,
		//	}

		//		}

		err := u.OrderProductRepository.InsertToOrdersProducts(req)
		if err != nil {
			log.Println(err)
		}
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
