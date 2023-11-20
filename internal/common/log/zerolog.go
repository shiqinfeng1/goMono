package log

import (
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/rs/zerolog"
	"github.com/shiqinfeng1/goMono/internal/common/vars"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	timeFormat = time.RFC3339Nano
)

// zerolog的屏幕输出
func newConsoleWriter() io.Writer {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: timeFormat}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	consoleWriter.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s;", i)
	}
	return consoleWriter
}

// zerolog的文件输出
func newFileWriter(fileName string) io.Writer {
	return &lumberjack.Logger{
		Filename:   fmt.Sprintf("./log/%v.log", fileName),
		MaxSize:    100,  // megabytes
		MaxBackups: 3,    // file numbers
		MaxAge:     28,   // days
		Compress:   true, // disabled by default
	}
}

type fluentWriteSyncer struct {
	output *fluent.Fluent
}

func newFluentWriteSyncer(addr string) (*fluentWriteSyncer, error) {
	c := fluent.Config{
		Async:              true,
		ForceStopAsyncSend: true,
	}
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	switch u.Scheme {
	case "tcp":
		host, port, err2 := net.SplitHostPort(u.Host)
		if err2 != nil {
			return nil, err2
		}
		if c.FluentPort, err = strconv.Atoi(port); err != nil {
			return nil, err
		}
		c.FluentNetwork = u.Scheme
		c.FluentHost = host
	case "unix":
		c.FluentNetwork = u.Scheme
		c.FluentSocketPath = u.Path
	default:
		return nil, fmt.Errorf("unknown network: %s", u.Scheme)
	}
	// 检查地址是否存在服务
	_, err = net.Dial(u.Scheme, u.Host)
	if err != nil {
		return nil, err
	}
	fl, err := fluent.New(c)
	if err != nil {
		return nil, err
	}
	return &fluentWriteSyncer{
		output: fl,
	}, nil
}

func (x *fluentWriteSyncer) Write(data []byte) (n int, err error) {
	if err := x.output.Post("zerolog", data); err != nil {
		return 0, err
	}
	return len(data), nil
}

// zerolog的日志监控输出
func newMonitorWriter(endpoint ...string) io.Writer {
	// EFK: Elasticsearch + Fluentd + Kibana
	var addr string
	if len(endpoint) != 0 {
		addr = endpoint[0]
	} else {
		addr = "tcp://127.0.0.1:24224"
	}
	logger, err := newFluentWriteSyncer(addr)
	if err != nil {
		return klog.NewWriter(klog.DefaultLogger, klog.WithWriteMessageKey(fmt.Sprintf("(NO Monitor: %v)", err)))
	}
	return logger
}

// 生成一个zero的日志器，支持输出到屏幕、日志文件、远端日志服务
func NewZeroLogger(svcID, svcName string, lvl zerolog.Level, endpoint ...string) *zerolog.Logger {
	zerolog.TimeFieldFormat = timeFormat
	m, _ := vars.NewModeFromString(os.Getenv("MODE"))
	if !m.IsValid() {
		panic(m.ErrInvaild())
	}
	fileName := svcID + "-" + svcName
	var l zerolog.Logger
	if m.Is(vars.ModeDevelop) {
		multi := zerolog.MultiLevelWriter(newConsoleWriter(), newFileWriter(fileName))
		l = zerolog.New(multi).With().Timestamp().Caller().Stack().Logger().Level(lvl)
	}
	if m.Is(vars.ModeTest) {
		multi := zerolog.MultiLevelWriter(newFileWriter(fileName), newMonitorWriter(endpoint...))
		l = zerolog.New(multi).With().Timestamp().Caller().Stack().Logger().Level(lvl)
	}
	if m.Is(vars.ModeProduct) {
		multi := zerolog.MultiLevelWriter(newFileWriter(fileName), newMonitorWriter(endpoint...))
		l = zerolog.New(multi).With().Timestamp().Logger().Level(lvl)
	}
	return &l
}
