package net

import (
	"fmt"
	"net"
	"os"
)

// GetIP get IP with options
func GetIP(opts ...Option) string {
	option := decodeOpts(opts)
	envIP := os.Getenv(option.env)
	if len(envIP) > 0 {
		return envIP
	}

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addresses {
		if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				return inet.IP.String()
			}
		}
	}
	return ""
}
