package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	"github.com/CajeroAutomatico/client/models"
)

func main() {
	PORT := "8080"
	conn, err := net.Dial("tcp", ":"+PORT)
	if err != nil {
		fmt.Println("Error al conectar con el servidor")
		return
	}
	var noCuenta int
	var nip int
	fmt.Println("Ingrese su numero de cuenta:")
	fmt.Scan(&noCuenta)
	fmt.Println("Ingrese su nip:")
	fmt.Scan(&nip)
	var cuenta = models.Cuenta{
		NoCuenta: noCuenta,
		NIP:      nip,
	}
	err = gob.NewEncoder(conn).Encode(cuenta) //Primera envio de informacion
	if err != nil {
		fmt.Println("Error al realizar codificacion")
		return
	}
	// Recibimos la respuesta en caso de que las credenciales sean
	// validas o no
	var login models.Login
	err = gob.NewDecoder(conn).Decode(&login) //Sgundo envio de informacion
	if err != nil {
		fmt.Println("Error al validar los datos")
	}
	if !login.Acceso {
		fmt.Println("Error en los datos")
		return
	}
	fmt.Println(login)
	// Creamos un ciclo el cual permitira al cliente hacer varia operaciones
	var operacion int
	var dato int
	var transaccion models.Transaccion
	// Limpiamos el buffer
	err = gob.NewEncoder(conn).Encode("Limpiamos buffer")
	for login.Acceso {
		// Leemos los datos de la cuenta
		err = gob.NewDecoder(conn).Decode(&cuenta)
		if err != nil {
			fmt.Println("Error al realizar decodificacion de los datos de la cuenta")
			return
		}
		// Imprimimos informacion de la cuenta
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("Cliente:", cuenta.Nombre)
		fmt.Println("NoCuenta:", cuenta.NoCuenta)
		fmt.Println("Saldo:", cuenta.Saldo)
		// Menu de opciones
		fmt.Println()
		fmt.Println("Opciones")
		fmt.Println()
		fmt.Println("1.-Retirar")
		fmt.Println("2.-Cambiar nip")
		fmt.Println("3.-Salir")
		fmt.Scan(&operacion)
		transaccion.Operacion = operacion
		switch transaccion.Operacion {
		case 1:
			// Retiro
			// Enviamos la transaccion al servidor
			fmt.Println("Ingrse el monto a retirar")
			fmt.Scan(&dato)
			if dato > cuenta.Saldo {
				fmt.Println("Saldo insuficiente")
				time.Sleep(time.Second)
			}
			transaccion.Dato = dato
			err = gob.NewEncoder(conn).Encode(transaccion)
			if err != nil {
				fmt.Println("Error al realizar codificacion de la transaccion")
				return
			}
		case 2:
			// CAmbio nip
			// Enviamos la transaccion al servidor
			fmt.Println("Ingrse el nuevo nip de 4 digitos")
			fmt.Scan(&dato)
			if dato < 1000 || dato > 9999 {
				fmt.Println("NIP invalido")
				time.Sleep(time.Second)
			}
			transaccion.Dato = dato
			err = gob.NewEncoder(conn).Encode(transaccion)
			if err != nil {
				fmt.Println("Error al realizar codificacion de la transaccion")
				return
			}
		case 3:
			transaccion.Dato = 0
			fmt.Println(transaccion)
			// Enviamos la transaccion al servidor
			err = gob.NewEncoder(conn).Encode(transaccion)
			if err != nil {
				fmt.Println("Error al realizar codificacion de la transaccion")
				return
			}

			// Terminamos el ciclo de transaccion
			login.Acceso = false
		default:
			fmt.Println("Opcion invalida")
			err = gob.NewEncoder(conn).Encode(transaccion)
			if err != nil {
				fmt.Println("Error al realizar codificacion de la transaccion")
				return
			}
		}
	}
}
