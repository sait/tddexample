package nota_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNota(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Nota Test Suite")
}

var _ = Describe("Una nota de compra", func() {

	Context("Inicialmente Tiene ", func() {
		It("cliente vacio", func() {})
		It("0 partidas", func() {})
		Specify("el total es de 0.00", func() {})
	})

	Context("Cuando se define el cliente", func() {
		It("la nota tiene definido el valor dado", func() {})
	})

	Context("Cuando se agrega un articulo", func() {
		It("la nota tiene 1 partida mas", func() {})
		It("el importe de la partida se calcula", func() {})
		Specify("se acumula el importe de la factura al total de la nota", func() {})
	})

	Context("Cuando se modifica la partida", func() {
		It("si la cantidad es cero, la partida se elimina", func() {})
		It("si la cantidad se modifica y es mayor a cero, se re calcula el importe", func() {})
		Specify("si el importe cambia, se actualiza el total de la nota", func() {})
	})

	Context("Cuando se elimina una partida", func() {
		It("la nota debe de tener una partida menos", func() {})
		Specify("el total de la nota se actualiza", func() {})
	})

})
