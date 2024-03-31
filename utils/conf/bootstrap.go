package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	kcfg "github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/shiqinfeng1/goMono/utils/types"
)

var (
	ErrFileNotExist = func(f string) error { return errors.New(fmt.Sprintf("config file not exist:%v", f)) }
	ErrNoCfgSource  = errors.New("not found config source")
)

// 判断文件是否存在
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// 获取当前可执行程序的所在路径
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
		File   string // 配置文件名
		Field  string // 配置字段名
		Target interface{}
	}
)

// 根据运行环境，加载指定的配置文件cfgFile， 并解析到指定的结构scanTarget，注册配置更新回调函数onChanges
func Bootstrap(scanTarget []ScanTarget, onChanges CfgOnChanges) {
	// 检查运行模式： product,develop,test
	runMode := types.NewModeFromString(os.Getenv("MODE"))
	if !runMode.IsValid() {
		panic(runMode.ErrInvaild())
	}
	// 根据运行环境，获取配置源
	source := func() []kcfg.Source {
		ks := make([]kcfg.Source, 0)
		// 处理多个配置文件
		for _, f := range scanTarget {
			// 获取配置文件的路径，如果当前路径不存在，那么查询上一级的config文件夹
			fpath := filepath.Join(getExecutablePath(), "config", f.File)
			if !fileExists(fpath) {
				fpath = filepath.Join(getExecutablePath(), "..", "config", f.File)
			}
			// 根据运行环境获取配资源
			switch {

			// product生产环境：只从配置中心获取配置，如果获取不到则panic
			case runMode.Is(types.ModeProduct):
				ks = append(ks, MustNewNacosSource(runMode.String(), f.File))

			// test测试环境：从配置中心或本地文件获取，
			case runMode.Is(types.ModeTest):
				nacos, err := NewNacosSource(runMode.String(), fpath)
				if err == nil { // 配资源存在，则导入
					ks = append(ks, nacos)
				}
				if fileExists(fpath) { // 本地配置存在，则导入
					ks = append(ks, file.NewSource(fpath))
				}
				// 如果2个源都没有获取到配置，那么panic
				if len(ks) == 0 {
					panic(ErrNoCfgSource)
				}

			// 开发环境：只从本地路径获取配置文件
			case runMode.Is(types.ModeDevelop):
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
	// 解析所有配置数据
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
	// 注册配置动态更新的回调函数
	for item, onChange := range onChanges {
		if err := c.Watch(item, onChange); err != nil {
			panic(err)
		}
	}
}
