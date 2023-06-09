package config

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var GlobalCfg *Config

type Config struct {
	Debug              bool
	Demo               bool
	Container          bool
	DB                 string
	Mysql              *Mysql
	Redis              *Redis
	ResetPassword      string
	EncryptionKey      string
	EncryptionPassword []byte
	NewEncryptionKey   string
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
		ResetPassword:    viper.GetString("reset-password"),
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

	return config, nil

}

func init() {
	var err error
	GlobalCfg, err = SetupConfig()
	if err != nil {
		panic(err)
	}
}
