package log

import (
	"sync"

	zlog "github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/rs/zerolog"
	"github.com/shiqinfeng1/goMono/internal/common/config"
	"github.com/shiqinfeng1/goMono/internal/common/types"
)

type kloggerWrap struct {
	l        klog.Logger
	svcInfo  *types.SrvInfo
	endpoint string
	lvl      zerolog.Level
}

func (k *kloggerWrap) Log(level klog.Level, keyvals ...interface{}) error {
	return k.l.Log(level, keyvals)
}

var once sync.Once
var klogger *kloggerWrap

func newKLogger(srvInfo *types.SrvInfo, lvl zerolog.Level, endpoint string) klog.Logger {
	zl := newZeroLogger(srvInfo.ID, srvInfo.Name, lvl, endpoint)
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
func New(svcInfo *types.SrvInfo, log *config.Log) klog.Logger {
	once.Do(func() {
		lvl, err := zerolog.ParseLevel(log.Level)
		if err != nil {
			lvl = zerolog.DebugLevel
		}
		klogger = &kloggerWrap{
			svcInfo: svcInfo,
			lvl:     lvl,
			l:       newKLogger(svcInfo, lvl, log.Endpoint),
		}
	})
	return klogger
}

func SetLevel(lvlStr string) {
	if klogger == nil {
		return
	}
	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		return
	}
	// 因为zerolog设置level后会返回一个新的logger，导致无法修改原logger的level，因此对于新的level直接new一个新的
	if klogger.lvl != lvl {
		klogger.l = newKLogger(klogger.svcInfo, lvl, klogger.endpoint)
	}
}
