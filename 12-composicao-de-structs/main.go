package main

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

func main() {
	client := Cliente{
		Nome:  "Hyago",
		Idade: 20,
		Ativo: true,
	}

	client.Endereco.Cidade = "Rio de janeiro"

}
