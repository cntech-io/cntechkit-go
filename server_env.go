package cntechkitgo

import (
	"fmt"

	"github.com/cntech-io/cntechkit-go/utils"
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
		DebugModeFlag:  utils.GetBooleanEnv(string(DEBUG_MODE_FLAG), false),
		AppPort:        utils.GetStringEnv(string(APP_PORT), false),
		TrustedProxies: utils.GetStringArrayEnv(string(TRUSTED_PROXIES), ",", false),
	}
}
