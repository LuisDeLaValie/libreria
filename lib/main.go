package main

import (
	"github.com/TDTxLE/libreria/api"
)

func main() {

	api.Start()

	api.Run() // listen and serve on 0.0.0.0:8080
}
