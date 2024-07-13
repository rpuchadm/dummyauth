package ini

import (
	"myproject/psql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouteDrop(c *gin.Context) {

	db, err := psql.PsqlSetup()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Eliminar la tabla
	msg, err := DropTableAplicaciones(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func RouteInit(c *gin.Context) {

	db, err := psql.PsqlSetup()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	// Crear la tabla
	msg, err := CreateTableAplicaciones(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}
