package libroshandlers

import (
	"log"
	"time"

	librosmodels "github.com/TDTxLE/libreria/models/libros.models"
	"github.com/gin-gonic/gin"
)

func ListarHandler(c *gin.Context) {

	res, err := librosmodels.ListarLibros()
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, res)
}
func ObetenerHandler(c *gin.Context) {
	id := c.Param("id")
	res, err := librosmodels.ObtenerLibro(id)
	if err != nil {
		c.JSON(500, err)
	}
	log.Printf("key:%s\nlibro:%v\n", id, res)
	c.JSON(200, res)
}
func CrearHandler(c *gin.Context) {
	var crearlibro librosmodels.LibroModelForm
	if err := c.ShouldBindJSON(&crearlibro); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	noww := time.Now()
	crearlibro.Creado = noww
	res, err := librosmodels.CrearLibro(crearlibro)
	if err != nil {
		c.JSON(500, err)
		return
	}

	nuevolibro, err := librosmodels.ObtenerLibro(res.Hex())
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nuevolibro)
}
func ActualizarHandler(c *gin.Context) {

	id := c.Param("id")
	var actualizar librosmodels.LibroModelForm
	if err := c.ShouldBindJSON(&actualizar); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	noww := time.Now()
	actualizar.Actualizado = noww
	err := librosmodels.ActualizarLibro(id, actualizar)
	if err != nil {
		c.JSON(500, err)
		return
	}

	nuevolibro, err := librosmodels.ObtenerLibro(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nuevolibro)
}
func EliminarHandler(c *gin.Context) {
	id := c.Param("id")
	err := librosmodels.EliminarLibro(id)
	if err != nil {
		c.JSON(500, err)
	}
	c.String(200, "Se elimino el libro")
}
