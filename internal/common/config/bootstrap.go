package config

import (
	"flag"
	"os"

	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/shiqinfeng1/goMono/internal/common/types"
)

func Bootstrap(srvCfgFile string) (c kcfg.Config) {

	m, _ := types.NewModeFromString(os.Getenv("MODE"))
	if !m.IsValid() {
		panic(m.ErrInvaild())
	}

	if m.Is(types.ModeProduct) || m.Is(types.ModeTest) {
		c = kcfg.New(
			kcfg.WithSource(
				NewNacosSource(m.String(), srvCfgFile),
			),
		)
	}
	if m.Is(types.ModeDevelop) {
		flag.Parse()
		var flagconf string
		flag.StringVar(&flagconf, "conf", "./"+srvCfgFile, "config path, eg: -conf config.yaml")
		c = kcfg.New(
			kcfg.WithSource(
				file.NewSource(flagconf),
			),
		)
	}
	if err := c.Load(); err != nil {
		panic(err)
	}
	return
}
