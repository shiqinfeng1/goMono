package cmd

import (
	"os"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/shiqinfeng1/goMono/utils/types"
)

// go build -ldflags "-X cmd.Version=x.y.z"
var (
	Name    = "training"    // Name is the name of the compiled software.
	Version string          // Version is the version of the compiled software.
	ID, _   = os.Hostname() // 主机信息
	svcInfo = &types.SrvInfo{
		ID:      ID,
		Name:    Name,
		Version: Version,
	}
)

func init() {
	Bootstrap()
}

var Main = gcmd.Command{
	Name:  "main",
	Usage: "main <sub-command>",
	Brief: "this is main command, please specify a sub command",
}
