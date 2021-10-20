package handler

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	info := fmt.Sprintf("welcome! 您当前的ip: %serverMux", r.RemoteAddr)
	w.Write([]byte(info))
}
