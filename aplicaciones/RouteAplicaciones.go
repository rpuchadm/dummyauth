package aplicaciones

import (
	"fmt"
	"log"
	"math/rand/v2"
	"myproject/psql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Aplicaciones struct {
	Id              int        `json:"id"`
	Nombre          string     `json:"nombre"`
	Token           *string    `json:"token"`
	TokenOld        *string    `json:"token_old"`
	TokenLastUpdate *time.Time `json:"token_last_update"`
}

func listAplicaciones() ([]Aplicaciones, error) {
	db, err := psql.PsqlSetup()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	const query = `
		SELECT
			id,nombre,token,token_old,token_last_update
		FROM APLICACIONES`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Aplicaciones
	for rows.Next() {
		var item Aplicaciones
		err := rows.Scan(&item.Id, &item.Nombre,
			&item.Token, &item.TokenOld,
			&item.TokenLastUpdate)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, nil
}

func creaToken32() string {
	token := ""
	for len(token) < 32 {
		intval := rand.Int64()
		if intval < 0 {
			intval = -intval
		}
		token = fmt.Sprintf("%s%x", token, intval)
	}
	if len(token) > 32 {
		token = token[:32]
	}
	return token
}

func creaAplicacion(nombre string) (int, string, error) {
	db, err := psql.PsqlSetup()
	if err != nil {
		return 0, "", err
	}
	defer db.Close()

	token := creaToken32()
	const query = `
		INSERT INTO APLICACIONES (nombre, token)
		VALUES ($1, $2)
		RETURNING id`

	var id int
	err = db.QueryRow(query, nombre, token).Scan(&id)
	if err != nil {
		return 0, "", err
	}
	return id, token, nil
}

func getAplicacion(id int) (*Aplicaciones, error) {
	db, err := psql.PsqlSetup()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	const query = `
		SELECT
			id,nombre,token,token_old,token_last_update
		FROM APLICACIONES
		WHERE id = $1`

	var item Aplicaciones
	err = db.QueryRow(query, id).Scan(&item.Id, &item.Nombre,
		&item.Token, &item.TokenOld,
		&item.TokenLastUpdate)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func actualizaTokenAplicacion(id int) (string, error) {
	db, err := psql.PsqlSetup()
	if err != nil {
		return "", err
	}
	defer db.Close()

	aplicacion, err := getAplicacion(id)
	if err != nil {
		return "", err
	}

	if aplicacion.TokenLastUpdate != nil {
		now := time.Now()
		if now.Sub(*aplicacion.TokenLastUpdate) < 5*time.Minute {
			return "", fmt.Errorf("token updated less than 5 minutes ago")
		}
	}

	token := creaToken32()
	const query = `
		UPDATE APLICACIONES
		SET token_old = token, token = $1,
			token_last_update = NOW()
		WHERE id = $2`

	res, err := db.Exec(query, token, id)
	if err != nil {
		return "", err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rows == 0 {
		return "", fmt.Errorf("no rows affected")
	}
	return token, nil
}

func deleteAplicacion(id int) error {
	db, err := psql.PsqlSetup()
	if err != nil {
		return err
	}
	defer db.Close()

	const query = `
		DELETE FROM APLICACIONES
		WHERE id = $1`

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

type NombreJson struct {
	Nombre string `json:"nombre"`
}

type IdJson struct {
	Id string `json:"id"`
}

func RouteAplicaciones(c *gin.Context) {

	if c.Request.Method == "GET" {

		list, err := listAplicaciones()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, list)

	} else if c.Request.Method == "POST" {

		nombreJson := NombreJson{}
		if err := c.ShouldBindJSON(&nombreJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if nombreJson.Nombre == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nombre is required"})
			return
		}

		id, token, err := creaAplicacion(nombreJson.Nombre)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id, "token": token})

	} else if c.Request.Method == "PUT" {

		idJson := IdJson{}
		if err := c.ShouldBindJSON(&idJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		iid, error := strconv.Atoi(idJson.Id)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be an integer"})
			return
		}

		newtoken, err := actualizaTokenAplicacion(iid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"newtoken": newtoken})

	} else if c.Request.Method == "DELETE" {

		idJson := IdJson{}
		if err := c.ShouldBindJSON(&idJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		iid, error := strconv.Atoi(idJson.Id)
		if error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id must be an integer"})
			return
		}

		err := deleteAplicacion(iid)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Aplicacion deleted"})

	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
	}

}
