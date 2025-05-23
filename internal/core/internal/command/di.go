package command

import (
	"github.com/harluo/boot/internal/core/internal/command/internal"
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(
		newInfo,
		newVersion,
		func(info *Info, version *Version) internal.Put {
			return internal.Put{
				Info:    info,
				Version: version,
			}
		},
	).Build().Apply()
}
