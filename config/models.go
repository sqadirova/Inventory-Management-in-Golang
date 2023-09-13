package config

import (
	"fmt"
	"time"
)

type ProfileConfigurations struct {
	Profile Profile
}

type Profile struct {
	Active string
}

type Configurations struct {
	App
	Database
	Jwt
	Swagger
	Cookie
	KeycloakConfig
}

type App struct {
	Host string
	Port int
}

type Swagger struct {
	Host   string
	Scheme string
}

type Database struct {
	Host                string
	Port                int
	Dialect             string
	User                string
	DBName              string
	Password            string
	MaxIdleConn         int
	MaxOpenConn         int
	MaxConnLifetimeHour int
	SSLMode             string
	Schema              string
}

func (dbConfig *Database) URL() string {
	dbSource := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSLMode)
	return dbSource
}

type Jwt struct {
	SecretKey         string
	PublicKey         string
	PrivateKey        string
	SaltKey           string
	AccessTokenExpire int64
}

type Cookie struct {
	Name     string
	HttpOnly bool
	Secure   bool
	Expires  time.Duration
	SameSite string
}

type KeycloakConfig struct {
	Path            string
	ClientName      string
	ClientSecretKey string
	RealmName       string
}
