package api

import (
	"fmt"
	"io"
	"os"
	"time"

	libroshandlers "github.com/TDTxLE/libreria/api/handlers"

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
	libros := r.Group("/api/libros")
	{
		libros.GET("/", libroshandlers.ListarHandler)
		libros.GET("/:id", libroshandlers.ObetenerHandler)
		libros.POST("/", libroshandlers.CrearHandler)
		libros.PUT("/:id", libroshandlers.ActualizarHandler)
		libros.DELETE("/:id", libroshandlers.EliminarHandler)
	}
}
