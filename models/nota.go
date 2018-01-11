package models

import (
	"encoding/json"
)

type Nota struct {
	Cliente  string
	Partidas []*Partida //arreglo de apuntadores
	Total    float64
}

type Partida struct {
	Id          int
	Descripcion string
	Cantidad    float64
	Precio      float64
	Importe     float64
}

func (n *Nota) SetCliente(cliente string) {
	n.Cliente = cliente
}

func (n *Nota) GetCliente() string {
	return n.Cliente
}

func (n *Nota) GetPartida(id int) *Partida {
	finded := &Partida{}
	for _, pa := range n.Partidas {
		if id == pa.Id {
			finded = pa
		}
	}
	return finded
}

func (n *Nota) AddPartida(p *Partida) {
	p.Calcular()
	n.Partidas = append(n.Partidas, p)
	n.Calcular()
}

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

func (n *Nota) CountPartidas() int {
	return len(n.Partidas)
}

func (n *Nota) Exist(id int) bool {
	finded := false
	for _, pa := range n.Partidas {
		if id == pa.Id {
			finded = true
		}
	}
	return finded
}

func (n *Nota) Calcular() float64 {
	n.Total = 0
	for _, p := range n.Partidas {
		n.Total += p.Importe
		// si hay mas calculos como descuento, etc, se hacen aqui
	}
	return n.Total
}

func (n *Nota) GetTotal() float64 {
	return n.Total
}

func (p *Partida) GetImporte() float64 {
	return p.Importe
}
func (p *Partida) Calcular() float64 {
	p.Importe = 0
	p.Importe = p.Cantidad * p.Precio
	//si hay mas calculos por partida se hacen aqui
	return p.Importe
}

func (n *Nota) ToJSON() (error, string) {
	b, err := json.Marshal(n)
	if err != nil {
		return err, ""
	}
	return err, string(b)
}
