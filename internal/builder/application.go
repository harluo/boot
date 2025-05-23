package builder

import (
	"github.com/harluo/boot/internal"
	"github.com/harluo/boot/internal/builder/internal/function"
	"github.com/harluo/boot/internal/builder/internal/once"
	"github.com/harluo/boot/internal/core"
	"github.com/harluo/boot/internal/internal/application"
	"github.com/harluo/boot/internal/internal/config"
)

var shadow *Application

type Application struct {
	params  *config.Application
	timeout *Timeout
	banner  *Banner
	code    *Code
	help    *Help
}

// NewApplication !基于单例实现，保证每次只生成一个可配置项
func NewApplication() (application *Application) {
	once.Application.Do(newApplication)
	application = shadow

	return
}

func newApplication() {
	shadow = new(Application)
	shadow.params = config.NewApplication()

	// !预创建，保证单例
	shadow.timeout = newTimeout(shadow)
	shadow.banner = newBanner(shadow)
	shadow.help = newHelp(shadow)
	shadow.code = newCode(shadow)

	return
}

func (a *Application) Validate() *Application {
	return a.set(func() {
		a.params.Validate = true
	})
}

func (a *Application) Invalidate() *Application {
	return a.set(func() {
		a.params.Validate = false
	})
}

func (a *Application) Author(name string, email string) *Application {
	return a.set(func() {
		a.params.Authors = append(a.params.Authors, application.NewAuthor(name, email))
	})
}

func (a *Application) Copyright(copyright string) *Application {
	return a.set(func() {
		a.params.Copyright = copyright
	})
}

func (a *Application) Description(description string) *Application {
	return a.set(func() {
		a.params.Description = description
	})
}

func (a *Application) Usage(usage string) *Application {
	return a.set(func() {
		a.params.Usage = usage
	})
}

func (a *Application) Metadata(key string, value any) *Application {
	return a.set(func() {
		a.params.Metadata[key] = value
	})
}

func (a *Application) Name(name string) *Application {
	return a.set(func() {
		internal.Name = name
	})
}

func (a *Application) Timeout() *Timeout {
	return a.timeout
}

func (a *Application) Banner() *Banner {
	return a.banner
}

func (a *Application) Code() *Code {
	return a.code
}

func (a *Application) Help() *Help {
	return a.help
}

func (a *Application) Instance() *core.Application {
	return core.New(a.params)
}

func (a *Application) set(set function.Set) (application *Application) {
	set()
	application = a

	return
}
