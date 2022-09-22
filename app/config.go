package app

import (
	"github.com/joye61/gotpl/def"
)

// 任意环境都通用的配置
var CommonConfig = def.AppConfig{
	Port: 3003,
}

// 不同环境的相关配置，如果没有找到，则选择通用配置
var EnvConfig = map[string]def.AppConfig{
	"dev":  {},
	"test": {},
	"prod": {},
}
