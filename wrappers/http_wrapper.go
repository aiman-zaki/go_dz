package wrappers

import (
	"fmt"
	"net/http"
	"strings"
)
// t
func cleanRemoteIP(s string) string {
	idx := strings.LastIndex(s, ":")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

func getRemoteIP(r *http.Request) string {
	hdr := r.Header
	hdrRealIP := hdr.Get("X-Real-Ip")
	hdrForwardedFor := hdr.Get("X-Forwarded-For")
	hdrForwardedIP := ""
	if hdrForwardedIP == "" && hdrForwardedFor == "" {
		return cleanRemoteIP(r.RemoteAddr)
	}
	if hdrForwardedFor != "" {
		parts := strings.Split(hdrForwardedFor, ",")
		for i, p := range parts {
			parts[i] = strings.TrimSpace(p)
		}
		return parts[0]
	}
	return hdrRealIP

}

//LogRequest : get ip of incoming connection and log
func LogRequest(r *http.Request, message string) {
	r.Header.Get("X-FORWARDED-FOR")

	fmt.Println(message + r.Method + ":" + r.RemoteAddr)
}
