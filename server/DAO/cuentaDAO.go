package DAO

import (
	"context"
	"log"
	"time"

	"github.com/CajeroAutomatico/server/database"
	"github.com/CajeroAutomatico/server/models"
)

// CuentaDAO es la clase que se utilizara para la mainpulacion de la bases de datos
type CuentaDAO struct{}

// Retiro es el metodo encargado de realizar la acutalizacion del saldo en la base de datos
func (cDAO *CuentaDAO) Retiro(cuenta models.Cuenta) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := "UPDATE cuenta SET saldo=? WHERE noCuenta=?"
	_, err := database.DBConn.QueryContext(
		ctx,
		query,
		cuenta.Saldo,
		cuenta.NoCuenta,
	)
	if err != nil {
		log.Println("Error al realizar retiro : " + err.Error())
		return err
	}
	return nil

}

// CambioNip es el metodo encargado de realizar el cambio del nip del cliente en la base de datos
func (cDAO *CuentaDAO) CambioNip(cuenta models.Cuenta) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := "UPDATE cuenta SET nip=? WHERE noCuenta=?"
	_, err := database.DBConn.QueryContext(
		ctx,
		query,
		cuenta.NIP,
		cuenta.NoCuenta,
	)
	if err != nil {
		log.Println("Error al cambiar nip retiro : " + err.Error())
		return err
	}
	return nil

}
