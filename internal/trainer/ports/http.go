package ports

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	v1 "github.com/shiqinfeng1/goMono/api/trainer/v1"
	"github.com/shiqinfeng1/goMono/internal/trainer/conf"
	"github.com/shiqinfeng1/goMono/internal/trainer/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, logger log.Logger, tp *trace.TracerProvider, s *service.HttpService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp),
			),
			logging.Server(logger),
			selector.Server(
				jwt.Server(
					func(token *jwt2.Token) (interface{}, error) {
						return []byte(ac.ApiKey), nil
					},
					jwt.WithSigningMethod(jwt2.SigningMethodHS256),
					jwt.WithClaims(func() jwt2.Claims {
						return &jwt2.MapClaims{}
					})),
			).Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	h := openapiv2.NewHandler()
	srv.HandlePrefix("/openapi/", h)
	v1.RegisterTrainerServiceHTTPServer(srv, s)
	return srv
}
