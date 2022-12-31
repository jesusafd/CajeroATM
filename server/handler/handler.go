package handler

import (
	"encoding/gob"
	"log"
	"net"

	"github.com/CajeroAutomatico/server/DAO"
	"github.com/CajeroAutomatico/server/models"
)

func Handler(conn net.Conn, cuenta models.Cuenta, login models.Login) {
	var buffer string
	// Limpiamos buffer
	err := gob.NewDecoder(conn).Decode(&buffer)
	if err != nil {
		log.Println("Error al limpiar buffer:", err.Error())
		conn.Close()
		return
	}
	// Enviamos los datos de la cuenta al cliente
	err = gob.NewEncoder(conn).Encode(cuenta)
	if err != nil {
		log.Println("Error al codifcar los datos de la cuenta:", err.Error())
		conn.Close()
		return
	}
	var transaccion models.Transaccion
	var cuentaDAO DAO.CuentaDAO
	for {
		// Leemos las transacciones del cliente
		err = gob.NewDecoder(conn).Decode(&transaccion)
		if err != nil {
			log.Println("Error al decodificar la transaccion :", err.Error())
			return
		}
		log.Println(transaccion)
		switch transaccion.Operacion {
		case 1:
			// Retiro
			if transaccion.Dato > cuenta.Saldo {
				log.Println("Monto a retirar invalido")
			} else {
				cuenta.Saldo -= transaccion.Dato
				err = cuentaDAO.Retiro(cuenta)
				if err != nil {
					log.Println("Error al realizar el retiro : " + err.Error())
					return
				}
			}
			err = gob.NewEncoder(conn).Encode(cuenta)
			if err != nil {
				log.Println("Error al realizar la codificacion de los datos:", err.Error())
				return
			}
		case 2:
			// Cambio nip
			if transaccion.Dato < 1000 || transaccion.Dato > 9999 {
				log.Println("NIP invalido")
			} else {
				cuenta.NIP = transaccion.Dato
				err := cuentaDAO.CambioNip(cuenta)
				if err != nil {
					log.Println("Error al realizar el cambio de nip : " + err.Error())
					return
				}
			}
			err = gob.NewEncoder(conn).Encode(cuenta)
			if err != nil {
				log.Println("Error al realizar la codificacion de los datos:", err.Error())
				return
			}
		case 3:
			// Salir
			login.Acceso = false
			log.Println("El cliente", cuenta.Nombre, " con numero de cuenta:", cuenta.NoCuenta, "termino su transaccion")
			conn.Close()
			return
		default:
			err = gob.NewEncoder(conn).Encode(cuenta)
			if err != nil {
				log.Println("Error al realizar la codificacion de los datos:", err.Error())
				return
			}
		}
	}
}
