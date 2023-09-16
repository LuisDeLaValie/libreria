package libroshandlers

import (
	"time"

	librosmodels "github.com/TDTxLE/libreria/models/libros.models"
	"github.com/TDTxLE/libreria/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListarHandler(c *gin.Context) {
	// Obtener todos los query parameters
	queryParams := c.Request.URL.Query()
	// Convertir los query parameters a un filtro BSON
	filter := utils.QueryParamsToBson(queryParams)

	res, err := librosmodels.ListarLibros(filter)
	if err != nil {
		c.JSON(500, err)
	}
	response := map[string]interface{}{
		"count":  len(res),
		"libros": res,
	}
	c.JSON(200, response)
}
func ObetenerHandler(c *gin.Context) {
	id := c.Param("id")
	res, err := librosmodels.ObtenerLibro(id)
	if err != nil {
		c.JSON(500, err)
	}
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
func CrearVariosHandler(c *gin.Context) {
	var crearlibros []librosmodels.LibroModelForm
	if err := c.ShouldBindJSON(&crearlibros); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := librosmodels.CrearvariosLibros(crearlibros)
	if err != nil {
		c.JSON(500, err)
		return
	}

	respnse, err := librosmodels.ListarLibros(&bson.M{"_id": bson.M{"$in": res}})
	if err != nil {
		c.JSON(500, err)
		return
	}

	response2 := map[string]interface{}{
		"count":  len(respnse),
		"libros": respnse,
	}
	c.JSON(200, response2)

}
