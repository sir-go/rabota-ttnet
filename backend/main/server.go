package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct{ srv *http.Server }

func (s *Server) Run() {
	gin.SetMode(CFG.Service.RunMode)
	r := gin.New()

	r.Use(
		gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s %s%3d%s %s%-7s%s %s [%d %s]\n%s",
				param.TimeStamp.Format("06/01/02 15:04:05"),
				param.StatusCodeColor(), param.StatusCode, param.ResetColor(),
				param.MethodColor(), param.Method, param.ResetColor(),
				param.Path,
				param.BodySize,
				param.Latency,
				param.ErrorMessage,
			)
		}),
		gin.Recovery(),
	)
	if CFG.Service.RunMode == gin.DebugMode {
		r.Use(cors.Default())
	}

	rApi := r.Group("api")
	rV1 := rApi.Group("v1")
	rProfile := rV1.Group("profile")
	rProfile.POST("/", SRV.Add)

	addr := fmt.Sprintf("%s:%d", CFG.Service.Host, CFG.Service.Port)
	LOG.Println("run web-server on http://" + addr)

	s.srv = &http.Server{
		Addr:           addr,
		WriteTimeout:   CFG.Service.Timeouts.Write.Duration,
		ReadTimeout:    CFG.Service.Timeouts.Read.Duration,
		IdleTimeout:    CFG.Service.Timeouts.Idle.Duration,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		eh(err)
	}
}

func (s *Server) Shutdown() {
	LOG.Println("shutdown web-server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	eh(s.srv.Shutdown(ctx))
}
