package core

import (
	"flag"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joye61/gotpl/app"
)

type Application struct {
	/// 当前环境，一般值dev|test|prod,默认dev
	Env string
}

/// 获取配置信息，值需要单独断言
func (ctx *Application) Config(key string) any {

	cc := reflect.ValueOf(app.CommonConfig)
	ec := reflect.ValueOf(app.EnvConfig[ctx.Env])

	value := ec.FieldByName(strings.ToUpper(key[:1]) + key[1:])
	if value.IsZero() || value.IsNil() {
		value = cc.FieldByName(strings.ToUpper(key[:1]) + key[1:])
	}

	return value.Interface()
}

var App *Application

/// 初始化配置
/// 设置日志等信息
/// 初始化数据库连接
/// 实例化Gin
/// 初始化注册的路由信息
func Run() {
	// 默认是开发环境
	App = &Application{}

	// 获取当前开发环境
	flag.StringVar(&App.Env, "env", "dev", "Current environment")
	flag.Parse()

	r := gin.Default()
	app.RouterSet(r)
	r.Run(":3000")
}
