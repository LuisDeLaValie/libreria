package coleccionhandler

import (
	"time"

	collecionmodels "github.com/TDTxLE/libreria/models/colleccion.models"
	"github.com/gin-gonic/gin"
)

func ListarHandler(c *gin.Context) {

	res, err := collecionmodels.ListarColecciones()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}
func ObetenerHandler(c *gin.Context) {
	id := c.Param("id")
	res, err := collecionmodels.ObtenerColeccion(id)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}
func CrearHandler(c *gin.Context) {
	var crearlibro collecionmodels.ColleccionModel
	if err := c.ShouldBindJSON(&crearlibro); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	noww := time.Now()
	crearlibro.Creado = noww
	res, err := collecionmodels.CrearColeccion(crearlibro)
	if err != nil {
		c.JSON(500, err)
		return
	}

	nuevolibro, err := collecionmodels.ObtenerColeccion(res.Hex())
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nuevolibro)
}
func ActualizarHandler(c *gin.Context) {

	id := c.Param("id")
	var actualizar collecionmodels.ColleccionModel
	if err := c.ShouldBindJSON(&actualizar); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	noww := time.Now()
	actualizar.Actualizado = noww
	err := collecionmodels.ActualizarColeccion(id, actualizar)
	if err != nil {
		c.JSON(500, err)
		return
	}

	nuevolibro, err := collecionmodels.ObtenerColeccion(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nuevolibro)
}
func EliminarHandler(c *gin.Context) {
	id := c.Param("id")
	err := collecionmodels.EliminarColeccion(id)
	if err != nil {
		c.JSON(500, err)
	}
	c.String(200, "Se elimino el libro")
}
