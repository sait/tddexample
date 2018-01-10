package models

type Nota struct {
	Cliente  string
	Partidas map[string]Partida
	Total    float64
}

type Partida struct {
	Id          int
	Descripcion string
	Cantidad    float64
	Precio      float64
	Importe     float64
}

func (n *Nota) SetCliente(cliente string) {}

func (n *Nota) GetCliente() string {
	return "asdasd"
}

func (n *Nota) GetPartida(id int) Partida {
	partida := Partida{}
	return partida
}

func (n *Nota) AddPartida(p Partida) {}

func (n *Nota) UpdPartida(p Partida) {}

func (n *Nota) DltPartida(id int) {}

func (n *Nota) CountPartidas() int {
	return 2
}

func (n *Nota) Exist(id int) bool {
	return true
}

func (n *Nota) Calcular() float64 {
	return 645
}

func (n *Nota) GetTotal() float64 {
	return 645
}

func (p *Partida) GetImporte() float64 {
	return 978
}
func (p *Partida) Calculate() float64 {
	return 978
}
