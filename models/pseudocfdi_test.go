package models

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sait/tddexample/models"
	"fmt"
)

func TestPsuedoCFDI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PseudoCfdi Test Suite")
}

var _ = Describe("Un cfdi tipo ingreso", func() {

	//se crea una nota inicial
	nota := new(Nota)
	nota.SetCliente("Epifanio de la Torre")
	partida1 := &Partida{Id: 1, Descripcion: "MATABICHOS 400ml", Cantidad: 1, Precio: 25}
	partida2 := &Partida{Id: 2, Descripcion: "TALLADOR TURBO", Cantidad: 50, Precio: 10}
	nota.AddPartida(partida1)
	nota.AddPartida(partida2)

	_,data:=nota.ToXML()
	fmt.Println(data)

	Context("Para poder facturar debe de tener por lo menos ", func() {
		It("cliente", func() {
			Expect(nota.GetCliente()).Should(Not(BeEmpty()))
		})

		It("tener 1 o mas partidas", func() {
			Expect(nota.CountPartidas()).Should(BeNumerically(">", 0))
		})

		Specify("el total debe de ser mayor a 0", func() {
			Expect(nota.GetTotal()).Should(Not(BeZero()))
		})
	})

	Context("En la estructura del CFDI XML debe existir", func() {

		Context("un nodo Comprobante", func() {
			It("debe de tener un attributo version con el valor 3.3", func() {})
			It(" debe de tener un attributo total", func() {})
			Specify(" debe de tener el total de 525", func() {})
		})

		Context("un nodo Receptor con el nombre del cliente", func() {})

		Context("nodos Concepto", func() {
			Specify("deben de ser 2 nodos", func() {})

			Specify("el nodo con indice 0", func() {
				It("la descripcion debe de ser", func() {})
				It("la cantidad debe de ser", func() {})
				It("el precio debe de ser", func() {})
			})

			Specify("el nodo con indice 1", func() {
				It("la descripcion debe de ser", func() {})
				It("la cantidad debe de ser", func() {})
				It("el precio debe de ser", func() {})
			})

			_,data:=nota.ToXML()
			fmt.Println(data)
		})
	})
})
