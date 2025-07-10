package utils

import "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

func CreatePattern(method string, pathSegments ...string) (string, runtime.Pattern) {
	pattern := runtime.MustPattern(
		runtime.NewPattern(1, generatePatternIndexes(len(pathSegments)), pathSegments, ""),
	)
	return method, pattern
}

// generatePatternIndexes membantu membuat pola angka yang sesuai dengan jumlah segment
func generatePatternIndexes(segmentCount int) []int {
	indexes := []int{}
	for i := 0; i < segmentCount; i++ {
		indexes = append(indexes, 2, i)
	}
	return indexes
}
