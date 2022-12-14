package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	zlog "github.com/rs/zerolog/log"

	"main/internal/store"
)

// Server has one REST endpoint /api/v1/profile,
// POST request pushes a JSON data to the table work.profile

func Run(srvConfig Config, dbConfig store.Config) (*http.Server, error) {
	gin.SetMode(srvConfig.RunMode)
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
	if srvConfig.RunMode == gin.DebugMode {
		r.Use(cors.Default())
	}

	rApi := r.Group("api")
	rV1 := rApi.Group("v1")
	rProfile := rV1.Group("profile")
	rProfile.POST("/", func(context *gin.Context) {
		context.Set("dbConfig", dbConfig)
	}, store.AddProfile)

	addr := fmt.Sprintf("%s:%d", srvConfig.Host, srvConfig.Port)
	//goland:noinspection HttpUrlsUsage
	zlog.Info().Str("address", "http://"+addr)

	srv := &http.Server{
		Addr:           addr,
		WriteTimeout:   srvConfig.Timeouts.Write,
		ReadTimeout:    srvConfig.Timeouts.Read,
		IdleTimeout:    srvConfig.Timeouts.Idle,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return nil, err
	}
	return srv, nil
}
