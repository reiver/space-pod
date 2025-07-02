package cfg

import (
	"fmt"

	"github.com/reiver/space-pod/env"
)

func WebServerTCPAddress() string {
	return fmt.Sprintf(":%s", env.TcpPort)
}
