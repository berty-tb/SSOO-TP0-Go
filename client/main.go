package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	log.Println("Hola soy un log")
	globals.ClientConfig = utils.IniciarConfiguracion("config.json")
	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración")
	}
	// loggeamos el valor de la config
	log.Println(globals.ClientConfig.Mensaje)

	/////////////////////////////////////////////////////////////////////////////////////
	//acá pongo la parte en la que se lee hasta tocar un enter
	/*
		reader := bufio.NewReader(os.Stdin) //el new reader le asigna a la variable reader el ser un buffer que guarda lo ingresado por el argumento de new reader
		log.Println("Ingrese el texto a enviar junto a la clave:")
		text, _ := reader.ReadString('\n') //guarda en reader toda la cadena ingresada hasta encontrar un \n
		log.Println("El texto ingresado fue: ")
		log.Print(text)
		todo esto ya está incluido en la utils.LeerConsola*/
	/////////////////////////////////////////////////////////////////////////////////////
	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)
	// leer de la consola el mensaje
	// utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	utils.GenerarYEnviarPaquete()
	log.Println("Finaliza el cliente")
}
