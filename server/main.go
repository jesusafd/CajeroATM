package main

import (
	"log"
	"net"

	"github.com/CajeroAutomatico/server/database"
	"github.com/CajeroAutomatico/server/handler"
	"github.com/CajeroAutomatico/server/middlewares"
)

func main() {
	// Establecemos la conexion con la bases de datos
	database.NewConnection()
	defer database.DBConn.Close()
	if database.CheckConnection() {
	}
	// Usaraemos el purto 8080 para poner a la escucha el servidor
	PORT := "8080"
	// Creamos el socket
	listener, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Println("No se pudo crear el socket : " + err.Error())
		return
	}
	// hacemos un ciclo infinito para poder acpetar varias conexiones con el servidor
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("No se pudo establecer conexion : " + err.Error())
			continue
		} else {
		}
		// Creamos una goroutine para ejecutar las transacciones del usurio
		go func() {
			middlewares.Validacion(conn, handler.Handler)
			defer conn.Close()
		}()

	}
}
