package main

import (
	"net/http"

	routes "github.com/YuriBertoldi/Go-exloja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
