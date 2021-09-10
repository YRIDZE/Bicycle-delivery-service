package requests

import (
	"net/http"
	"strconv"
	"strings"
)

func Params(r *http.Request) (int, error) {
	fields := strings.Split(r.URL.String(), "/")
	return strconv.Atoi(fields[len(fields)-1])
}
