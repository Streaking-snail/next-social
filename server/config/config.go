package config

import (
	"github.com/spf13/viper"
	"github.com/spf13/viper"
)

type Config struct {
	Debug              bool
	Demo               bool
	Container          bool
	DB                 string
	//Server             *Server
	Mysql              *Mysql
	Redis              *Redis
	//Sqlite             *Sqlite
	ResetPassword      string
	ResetTotp          string
	EncryptionKey      string
	EncryptionPassword []byte
	NewEncryptionKey   string
	//Guacd              *Guacd
	//Sshd               *Sshd
}

type Mysql struct {
	Hostname string
	Port     int
	Username string
	Password string
	Database string
}

type Redis struct {
	Name    string
	Type    string
	Address string
	Auth    string
}

func SetupConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("/etc/next-social/")
	viper.AddConfigPath("$HOME/.next-social")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	pflag.String("db", "mysql", "db mode")
	//pflag.String("sqlite.file", path.Join("/usr/local/next-social/data", "sqlite", "next-social.db"), "sqlite db file")
	//MySQL
	pflag.String("mysql.hostname", "127.0.0.1", "mysql hostname")
	pflag.Int("mysql.port", 3306, "mysql port")
	pflag.String("mysql.username", "mysql", "mysql username")
	pflag.String("mysql.password", "mysql", "mysql password")
	pflag.String("mysql.database", "next-social", "mysql database")
	//Redis
	pflag.String("redis.name", "redis", "redis name")
	pflag.String("redis.type", "tcp", "redis type")
	pflag.String("redis.address", "127.0.0.1:6379", "redis address")
	pflag.String("redis.auth", "mysql", "redis auth")

	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return nil, err
	}
	_ = viper.ReadInConfig()

	var config = &Config{
		DB: viper.GetString("db"),
		Mysql: &Mysql{
			Hostname: viper.GetString("mysql.hostname"),
			Port:     viper.GetInt("mysql.port"),
			Username: viper.GetString("mysql.username"),
			Password: viper.GetString("mysql.password"),
			Database: viper.GetString("mysql.database"),
		},
		// Sqlite: &Sqlite{
		// 	File: viper.GetString("sqlite.file"),
		// },
		ResetPassword:    viper.GetString("reset-password"),
		ResetTotp:        viper.GetString("reset-totp"),
		Debug:            viper.GetBool("debug"),
		Demo:             viper.GetBool("demo"),
		Container:        viper.GetBool("container"),
		EncryptionKey:    viper.GetString("encryption-key"),
		NewEncryptionKey: viper.GetString("new-encryption-key"),
	}

	if config.EncryptionKey == "" {
		config.EncryptionKey = "next-social"
	}

	md5Sum := fmt.Sprintf("%x", md5.Sum([]byte(config.EncryptionKey)))
	config.EncryptionPassword = []byte(md5Sum)

	
}

func init() {
	var err error
	GlobalCfg, err = SetupConfig()
	if err != nil {
		panic(err)
	}
}


