package handlers

import (
	"fmt"
	"net/http"
)

func HandleURLShortenerRequest(w http.ResponseWriter, req *http.Request, urlMap map[string]string) {
	if val, ok := urlMap[req.URL.Path[1:]]; ok {
		http.Redirect(w, req, val, http.StatusSeeOther)
	} else {
		fmt.Fprintf(w, "not found")
	}
}
