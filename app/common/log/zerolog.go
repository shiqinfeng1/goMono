package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/pkg/errors"

	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/rs/zerolog"
	"github.com/shiqinfeng1/goMono/app/common/client"
	"github.com/shiqinfeng1/goMono/app/common/conf"
	"github.com/shiqinfeng1/goMono/app/common/types"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	timeFormat                = time.RFC3339Nano
	ErrInvalidMoitorAddr      = errors.New("invalid log monitor endpoint")
	ErrFailCreateMoitorClient = func(err error) error { return fmt.Errorf("fail to create monitor client:%w", err) }
)

// zerolog的屏幕输出
func newConsoleWriter() io.Writer {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	return consoleWriter
}

// zerolog的文件输出
func newFileWriter(logPath string, fcfg *conf.File) io.Writer {
	return &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    int(fcfg.GetMaxSize()),    // megabytes
		MaxBackups: int(fcfg.GetMaxBackups()), // file numbers
		MaxAge:     int(fcfg.GetMaxAge()),     // days
		Compress:   fcfg.GetCompress(),        // disabled by default
	}
}

// zerolog的日志监控输出
func mustNewMonitorWriter(endpoint string) io.Writer {
	// EFK: Elasticsearch + Fluentd + Kibana
	if endpoint == "" {
		panic(ErrInvalidMoitorAddr)
	}
	logger, err := client.NewFluentWriteSyncer(endpoint)
	if err != nil { // 如果远程日志服务器链接失败，那么重定向到默认输出（屏幕）
		panic(ErrFailCreateMoitorClient(err))
	}
	return logger
}
func newMonitorWriter(endpoint string) io.Writer {
	logger, err := client.NewFluentWriteSyncer(endpoint)
	if err != nil { // 如果远程日志服务器链接失败，那么重定向到默认输出（屏幕）
		return klog.NewWriter(klog.DefaultLogger, klog.WithWriteMessageKey(fmt.Sprintf("(NO LogMonitor: %v)", err)))
	}
	return logger
}

// 生成一个zero的日志器，支持输出到屏幕、日志文件、远端日志服务
func newZeroLogger(fileName string, lvl zerolog.Level, fcfg *conf.File, mcfg *conf.Monitor) *zerolog.Logger {
	zerolog.TimeFieldFormat = timeFormat
	m := types.NewModeFromString(os.Getenv("MODE"))
	if !m.IsValid() {
		panic(m.ErrInvaild())
	}
	var l zerolog.Logger
	if m.Is(types.ModeDevelop) {
		multi := zerolog.MultiLevelWriter(newConsoleWriter(), newFileWriter(fileName, fcfg))
		l = zerolog.New(multi).With().Timestamp().Stack().Logger().Level(lvl)
	}
	if m.Is(types.ModeTest) {
		multi := zerolog.MultiLevelWriter(newFileWriter(fileName, fcfg), newMonitorWriter(mcfg.GetEndpoint()))
		l = zerolog.New(multi).With().Timestamp().Stack().Logger().Level(lvl)
	}
	if m.Is(types.ModeProduct) {
		multi := zerolog.MultiLevelWriter(mustNewMonitorWriter(mcfg.GetEndpoint()))
		l = zerolog.New(multi).With().Timestamp().Logger().Level(lvl)
	}
	return &l
}
