package limiter

import (
	"github.com/labstack/echo"
)
type Limiter struct {
	rules []Rule
	matcher *matcher
}

// Allow ...
func (limiter *Limiter) Allow(path string, metadata map[string]string) bool {
	node := limiter.matcher.find(path, metadata)
	if node == nil {
		return true
	}
	return node.fn.(*counter).allow()
}

func (limiter *Limiter) update() {
	var rules =make([]Rule, 10)
	for _, rule := range rules {
		limiter.matcher.add(rule.Path, rule.Metadata, newCounter(&rule))
	}
}

func Filter() echo.MiddlewareFunc {
	var limiter = new(Limiter)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			var metadata = map[string]string{
				"method": ctx.Request().Method,
				"aid": ctx.Request().Header.Get("Aid"),
				"schema": ctx.Scheme(),
			}
			if ok := limiter.Allow(ctx.Path(), metadata); ok {
				return ctx.JSON(419, nil)
			}
			return
		}
	}
}

/*
method: GET|POST|DELETE ...
header: AID:*
url: /api/config/*
 */
type Rule struct {
	Path string // 路径
	Metadata map[string][]string
	Threshold uint64 // 每秒的限制数
}
// GET:AID:1020: