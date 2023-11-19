package config

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/shiqinfeng1/goMono/internal/common/vars"
)

func Bootstrap(srvCfgFile string) (c config.Config) {

	m, _ := vars.NewModeFromString(os.Getenv("MODE"))
	if !m.IsValid() {
		panic(m.ErrInvaild())
	}

	if m.Is(vars.ModeProduct) || m.Is(vars.ModeTest) {
		c = config.New(
			config.WithSource(
				NewNacosSource(m.String(), srvCfgFile),
			),
		)
	}
	if m.Is(vars.ModeDevelop) {
		flag.Parse()
		var flagconf string
		flag.StringVar(&flagconf, "conf", "./"+srvCfgFile, "config path, eg: -conf config.yaml")
		c = config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
	}
	if err := c.Load(); err != nil {
		panic(err)
	}
	return
}
