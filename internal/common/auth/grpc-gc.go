package auth

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/compute/metadata"
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"
)

// 基于谷歌云的grpc安全通信用的token，需要实现PerRPCCredentials接口
// 如果服务部署在谷歌云上，则使用该grpc安全通信选项
type metadataServerToken struct {
	serviceURL string
}

func NewGcGrpcServerToken(grpcAddr string) credentials.PerRPCCredentials {
	// based on https://cloud.google.com/run/docs/authenticating/service-to-service#go
	// service need to have https prefix without port
	serviceURL := "https://" + strings.Split(grpcAddr, ":")[0]

	return metadataServerToken{
		serviceURL: serviceURL,
	}
}

// GetRequestMetadata is called on every request, so we are sure that token is always not expired
func (t metadataServerToken) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	// based on https://cloud.google.com/run/docs/authenticating/service-to-service#go
	tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", t.serviceURL)
	idToken, err := metadata.Get(tokenURL)
	if err != nil {
		return nil, errors.Wrap(err, "cannot query id token for gRPC")
	}

	return map[string]string{
		"authorization": "Bearer " + idToken,
	}, nil
}

func (metadataServerToken) RequireTransportSecurity() bool {
	return true
}
