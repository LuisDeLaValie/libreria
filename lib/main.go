package main

import (
	"github.com/TDTxLE/libreria/api"
)

func main() {

	/* // Abrir el archivo de registro en modo de añadir (append)
	var err error
	logFile, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Configurar el logger para que escriba en el archivo de registro
	log.SetOutput(logFile)

	api.Start()

	api.Run() // listen and serve on 0.0.0.0:8080

	// Cerrar el archivo de registro
	logFile.Close()

	// Finalizar los tests
	os.Exit(0) */

	api.Start()

	api.Run() // listen and serve on 0.0.0.0:8080
}
