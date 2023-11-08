package client

import (
	"net"
	"time"
)

// 等待指定的服务地址可用
func waitForPort(addr string, timeout time.Duration) bool {
	portAvailable := make(chan struct{})
	timeoutCh := time.After(timeout)

	go func() {
		for {
			select {
			case <-timeoutCh:
				return
			default:
				// continue
			}

			_, err := net.Dial("tcp", addr)
			if err == nil {
				close(portAvailable)
				return
			}

			time.Sleep(time.Millisecond * 200)
		}
	}()

	select {
	case <-portAvailable:
		return true
	case <-timeoutCh:
		return false
	}
}
