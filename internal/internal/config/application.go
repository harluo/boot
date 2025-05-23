package config

import (
	"sync"

	"github.com/harluo/boot/internal/internal/application"
	"github.com/harluo/boot/internal/internal/constant"
)

var once sync.Once
var shadow = new(Application)

type Application struct {
	// 帮助
	Help *Help
	// 徽标
	Banner *Banner
	// 状态码
	Code *Code
	// 超时
	Timeout *Timeout

	// 合法性验证，包括
	// 启动器构造方法合法性验证
	// 构造方法合法性验证
	Validate bool
	// 应用描述
	Description string
	// 使用方式
	Usage string

	// 版权
	Copyright string
	// 作者
	Authors application.Authors
	// 元数据
	Metadata map[string]any
}

func NewApplication() (application *Application) {
	once.Do(func() {
		shadow = newApplication()
	})
	application = shadow

	return
}

func newApplication() *Application {
	return &Application{
		Help:    newHelp(),
		Banner:  newBanner(),
		Code:    newCode(),
		Timeout: NewTimeout(),

		Validate:    true,
		Description: "一个使用github.com/harluo/boot构建的应用程序，可以使用应用程序提供的命令来使用本程序",
		Usage:       "一个使用github.com/harluo/boot构建的应用程序",

		Copyright: constant.Copyright,
		Authors: application.Authors{{
			Name:  constant.AuthorName,
			Email: constant.AuthorEmail,
		}},
		Metadata: make(map[string]any),
	}
}
