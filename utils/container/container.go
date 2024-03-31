package container

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/client"
)

var (
	ErrGetDockerClientFail  = func(err error) error { return fmt.Errorf("get container client:%w", err) }
	ErrContainerInspectFail = func(err error) error { return fmt.Errorf("container inspect:%w", err) }
	ErrContainerNameInvalid = func(name string) error { return fmt.Errorf("container name invalid:%v", name) }
)

// 容器客户端
func getDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
		client.WithHost("unix:///run/docker.sock"),
	)
	if err != nil {
		return nil, ErrGetDockerClientFail(err)
	}
	return cli, nil
}

// 获取容器名称
func GetName() (string, error) {
	cli, err := getDockerClient()
	if err != nil {
		return "", err
	}
	// HOSTNAME 就是docker yaml中的hostname值
	container, err := cli.ContainerInspect(context.Background(), os.Getenv("HOSTNAME"))
	if err != nil {
		return "", ErrContainerInspectFail(err)
	}
	name := strings.TrimPrefix(container.Name, "/")
	return name, nil
}

// 获取容器名称的前面部分，去除容器名称中系统追加的后缀
func GetNameLite() (string, error) {
	name, err := GetName()
	if err != nil {
		return "", err
	}
	names := strings.Split(name, ".")
	if len(names) != 3 {
		return "", ErrContainerNameInvalid(name)
	}
	return strings.Join(names[:2], "."), nil
}
