package models

// Cuenta es el modelo que se usara para el acceso a la base de datos
type Cuenta struct {
	Nombre   string
	NoCuenta int
	Saldo    int
	NIP      int
}
