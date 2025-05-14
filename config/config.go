package config

type ConfigDatabase struct {
	AppName  string `env:"APP_NAME"`
	AppEnv   string `env:"APP_ENV"`
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	LogLevel string `env:"LOG_LEVEL"`
}

var Cfg = ConfigDatabase{
	AppName:  "TRONICS",
	AppEnv:   "DEV",
	Port:     "8081",
	Host:     "localhost",
	LogLevel: "ERROR",
}
