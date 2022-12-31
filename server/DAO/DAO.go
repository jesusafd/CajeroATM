package DAO

import "github.com/CajeroAutomatico/server/models"

// DAO es la interface que implementara la clase CuentaDAO
type DAO interface {
	Retiro(c models.Cuenta, retiro int) error
	CambiarNip(c models.Cuenta, nip int) error
}
