package config

import (
	"os"
	"path/filepath"

	"github.com/cockroachdb/errors"
	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

var (
	ErrFileNotExist = func(f string) error { return errors.Newf("config file not exist:%v", f) }
	ErrNoCfgSource  = errors.Newf("no found config source")
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func Bootstrap(cfgFile []string, scanTarget []interface{}, onChanges map[string]func(key string, value kcfg.Value)) {
	runMode, _ := types.NewModeFromString(os.Getenv("MODE"))
	if !runMode.IsValid() {
		panic(runMode.ErrInvaild())
	}
	source := func() []kcfg.Source {
		ks := make([]kcfg.Source, 0)
		for _, f := range cfgFile {
			f := filepath.Join("./config", f)
			switch {
			case runMode.Is(types.ModeProduct):
				ks = append(ks, MustNewNacosSource(runMode.String(), f))
			case runMode.Is(types.ModeTest):
				nacos, err := NewNacosSource(runMode.String(), f)
				if err == nil {
					ks = append(ks, nacos)
				}
				if fileExists(f) {
					ks = append(ks, file.NewSource(f))
				}
				ks = append(ks, file.NewSource(f))
				if len(ks) == 0 {
					panic(ErrNoCfgSource)
				}
			case runMode.Is(types.ModeDevelop):
				if !fileExists(f) {
					panic(ErrFileNotExist(f))
				}
				ks = append(ks, file.NewSource(f))
			}
		}
		return ks
	}()
	c := kcfg.New(kcfg.WithSource(source...))
	if err := c.Load(); err != nil {
		panic(err)
	}
	for _, t := range scanTarget {
		if err := c.Scan(t); err != nil {
			panic(err)
		}
	}
	for item, onChange := range onChanges {
		if err := c.Watch(item, onChange); err != nil {
			panic(err)
		}
	}
}
