package models

import (
	"encoding/json"
)

//Estructura Raiz de Nota
type Nota struct {
	Cliente  string
	Partidas []*Partida //arreglo de apuntadores
	Total    float64
}

//Estructura para representar una partida
type Partida struct {
	Id          int
	Descripcion string
	Cantidad    float64
	Precio      float64
	Importe     float64
}

//Metodo para definir valor del atributo cliente
func (n *Nota) SetCliente(cliente string) {
	n.Cliente = cliente
}

//Metodo para obtener el valor del atributo cliente
func (n *Nota) GetCliente() string {
	return n.Cliente
}

//Metodo para obtener el valor de una partida por Id
func (n *Nota) GetPartida(id int) *Partida {
	finded := &Partida{}
	for _, pa := range n.Partidas {
		if id == pa.Id {
			finded = pa
		}
	}
	return finded
}

//Metodo para agregar una partida a una nota
func (n *Nota) AddPartida(p *Partida) {
	p.Calcular()
	n.Partidas = append(n.Partidas, p)
	n.Calcular()
}

//Metodo para actualizar los valores de una partida
func (n *Nota) UpdPartida(p *Partida) {
	if p.Cantidad == 0 {
		n.DltPartida(p.Id)
	}
	for _, pa := range n.Partidas {
		if p.Id == pa.Id {
			p.Calcular()
			pa = p
		}
	}
	n.Calcular()
}

//Metodo para eliminar una partida
func (n *Nota) DltPartida(id int) {
	for i, pa := range n.Partidas {
		if id == pa.Id {
			copy(n.Partidas[i:], n.Partidas[i+1:])
			n.Partidas[len(n.Partidas)-1] = nil // or the zero value of T
			n.Partidas = n.Partidas[:len(n.Partidas)-1]
		}
	}
	n.Calcular()
}

//Metodo para obtener el numero de partidas de una nota
func (n *Nota) CountPartidas() int {
	return len(n.Partidas)
}

//Metodo para saber si existe una partida en una nota
func (n *Nota) Exist(id int) bool {
	finded := false
	for _, pa := range n.Partidas {
		if id == pa.Id {
			finded = true
		}
	}
	return finded
}

//Metodo para calcular los montos de una partida
func (n *Nota) Calcular() float64 {
	n.Total = 0
	for _, p := range n.Partidas {
		n.Total += p.Importe
		// si hay mas calculos como descuento, etc, se hacen aqui
	}
	return n.Total
}

//Metodo para obtener el monto total de una nota
func (n *Nota) GetTotal() float64 {
	return n.Total
}

//Metodo para obtener el importe de una partida
func (p *Partida) GetImporte() float64 {
	return p.Importe
}

//Metodo para calcular los montos de una partida
func (p *Partida) Calcular() float64 {
	p.Importe = 0
	p.Importe = p.Cantidad * p.Precio
	//si hay mas calculos por partida se hacen aqui
	return p.Importe
}

//Metodo para convertir una nota a formato JSON
func (n *Nota) ToJSON() (error, string) {
	b, err := json.Marshal(n)
	if err != nil {
		return err, ""
	}
	return err, string(b)
}
