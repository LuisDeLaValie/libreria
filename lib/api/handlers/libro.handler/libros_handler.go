package libroshandlers

import (
	"time"

	librosmodels "github.com/TDTxLE/libreria/models/libros.models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListarHandler(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	filtro := bson.M{}
	// Obtener todos los query parameters
	for key, values := range queryParams {
		if key == "id" || key == "coleccion" {
			objectIDs := make([]primitive.ObjectID, len(values))
			for i, id := range values {
				objID, err := primitive.ObjectIDFromHex(id)
				if err != nil {
					c.String(500, "Error al convertir el ID %s a ObjectID: %s", id, err)
				}
				objectIDs[i] = objID
			}

			if key == "id" {
				key = "_id"
			}

			if len(objectIDs) > 1 {
				filtro[key] = bson.M{"$in": objectIDs}
			} else {
				filtro[key] = objectIDs[0]
			}

		}

		if len(values) > 1 {
			filtro[key] = bson.M{"$in": values}
		} else {
			filtro[key] = values[0]
		}
	}
	res, err := librosmodels.ListarLibros(filtro)
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

	respnse, err := librosmodels.ListarLibros(bson.M{"_id": res})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(500, respnse)

}
