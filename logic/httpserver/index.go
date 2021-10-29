package httpserver

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintf("welcome! 您当前的ip: %s", r.RemoteAddr)
	_, _ = w.Write([]byte(info))
}
