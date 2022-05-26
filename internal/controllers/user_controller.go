package controllers

import (
	"delivery/internal/auth/services"
	"delivery/internal/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (u UserController) Profile(w http.ResponseWriter, r *http.Request) {
	//allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token"

	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		log.Println("Profile start")
		log.Println(r.Header.Get("Authorization"))
		tokenString := services.GetTokenFromBearerString(r.Header.Get("Authorization"))
		log.Println(tokenString)
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
		log.Println("profile end", claims.ID)

	default:
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
}

func (u UserController) Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		log.Println("ref start")

		req := new(models.UserRequestPairTokens)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Println(err)
			log.Println(err)
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cfg := u.ConfigController.Config
		refreshTokenString := services.GetTokenFromBearerString(req.RefreshToken)

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
		resp := models.UserResponsePairTokens{
			UserID:       claims.ID,
			AccessToken:  newAccessTokenString,
			RefreshToken: newRefreshTokenString,
		}

		log.Println("AAA", newAccessTokenString, "\nRRR", newRefreshTokenString)
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(resp)
		w.Write(data)
		//json.NewEncoder(w).Encode(newAccessTokenString)
		//json.NewEncoder(w).Encode(newRefreshTokenString)
		log.Println("refresh end", claims.ID)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (u UserController) GetProductsInCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
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

func (u UserController) AddProductsFromCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//req := new([]models.OrderProducts)
		req := new(models.RequestOrderProducts)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println(req.Cart, req.PaymentMethod, req.ContactData)
		var fullPrise float32
		for _, v := range req.Cart {
			fullPrise += v.PurchasePrice * float32(v.NumbersOfProduct)
		}
		status := ""
		if req.PaymentMethod == "Cash" {
			status = "cash"
		} else {
			status = "not paid"
		}
		order := models.Order{
			UserID:        req.ContactData.UserID,
			Price:         fullPrise,
			Status:        status,
			PaymentMethod: req.PaymentMethod,
		}
		orderId, err := u.OrderRepository.InsertToOrders(order)
		if err != nil {
			log.Println(err)
		}
		var orderProducts []models.OrderProducts
		for _, v := range req.Cart {
			orderProduct := models.OrderProducts{
				ProductId:        v.ProductId,
				OrderId:          orderId,
				NumbersOfProduct: v.NumbersOfProduct,
				PurchasePrice:    v.PurchasePrice,
			}
			orderProducts = append(orderProducts, orderProduct)
		}

		_ = u.OrderProductRepository.InsertToOrdersProducts(orderProducts)
		w.WriteHeader(http.StatusOK)
		data, _ := json.Marshal(orderId)
		w.Write(data)
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (u UserController) SimulationCardPay(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//simulation
		req := new(models.RequestCardPay)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		log.Println(req.Card, req.LastOrderId, req.Status)
		w.WriteHeader(http.StatusOK)
		err := u.OrderRepository.UpdateOrdersByID(req)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}

	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

func (u UserController) GetOldOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := r.URL.Path
		parts := strings.Split(path, "/")
		userId, _ := strconv.Atoi(parts[2])
		//log.Println(userId)
		orders, _ := u.OrderRepository.GetOldOrdersByUserID(userId)
		//log.Println(orders)
		json.NewEncoder(w).Encode(orders)
		log.Println(orders)

	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
func (u UserController) A(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
	default:
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}

//func (u UserController) A(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "POST":
//	default:
//		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
//	}
//}
//
//func (u UserController) UserData(w http.ResponseWriter, r *http.Request) {
//	switch r.Method {
//	case "POST":
//		w.Header().Add("Access-Control-Allow-Origin", "*")
//	default:
//		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
//	}
//}
