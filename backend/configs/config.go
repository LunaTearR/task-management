package configs

import (
	"fmt"
	"strings"

	v "github.com/spf13/viper"
)

type Config struct {
	XThreads      int
	SmallXThreads int
	TinyThreads   int
	App           App
	Database      Database
	Jwt           Jwt
}

type App struct {
	Name                  string
	Port                  string
	Prefork               bool
	DisableStartupMessage bool
	Env                   string
	Cors                  Cors
}

type Cors struct {
	AllowedOrigins []string
}

type Database struct {
	Connection      string
	DbName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	ConnMaxIdleTime int
}

type Jwt struct {
	SecretKey string
}

func NewConfig() *Config {
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")
	v.AddConfigPath("/etc/config")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetDefault("app.port", 3000)
	v.SetDefault("app.name", "task-management-backend")
	v.SetDefault("app.env", "development")
	v.SetDefault("app.prefork", false)
	v.SetDefault("app.disable_startup_message", false)
	v.SetDefault("app.enable_trusted_proxy_check", false)
	v.SetDefault("app.cors.allow_origins", "*")

	v.SetDefault("app.xthreads", 32)
	v.SetDefault("app.small_xthreads", 16)
	v.SetDefault("app.tiny_xthreads", 8)

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("error reading config file: ", err)
	}

	var cfg Config
	cfg.XThreads = v.GetInt("app.xthreads")
	cfg.SmallXThreads = v.GetInt("app.small_xthreads")
	cfg.TinyThreads = v.GetInt("app.tiny_xthreads")

	cfg.App = App{
		Name:                  v.GetString("app.name"),
		Port:                  v.GetString("app.port"),
		Prefork:               v.GetBool("app.prefork"),
		DisableStartupMessage: v.GetBool("app.disable_startup_message"),
		Env:                   v.GetString("app.env"),
		Cors: Cors{
			AllowedOrigins: v.GetStringSlice("app.cors.allow_origins"),
		},
	}

	cfg.Database = Database{
		Connection:      v.GetString("app.database.connection"),
		DbName:          v.GetString("app.database.db_name"),
		MaxOpenConns:    v.GetInt("app.database.max_open_conns"),
		MaxIdleConns:    v.GetInt("app.database.max_idle_conns"),
		ConnMaxLifetime: v.GetInt("app.database.conn_max_lifetime"),
		ConnMaxIdleTime: v.GetInt("app.database.conn_max_idle_time"),
	}

	cfg.Jwt = Jwt{
		SecretKey: v.GetString("app.jwt.secret_key"),
	}

	return &cfg
}
