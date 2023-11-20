package log

import (
	"sync"

	zlog "github.com/go-kratos/kratos/contrib/log/zerolog/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/rs/zerolog"
)

type kloggerWrap struct {
	l                      klog.Logger
	svcID, svcName, svcVer string
	lvl                    zerolog.Level
}

func (k *kloggerWrap) Log(level klog.Level, keyvals ...interface{}) error {
	return k.l.Log(level, keyvals)
}

var once sync.Once
var klogger *kloggerWrap

func newKLogger(svcID, svcName, svcVer string, lvl zerolog.Level, endpoint ...string) klog.Logger {
	zl := NewZeroLogger(svcID, svcName, lvl)
	zlogger := zlog.NewLogger(zl)
	return klog.With(zlogger,
		"svc.id", svcID,
		"svc.name", svcName,
		"svc.version", svcVer,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
}

// 全局初始化一次
func New(svcID, svcName, svcVer string, lvlStr string, endpoint ...string) klog.Logger {
	once.Do(func() {
		lvl, err := zerolog.ParseLevel(lvlStr)
		if err != nil {
			lvl = zerolog.DebugLevel
		}
		klogger = &kloggerWrap{
			svcID:   svcID,
			svcName: svcName,
			svcVer:  svcVer,
			lvl:     lvl,
			l:       newKLogger(svcID, svcName, svcVer, lvl, endpoint...),
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
	// 因为zerolog设置level后会返回一个新的logger，无法修改原logger的level，因此对于新的level直接new一个新的
	if klogger.lvl != lvl {
		klogger.l = newKLogger(klogger.svcID, klogger.svcName, klogger.svcVer, lvl)
	}
}
