package api

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Start() {
	r = gin.New()

	iniciarHandlers()

	generarLogs()

}

func Run() {
	r.Run()
}

func generarLogs() {

	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		log := fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)

		// Logging to a file.
		ginlog, _ := os.Create("gin.log")
		defer ginlog.Close()
		io.MultiWriter(ginlog).Write([]byte(log))

		return log
	}))

	r.Use(gin.Recovery())

}

func iniciarHandlers() {
	v1 := r.Group("/prueba/v1")
	{
		v1.GET("/ping", Prubea)
	}
	v2 := r.Group("/prueba/v2")
	{
		v2.GET("/ping", Prubea)
	}
}
