package models

import (
	"encoding/xml"
	"strconv"
	"time"
	"fmt"
	"io/ioutil"
)

type CFDI struct {
	XMLName     xml.Name    `xml:" Comprobante,omitempty"`
	Version     string      `xml:" Version,attr,omitempty"`
	Total       string      `xml:" Total,attr,omitempty"`
	Receptor    *Receptor   `xml:" Receptor,omitempty"`
	Conceptos   *Conceptos  `xml:" Conceptos,omitempty"`
	Complemento Complemento `xml:" Complemento,omitempty"`
}

type Receptor struct {
	Nombre string `xml:" Nombre,attr,omitempty"`
}

type Conceptos struct {
	Concepto []*Concepto `xml:" Concepto,omitempty"`
	XMLName  xml.Name    `xml:" Conceptos,omitempty"`
}

type Concepto struct {
	XMLName          xml.Name `xml:" Concepto,omitempty"`
	NoIdentificacion string   `xml:" NoIdentificacion,attr,omitempty"`
	Descripcion      string   `xml:" Descripcion,attr,omitempty"`
	Cantidad         string   `xml:" Cantidad,attr,omitempty"`
	ValorUnitario    string   `xml:" ValorUnitario,attr,omitempty"`
	Importe          string   `xml:" Importe,attr,omitempty"`
}

type Complemento struct {
	TimbreFiscalDigital TimbreFiscalDigital `xml:" TimbreFiscalDigital,omitempty"`
	XMLName             xml.Name            `xml:" Complemento,omitempty"`
}

type TimbreFiscalDigital struct {
	FechaTimbrado    string   `xml:" FechaTimbrado,attr, omitempty"`
	NoCertificadoSAT string   `xml:" NoCertificadoSAT,attr,omitempty"`
	SelloCFD         string   `xml:" SelloCFD,attr, omitempty"`
	SelloSAT         string   `xml:" SelloSAT,attr, omitempty"`
	UUID             string   `xml:" UUID,attr, omitempty"`
	Version          string   `xml:" Version,attr, omitempty"`
	XMLName          xml.Name `xml:" TimbreFiscalDigital, omitempty"`
}

func (n *Nota) ToCFDI() *CFDI {

	cfdi := &CFDI{}
	cfdi.Version = "3.3"
	cfdi.Total = floattostr(n.Total)
	cfdi.Receptor = &Receptor{n.Cliente}

	conceptos := new(Conceptos)
	for _, pa := range n.Partidas {
		concepto := new(Concepto)
		concepto.NoIdentificacion = strconv.Itoa(pa.Id)
		concepto.Descripcion = pa.Descripcion
		concepto.ValorUnitario = floattostr(pa.Precio)
		concepto.Cantidad = floattostr(pa.Cantidad)
		concepto.Importe = floattostr(pa.Importe)
		conceptos.Concepto = append(conceptos.Concepto, concepto)
	}
	cfdi.Conceptos = conceptos

	return cfdi
}

func (c *CFDI) ToXML() (string, error) {
	b, err := xml.MarshalIndent(&c, "  ", "   ")
	if err != nil {
		return "", err
	}
	return string(b), err
}

func (c *CFDI) Timbrar() {
	//mock timbrado
	timbre := TimbreFiscalDigital{}
	timbre.Version = "1.2"
	timbre.FechaTimbrado = time.Now().UTC().Format(time.RFC3339)
	timbre.NoCertificadoSAT = "00254556"
	timbre.SelloCFD = "izUxYfQCQMDPTMlYFRUNTaOPBYte1NTw10XdiWFhjQ0X7cG1kPIsyFsi9bhS5oeBj7clOazzGpNUh7RxUsuXF3b2Wlp"
	timbre.UUID = "dc9076e9-2fda-4019-bd2c-900a8284b9c4"
	//
	c.Complemento.TimbreFiscalDigital = timbre

	//se convierte a xml la estructura
	b, _ := c.ToXML()
	//se graba en archivo
	c.ToFile(b, "nota.xml")
}

func (c *CFDI) ToFile(sxml string, filename string) {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	fmt.Println(err)
	//}

	err := ioutil.WriteFile(filename, []byte(sxml), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func floattostr(fv float64) string {
	return strconv.FormatFloat(fv, 'f', 2, 64)
}
