package models

// Transaccion es el modelo que se usara para el envio de transaccione entre clite y servidor
type Transaccion struct {
	Operacion int
	Dato      int
}
