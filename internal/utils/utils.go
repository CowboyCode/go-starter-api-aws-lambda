package utils

import (
	"net"
	"net/http"
)

func GetRemoteAddress(r *http.Request) string {
	// grab ip address
	an, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	return an
}
