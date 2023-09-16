package api

import (
	"fmt"
	// "io"
	// "os"
	"time"

	autoreshandlers "github.com/TDTxLE/libreria/api/handlers/autor.handler"
	coleccionhandler "github.com/TDTxLE/libreria/api/handlers/coleccion.handler"
	libroshandlers "github.com/TDTxLE/libreria/api/handlers/libro.handler"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Start() {
	r = gin.New()

	generarLogs()

	iniciarHandlers()

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

		// // Logging to a file.
		// ginlog, _ := os.Create("gin.log")
		// defer ginlog.Close()
		// io.MultiWriter(ginlog).Write([]byte(log))

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
		libros.POST("/crearvarios", libroshandlers.CrearVariosHandler)
	}

	autores := r.Group("/api/autores")
	{
		autores.GET("/", autoreshandlers.ListarHandler)
		autores.GET("/:id", autoreshandlers.ObetenerHandler)
		autores.POST("/", autoreshandlers.CrearHandler)
		autores.PUT("/:id", autoreshandlers.ActualizarHandler)
		autores.DELETE("/:id", autoreshandlers.EliminarHandler)
	}

	colecciones := r.Group("/api/colecciones")
	{
		colecciones.GET("/", coleccionhandler.ListarHandler)
		colecciones.GET("/:id", coleccionhandler.ObetenerHandler)
		colecciones.POST("/", coleccionhandler.CrearHandler)
		colecciones.PUT("/:id", coleccionhandler.ActualizarHandler)
		colecciones.DELETE("/:id", coleccionhandler.EliminarHandler)
		colecciones.PUT("/agregarlibros/:colid", coleccionhandler.AgregarLiroHandler)
		colecciones.PUT("/removerlibros/:colid", coleccionhandler.RemoverLirosHandler)
	}
}
