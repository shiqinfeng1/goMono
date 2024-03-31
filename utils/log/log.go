package log

import (
	"fmt"

	"github.com/pkg/errors"

	zlog "github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/rs/zerolog"
	cconf "github.com/shiqinfeng1/goMono/utils/conf"
	"github.com/shiqinfeng1/goMono/utils/container"
	"github.com/shiqinfeng1/goMono/utils/types"
)

var (
	ErrNotInitedLogger = errors.New("not initialized logger")
	ErrInvalidLogLevel = func(err error) error { return fmt.Errorf("invalid log level:%w", err) }
)

// 对kratos的log的封装结构
type KloggerWrap struct {
	zlogger *zerolog.Logger // zerolog的实例
	klogger klog.Logger     // kratos的log的接口
	svcInfo *types.SrvInfo  // 服务信息
	f       *cconf.File     // 日志文件的配置
	m       *cconf.Monitor  // 日志服务器的地址
	oldLvl  zerolog.Level   // 记录当前日志等级，在动态更新等级时，该值作为旧等级使用
}

var kloggerWrap *KloggerWrap

// 全局初始化一次
func New(svcInfo *types.SrvInfo, logCfg *cconf.Log) klog.Logger {
	var fileName string
	// 读取日志等级
	lvl, err := zerolog.ParseLevel(logCfg.Level)
	if err != nil {
		lvl = zerolog.DebugLevel
	}

	kloggerWrap = &KloggerWrap{
		svcInfo: svcInfo,
		oldLvl:  lvl,
		f:       logCfg.File,
		m:       logCfg.Monitor,
	}
	// 以容器名称作为日志文件名称，用于多副本时，区分日志文件，防止写入冲突
	name, err := container.GetNameLite()
	if err != nil { // 不是容器运行时，使用配置的名称命名
		klog.Errorf("get myself container name: %v", err)
		fileName = fmt.Sprintf("./log/%v-%v.log", svcInfo.Name, svcInfo.ID)
	} else {
		fileName = fmt.Sprintf("./log/%v.log", name)
	}
	// 生产一个zerolog实例
	kloggerWrap.zlogger = newZeroLogger(fileName, logCfg.File, logCfg.Monitor)
	// 用kratos对zerolog进行一层包装， 带有一个全局的输出字段svc.id/name/ver"
	kloggerWrap.klogger = klog.With(zlog.NewLogger(kloggerWrap.zlogger),
		"svc.id/name/ver", fmt.Sprintf("%v/%v/%v", svcInfo.ID, svcInfo.Name, svcInfo.Version),
	)
	// 设置zerolog的服务等级
	kloggerWrap.zlogger.Level(lvl)
	return kloggerWrap.klogger
}

func SetLevel(lvlStr string) {
	if kloggerWrap == nil {
		panic(ErrNotInitedLogger)
	}
	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		kloggerWrap.klogger.Log(klog.LevelError, "err", ErrInvalidLogLevel(err))
		return
	}
	if kloggerWrap.oldLvl != lvl {
		kloggerWrap.zlogger.Level(lvl)
	}
	kloggerWrap.klogger.Log(klog.LevelWarn, "msg", fmt.Sprintf("change log level '%v' to '%v'", kloggerWrap.oldLvl, lvl))
	kloggerWrap.oldLvl = lvl
}
