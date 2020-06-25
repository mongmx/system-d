package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opencensus.io/trace"
	"strconv"
)

// TraceMiddleware init
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, span := trace.StartSpan(c, c.Request.URL.String())
		defer func() {
			span.AddAttributes(
				trace.StringAttribute("/http/status_code", strconv.Itoa(c.Writer.Status())),
				trace.StringAttribute("/http/method", c.Request.Method),
				trace.StringAttribute("/http/path", c.Request.RequestURI),
				trace.StringAttribute("/http/host", c.Request.Host),
			)
			span.End()
		}()
		c.Set("span", span.SpanContext())
		c.Next()
	}
}

// GetSpan extracts span from context.
func GetSpan(ctx *gin.Context) (span trace.SpanContext, exists bool) {
	spanI, _ := ctx.Get("span")
	span, ok := spanI.(trace.SpanContext)
	exists = span != trace.SpanContext{} && ok
	return
}
