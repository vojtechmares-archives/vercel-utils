package handler

import (
	"fmt"
	"net/http"
	"strings"
)

// Handler Returns remote IP address
func Handler(w http.ResponseWriter, r *http.Request) {
	addr := r.RemoteAddr
	fmt.Fprintf(w, strings.Split(addr, ":")[0])
}
