package main

import (
	"html/template"
	"log"
	"net/http"
	// "github.com/gorilla/mux"
)

type Usuario struct {
	Name          string
	Activo        bool
	Administrador bool
	Edad          int
	Array         []string
}

func (this Usuario) Metodo(llave string) bool {
	return this.Activo && this.Administrador && llave == "si"
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("Template/home.html", "Template/footer.html"))
		w.Header().Set("Content-Type", "text/html")

		array := []string{"elkin", "miller", "nore", "abraham", "isa"}

		user := &Usuario{
			Name:          "Abraham",
			Administrador: true,
			Activo:        true,
			Edad:          22,
			Array:         array,
		}

		template.Execute(w, user)
	})

	staticFile := http.FileServer(http.Dir("Assets"))

	http.Handle("/Assets/", http.StripPrefix("/Assets/", staticFile))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
