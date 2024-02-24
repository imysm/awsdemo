/**
 * @Author: aesoper
 * @Description:
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 2020/5/20 0:01
 */

package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	LogMap map[string]interface{}
	// ResponseWriter 用于读取返回信息
	ResponseWriter struct {
		gin.ResponseWriter
		body *bytes.Buffer
	}

	// AfterResponseFunc NOTICE:
	AfterResponseFunc func(ctx *gin.Context, l LogMap)
)

func (w ResponseWriter) Write(b []byte) (int, error) {
	if n, err := w.body.Write(b); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriter) WriteString(b string) (int, error) {
	if n, err := w.body.WriteString(b); err != nil {
		return n, err
	}
	return w.ResponseWriter.WriteString(b)
}

func DefaultLogger() gin.HandlerFunc {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return NewLogger(func(ctx *gin.Context, aLog LogMap) {
		if ctx.Writer.Status() == http.StatusOK {
			logger.WithFields(logrus.Fields(aLog)).Info()
		} else {
			logger.WithFields(logrus.Fields(aLog)).Error()
		}
	})
}

func NewLogger(fs ...AfterResponseFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responseWriter := &ResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = responseWriter

		startTime := time.Now()

		// 读取body之后记得重写body,否则下个处理链就获取不到上下文了
		body, err := ioutil.ReadAll(ctx.Request.Body)
		defer ctx.Request.Body.Close()
		if err == nil {
			buf := bytes.NewBuffer(body)
			ctx.Request.Body = ioutil.NopCloser(buf)
		}

		ctx.Next()

		endTime := time.Now()
		timeConsuming := time.Since(startTime).Nanoseconds() / 1e6

		requestId := ctx.Request.Header.Get("X-Request-ID")
		if requestId == "" {
			requestId = ctx.Writer.Header().Get("X-Request-ID")
		}

		accessLogMap := map[string]interface{}{
			"request_time":        startTime,
			"request_method":      ctx.Request.Method,
			"request_uri":         ctx.Request.RequestURI,
			"request_proto":       ctx.Request.Proto,
			"request_ua":          ctx.Request.UserAgent(),
			"request_referer":     ctx.Request.Referer(),
			"request_post_data":   ctx.Request.PostForm.Encode(),
			"request_query_data":  ctx.Request.URL.Query(),
			"request_client_ip":   ctx.ClientIP(),
			"request_body":        string(body),
			"request_header":      ctx.Request.Header,
			"request_id":          requestId,
			"request_path":        ctx.Request.URL.Path,
			"request_host":        ctx.Request.Host,
			"request_remote_addr": ctx.Request.RemoteAddr,

			"response_time": endTime,
			//"response_data":   responseWriter.body.String(),
			"response_status": ctx.Writer.Status(),
			"response_size":   ctx.Writer.Size(),
			"cost_time":       timeConsuming,
		}

		for _, f := range fs {
			f(ctx, accessLogMap)
		}
	}
}
