package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// DBConn es el objeto conexion con la base de datos
var DBConn *sqlx.DB

// NewConexion Funcion encargada de crear la conexion con la
// base de datos
func NewConnection() error {
	settingsDB, err := SetSettings()
	if err != nil {
		// En caso de haber un error la funcion SetSettings
		// lo notificara en la consola
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var infoDB string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", settingsDB.User, settingsDB.Password, settingsDB.Host, settingsDB.Port, settingsDB.Name)

	var db *sqlx.DB
	// Creamos la conexion con la BD
	db, err = sqlx.ConnectContext(ctx, "mysql", infoDB)
	if err != nil {
		log.Println("Error al crear la conexion con la BD\n" + err.Error())
		return err
	}
	// Seteamos el numero maximo de conexiones
	db.SetMaxOpenConns(5)
	// La creacion de la conexion fue exitosa
	DBConn = db
	return nil
}

// CheckConnection permite comprobar el estado de la conexion
func CheckConnection() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Hacemos un ping a la bd para saber el estado de esta
	err := DBConn.PingContext(ctx)
	if err != nil {
		log.Println("Conexion con la base de datos fallida" + err.Error())
		return false
	}
	log.Println("Conexion exitosa con la BD")
	return true
}
