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
	UserService   string
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
	AllowOrigins string
}

type Database struct {
	Connection      string
	DbName          string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	ConnMaxIdleTime int
}


func NewConfig() *Config {

	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./configs")
	v.AddConfigPath("/etc/config")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// # Default value ////////////////////////////
	v.SetDefault("app.port", 3000)
	v.SetDefault("app.name", "taskmanagement-backend")
	v.SetDefault("app.env", "development")
	v.SetDefault("app.prefork", false)
	v.SetDefault("app.disable_startup_message", false)
	v.SetDefault("app.enable_trusted_proxy_check", false)
	v.SetDefault("app.cors.allow_origins", "*")
	// # //////////////////////////////////////////

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
			AllowOrigins: v.GetString("app.cors.allow_origins"),
		},
	}
	cfg.Database = Database{
		DbName: v.GetString("app.database.dbname"),
		Connection: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			v.GetString("app.database.dbhost"),
			v.GetString("app.database.dbuser"),
			v.GetString("app.database.dbpassword"),
			v.GetString("app.database.dbname"),
			v.GetString("app.database.dbport"),
			v.GetString("app.database.sslmode"),
		),
	}
	cfg.UserService = v.GetString("app.user_service.url")

	return &cfg
}
