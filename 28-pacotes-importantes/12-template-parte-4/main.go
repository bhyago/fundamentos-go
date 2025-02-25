package main

import (
	"log"
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

// Carrega o template uma única vez
var tmpl = template.Must(template.ParseFiles("template.html"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, Cursos{
			{"Go", 40},
			{"Java", 20},
			{"Python", 10},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Erro ao executar o template:", err)
		}
	})

	log.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
