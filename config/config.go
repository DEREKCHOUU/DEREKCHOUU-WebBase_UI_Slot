package config

import (
	"fmt"

	_env "github.com/ilyakaznacheev/cleanenv"
)

type (
	Global struct {
		Os  `yaml:"os"`
		App `yaml:"app"`
		DB  `yaml:"db"`
		Log `yaml:"log"`
	}

	Os struct {
		Selectedos string `env-required:"true" yaml:"selectos" env:"OS_NAME"`
		Ip         string `env-default:"192.168.10.10" yaml:"ip" env:"OS_IP"`
		Subnet     int    `env-default:"24" yaml:"subnet" env:"OS_SUBNET"`
		Gateway    string `yaml:"192.168.10.1" env:"OS_GATEWAY"`
		DHCP       bool   `env-default:"true" yaml:"dhcp" env:"OS_DHCP"`
	}

	App struct {
		Uiport int    `env-default:"4041" yaml:"uiport" env:"UI_PORT"`
		UTC    string `env-default:"Asia/Hong_Kong" yaml:"UTC" env:"UI_UTC"`
		NTP    string `env-default:"192.168.0.1" yaml:"NTP" env:"UI_NTP"`
	}

	DB struct {
		SelectDB string `env-default:"mysql" yaml:"selectdb" env:"UI_DB_NAME"`
		DBac     string `env-default:"admin" yaml:"dbac" env:"UI_DB_AC"`
		DBpw     string `env-default:"IWantToDie" yaml:"dbpw" env:"UI_DB_PW"`
		DBip     string `env-default:"localhost" yaml:"dbip" env:"UI_DB_IP"`
		DBport   int    `env-default:"2043" yaml:"dbport" env:"UI_DB_PORT"`
	}

	Log struct {
		Level string `env-default:"debug" yaml:"level" env:"UI_DB_NAME"`
	}
)

func NewConfig() (*Global, error) {
	cfg := &Global{}

	err := _env.ReadConfig("./config/global.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = _env.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func UpdateConfig(cfg ...interface{}) error {
	err := _env.UpdateEnv(&cfg)
	if err != nil {
		return err
	}
	_env.ReadConfig("./config/global.yml", cfg)
	return err
}
