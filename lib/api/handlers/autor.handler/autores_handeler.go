package autoreshandlers

import (
	"time"

	autormodels "github.com/TDTxLE/libreria/models/autor.models"
	"github.com/gin-gonic/gin"
)

func ListarHandler(c *gin.Context) {

	res, err := autormodels.ListarAutores()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}
func ObetenerHandler(c *gin.Context) {
	id := c.Param("id")
	res, err := autormodels.ObtenerAutor(id)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}
func CrearHandler(c *gin.Context) {
	var crearautor autormodels.AutorModel
	if err := c.ShouldBindJSON(&crearautor); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	noww := time.Now()
	crearautor.Creado = &noww
	res, err := autormodels.CrearAutor(crearautor)
	if err != nil {
		c.JSON(500, err)
		return
	}

	nuevolibro, err := autormodels.ObtenerAutor(res.Hex())
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nuevolibro)
}
func ActualizarHandler(c *gin.Context) {

	id := c.Param("id")
	var actualizar autormodels.AutorModel
	if err := c.ShouldBindJSON(&actualizar); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	noww := time.Now()
	actualizar.Actualizado = &noww
	err := autormodels.ActualizarAutor(id, actualizar)
	if err != nil {
		c.JSON(500, err)
		return
	}

	nuevolibro, err := autormodels.ObtenerAutor(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nuevolibro)
}
func EliminarHandler(c *gin.Context) {
	id := c.Param("id")
	err := autormodels.EliminarAutor(id)
	if err != nil {
		c.JSON(500, err)
	}
	c.String(200, "Se elimino el libro")
}
