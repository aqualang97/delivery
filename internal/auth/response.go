package repositories

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

//type RefreshResponse struct {
//	NewAccessToken   string
//	NewRefreshToken  string
//	UserID           int
//	AccessExpiredAt  time.Time
//	RefreshExpiredAt time.Time
//	Expired          string
//}

type RegistrationResponse struct {
}

//type UserResponse struct {
//	ID    int
//	Email string
//	Name  string
//}
