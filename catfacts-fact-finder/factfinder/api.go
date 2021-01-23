package factfinder

import (
	"fmt"
	"net/http"
)

func healthCheck(resp http.ResponseWriter, req *http.Request) {
	// add some health checks here if required
	fmt.Fprintf(resp, "Ok")
}
