package main

// Esse arquivo esta chamando o arquivo view.html. No caso, fiz uma rota para que chama a pagina html e exiba na rota respectivamente.

import (
	"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template

type usuarios struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		u := usuarios{
			"João",
			"asgdvasd@gmail.com",
		}

		templates.ExecuteTemplate(w, "view.html", u)
	})

	fmt.Println("Rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
