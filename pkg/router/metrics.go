package router

import (
	"fmt"
	"github.com/Gebes/there/v2"
	"runtime"
)

func GetMetrics(request there.HttpRequest) there.HttpResponse {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	return there.Json(there.StatusOK, map[string]string{
		"alloc":      fmt.Sprint(bToMb(m.Alloc), "mb"),
		"TotalAlloc": fmt.Sprint(bToMb(m.TotalAlloc), "mb"),
		"Sys":        fmt.Sprint(bToMb(m.Sys), "mb"),
		"NumGC":      fmt.Sprint(m.NumGC),
		"GoVersion":  runtime.Version(),
	})
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
