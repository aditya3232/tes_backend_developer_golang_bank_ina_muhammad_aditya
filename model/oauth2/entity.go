package oauth2

// struct for dto not model field
type UserInfo struct {
	AccessTokenToken string `json:"access_token"`
	Email            string `json:"email"`
	Expiry           string `json:"expiry"`
}
