package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	AdminApi       adminApi
	Log            log
}

type adminApi rest.RestConf

type log struct {
	Level    string
	Path     string
	KeepDays int8
}
