package coleccionhandler

import (
	"log"
	"time"

	"github.com/TDTxLE/libreria/models"
	collecionmodels "github.com/TDTxLE/libreria/models/colleccion.models"
	librosmodels "github.com/TDTxLE/libreria/models/libros.models"
	utilsmongo "github.com/TDTxLE/libreria/utils/mongoTogolang"
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

func AgregarLiroHandler(c *gin.Context) {

	/* defer func() {
		err := recover()
		if err != nil {
			c.JSON(500, err)
		}
	}() */

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
	log.Printf("ids:%v\nactualizar:%s", idsLibros, lActualizar.ToJson())
	if err := librosmodels.ActualizarVariosLibros(idsLibros, lActualizar); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.String(200, "se agregaron los libros")
}

func RemoverLirosHandler(c *gin.Context) {

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
	if err := librosmodels.EliminarVariosLibros(idsLibros); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.String(200, "se agregaron los libros")
}
