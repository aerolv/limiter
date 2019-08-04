package limiter

import (
	"fmt"
	"testing"
)

func TestLimiter_Allow(t *testing.T) {
	limiter := Limiter{
		rules: make([]Rule, 0),
		matcher:newRouter(),
	}

	rule := Rule{
		Path:      "/api/*/config",
		Metadata:  map[string][]string{
			"aid": []string{"1000", "1001"},
		},
		Threshold: 1000,
	}

	limiter.matcher.add(rule.Path, rule.Metadata, newCounter(&rule))
	for _, route := range limiter.matcher.routes {
		fmt.Printf("route = %+v\n", route)
	}

	ret := limiter.Allow("/api/v1/config3", map[string]string{
		"aid":"1000",
	})
	fmt.Printf("ret = %+v\n", ret)
}