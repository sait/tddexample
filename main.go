package main

import (
	"fmt"
	"github.com/sait/tddexample/models"
)

func main() {
	fmt.Println("= Hola Mundo TDD =")

	nota := new(models.Nota)

	//Cliente
	nota.SetCliente("Maximo Montes")
	fmt.Println("Cliente Definido :", nota.Cliente)

	//Partida1
	partida1 := new(models.Partida)
	partida1.Id = 1
	partida1.Descripcion = "PODADORA TERMINATOR"
	partida1.Cantidad = 5
	partida1.Precio = 525
	nota.AddPartida(partida1)
	//Mensaje
	fmt.Println("Partida1 Agregada ", partida1)
	fmt.Printf("Total : $ %.2f \n", nota.GetTotal())

	//Partida2
	partida2 := &models.Partida{Id: 2, Descripcion: "ABRECUBETAS 4000 XL", Cantidad: 1, Precio: 75}
	nota.AddPartida(partida2)
	//Mensaje
	fmt.Println("Partida2 Agregada ", partida2)
	fmt.Printf("Total : $ %.2f \n", nota.Total)

	//Partida3 Captura
	fmt.Println("\n ~ Capturar Partida ~")
	partida3 := new(models.Partida)
	fmt.Print("Id: ")
	fmt.Scanf("%d", &partida3.Id)
	fmt.Print("Descripcion: ")
	fmt.Scanf("%s", &partida3.Descripcion)
	fmt.Print("Cantidad: ")
	fmt.Scanf("%f", &partida3.Cantidad)
	fmt.Print("Precio: ")
	fmt.Scanf("%f", &partida3.Precio)
	nota.AddPartida(partida3)
	//Mensaje
	fmt.Println("Partida3 Agregada ", partida3)
	fmt.Printf("Total : $ %.2f \n", nota.Total)

	//Imprimir Nota en JSON
	fmt.Println("----------------JSON-----------------")
	_, snota := nota.ToJSON()
	fmt.Println(snota)
}
