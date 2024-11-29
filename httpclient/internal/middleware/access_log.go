package middleware

import (
	"bytes"
	"httpclient/global"
	"httpclient/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w AccessLogWriter) Write(b []byte) (int, error) {
	if n, err := w.Body.Write(b); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(b)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.Body.String(),
		}
		global.Logger.WithFields(fields).Infof(c, "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
