# Cajero automatico

## Descripcion del proyecto
El programa trata de simular el funcionamiento de un sistema de un cajero automatio de acuerdo al entendimiento del programador.

El sistema considera cuestiones como que los datos del cliente solo pueden ser modificados en las sucursales, al igual que las bajas de las cuentas

El progrma solo realiza conexiones con el sismtema (servidor) que permite realizar las transacciones entre las cuales se encuentran:
* Consultar saldo
* Realizar retiro
* Cambio de NIP
El sistema envia correos de la transaccion realizada al correo del usuario


### Aprendizajes durante la realizacion de la practica
* Al ingresar un caso 0 en el menu del lado del servidor ocacionaba un error, se cree que es debido a que el paquete 
gob toma ese valor como nulo lo cual genera un error al decodificar la transaccion del lado del servidor