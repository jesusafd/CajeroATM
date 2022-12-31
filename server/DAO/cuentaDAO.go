package DAO

import (
	"context"
	"log"
	"time"

	"github.com/CajeroAutomatico/server/database"
	"github.com/CajeroAutomatico/server/models"
)

type CuentaDAO struct {
}

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
