package main

import (
	"flag"
	"fmt"


type Params struct {
	ConfigPath string
}

const AppConfigKey string = "FileHandlerContext"

type AppConfig struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func (appConfig AppConfig) GetKey() string {
	return AppConfigKey
}

func GetAppConfigFromContext(ctx context_handler.Icontext) *AppConfig {
	if ok := ctx.CheckHandlerIsSetInContext(AppConfigKey); !ok {
		// TODO: add error handler to context
		return nil
	}
	return ctx.GetHandlerFromContext(AppConfigKey).(*AppConfig)
}

func main() {
	// main
	ctx := CreateNewContext()

	// main -> read command params
	prms := Params{}
	flag.StringVar(&prms.ConfigPath, "host", "configs.json", "server host")
	flag.Parse()

	ctx.AddAppConfigToContext(prms.ConfigPath)

	// set and execute load method
	ctx.LoadUtils()

	// in constructors as a DI
	config := GetAppConfigFromContext(ctx)
	fmt.Println(config)
}

// application core
type NewContext struct {
	context_handler.Icontext
}

func CreateNewContext() *NewContext {
	return &NewContext{
		context_handler.CreateContext(),
	}
}
func (ctx *NewContext) AddAppConfigToContext(configPath string) {
	cfg := &AppConfig{}
	ctx.Icontext.AddHandlerToContext(cfg)
	ctx.Icontext.AddUtil(config.InitConfig(cfg, configPath))
}
