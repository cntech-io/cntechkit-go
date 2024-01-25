package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

type ServerEnv struct {
	DebugModeFlag  bool
	AppPort        string
	TrustedProxies []string
}

type ServerEnvName string

const (
	DEBUG_MODE_FLAG ServerEnvName = "DEBUG_MODE_FLAG"
	APP_PORT        ServerEnvName = "APP_PORT"
	TRUSTED_PROXIES ServerEnvName = "TRUSTED_PROXIES"
)

func NewServerEnv() *ServerEnv {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	return &ServerEnv{
		DebugModeFlag:  GetBoolean(string(DEBUG_MODE_FLAG), false),
		AppPort:        GetString(string(APP_PORT), false),
		TrustedProxies: GetStringArray(string(TRUSTED_PROXIES), ",", false),
	}
}
