package pprof

import (
	"net/http"
	_ "net/http/pprof"
	"testing"
)

func TestPprof(t *testing.T) {
	http.ListenAndServe("0.0.0.0:6666", nil)
}
