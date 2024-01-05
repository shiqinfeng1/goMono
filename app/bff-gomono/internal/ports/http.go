package ports

import (
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
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v1 "github.com/shiqinfeng1/goMono/api/bff/v1"
	conf "github.com/shiqinfeng1/goMono/app/bff-gomono/internal/conf"
	"github.com/shiqinfeng1/goMono/app/bff-gomono/internal/service"
	"github.com/shiqinfeng1/goMono/app/common/client"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(httpAddr *conf.HTTP, ac *conf.Auth, logger log.Logger, s *service.HttpService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(log.With(logger,
				"layer", "ports",
			)),
			selector.Server(
				kjwt.Server(
					func(token *jwt.Token) (interface{}, error) {
						return []byte(ac.ApiKey), nil
					},
					kjwt.WithSigningMethod(jwt.SigningMethodHS256),
					kjwt.WithClaims(func() jwt.Claims {
						return &jwt.MapClaims{}
					})),
			).Build(),
			kmetrics.Server(
				kmetrics.WithSeconds(prom.NewHistogram(client.MetricsSeconds)),
				kmetrics.WithRequests(prom.NewCounter(client.MetricsRequests)),
			),
			tracing.Server(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if httpAddr.Network != "" {
		opts = append(opts, http.Network(httpAddr.Network))
	}
	if httpAddr.Addr != "" {
		opts = append(opts, http.Address(httpAddr.Addr))
	}
	if httpAddr.Timeout != nil {
		opts = append(opts, http.Timeout(httpAddr.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/openapi/", h)
	v1.RegisterBFFHTTPServer(srv, s)
	return srv
}
