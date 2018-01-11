package models

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sait/tddexample/models"
)

func TestNota(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nota Test Suite")
}

var _ = Describe("Una nota de compra", func() {

	Context("Inicialmente Tiene ", func() {
		//se crea una nueva nota
		nota := new(models.Nota)

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
		//se crea una nueva nota
		nota := new(models.Nota)
		//se define el valor del cliente
		nota.SetCliente("Jose Perez Leon")

		It("el atributo de cliente debe de ser igual al asignado", func() {
			Expect(nota.Cliente).Should(Equal("Jose Perez Leon"))
		})
	})

	Context("Cuando se agrega un articulo", func() {
		//se crea una nueva nota
		nota := new(models.Nota)
		//se obtiene total antes de agregar y el no partidas
		totalinicial := nota.GetTotal()
		nopartidas := nota.CountPartidas()

		//se agrega una partida a la nota
		partida1 := &models.Partida{Id: 1, Descripcion: "ABRECUBETAS MAX", Cantidad: 3, Precio: 40.50}
		nota.AddPartida(partida1)

		It(" tiene 1 partida mas", func() {
			Expect(nota.CountPartidas()).Should(Equal(nopartidas + 1))
		})

		It("y el importe de la partida se calcula", func() {
			Expect(partida1.Importe).To(Equal(partida1.Precio * partida1.Cantidad))
		})

		Specify("se acumula el importe de la factura al total de la nota", func() {
			Expect(nota.GetTotal()).Should(Equal(totalinicial + partida1.GetImporte()))
		})
	})

	Context("Cuando se modifica la cantidad de la partida", func() {
		//se crea una nueva nota
		nota := new(models.Nota)
		//se obtiene total antes de agregar
		totalinicial := nota.GetTotal()
		//se agrega una partida a la nota
		partida1 := &models.Partida{Id: 1, Descripcion: "ABRECUBETAS MAX", Cantidad: 3, Precio: 40.50}
		nota.AddPartida(partida1)
		//se obtiene partida a modificar
		partida2 := nota.GetPartida(1)
		partida2.Cantidad = 5
		nota.UpdPartida(partida2)

		It("se re calcula el importe", func() {
			Expect(partida2.GetImporte()).Should(Equal(partida2.Cantidad * partida2.Precio))
		})

		It("se actualiza el total de la nota", func() {
			Expect(nota.GetTotal()).Should(Equal(totalinicial + partida2.GetImporte()))
		})

		Context("si es 0 cero", func() {
			nota.AddPartida(&models.Partida{Id: 2, Descripcion: "LAVAPLATOS 3000 TURBO", Cantidad: 8, Precio: 10})
			partida3 := nota.GetPartida(2) //se obtiene partida a modificar
			partida3.Cantidad = 0
			nota.UpdPartida(partida3)

			It("la partida se elimina", func() {
				Expect(nota.Exist(partida3.Id)).Should(BeFalse())
			})
		})
	})

	Context("Cuando se elimina una partida", func() {
		//se crea una nueva nota
		nota := new(models.Nota)
		//se agrega una partida a la nota
		partida1 := &models.Partida{Id: 1, Descripcion: "ABRECUBETAS MAX", Cantidad: 3, Precio: 40.50}
		nota.AddPartida(partida1)
		//se cuentan las partidas iniciales
		totalinicial := nota.GetTotal()
		partidasiniciales := nota.CountPartidas()
		importepartida := partida1.GetImporte()
		//se elimina partida id=1
		nota.DltPartida(1)

		It("debe de tener una partida menos", func() {
			Expect(nota.CountPartidas()).Should(Equal(partidasiniciales - 1))
		})

		Specify("el total de la nota se actualiza", func() {
			Expect(nota.GetTotal()).To(Equal(totalinicial - importepartida))
		})
	})

})
