package config

import (
	"github.com/MamushevArup/server/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Debug  *bool `yaml:"is_debug"`
	Listen `yaml:"listen"`
}
type Listen struct {
	BindIp string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

var ins *Config

func ReadConfig() *Config {
	lg := logger.NewLogger()
	lg.Info("Read configuration")
	ins = &Config{}
	if err := cleanenv.ReadConfig("config.yml", ins); err != nil {
		desc, _ := cleanenv.GetDescription(ins, nil)
		lg.Info(desc)
		lg.Fatal(err)
	}

	return ins
}
