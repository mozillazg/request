package request

import (
	"fmt"
	"net/http"
)

var DefaultRedirectLimit = 10

func defaultCheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) > DefaultRedirectLimit {
		return fmt.Errorf("stopped after %d redirects", len(via))
	}
	if len(via) == 0 {
		return nil
	}
	// Redirect requests with the first Header
	for key, val := range via[0].Header {
		// Don't copy Referer Header
		if key != "Referer" {
			req.Header[key] = val
		}
	}
	return nil
}

func applyCheckRdirect(a *Args) {
	if a.Client.CheckRedirect == nil {
		a.Client.CheckRedirect = defaultCheckRedirect
	}
}
