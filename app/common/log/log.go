package log

import (
	"fmt"

	"github.com/pkg/errors"

	zlog "github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/rs/zerolog"
	cconf "github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/container"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

var (
	ErrNotInitedLogger = errors.New("not initialized logger")
	ErrInvalidLogLevel = func(err error) error { return fmt.Errorf("invalid log level:%w", err) }
)

type KloggerWrap struct {
	zlogger *zerolog.Logger
	klogger klog.Logger
	svcInfo *types.SrvInfo
	f       *cconf.File
	m       *cconf.Monitor
	oldLvl  zerolog.Level
}

var (
	klogger *KloggerWrap
)

// 全局初始化一次
func New(svcInfo *types.SrvInfo, logCfg *cconf.Log) klog.Logger {
	var fileName string
	lvl, err := zerolog.ParseLevel(logCfg.Level)
	if err != nil {
		lvl = zerolog.DebugLevel
	}

	klogger = &KloggerWrap{
		svcInfo: svcInfo,
		oldLvl:  lvl,
		f:       logCfg.File,
		m:       logCfg.Monitor,
	}
	name, err := container.GetNameLite()
	if err != nil {
		klog.Errorf("get myself container name: %v", err)
		fileName = fmt.Sprintf("./log/%v-%v.log", svcInfo.Name, svcInfo.ID)
	} else {
		fileName = fmt.Sprintf("./log/%v.log", name)
	}
	klogger.zlogger = newZeroLogger(fileName, logCfg.File, logCfg.Monitor)
	klogger.klogger = klog.With(zlog.NewLogger(klogger.zlogger),
		"svc.id/name/ver", fmt.Sprintf("%v/%v/%v", svcInfo.ID, svcInfo.Name, svcInfo.Version),
	)
	klogger.zlogger.Level(lvl)
	return klogger.klogger
}

func SetLevel(lvlStr string) {
	if klogger == nil {
		panic(ErrNotInitedLogger)
	}
	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		klogger.klogger.Log(klog.LevelError, "err", ErrInvalidLogLevel(err))
		return
	}
	// 因为zerolog设置level后会返回一个新的logger，导致无法修改原logger的level，因此对于新的level直接new一个新的
	if klogger.oldLvl != lvl {
		klogger.zlogger.Level(lvl)
	}
	klogger.klogger.Log(klog.LevelWarn, "msg", fmt.Sprintf("change log level '%v' to '%v'", klogger.oldLvl, lvl))
	klogger.oldLvl = lvl
}
