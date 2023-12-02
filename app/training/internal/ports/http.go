package ports

import (
	"context"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	kjwt "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	kmetrics "github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	v1 "github.com/shiqinfeng1/goMono/api/training/v1"
	"github.com/shiqinfeng1/goMono/app/common/client"
	conf "github.com/shiqinfeng1/goMono/app/common/config/training"
	"github.com/shiqinfeng1/goMono/app/training/internal/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

// 设置白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/shop.interface.v1.ShopInterface/Login"] = struct{}{}
	whiteList["/shop.interface.v1.ShopInterface/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.HTTP, ac *conf.Auth, logger log.Logger, tp *trace.TracerProvider, s *service.HttpService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp),
			),
			logging.Server(logger),
			selector.Server(
				kjwt.Server(
					func(token *jwt.Token) (interface{}, error) {
						return []byte(ac.ApiKey), nil
					},
					kjwt.WithSigningMethod(jwt.SigningMethodHS256),
					kjwt.WithClaims(func() jwt.Claims {
						return &jwt.MapClaims{}
					})),
			).Match(NewWhiteListMatcher()).Build(),
			kmetrics.Server(
				kmetrics.WithSeconds(prom.NewHistogram(client.MetricsSeconds)),
				kmetrics.WithRequests(prom.NewCounter(client.MetricsRequests)),
			),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Network != "" {
		opts = append(opts, http.Network(c.Network))
	}
	if c.Addr != "" {
		opts = append(opts, http.Address(c.Addr))
	}
	if c.Timeout != nil {
		opts = append(opts, http.Timeout(c.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/openapi/", h)
	v1.RegisterTrainingServiceHTTPServer(srv, s)
	return srv
}
