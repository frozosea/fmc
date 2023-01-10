package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"os"
)

type PostgresConfig struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
}

func (p *PostgresConfig) Url() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Host,
		p.Port,
		p.Username,
		p.Password,
		p.DatabaseName)
}

type RedisConfig struct {
	Url string
	Ttl string
}

type EnvVariables struct {
	*PostgresConfig
	*RedisConfig
	RedisUrl             string
	RedisTtl             string
	TwoCaptchaApiKey     string
	SitcServiceUsername  string
	SitcServicePassword  string
	SitcServiceBasicAuth string
	SitcAccessToken      string
	AltsKey              string
}

func getDatabase(config *PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.Url())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getEnvVariable(variableName string) (string, error) {
	variable := os.Getenv(variableName)
	if variable == "" {
		return "", errors.New(fmt.Sprintf(`no %s env variable`, variableName))
	}
	return variable, nil
}

func getEnvVariables() (*EnvVariables, error) {
	variables := map[string]string{
		"POSTGRES_HOST":           "",
		"POSTGRES_USERNAME":       "",
		"POSTGRES_PASSWORD":       "",
		"POSTGRES_DATABASE_NAME":  "",
		"POSTGRES_PORT":           "",
		"REDIS_URL":               "",
		"REDIS_TTL":               "",
		"TWO_CAPTCHA_API_KEY":     "",
		"SITC_SERVICE_USERNAME":   "",
		"SITC_SERVICE_PASSWORD":   "",
		"SITC_SERVICE_BASIC_AUTH": "",
		"SITC_ACCESS_TOKEN":       "",
		"ALTS_KEY":                "",
	}
	for name := range variables {
		v, err := getEnvVariable(name)
		if err != nil {
			return nil, err
		}
		variables[name] = v
	}
	return &EnvVariables{
		PostgresConfig: &PostgresConfig{
			Host:         variables["POSTGRES_HOST"],
			Username:     variables["POSTGRES_USERNAME"],
			Password:     variables["POSTGRES_PASSWORD"],
			DatabaseName: variables["POSTGRES_DATABASE_NAME"],
			Port:         variables["POSTGRES_PORT"],
		},
		RedisConfig: &RedisConfig{
			Url: variables["REDIS_URL"],
			Ttl: variables["REDIS_TTL"],
		},
		TwoCaptchaApiKey:     variables["TWO_CAPTCHA_API_KEY"],
		SitcServiceUsername:  variables["SITC_SERVICE_USERNAME"],
		SitcServicePassword:  variables["SITC_SERVICE_PASSWORD"],
		SitcServiceBasicAuth: variables["SITC_SERVICE_BASIC_AUTH"],
		SitcAccessToken:      variables["SITC_ACCESS_TOKEN"],
		AltsKey:              variables["ALTS_KEY"],
	}, nil
}
