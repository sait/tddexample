package models

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPsuedoCFDI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PseudoCfdi Test Suite")
}

var _ = Describe("Un cfdi tipo ingreso", func() {

	//se crea una nota inicial
	nota := new(Nota)
	nota.SetCliente("Epifanio de la Torre")
	partida1 := &Partida{Id: 1, Descripcion: "MATABICHOS 400ML", Cantidad: 1, Precio: 25}
	partida2 := &Partida{Id: 2, Descripcion: "TALLADOR TURBO", Cantidad: 50, Precio: 10}
	nota.AddPartida(partida1)
	nota.AddPartida(partida2)

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

		cfdi := nota.ToCFDI()

		Context("un nodo Comprobante", func() {
			It("debe de tener un attributo version con el valor 3.3", func() {

				Expect(cfdi.Version).Should(Equal("3.3"))
			})
			It(" debe de tener un attributo total", func() {})
			Specify(" debe de tener el total de 525", func() {
				Expect(cfdi.Total).Should(Equal("525.00"))
			})
		})

		Context("un nodo Receptor con el nombre del cliente", func() {
			It(" con el valor", func() {
				Expect(cfdi.Receptor.Nombre).To(Equal("Epifanio de la Torre"))
			})
		})

		Context("nodos Concepto", func() {

			Specify("deben de ser 2 nodos", func() {
				Expect(len(cfdi.Conceptos.Concepto)).Should(Equal(2))
			})

			Context("el nodo con indice 0", func() {
				It("la descripcion debe de ser", func() {
					Expect(cfdi.Conceptos.Concepto[0].Descripcion).Should(Equal("MATABICHOS 400ML"))
				})
				It("la cantidad debe de ser", func() {
					Expect(cfdi.Conceptos.Concepto[0].Cantidad).Should(Equal("1.00"))
				})
				It("el precio debe de ser", func() {
					Expect(cfdi.Conceptos.Concepto[0].ValorUnitario).Should(Equal("25.00"))
				})
			})

			Context("el nodo con indice 1", func() {
				It("la descripcion debe de ser", func() {
					Expect(cfdi.Conceptos.Concepto[1].Descripcion).Should(Equal("TALLADOR TURBO"))
				})
				It("la cantidad debe de ser", func() {
					Expect(cfdi.Conceptos.Concepto[1].Cantidad).Should(Equal("50.00"))
				})
				It("el precio debe de ser", func() {
					Expect(cfdi.Conceptos.Concepto[1].ValorUnitario).Should(Equal("10.00"))
				})
			})
		})
	})
})
