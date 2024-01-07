package oauth2

type LoginFormatter struct {
	AccessTokenToken string `json:"access_token"`
	Email            string `json:"email"`
	Expiry           string `json:"expiry"`
}

func LoginFormat(userInfo UserInfo) LoginFormatter {
	var formatter LoginFormatter

	formatter.AccessTokenToken = userInfo.AccessTokenToken
	formatter.Email = userInfo.Email
	formatter.Expiry = userInfo.Expiry

	return formatter
}
