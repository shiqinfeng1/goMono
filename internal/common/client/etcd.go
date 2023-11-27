package client

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

// []string{"127.0.0.1:2379"}
func NewEtcd(endpoints []string) (*clientv3.Client, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}
