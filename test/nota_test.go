package nota_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"NotaFactura/models"
)

func TestNota(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nota Test Suite")
}

var _ = Describe("Una nota de compra", func() {

	//se crea una nueva nota
	nota := &models.Nota{}

	Context("Inicialmente Tiene ", func() {

		It("cliente vacio", func() {
			Expect(nota.GetCliente()).Should(BeEmpty())
		})

		It("0 partidas", func() {
			Expect(nota.CountPartidas()).Should(BeZero())
		})

		Specify("el total es de 0.00", func() {
			Expect(nota.GetTotal()).Should(BeZero())
		})
	})

	Context("Cuando se define el cliente", func() {

		//se define el valor del cliente
		nota.SetCliente("Jose Perez Leon")

		It("la nota tiene definido el valor dado", func() {
			Expect(nota.Cliente).Should(Equal("Jose Perez Leon"))
		})
	})

	Context("Cuando se agrega un articulo", func() {

		//se obtiene total antes de agregar
		totalnota := nota.GetTotal()

		//se agrega una partida a la nota
		partida1 := models.Partida{Id: 1, Descripcion: "ABRECUBETAS MAX", Cantidad: 3, Precio: 40.50}
		nota.AddPartida(partida1)

		It("la nota tiene 1 partida mas", func() {
			Expect(nota.CountPartidas()).Should(Equal(nota.CountPartidas() + 1))
		})

		It("el importe de la partida se calcula", func() {})
		Specify("se acumula el importe de la factura al total de la nota", func() {
			Expect(nota.GetTotal()).Should(Equal(totalnota + partida1.GetImporte()))
		})
	})

	Context("Cuando se modifica la cantidad de la partida", func() {
		//se  obtiene el total antes de actualizar
		totalnota := nota.GetTotal()

		partida2 := nota.GetPartida(1) //se obtiene partida a modificar
		partida2.Cantidad = 6
		nota.UpdPartida(partida2)

		//se calcula importe que debe de dar
		importe := partida2.Cantidad * partida2.Precio

		It("se re calcula el importe", func() {
			Expect(partida2.Importe).Should(Equal(importe))
		})

		Specify("se actualiza el total de la nota", func() {
			Expect(nota.GetTotal()).Should(Equal(totalnota + importe))
		})

		Context("si es 0 cero", func() {

			partida2 := nota.GetPartida(1) //se obtiene partida a modificar
			partida2.Cantidad = 0
			nota.UpdPartida(partida2)

			It("la partida se elimina", func() {
				Expect(nota.Exist(partida2.Id)).Should(BeFalse())
			})
		})
	})

	Context("Cuando se elimina una partida", func() {

		totalpartidas := nota.CountPartidas()
		nota.DltPartida(1)

		It("la nota debe de tener una partida menos", func() {
			Expect(nota.CountPartidas()).Should(Equal(totalpartidas - 1))
		})

		Specify("el total de la nota se actualiza", func() {})
	})

})
