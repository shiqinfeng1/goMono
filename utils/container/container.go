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

func GetName() (string, error) {
	cli, err := getDockerClient()
	if err != nil {
		return "", err
	}
	container, err := cli.ContainerInspect(context.Background(), os.Getenv("HOSTNAME"))
	if err != nil {
		return "", ErrContainerInspectFail(err)
	}
	name := strings.TrimPrefix(container.Name, "/")
	return name, nil
}

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
