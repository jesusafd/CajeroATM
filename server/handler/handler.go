package handler

import (
	"encoding/gob"
	"log"
	"net"

	"github.com/CajeroAutomatico/server/DAO"
	"github.com/CajeroAutomatico/server/models"
)

// Handler es el manejador el cual se encarga de direccionar las transacciones a sus respectivas fuciones
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
	// Creamos un ciclo para que el usario pueda realizar distintas transacciones
	for {
		// Leemos las transacciones del cliente
		err = gob.NewDecoder(conn).Decode(&transaccion)
		if err != nil {
			log.Println("Error al decodificar la transaccion :", err.Error())
			return
		}
		switch transaccion.Operacion {
		case 1:
			// Retiro
			// verificamos si el monto a retirar es mayor al saldo, de ser asi no sera posible realizar el retiro
			// en caso contrario si se realizara
			if transaccion.Dato > cuenta.Saldo {
				log.Println("Monto a retirar invalido")
			} else {
				cuenta.Saldo -= transaccion.Dato
				// Realizamos el retiro, en caso de no existir error el retiro fue exitoso
				err = cuentaDAO.Retiro(cuenta)
				if err != nil {
					log.Println("Error al realizar el retiro : " + err.Error())
					// Si no se realizo el retiro en bd se suma el dinero a la cuenta de nuevo
					cuenta.Saldo += transaccion.Dato
					return
				}
			}
			// Enviamos los datos de la cuenta al cliente
			err = gob.NewEncoder(conn).Encode(cuenta)
			if err != nil {
				log.Println("Error al realizar la codificacion de los datos:", err.Error())
				return
			}
		case 2:
			// Cambio nip
			// Los nip son entros de 4 digitos, solo son admitibles este tipo de nip, caso contrario no se permitira realizar el cambio
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
			// Enviamos los datos de la cuenta al cliente
			err = gob.NewEncoder(conn).Encode(cuenta)
			if err != nil {
				log.Println("Error al realizar la codificacion de los datos:", err.Error())
				return
			}
		case 3:
			// Salir
			log.Println("El cliente", cuenta.Nombre, " con numero de cuenta:", cuenta.NoCuenta, "termino su transaccion")
			conn.Close()
			return
		default:
			// En caso de ingresar una operacion invalida no se continua con el ciclo
			// Se envian los datos de la cuenta al cliente ya que de no hacerlo el cliente
			// se quedara colgado
			err = gob.NewEncoder(conn).Encode(cuenta)
			if err != nil {
				log.Println("Error al realizar la codificacion de los datos:", err.Error())
				return
			}
		}
	}
}
