package checker

import (
	"github.com/harluo/boot/internal/application"
)

type Arguments interface {
	Arguments() []application.Argument
}
