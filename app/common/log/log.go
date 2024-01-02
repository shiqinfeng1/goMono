package log

import (
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"

	zlog "github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/rs/zerolog"
	cconf "github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/types"
)

var (
	ErrNotInitedLogger = errors.New("not initialized logger")
	ErrInvalidLogLevel = func(err error) error { return fmt.Errorf("invalid log level:%w", err) }
)

type kloggerWrap struct {
	l       klog.Logger
	svcInfo *types.SrvInfo
	f       *cconf.File
	m       *cconf.Monitor
	lvl     zerolog.Level
}

func (k *kloggerWrap) Log(level klog.Level, keyvals ...interface{}) error {
	return k.l.Log(level, keyvals)
}

var (
	once    sync.Once
	klogger *kloggerWrap
)

func newKLogger(srvInfo *types.SrvInfo, lvl zerolog.Level, fcfg *cconf.File, mcfg *cconf.Monitor) klog.Logger {
	fileName := time.Now().Format(fmt.Sprintf("./log/%v-%v-20060102.log", srvInfo.Name, srvInfo.ID))
	zl := newZeroLogger(fileName, lvl, fcfg, mcfg)
	zlogger := zlog.NewLogger(zl)
	return klog.With(zlogger,
		"svc.id", srvInfo.ID,
		"svc.name", srvInfo.Name,
		"svc.version", srvInfo.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
}

// 全局初始化一次
func New(svcInfo *types.SrvInfo, log *cconf.Log) klog.Logger {
	once.Do(func() {
		lvl, err := zerolog.ParseLevel(log.Level)
		if err != nil {
			lvl = zerolog.DebugLevel
		}
		klogger = &kloggerWrap{
			svcInfo: svcInfo,
			lvl:     lvl,
			f:       log.File,
			m:       log.Monitor,
			l:       newKLogger(svcInfo, lvl, log.File, log.Monitor),
		}
	})
	return klogger
}

func SetLevel(lvlStr string) {
	if klogger == nil {
		klogger.l.Log(klog.LevelError, "SetLevel", ErrNotInitedLogger)
		return
	}
	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		klogger.l.Log(klog.LevelError, "SetLevel", ErrInvalidLogLevel(err))
		return
	}
	// 因为zerolog设置level后会返回一个新的logger，导致无法修改原logger的level，因此对于新的level直接new一个新的
	oldlvl := klogger.lvl
	if klogger.lvl != lvl {
		klogger.l = newKLogger(klogger.svcInfo, lvl, klogger.f, klogger.m)
	}
	klogger.l.Log(klog.LevelWarn, "SetLevel", fmt.Sprintf("change log level '%v' to '%v'", oldlvl, lvl))
}
