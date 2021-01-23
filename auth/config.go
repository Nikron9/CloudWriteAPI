package auth

type tokenConfigModel struct {
	tokenSecret string
	tokenExp    string
}

var tokenConfig = tokenConfigModel{
	tokenSecret: "secret",
	tokenExp:    "1h",
}
