package app_config

import "os"

var Port = ":8080"

var STATIC_ROUTE = "/public"
var STATIC_DIR = "/public"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")
	if portEnv != "" {
		Port = portEnv
	}

	statisRouteEnv := os.Getenv("STATIC_ROUTE")
	if statisRouteEnv != "" {
		STATIC_ROUTE = statisRouteEnv
	}

	statisDirEnv := os.Getenv("STATIC_DIR")
	if statisDirEnv != "" {
		STATIC_DIR = statisDirEnv
	}
}
