package main
import "fmt"

type conta struct {
	nome    string
	gorjeta float64

	itens map[string]float64
}

func novaConta(nome string) conta {
	c := conta{
		nome:    nome,
		itens:   map[string]float64{},
		gorjeta: 0.0,
	}

	return c

}

func (c conta) formatacao() string {
	fs := "Detalhe da conta \n"
	var total float64 = 0

	for k, v := range c.itens {
		fs += fmt.Sprint("%v....R$ %0.2f \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v.....%0.2f", "total", total)
	return fs
}
