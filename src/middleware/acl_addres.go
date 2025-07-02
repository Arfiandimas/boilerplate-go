package middleware

import (
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/consts"
)

// IPLocation ...
func IPLocation(r *http.Request) string {
	r.RemoteAddr = r.Header.Get("X-Real-IP")
	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	if xff := strings.Trim(r.Header.Get("X-Forward-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			remoteIP = ip.String()
		}
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			remoteIP = ip.String()
		}
	}
	if remoteIP != "::1" {
		aclsIP := strings.Split(os.Getenv("ACL_ADDRESS"), ",")
		for _, i := range aclsIP {
			_, netIP, _ := net.ParseCIDR(i)
			ipB := net.ParseIP(remoteIP)
			if netIP.Contains(ipB) {
				return consts.MiddlewarePassed
			}
			return consts.AuthorizationFailure
		}

	}
	return consts.MiddlewarePassed

}
