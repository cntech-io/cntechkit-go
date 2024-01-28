package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

type ServerEnv struct {
	DebugModeFlag              bool
	AppPort                    string
	TrustedProxies             []string
	AccessTokenSecret          string
	RefreshTokenSecret         string
	AccessTokenExpireDuration  int
	RefreshTokenExpireDuration int
}

type ServerEnvName string

const (
	DEBUG_MODE_FLAG               ServerEnvName = "DEBUG_MODE_FLAG"
	APP_PORT                      ServerEnvName = "APP_PORT"
	TRUSTED_PROXIES               ServerEnvName = "TRUSTED_PROXIES"
	ACCESS_TOKEN_SECRET           ServerEnvName = "ACCESS_TOKEN_SECRET"
	REFRESH_TOKEN_SECRET          ServerEnvName = "REFRESH_TOKEN_SECRET"
	ACCESS_TOKEN_EXPIRE_DURATION  ServerEnvName = "ACCESS_TOKEN_EXPIRE_DURATION"
	REFRESH_TOKEN_EXPIRE_DURATION ServerEnvName = "REFRESH_TOKEN_EXPIRE_DURATION"
)

func NewServerEnv() *ServerEnv {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	return &ServerEnv{
		DebugModeFlag:              GetBoolean(string(DEBUG_MODE_FLAG), false),
		AppPort:                    GetString(string(APP_PORT), false),
		TrustedProxies:             GetStringArray(string(TRUSTED_PROXIES), ",", false),
		AccessTokenSecret:          GetString(string(ACCESS_TOKEN_SECRET), false),
		RefreshTokenSecret:         GetString(string(REFRESH_TOKEN_SECRET), false),
		AccessTokenExpireDuration:  GetNumber(string(ACCESS_TOKEN_EXPIRE_DURATION), false),
		RefreshTokenExpireDuration: GetNumber(string(REFRESH_TOKEN_EXPIRE_DURATION), false),
	}
}
