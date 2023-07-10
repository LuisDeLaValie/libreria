package handlers

import (
	"github.com/TDTxLE/libreria/models"
	librosmodels "github.com/TDTxLE/libreria/models/libros.models"
	utilsmongo "github.com/TDTxLE/libreria/utils/mongoTogolang"
	"github.com/gin-gonic/gin"
)

func AgregarLibroColeccionHandler(c *gin.Context) {

	// obtener el parametro colid
	coleid := c.Param("colid")
	id, err := utilsmongo.ValidarOID(coleid)
	if err != nil {
		c.JSON(500, models.ResposeError{
			Status:  "id no valid",
			Message: "Error al obtener el id",
			Detalle: err,
		})
		return
	}

	// obtener la lista de los libros agregar a la coleccion
	var idsLibros []string
	if err := c.ShouldBindJSON(&idsLibros); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// actualizr los libros
	var lActualizar librosmodels.LibroModelForm
	lActualizar.Collection = *id
	if err := librosmodels.ActualizarVariosLibros(idsLibros, lActualizar); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.String(200, "se agregaron los libros")
}
