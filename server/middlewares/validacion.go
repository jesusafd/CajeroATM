package middlewares

import (
	"encoding/gob"
	"log"
	"net"

	"github.com/CajeroAutomatico/server/database"
	"github.com/CajeroAutomatico/server/models"
)

func Validacion(conn net.Conn, handler func(net.Conn, models.Cuenta, models.Login)) {
	var cuenta models.Cuenta
	// leemos los datos que se resivieron de la conexion
	err := gob.NewDecoder(conn).Decode(&cuenta) //Primera envio de informacion
	if err != nil {
		log.Println("Error al decodificar los datos")
		return
	}
	// Verificamos si la cuenta y el nip son validos
	c, err := database.Chequeo(cuenta)
	var login models.Login
	if err != nil {
		login.Acceso = false
		// Enviamos una respuesta al cliente, para evitar que se quede colgado
		e := gob.NewEncoder(conn).Encode(login) //Segundo envio de informacion
		if e != nil {
			log.Println("Error al enviar login de error :", err.Error())
		}
		log.Println("Error al verificar los datos : " + err.Error())
		return
	}
	// En caso de no haber error quiere decir que las credenciales son validas
	// Ceramos un objeto login el cual indicara si el cliente tiene acceso y
	// cuando el cliente desee salir
	login.Acceso = true
	log.Println("Acceso consedido")
	// por lo que enviamos en objeto login con acceso valido
	err = gob.NewEncoder(conn).Encode(login) //Segundo envio de informacion
	handler(conn, c, login)
}
