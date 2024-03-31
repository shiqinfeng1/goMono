package conf

import (
	"os"
	"path/filepath"

	"github.com/cockroachdb/errors"
	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/shiqinfeng1/goMono/utils/types"
)

var (
	ErrFileNotExist = func(f string) error { return errors.Newf("config file not exist:%v", f) }
	ErrNoCfgSource  = errors.Newf("no found config source")
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func getExecutablePath() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

type (
	CfgOnChanges map[string]func(key string, value kcfg.Value)
	ScanTarget   struct {
		File   string
		Field  string
		Target interface{}
	}
)

// 根据运行环境，加载指定的配置文件cfgFile， 并解析到指定的结构scanTarget，注册配置更新回调函数onChanges
func Bootstrap(scanTarget []ScanTarget, onChanges CfgOnChanges) {
	// 运行环境从环境变量中获取
	runMode := types.NewModeFromString(os.Getenv("MODE"))
	if !runMode.IsValid() {
		panic(runMode.ErrInvaild())
	}
	// 根据运行环境，获取配置源
	source := func() []kcfg.Source {
		ks := make([]kcfg.Source, 0)
		for _, f := range scanTarget {
			fpath := filepath.Join(getExecutablePath(), "config", f.File)
			if !fileExists(fpath) {
				fpath = filepath.Join(getExecutablePath(), "..", "config", f.File)
			}
			switch {
			case runMode.Is(types.ModeProduct): // 生产环境：从配置中心获取配置，如果获取不到则panic
				ks = append(ks, MustNewNacosSource(runMode.String(), f.File))
			case runMode.Is(types.ModeTest): // 测试环境：从配置中心或本地文件获取，
				nacos, err := NewNacosSource(runMode.String(), fpath)
				if err == nil { // 存在则导入
					ks = append(ks, nacos)
				}
				if fileExists(fpath) { // 存在则导入
					ks = append(ks, file.NewSource(fpath))
				}
				// 如果2个源都没有获取到配置，那么panic
				if len(ks) == 0 {
					panic(ErrNoCfgSource)
				}
			case runMode.Is(types.ModeDevelop): // 开发环境：从本地路径获取配置文件
				if !fileExists(fpath) {
					panic(ErrFileNotExist(fpath))
				}
				ks = append(ks, file.NewSource(fpath))
			}
		}
		return ks
	}()
	// 配置配资源
	c := kcfg.New(kcfg.WithSource(source...))
	// 加载配置
	if err := c.Load(); err != nil {
		panic(err)
	}
	// 解析配置数据
	for _, t := range scanTarget {
		if t.Field != "" {
			if err := c.Value(t.Field).Scan(t.Target); err != nil {
				panic(err)
			}
		} else {
			if err := c.Scan(t.Target); err != nil {
				panic(err)
			}
		}
	}
	// 注册配置动态更新回调函数
	for item, onChange := range onChanges {
		if err := c.Watch(item, onChange); err != nil {
			panic(err)
		}
	}
}
