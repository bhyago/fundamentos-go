package main

import "fmt"

type Pessoa interface {
	Desativar()
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome     string
	Idade    int
	Ativo    bool
	Endereco Endereco // composição
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	client := Cliente{
		Nome:  "Hyago",
		Idade: 20,
		Ativo: true,
	}

	client.Endereco.Cidade = "Rio de janeiro"
	client.Desativar()

	minhaEmpresa := Empresa{}
	Desativacao(client)
	Desativacao(minhaEmpresa)

}
