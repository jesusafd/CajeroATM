package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/CajeroAutomatico/server/models"
)

// Chequeo es la funncion encargada de leer un registro en la bd
// tambien puede ser usado para saber si la cuenta existe
// en caso de que el usuario no exista la funcion devolvera un error
func Chequeo(cuenta models.Cuenta) (models.Cuenta, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var c models.Cuenta

	query := "SELECT * FROM cuenta WHERE noCuenta=?"
	rows, err := DBConn.QueryContext(
		ctx,
		query,
		cuenta.NoCuenta,
	)
	if err != nil {
		log.Println("Error al buscar registro en la base de datos : " + err.Error())
		return c, err
	}
	var nip int
	// Leemos el registro, si es que se encontro
	if rows.Next() {
		err = rows.Scan(&c.NoCuenta, &c.Nombre, &c.Saldo, &nip)
		if err != nil {
			log.Println("Error al leer datos extraidos de la bd : " + err.Error())
			return c, err
		}
	} else {
		err = errors.New("Cuenta no existe")
	}
	if nip != cuenta.NIP {
		log.Println("NIP incorrecto")
		return cuenta, errors.New("nip incorrecto")
	}
	return c, nil

}
